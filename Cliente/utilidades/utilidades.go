package utilidades

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "servidor.local/streaming-servidor/serviciosAudio"
)

func DecodificarReproducir(reader io.Reader, canalSincronizacion chan struct{}) {
	streamer, format, err := mp3.Decode(io.NopCloser(reader))
	if err != nil {
		log.Printf("error decodificando MP3: %v", err)
		close(canalSincronizacion)
		return
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/2))

	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		close(canalSincronizacion)
	})))
}

func RecibirAudio(
	stream pb.AudioService_AudioStreamClient,
	writer *io.PipeWriter,
	canalSincronizacion chan struct{}) {
	noFragmento := 0
	for {
		fragmento, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("Canción recibida completa.")
			writer.Close()
			break
		}
		if err != nil {
			if status.Code(err) == codes.Canceled {
				fmt.Println("Reproducción interrumpida por el usuario.")
			} else {
				log.Printf("Error recibiendo chunk: %v", err)
			}
			writer.CloseWithError(err)
			break
		}
		noFragmento++
		fmt.Printf("\nFragmento #%d recibido (%d bytes) reproduciendo ...", noFragmento, len(fragmento.Data))

		if _, err := writer.Write(fragmento.Data); err != nil {
			// Si el usuario detiene justo en este instante.
			break
		}
	}
	<-canalSincronizacion
	fmt.Println("Reproducción finalizada.")
}

func IniciarReproduccion(client pb.AudioServiceClient, nombreArchivo string) (func(), error) {
	ctx, cancel := context.WithCancel(context.Background())

	stream, err := client.AudioStream(ctx, &pb.AudioRequest{NombreArchivo: nombreArchivo})
	if err != nil {
		cancel()
		return nil, fmt.Errorf("error iniciando stream: %w", err)
	}

	pr, pw := io.Pipe()
	canalSincronizacion := make(chan struct{})

	go RecibirAudio(stream, pw, canalSincronizacion)
	go DecodificarReproducir(pr, canalSincronizacion)

	detener := func() {
		speaker.Clear() // corta el audio de inmediato
		cancel()        // stream.Recv() retornará error -> RecibirAudio cierra el pipe por sí solo
	}
	return detener, nil
}
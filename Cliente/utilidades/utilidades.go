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

	pb "servidor.local/streaming-servidor/serviciosAudio"
)

//! DecodificarReproducir y RecibirAudio: tal como en tu práctica de clase,
//! solo se cambió log.Fatalf -> log.Printf (ver explicación arriba).

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
			log.Printf("Error recibiendo chunk: %v", err)
			writer.CloseWithError(err)
			break
		}
		noFragmento++
		fmt.Printf("\nFragmento #%d recibido (%d bytes) reproduciendo ...", noFragmento, len(fragmento.Data))

		if _, err := writer.Write(fragmento.Data); err != nil {
			log.Printf("Error escribiendo en pipe: %v", err)
			break
		}
	}
	// Esperar hasta que termine la reproducción
	<-canalSincronizacion
	fmt.Println("Reproducción finalizada.")
}

//! IniciarReproduccion es lo único NUEVO: envuelve tus dos funciones tal cual
//! están, sin tocar su lógica interna, y agrega la capacidad de "detener" que
//! necesita vista5 (el usuario puede abandonar en cualquier momento).
func IniciarReproduccion(client pb.AudioServiceClient, nombreArchivo string) (func(), error) {
	ctx, cancel := context.WithCancel(context.Background())

	stream, err := client.AudioStream(ctx, &pb.AudioRequest{NombreArchivo: nombreArchivo})
	if err != nil {
		cancel()
		return nil, fmt.Errorf("error iniciando stream: %w", err)
	}

	pr, pw := io.Pipe()
	canalSincronizacion := make(chan struct{})

	// Función de recepción, corriendo en su propio hilo.
	go RecibirAudio(stream, pw, canalSincronizacion)

	// Función de decodificación/reproducción, corriendo en su propio hilo.
	go DecodificarReproducir(pr, canalSincronizacion)

	detener := func() {
		speaker.Clear()  // corta el audio de inmediato, sin esperar el callback
		cancel()         // stream.Recv() retorna error -> RecibirAudio cierra el pipe
		pr.Close()       // libera al decodificador si seguía esperando datos
	}
	return detener, nil
}
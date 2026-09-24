package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "servidor.local/streaming-servidor/serviciosAudio" // ajusta el import si lo pruebas por separado
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewAudioServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	stream, err := client.AudioStream(ctx, &pb.AudioRequest{NombreArchivo: "bohemian_rhapsody.mp3"})
	if err != nil {
		log.Fatal(err)
	}

	out, _ := os.Create("descarga_prueba.mp3")
	defer out.Close()

	for {
		chunk, err := stream.Recv()
		if err != nil {
			break // EOF esperado al terminar
		}
		out.Write(chunk.Data)
		log.Printf("Chunk #%d recibido (%d bytes)\n", chunk.Numero, len(chunk.Data))
	}
	log.Println("Descarga de prueba completa.")
}
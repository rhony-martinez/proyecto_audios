package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"cliente.local/cliente-audios/clienteCola"
	"cliente.local/cliente-audios/vistas"
	pb "servidor.local/streaming-servidor/serviciosAudio"
)

func main() {
	// Conexión gRPC al Servidor de Streaming
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	grpcClient := pb.NewAudioServiceClient(conn)

	// ?Conexión persistente al publicador de RabbitMQ (ajustar la IP de Windows)
	publisher, err := clienteCola.NewClientePublisher("amqp://admin:1234@192.168.0.108:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer publisher.Cerrar()

	vistas.MostrarMenuPrincipal(grpcClient, publisher)
}

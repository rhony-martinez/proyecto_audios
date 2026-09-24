package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"servidor.local/streaming-servidor/capaControladores"
	pb "servidor.local/streaming-servidor/serviciosAudio"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAudioServiceServer(grpcServer, &capaControladores.ControladorServidor{})

	log.Println("Servidor gRPC de Streaming escuchando en :50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
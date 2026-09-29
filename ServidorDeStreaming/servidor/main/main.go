package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"servidor.local/streaming-servidor/capaControladores"
	pb "servidor.local/streaming-servidor/serviciosAudio"
)

func main() {
	// main abre el puerto TCP y registra el controlador del servicio gRPC.
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
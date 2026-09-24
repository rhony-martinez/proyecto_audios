package capaControladores

import (
	"log"

	"servidor.local/streaming-servidor/capaFachada"
	pb "servidor.local/streaming-servidor/serviciosAudio"
)

type ControladorServidor struct {
	pb.UnimplementedAudioServiceServer
}

// AudioStream implementa el procedimiento remoto y delega a la fachada.
func (s *ControladorServidor) AudioStream(req *pb.AudioRequest, stream pb.AudioService_AudioStreamServer) error {
	log.Printf("Eco [gRPC]: AudioStream invocado con nombreArchivo=%s\n", req.NombreArchivo)
	return fachada.EnviarFragmentosAudio(req.NombreArchivo, stream)
}
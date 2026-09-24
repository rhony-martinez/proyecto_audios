package fachada

import (
	"fmt"
	"io"
	"log"

	"servidor.local/streaming-servidor/capaAccesoDatos"
	pb "servidor.local/streaming-servidor/serviciosAudio"
)

// EnviarFragmentosAudio lee el archivo en chunks de 32KB y los envía por el stream gRPC.
func EnviarFragmentosAudio(nombreArchivo string, stream pb.AudioService_AudioStreamServer) error {
	log.Printf("Eco [fachada]: EnviarFragmentosAudio llamado con nombreArchivo=%s\n", nombreArchivo)

	file, err := capaAccesoDatos.AbrirArchivoAudio(nombreArchivo)
	if err != nil {
		return err
	}
	defer file.Close()

	buf := make([]byte, 32*1024) // 32KB por chunk
	chunkNum := int32(0)

	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			log.Println("Eco [fachada]: audio enviado completo:", nombreArchivo)
			break
		}
		if err != nil {
			return fmt.Errorf("error leyendo archivo: %w", err)
		}

		chunkNum++
		if n > 0 {
			objChunk := &pb.AudioChunk{Data: buf[:n], Numero: chunkNum}
			if err := stream.Send(objChunk); err != nil {
				return fmt.Errorf("error enviando chunk #%d: %w", chunkNum, err)
			}
			log.Printf("Eco [fachada]: chunk #%d enviado (%d bytes)\n", chunkNum, n)
		}
	}
	return nil
}
package main

import (
	"github.com/gin-gonic/gin"
	"servidor.local/metadata-servidor/capaAccesoDatos/repository"
	"servidor.local/metadata-servidor/capaControladores"
	"servidor.local/metadata-servidor/capaFachadaServices/fachada"
)

func main() {
	repo := repository.NewMetadataRepository()
	fach := fachada.NewMetadataFachada(repo)
	controller := capaControladores.NewMetadataController(fach)

	router := gin.Default()
	router.GET("/tipos", controller.ListarTipos)
	router.GET("/tipos/:idTipo/audios", controller.ListarAudiosPorTipo)
	router.GET("/audios/musica/:titulo", controller.ConsultarDetalleMusica)
	router.GET("/audios/podcast/:titulo", controller.ConsultarDetallePodcast)
	// router.GET("/audios/audiolibro/:titulo", controller.ConsultarDetalleAudiolibro)
	// router.GET("/audios/ruidoblanco/:titulo", controller.ConsultarDetalleRuidoBlanco)

	router.Run(":8081") // 8081 para no chocar con el 8080 del ejemplo de clase
}

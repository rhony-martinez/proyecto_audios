package main

import (
	"github.com/gin-gonic/gin"
	"servidor.local/audio-servidor/capaAccesoDatos"
	"servidor.local/audio-servidor/capaControladores"
	"servidor.local/audio-servidor/capaFachadaServices/fachada"
)

func main() {
	// main construye las capas del servicio y registra la ruta HTTP para cargar audios.
	repo := capaAccesoDatos.GetRepositorioAudios()
	fach := fachada.NewFachadaAudios(repo)
	controller := capaControladores.NewAudioController(fach)

	router := gin.Default()
	router.POST("/audios/upload", controller.SubirAudio)

	router.Run(":8082")
}
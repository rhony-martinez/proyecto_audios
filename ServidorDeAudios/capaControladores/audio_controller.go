package capaControladores

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"servidor.local/audio-servidor/capaFachadaServices/fachada"
)

type AudioController struct {
	// AudioController adapta las solicitudes HTTP de carga a la fachada de audios.
	fachada *fachada.FachadaAudios
}

func NewAudioController(f *fachada.FachadaAudios) *AudioController {
	// NewAudioController crea el controlador con la fachada que ejecuta las operaciones.
	return &AudioController{fachada: f}
}

// SubirAudio - POST /audios/upload (multipart/form-data, campo "archivo")
// Lo invoca el administrador para almacenar un nuevo mp3.
func (c *AudioController) SubirAudio(ctx *gin.Context) {
	log.Println("Eco [REST]: POST /audios/upload invocado")

	fileHeader, err := ctx.FormFile("archivo")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"mensaje": "no se recibió el archivo: " + err.Error()})
		return
	}

	rutaDestino, err := c.fachada.ValidarYConstruirRuta(fileHeader.Filename)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"mensaje": err.Error()})
		return
	}

	if err := ctx.SaveUploadedFile(fileHeader, rutaDestino); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"mensaje": "error al guardar el archivo: " + err.Error()})
		return
	}

	c.fachada.ConfirmarAlmacenamiento(fileHeader.Filename)
	ctx.JSON(http.StatusCreated, gin.H{
		"mensaje": "Audio almacenado correctamente",
		"archivo": fileHeader.Filename,
	})
}
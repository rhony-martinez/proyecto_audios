package capaControladores

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"servidor.local/metadata-servidor/capaFachadaServices/fachada"
)

type MetadataController struct {
	fachada *fachada.MetadataFachada
}

func NewMetadataController(f *fachada.MetadataFachada) *MetadataController {
	return &MetadataController{fachada: f}
}

// GET /tipos
func (c *MetadataController) ListarTipos(ctx *gin.Context) {
	log.Println("Eco [REST]: GET /tipos invocado")
	ctx.JSON(http.StatusOK, c.fachada.ObtenerTipos())
}

// GET /tipos/:idTipo/audios
func (c *MetadataController) ListarAudiosPorTipo(ctx *gin.Context) {
	idTipo, err := strconv.Atoi(ctx.Param("idTipo"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"mensaje": "idTipo inválido"})
		return
	}
	log.Printf("Eco [REST]: GET /tipos/%d/audios invocado\n", idTipo)
	ctx.JSON(http.StatusOK, c.fachada.ObtenerAudiosPorTipo(idTipo))
}

// GET /audios/musica/:titulo
func (c *MetadataController) ConsultarDetalleMusica(ctx *gin.Context) {
	titulo := ctx.Param("titulo")
	log.Printf("Eco [REST]: GET /audios/musica/%s invocado\n", titulo)
	detalle, encontrado := c.fachada.ObtenerDetalleMusica(titulo)
	if !encontrado {
		ctx.JSON(http.StatusNotFound, gin.H{"mensaje": "Audio no encontrado"})
		return
	}
	ctx.JSON(http.StatusOK, detalle)
}
// ConsultarDetallePodcast, ConsultarDetalleAudiolibro, ConsultarDetalleRuidoBlanco
// siguen el mismo patrón, cada uno bajo su propia ruta: /audios/podcast/:titulo, etc.
// GET /audios/podcast/:titulo
func (c *MetadataController) ConsultarDetallePodcast(ctx *gin.Context) {
	titulo := ctx.Param("titulo")
	log.Printf("Eco [REST]: GET /audios/podcast/%s invocado\n", titulo)
	detalle, encontrado := c.fachada.ObtenerDetallePodcast(titulo)
	if !encontrado {
		ctx.JSON(http.StatusNotFound, gin.H{"mensaje": "Audio no encontrado"})
		return
	}
	ctx.JSON(http.StatusOK, detalle)
}

// GET /audios/audiolibro/:titulo
func (c *MetadataController) ConsultarDetalleAudiolibro(ctx *gin.Context) {
	titulo := ctx.Param("titulo")
	log.Printf("Eco [REST]: GET /audios/audiolibro/%s invocado\n", titulo)
	detalle, encontrado := c.fachada.ObtenerDetalleAudiolibro(titulo)
	if !encontrado {
		ctx.JSON(http.StatusNotFound, gin.H{"mensaje": "Audio no encontrado"})
		return
	}
	ctx.JSON(http.StatusOK, detalle)
}

// GET /audios/ruidoblanco/:titulo
func (c *MetadataController) ConsultarDetalleRuidoBlanco(ctx *gin.Context) {
	titulo := ctx.Param("titulo")
	log.Printf("Eco [REST]: GET /audios/ruidoblanco/%s invocado\n", titulo)
	detalle, encontrado := c.fachada.ObtenerDetalleRuidoBlanco(titulo)
	if !encontrado {
		ctx.JSON(http.StatusNotFound, gin.H{"mensaje": "Audio no encontrado"})
		return
	}
	ctx.JSON(http.StatusOK, detalle)
}
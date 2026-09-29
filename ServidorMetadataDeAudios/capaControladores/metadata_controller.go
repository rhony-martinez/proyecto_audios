package capaControladores

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"servidor.local/metadata-servidor/capaFachadaServices/dto"
	"servidor.local/metadata-servidor/capaFachadaServices/fachada"
)

type MetadataController struct {
	// La fachada contiene las reglas y conversiones fuera de la capa HTTP.
	fachada *fachada.MetadataFachada
}

func NewMetadataController(f *fachada.MetadataFachada) *MetadataController {
	// Se inyecta la fachada para que las rutas deleguen su trabajo en ella.
	return &MetadataController{fachada: f}
}

// GET /tipos
func (c *MetadataController) ListarTipos(ctx *gin.Context) {
	// Responde con las categorías disponibles en formato JSON.
	log.Println("Eco [REST]: GET /tipos invocado")
	ctx.JSON(http.StatusOK, c.fachada.ObtenerTipos())
}

// GET /tipos/:idTipo/audios
func (c *MetadataController) ListarAudiosPorTipo(ctx *gin.Context) {
	// El parámetro de ruta debe ser numérico antes de consultar la fachada.
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
	// Si no hay coincidencia, la API comunica el resultado con HTTP 404.
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
	// Si no hay coincidencia, la API comunica el resultado con HTTP 404.
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
	// Si no hay coincidencia, la API comunica el resultado con HTTP 404.
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
	// Si no hay coincidencia, la API comunica el resultado con HTTP 404.
	titulo := ctx.Param("titulo")
	log.Printf("Eco [REST]: GET /audios/ruidoblanco/%s invocado\n", titulo)
	detalle, encontrado := c.fachada.ObtenerDetalleRuidoBlanco(titulo)
	if !encontrado {
		ctx.JSON(http.StatusNotFound, gin.H{"mensaje": "Audio no encontrado"})
		return
	}
	ctx.JSON(http.StatusOK, detalle)
}

// Registrar
func (c *MetadataController) RegistrarMusica(ctx *gin.Context) {
	// El cuerpo JSON se valida antes de crear la entidad en la fachada.
	var body dto.MusicaRegistroDTO
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"mensaje": "datos inválidos: " + err.Error()})
		return
	}
	log.Println("Eco [REST]: POST /audios/musica invocado")
	c.fachada.RegistrarMusica(body)
	ctx.JSON(http.StatusCreated, gin.H{"mensaje": "Audio registrado correctamente"})
}

func (c *MetadataController) RegistrarPodcast(ctx *gin.Context) {
	// El cuerpo JSON se valida antes de crear la entidad en la fachada.
	var body dto.PodcastRegistroDTO
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"mensaje": "datos inválidos: " + err.Error()})
		return
	}
	log.Println("Eco [REST]: POST /audios/podcast invocado")
	c.fachada.RegistrarPodcast(body)
	ctx.JSON(http.StatusCreated, gin.H{"mensaje": "Audio registrado correctamente"})
}

func (c *MetadataController) RegistrarAudiolibro(ctx *gin.Context) {
	// El cuerpo JSON se valida antes de crear la entidad en la fachada.
	var body dto.AudiolibroRegistroDTO
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"mensaje": "datos inválidos: " + err.Error()})
		return
	}
	log.Println("Eco [REST]: POST /audios/audiolibro invocado")
	c.fachada.RegistrarAudiolibro(body)
	ctx.JSON(http.StatusCreated, gin.H{"mensaje": "Audio registrado correctamente"})
}

func (c *MetadataController) RegistrarRuidoBlanco(ctx *gin.Context) {
	// El cuerpo JSON se valida antes de crear la entidad en la fachada.
	var body dto.RuidoBlancoRegistroDTO
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"mensaje": "datos inválidos: " + err.Error()})
		return
	}
	log.Println("Eco [REST]: POST /audios/ruidoblanco invocado")
	c.fachada.RegistrarRuidoBlanco(body)
	ctx.JSON(http.StatusCreated, gin.H{"mensaje": "Audio registrado correctamente"})
}
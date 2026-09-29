package clienteRest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

const baseURL = "http://localhost:8081"
// baseURL contiene la dirección del servicio REST que publica los metadatos.

type TipoAudioDTO struct {
	// TipoAudioDTO representa una categoría devuelta por el servicio de metadatos.
	IdTipo int    `json:"idTipo"`
	Nombre string `json:"nombre"`
}

type AudioResumenDTO struct {
	// AudioResumenDTO contiene el título que se muestra en una lista de resultados.
	Titulo string `json:"titulo"`
}

type MusicaDetalleDTO struct {
	// MusicaDetalleDTO reúne los datos enviados al consultar una canción.
	Titulo            string `json:"titulo"`
	ArtistaPrincipal  string `json:"artistaPrincipal"`
	Album             string `json:"album"`
	GeneroMusical     string `json:"generoMusical"`
	SelloDiscografico string `json:"selloDiscografico"`
	AnioLanzamiento   int    `json:"anioLanzamiento"`
	Archivo           string `json:"archivo"`
}

type PodcastDetalleDTO struct {
	// PodcastDetalleDTO reúne los datos enviados al consultar un podcast.
	NombrePodcast     string `json:"nombrePodcast"`
	TituloEpisodio    string `json:"tituloEpisodio"`
	Anfitrion         string `json:"anfitrion"`
	TemporadaEpisodio string `json:"temporadaEpisodio"`
	NotasShow         string `json:"notasShow"`
	Clasificacion     string `json:"clasificacion"`
	Archivo           string `json:"archivo"`
}

type AudiolibroDetalleDTO struct {
	// AudiolibroDetalleDTO reúne los datos enviados al consultar un audiolibro.
	TituloLibro string `json:"tituloLibro"`
	Autor       string `json:"autor"`
	Narrador    string `json:"narrador"`
	Editorial   string `json:"editorial"`
	Isbn        string `json:"isbn"`
	Capitulo    string `json:"capitulo"`
	Archivo     string `json:"archivo"`
}

type RuidoBlancoDetalleDTO struct {
	// RuidoBlancoDetalleDTO reúne los datos enviados al consultar ruido blanco.
	TipoSonido          string `json:"tipoSonido"`
	FuenteAudio         string `json:"fuenteAudio"`
	UsoSugerido         string `json:"usoSugerido"`
	ProveedorContenido  string `json:"proveedorContenido"`
	DuracionBucle       string `json:"duracionBucle"`
	FrecuenciaDominante string `json:"frecuenciaDominante"`
	Archivo             string `json:"archivo"`
}

func ObtenerTipos() ([]TipoAudioDTO, error) {
	// ObtenerTipos solicita al servicio REST las categorías de audio disponibles.
	log.Println("Eco [clienteRest]: ObtenerTipos invocado")
	resp, err := http.Get(baseURL + "/tipos")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tipos []TipoAudioDTO
	if err := json.NewDecoder(resp.Body).Decode(&tipos); err != nil {
		return nil, err
	}
	return tipos, nil
}

func ObtenerAudiosPorTipo(idTipo int) ([]AudioResumenDTO, error) {
	// ObtenerAudiosPorTipo pide los títulos asociados al identificador de categoría.
	log.Printf("Eco [clienteRest]: ObtenerAudiosPorTipo llamado con idTipo=%d\n", idTipo)
	resp, err := http.Get(fmt.Sprintf("%s/tipos/%d/audios", baseURL, idTipo))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var audios []AudioResumenDTO
	if err := json.NewDecoder(resp.Body).Decode(&audios); err != nil {
		return nil, err
	}
	return audios, nil
}

// segmentoPorTipo mapea idTipo a la ruta REST correspondiente en el
// Servidor de Metadatos (definida en capaControladores del ese servidor).
func segmentoPorTipo(idTipo int) string {
	switch idTipo {
	case 1:
		return "musica"
	case 2:
		return "podcast"
	case 3:
		return "audiolibro"
	case 4:
		return "ruidoblanco"
	}
	return ""
}

func obtenerDetalleGenerico(segmento, titulo string, destino interface{}) error {
	// obtenerDetalleGenerico consulta un título y decodifica la respuesta en el DTO indicado.
	log.Printf("Eco [clienteRest]: ObtenerDetalle llamado con titulo=%s\n", titulo)
	urlConsulta := fmt.Sprintf("%s/audios/%s/%s", baseURL, segmento, url.PathEscape(titulo))

	resp, err := http.Get(urlConsulta)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("audio no encontrado (código %d)", resp.StatusCode)
	}
	return json.NewDecoder(resp.Body).Decode(destino)
}

func ObtenerDetalleMusica(titulo string) (MusicaDetalleDTO, error) {
	// ObtenerDetalleMusica devuelve los metadatos completos de una canción.
	var d MusicaDetalleDTO
	err := obtenerDetalleGenerico("musica", titulo, &d)
	return d, err
}

func ObtenerDetallePodcast(titulo string) (PodcastDetalleDTO, error) {
	// ObtenerDetallePodcast devuelve los metadatos completos de un podcast.
	var d PodcastDetalleDTO
	err := obtenerDetalleGenerico("podcast", titulo, &d)
	return d, err
}

func ObtenerDetalleAudiolibro(titulo string) (AudiolibroDetalleDTO, error) {
	// ObtenerDetalleAudiolibro devuelve los metadatos completos de un audiolibro.
	var d AudiolibroDetalleDTO
	err := obtenerDetalleGenerico("audiolibro", titulo, &d)
	return d, err
}

func ObtenerDetalleRuidoBlanco(titulo string) (RuidoBlancoDetalleDTO, error) {
	// ObtenerDetalleRuidoBlanco devuelve los metadatos completos de un sonido.
	var d RuidoBlancoDetalleDTO
	err := obtenerDetalleGenerico("ruidoblanco", titulo, &d)
	return d, err
}
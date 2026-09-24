package dto

// TipoAudioDTO viaja al cliente cuando pide la lista de tipos.
type TipoAudioDTO struct {
	IdTipo int    `json:"idTipo"`
	Nombre string `json:"nombre"`
}

// AudioResumenDTO se usa para listar audios de un tipo (solo título).
type AudioResumenDTO struct {
	Titulo string `json:"titulo"`
}

// MusicaDetalleDTO es la respuesta completa al consultar una canción.
type MusicaDetalleDTO struct {
	Titulo            string `json:"titulo"`
	ArtistaPrincipal  string `json:"artistaPrincipal"`
	Album             string `json:"album"`
	GeneroMusical     string `json:"generoMusical"`
	SelloDiscografico string `json:"selloDiscografico"`
	AnioLanzamiento   int    `json:"anioLanzamiento"`
	Archivo           string `json:"archivo"`
}

// Análogos: PodcastDetalleDTO, AudiolibroDetalleDTO, RuidoBlancoDetalleDTO
// con los mismos campos que su entity correspondiente + tag json.
// PodcastDetalleDTO es la respuesta completa al consultar un podcast.
type PodcastDetalleDTO struct {
	NombrePodcast     string `json:"nombrePodcast"`
	TituloEpisodio    string `json:"tituloEpisodio"`
	Anfitrion         string `json:"anfitrion"`
	TemporadaEpisodio string `json:"temporadaEpisodio"`
	NotasShow         string `json:"notasShow"`
	Clasificacion     string `json:"clasificacion"`
	Archivo           string `json:"archivo"`
}

// AudiolibroDetalleDTO es la respuesta completa al consultar un audiolibro.
type AudiolibroDetalleDTO struct {
	TituloLibro string `json:"tituloLibro"`
	Autor       string `json:"autor"`
	Narrador    string `json:"narrador"`
	Editorial   string `json:"editorial"`
	Isbn        string `json:"isbn"`
	Capitulo    string `json:"capitulo"`
	Archivo     string `json:"archivo"`
}

// RuidoBlancoDetalleDTO es la respuesta completa al consultar ruido blanco.
type RuidoBlancoDetalleDTO struct {
	TipoSonido         string `json:"tipoSonido"`
	FuenteAudio        string `json:"fuenteAudio"`
	UsoSugerido        string `json:"usoSugerido"`
	ProveedorContenido string `json:"proveedorContenido"`
	DuracionBucle      string `json:"duracionBucle"`
	FrecuenciaDominante string `json:"frecuenciaDominante"`
	Archivo            string `json:"archivo"`
}

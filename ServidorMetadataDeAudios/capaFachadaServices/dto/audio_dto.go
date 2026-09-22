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

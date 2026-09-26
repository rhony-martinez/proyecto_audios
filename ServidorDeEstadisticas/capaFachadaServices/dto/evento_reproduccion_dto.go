package dto

// EventoReproduccionDTO es lo que el Cliente publicará de forma asíncrona
// cada vez que un usuario reproduzca un audio.
type EventoReproduccionDTO struct {
	Titulo    string `json:"titulo"`
	TipoAudio string `json:"tipoAudio"`
	FechaHora string `json:"fechaHora"`
}
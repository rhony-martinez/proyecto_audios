package entity

// TipoAudio representa una categoría de audio: id entero + nombre.
type TipoAudio struct {
	idTipo int
	nombre string
}

func NewTipoAudio(idTipo int, nombre string) TipoAudio {
	return TipoAudio{idTipo: idTipo, nombre: nombre}
}

func (t *TipoAudio) GetIdTipo() int    { return t.idTipo }
func (t *TipoAudio) GetNombre() string { return t.nombre }
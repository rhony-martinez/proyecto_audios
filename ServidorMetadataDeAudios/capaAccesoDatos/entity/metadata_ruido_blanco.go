package entity

type MetadataRuidoBlanco struct {
	tipoSonido         string
	fuenteAudio        string
	usoSugerido        string
	proveedorContenido string
	duracionBucle      string
	frecuenciaDominante string
	archivo            string
}

func NewMetadataRuidoBlanco(tipoSonido, fuente, uso, proveedor, duracion, frecuencia, archivo string) MetadataRuidoBlanco {
	return MetadataRuidoBlanco{
		tipoSonido: tipoSonido, fuenteAudio: fuente, usoSugerido: uso,
		proveedorContenido: proveedor, duracionBucle: duracion,
		frecuenciaDominante: frecuencia, archivo: archivo,
	}
}

func (r *MetadataRuidoBlanco) GetTipoSonido() string          { return r.tipoSonido }
func (r *MetadataRuidoBlanco) GetFuenteAudio() string         { return r.fuenteAudio }
func (r *MetadataRuidoBlanco) GetUsoSugerido() string         { return r.usoSugerido }
func (r *MetadataRuidoBlanco) GetProveedorContenido() string  { return r.proveedorContenido }
func (r *MetadataRuidoBlanco) GetDuracionBucle() string       { return r.duracionBucle }
func (r *MetadataRuidoBlanco) GetFrecuenciaDominante() string { return r.frecuenciaDominante }
func (r *MetadataRuidoBlanco) GetArchivo() string             { return r.archivo }
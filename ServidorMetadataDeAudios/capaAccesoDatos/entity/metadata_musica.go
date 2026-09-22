package entity

// MetadataMusica encapsula los metadatos de una canción.
type MetadataMusica struct {
	titulo            string
	artistaPrincipal  string
	album             string
	generoMusical     string
	selloDiscografico string
	anioLanzamiento   int
	archivo           string // nombre del .mp3 para el servidor de streaming
}

func NewMetadataMusica(titulo, artista, album, genero, sello string, anio int, archivo string) MetadataMusica {
	return MetadataMusica{
		titulo: titulo, artistaPrincipal: artista, album: album,
		generoMusical: genero, selloDiscografico: sello,
		anioLanzamiento: anio, archivo: archivo,
	}
}

func (m *MetadataMusica) GetTitulo() string            { return m.titulo }
func (m *MetadataMusica) GetArtistaPrincipal() string  { return m.artistaPrincipal }
func (m *MetadataMusica) GetAlbum() string              { return m.album }
func (m *MetadataMusica) GetGeneroMusical() string      { return m.generoMusical }
func (m *MetadataMusica) GetSelloDiscografico() string  { return m.selloDiscografico }
func (m *MetadataMusica) GetAnioLanzamiento() int       { return m.anioLanzamiento }
func (m *MetadataMusica) GetArchivo() string            { return m.archivo }
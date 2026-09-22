package entity

type MetadataPodcast struct {
	nombrePodcast string
	tituloEpisodio string
	anfitrion     string
	temporadaEpisodio string
	notasShow     string
	clasificacion string
	archivo       string
}

func NewMetadataPodcast(nombre, tituloEp, host, tempEp, notas, clasif, archivo string) MetadataPodcast {
	return MetadataPodcast{
		nombrePodcast: nombre, tituloEpisodio: tituloEp, anfitrion: host,
		temporadaEpisodio: tempEp, notasShow: notas, clasificacion: clasif, archivo: archivo,
	}
}

func (p *MetadataPodcast) GetTituloEpisodio() string     { return p.tituloEpisodio }
func (p *MetadataPodcast) GetNombrePodcast() string      { return p.nombrePodcast }
func (p *MetadataPodcast) GetAnfitrion() string          { return p.anfitrion }
func (p *MetadataPodcast) GetTemporadaEpisodio() string  { return p.temporadaEpisodio }
func (p *MetadataPodcast) GetNotasShow() string          { return p.notasShow }
func (p *MetadataPodcast) GetClasificacion() string      { return p.clasificacion }
func (p *MetadataPodcast) GetArchivo() string            { return p.archivo }
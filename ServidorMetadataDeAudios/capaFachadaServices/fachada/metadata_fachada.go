package fachada

import (
	"log"

	"servidor.local/metadata-servidor/capaAccesoDatos/repository"
	"servidor.local/metadata-servidor/capaFachadaServices/dto"
)

type MetadataFachada struct {
	repo *repository.MetadataRepository
}

func NewMetadataFachada(repo *repository.MetadataRepository) *MetadataFachada {
	return &MetadataFachada{repo: repo}
}

func (f *MetadataFachada) ObtenerTipos() []dto.TipoAudioDTO {
	log.Println("Eco [fachada]: ObtenerTipos invocado")
	tipos := f.repo.ListarTipos()
	resultado := make([]dto.TipoAudioDTO, 0, len(tipos))
	for _, t := range tipos {
		resultado = append(resultado, dto.TipoAudioDTO{IdTipo: t.GetIdTipo(), Nombre: t.GetNombre()})
	}
	return resultado
}

// ObtenerAudiosPorTipo recibe el idTipo (1=Música,2=Podcasts,3=Audiolibros,4=RuidoBlanco)
// y retorna solo los títulos, tal como lo necesita la vista3 del cliente.
func (f *MetadataFachada) ObtenerAudiosPorTipo(idTipo int) []dto.AudioResumenDTO {
	log.Printf("Eco [fachada]: ObtenerAudiosPorTipo llamado con idTipo=%d\n", idTipo)
	resultado := []dto.AudioResumenDTO{}
	switch idTipo {
	case 1:
		for _, m := range f.repo.ListarMusica() {
			resultado = append(resultado, dto.AudioResumenDTO{Titulo: m.GetTitulo()})
		}
	case 2:
		for _, p := range f.repo.ListarPodcasts() {
			resultado = append(resultado, dto.AudioResumenDTO{Titulo: p.GetNombrePodcast()})
		}
	case 3:
		for _, a := range f.repo.ListarAudiolibros() {
			resultado = append(resultado, dto.AudioResumenDTO{Titulo: a.GetTituloLibro()})
		}
	case 4:
		for _, rb := range f.repo.ListarRuidoBlanco() {
			resultado = append(resultado, dto.AudioResumenDTO{Titulo: rb.GetTipoSonido()})
		}
	}
	return resultado
}

// ObtenerDetalleMusica retorna el detalle completo de una canción.
func (f *MetadataFachada) ObtenerDetalleMusica(titulo string) (dto.MusicaDetalleDTO, bool) {
	log.Printf("Eco [fachada]: ObtenerDetalleMusica llamado con titulo=%s\n", titulo)
	m, encontrado := f.repo.BuscarMusicaPorTitulo(titulo)
	if !encontrado {
		return dto.MusicaDetalleDTO{}, false
	}
	return dto.MusicaDetalleDTO{
		Titulo: m.GetTitulo(), ArtistaPrincipal: m.GetArtistaPrincipal(),
		Album: m.GetAlbum(), GeneroMusical: m.GetGeneroMusical(),
		SelloDiscografico: m.GetSelloDiscografico(), AnioLanzamiento: m.GetAnioLanzamiento(),
		Archivo: m.GetArchivo(),
	}, true
}

// ObtenerDetallePodcast, ObtenerDetalleAudiolibro, ObtenerDetalleRuidoBlanco
// siguen el mismo patrón que ObtenerDetalleMusica.
// ObtenerDetallePodcast retorna el detalle completo de un podcast.
func (f *MetadataFachada) ObtenerDetallePodcast(nombre string) (dto.PodcastDetalleDTO, bool) {
	log.Printf("Eco [fachada]: ObtenerDetallePodcast llamado con titulo=%s\n", nombre)
	p, encontrado := f.repo.BuscarPodcastPorNombre(nombre)
	if !encontrado {
		return dto.PodcastDetalleDTO{}, false
	}
	return dto.PodcastDetalleDTO{
		NombrePodcast:     p.GetNombrePodcast(),
		TituloEpisodio:    p.GetTituloEpisodio(),
		Anfitrion:         p.GetAnfitrion(),
		TemporadaEpisodio: p.GetTemporadaEpisodio(),
		NotasShow:         p.GetNotasShow(),
		Clasificacion:     p.GetClasificacion(),
		Archivo:           p.GetArchivo(),
	}, true
}

// ObtenerDetalleAudiolibro retorna el detalle completo de un audiolibro.
func (f *MetadataFachada) ObtenerDetalleAudiolibro(titulo string) (dto.AudiolibroDetalleDTO, bool) {
	log.Printf("Eco [fachada]: ObtenerDetalleAudiolibro llamado con titulo=%s\n", titulo)
	a, encontrado := f.repo.BuscarAudiolibroPorTitulo(titulo)
	if !encontrado {
		return dto.AudiolibroDetalleDTO{}, false
	}
	return dto.AudiolibroDetalleDTO{
		TituloLibro: a.GetTituloLibro(),
		Autor:       a.GetAutor(),
		Narrador:    a.GetNarrador(),
		Editorial:   a.GetEditorial(),
		Isbn:        a.GetIsbn(),
		Capitulo:    a.GetCapitulo(),
		Archivo:     a.GetArchivo(),
	}, true
}

// ObtenerDetalleRuidoBlanco retorna el detalle completo de un audio de ruido blanco.
func (f *MetadataFachada) ObtenerDetalleRuidoBlanco(titulo string) (dto.RuidoBlancoDetalleDTO, bool) {
	log.Printf("Eco [fachada]: ObtenerDetalleRuidoBlanco llamado con titulo=%s\n", titulo)
	rb, encontrado := f.repo.BuscarRuidoBlancoPorTitulo(titulo)
	if !encontrado {
		return dto.RuidoBlancoDetalleDTO{}, false
	}
	return dto.RuidoBlancoDetalleDTO{
		TipoSonido:          rb.GetTipoSonido(),
		FuenteAudio:         rb.GetFuenteAudio(),
		UsoSugerido:         rb.GetUsoSugerido(),
		ProveedorContenido:  rb.GetProveedorContenido(),
		DuracionBucle:       rb.GetDuracionBucle(),
		FrecuenciaDominante: rb.GetFrecuenciaDominante(),
		Archivo:             rb.GetArchivo(),
	}, true
}

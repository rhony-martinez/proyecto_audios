package repository

import (
	"log"

	"servidor.local/metadata-servidor/capaAccesoDatos/entity"
)

// MetadataRepository mantiene en memoria los tipos de audio y sus metadatos.
type MetadataRepository struct {
	tipos       []entity.TipoAudio
	musica      []entity.MetadataMusica
	podcasts    []entity.MetadataPodcast
	audiolibros []entity.MetadataAudiolibro
	ruidoBlanco []entity.MetadataRuidoBlanco
}

// NewMetadataRepository crea el repositorio y lo precarga con datos de ejemplo
// (al menos 2 audios por tipo, como exige el requerimiento).
func NewMetadataRepository() *MetadataRepository {
	r := &MetadataRepository{}
	r.cargarTipos()
	r.cargarMusica()
	r.cargarPodcasts()
	r.cargarAudiolibros()
	r.cargarRuidoBlanco()
	return r
}

func (r *MetadataRepository) cargarTipos() {
	r.tipos = []entity.TipoAudio{
		entity.NewTipoAudio(1, "Música"),
		entity.NewTipoAudio(2, "Podcasts"),
		entity.NewTipoAudio(3, "Audiolibros"),
		entity.NewTipoAudio(4, "Ruido Blanco"),
	}
}

func (r *MetadataRepository) cargarMusica() {
	r.musica = []entity.MetadataMusica{
		entity.NewMetadataMusica("Bohemian Rhapsody", "Queen", "A Night at the Opera", "Rock", "EMI", 1975, "bohemian_rhapsody.mp3"),
		entity.NewMetadataMusica("Blinding Lights", "The Weeknd", "After Hours", "Pop", "XO/Republic", 2020, "blinding_lights.mp3"),
	}
}

func (r *MetadataRepository) cargarPodcasts() {
	r.podcasts = []entity.MetadataPodcast{
		entity.NewMetadataPodcast("Radio Ambulante", "El silencio", "Daniel Alarcón", "T5 E1", "Historias de América Latina", "Para toda la familia", "radio_ambulante_ep1.mp3"),
		entity.NewMetadataPodcast("Cracks Podcast", "Innovación en LatAm", "Oso Trava", "T2 E10", "Charla sobre startups", "Explícito", "cracks_ep10.mp3"),
	}
}

func (r *MetadataRepository) cargarAudiolibros() {
	r.audiolibros = []entity.MetadataAudiolibro{
		entity.NewMetadataAudiolibro("Cien años de soledad", "Gabriel García Márquez", "Gustavo Bonfigli", "Penguin Random House", "978-0307474728", "Cap. 1", "cien_anios_soledad.mp3"),
		entity.NewMetadataAudiolibro("Harry Potter y la Piedra Filosofal", "J.K. Rowling", "Jim Dale", "Salamandra", "978-8478884452", "Cap. 1", "harry_potter_1.mp3"),
	}
}

func (r *MetadataRepository) cargarRuidoBlanco() {
	r.ruidoBlanco = []entity.MetadataRuidoBlanco{
		entity.NewMetadataRuidoBlanco("Ruido Blanco", "Lluvia", "Dormir", "Calm Sounds", "60 min", "Graves", "lluvia_loop.mp3"),
		entity.NewMetadataRuidoBlanco("Ruido Marrón", "Bosque", "Concentración", "Nature Loops", "45 min", "Graves", "bosque_loop.mp3"),
	}
}

// --- Operaciones de consulta ---

func (r *MetadataRepository) ListarTipos() []entity.TipoAudio {
	log.Println("Eco [capaAccesoDatos]: ListarTipos invocado")
	return r.tipos
}

func (r *MetadataRepository) ListarMusica() []entity.MetadataMusica {
	log.Println("Eco [capaAccesoDatos]: ListarMusica invocado")
	return r.musica
}
func (r *MetadataRepository) ListarPodcasts() []entity.MetadataPodcast {
	log.Println("Eco [capaAccesoDatos]: ListarPodcasts invocado")
	return r.podcasts
}
func (r *MetadataRepository) ListarAudiolibros() []entity.MetadataAudiolibro {
	log.Println("Eco [capaAccesoDatos]: ListarAudiolibros invocado")
	return r.audiolibros
}
func (r *MetadataRepository) ListarRuidoBlanco() []entity.MetadataRuidoBlanco {
	log.Println("Eco [capaAccesoDatos]: ListarRuidoBlanco invocado")
	return r.ruidoBlanco
}

func (r *MetadataRepository) BuscarMusicaPorTitulo(titulo string) (entity.MetadataMusica, bool) {
	for _, m := range r.musica {
		if m.GetTitulo() == titulo {
			return m, true
		}
	}
	return entity.MetadataMusica{}, false
}

// BuscarPodcastPorTitulo, BuscarAudiolibroPorTitulo, BuscarRuidoBlancoPorTitulo
// siguen exactamente el mismo patrón que BuscarMusicaPorTitulo de arriba.
func (r *MetadataRepository) BuscarPodcastPorNombre(nombre string) (entity.MetadataPodcast, bool) {
	for _, p := range r.podcasts {
		if p.GetNombrePodcast() == nombre {
			return p, true
		}
	}
	return entity.MetadataPodcast{}, false
}

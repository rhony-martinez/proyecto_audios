package repository

import (
	"log"
	"sync"

	"servidor.local/metadata-servidor/capaAccesoDatos/entity"
)

// MetadataRepository mantiene en memoria los tipos de audio y sus metadatos.
type MetadataRepository struct {
	mu          sync.Mutex
	tipos       []entity.TipoAudio
	musica      []entity.MetadataMusica
	podcasts    []entity.MetadataPodcast
	audiolibros []entity.MetadataAudiolibro
	ruidoBlanco []entity.MetadataRuidoBlanco
}

// NewMetadataRepository crea el repositorio y lo precarga con datos de ejemplo
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
		entity.NewMetadataMusica("La funa", "AlcolirykoZ", "sencillo", "Hip Hop", "El Arkeólogo", 2025, "La funa.mp3"),
		entity.NewMetadataMusica("El Remate", "AlcolirykoZ", "Anarcolirykoz", "Hip Hop", "El Arkeólogo", 2022, "El Remate.mp3"),
	}
}

func (r *MetadataRepository) cargarPodcasts() {
	r.podcasts = []entity.MetadataPodcast{
		entity.NewMetadataPodcast("Radio Ambulante", "Romper el silencio", "Daniel Alarcón", "Temporada 14", "Después de 20 años, se sientan a hablar. En octubre de 2001, Oswaldo Díaz fue secuestrado y esa tragedia atormentaría a su familia durante años. Pero cuando los victimarios firmaron un acuerdo de paz, la familia de Oswaldo se vio obligada a confrontarlos y a considerar el costo del perdón.", "Apto para todo público / Contenido periodístico", "romper_el_silencio.mp3"),
		entity.NewMetadataPodcast("The Wild Project", "Daniel Brunner (Ex FBI) | Así te hacen CONFESAR, Interrogar a psicópatas", "Jordi Wild", "Episodio 383", "Daniel Brunner, con 20 años en el FBI y experiencia en SWAT, revela las técnicas psicológicas usadas en interrogatorios para lograr confesiones de criminales y psicópatas.", "Mayores de 16 años / Contenido criminalístico", "the_wild_project_383.mp3"),
	}
}

func (r *MetadataRepository) cargarAudiolibros() {
	r.audiolibros = []entity.MetadataAudiolibro{
		entity.NewMetadataAudiolibro("Fábulas de Esopo", "Esopo", "Luis Ignacio González", "Penguin Random House Audio", "978-8491056621", "La zorra y las uvas", "fabulas_esopo.mp3"),
		entity.NewMetadataAudiolibro("Alí Babá y los cuarenta ladrones", "Anónimo", "Arturo López", "Kobo Audiolibros", "978-8424115005", "Ábrete sésamo", "alibaba_y_los_40_ladrones.mp3"),
	}
}

func (r *MetadataRepository) cargarRuidoBlanco() {
	r.ruidoBlanco = []entity.MetadataRuidoBlanco{
		entity.NewMetadataRuidoBlanco("Ruido Blanco Estático", "Frecuencias Profundas", "Relajación y Concentración", "TheMediaGuy", "5:00 min", "Graves / Subgraves", "soft_soothing_deep_white_noise_378857.mp3"),
		entity.NewMetadataRuidoBlanco("Ruido Marrón Ambientado", "Lluvia Urbana y Aves", "Dormir, Meditar y Bloquear Ruido", "WhiteNoiseSleepers", "9:41 min", "Graves Intensos / Frecuencias Bajas", "rainy_day_in_town_with_birds_singing_194011.mp3"),
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

func (r *MetadataRepository) BuscarPodcastPorNombre(nombre string) (entity.MetadataPodcast, bool) {
	for _, p := range r.podcasts {
		if p.GetNombrePodcast() == nombre {
			return p, true
		}
	}
	return entity.MetadataPodcast{}, false
}

func (r *MetadataRepository) BuscarAudiolibroPorTitulo(titulo string) (entity.MetadataAudiolibro, bool) {
	for _, a := range r.audiolibros {
		if a.GetTituloLibro() == titulo {
			return a, true
		}
	}
	return entity.MetadataAudiolibro{}, false
}

func (r *MetadataRepository) BuscarRuidoBlancoPorTitulo(titulo string) (entity.MetadataRuidoBlanco, bool) {
	for _, rb := range r.ruidoBlanco {
		if rb.GetTipoSonido() == titulo {
			return rb, true
		}
	}
	return entity.MetadataRuidoBlanco{}, false
}

// --- Operaciones de registro ---
func (r *MetadataRepository) RegistrarMusica(m entity.MetadataMusica) {
	r.mu.Lock()
	defer r.mu.Unlock()
	log.Println("Eco [capaAccesoDatos]: RegistrarMusica invocado, titulo=", m.GetTitulo())
	r.musica = append(r.musica, m)
}

func (r *MetadataRepository) RegistrarPodcast(p entity.MetadataPodcast) {
	r.mu.Lock()
	defer r.mu.Unlock()
	log.Println("Eco [capaAccesoDatos]: RegistrarPodcast invocado, titulo=", p.GetNombrePodcast())
	r.podcasts = append(r.podcasts, p)
}

func (r *MetadataRepository) RegistrarAudiolibro(a entity.MetadataAudiolibro) {
	r.mu.Lock()
	defer r.mu.Unlock()
	log.Println("Eco [capaAccesoDatos]: RegistrarAudiolibro invocado, titulo=", a.GetTituloLibro())
	r.audiolibros = append(r.audiolibros, a)
}

func (r *MetadataRepository) RegistrarRuidoBlanco(rb entity.MetadataRuidoBlanco) {
	r.mu.Lock()
	defer r.mu.Unlock()
	log.Println("Eco [capaAccesoDatos]: RegistrarRuidoBlanco invocado, titulo=", rb.GetTipoSonido())
	r.ruidoBlanco = append(r.ruidoBlanco, rb)
}
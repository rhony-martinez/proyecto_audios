package capaAccesoDatos

import (
	"log"
	"sync"

	"servidor.local/estadisticas-servidor/capaFachadaServices/dto"
)

type RepositorioEstadisticas struct {
	// RepositorioEstadisticas mantiene el total de reproducciones acumuladas por título.
	mu     sync.Mutex
	conteo map[string]int // titulo -> número de reproducciones acumuladas
}

func NewRepositorioEstadisticas() *RepositorioEstadisticas {
	// NewRepositorioEstadisticas inicia un repositorio con el contador vacío.
	return &RepositorioEstadisticas{conteo: make(map[string]int)}
}

// Registrar incrementa el contador de la canción y retorna el nuevo total.
func (r *RepositorioEstadisticas) Registrar(evento dto.EventoReproduccionDTO) int {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.conteo[evento.Titulo]++
	total := r.conteo[evento.Titulo]
	log.Printf("Eco [capaAccesoDatos]: %s ahora tiene %d reproducciones\n", evento.Titulo, total)
	return total
}
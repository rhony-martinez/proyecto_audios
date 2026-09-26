package capaAccesoDatos

import (
	"log"
	"sync"

	"servidor.local/estadisticas-servidor/capaFachadaServices/dto"
)

type RepositorioEstadisticas struct {
	mu      sync.Mutex
	eventos []dto.EventoReproduccionDTO
}

func NewRepositorioEstadisticas() *RepositorioEstadisticas {
	return &RepositorioEstadisticas{}
}

func (r *RepositorioEstadisticas) Registrar(evento dto.EventoReproduccionDTO) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.eventos = append(r.eventos, evento)
	log.Printf("Eco [capaAccesoDatos]: evento almacenado, total acumulado=%d\n", len(r.eventos))
}
package fachada

import (
	"encoding/json"
	"fmt"
	"log"

	"servidor.local/estadisticas-servidor/capaAccesoDatos"
	"servidor.local/estadisticas-servidor/capaFachadaServices/dto"
)

type FachadaEstadisticas struct {
	repo *capaAccesoDatos.RepositorioEstadisticas
}

func NewFachadaEstadisticas(repo *capaAccesoDatos.RepositorioEstadisticas) *FachadaEstadisticas {
	return &FachadaEstadisticas{repo: repo}
}

func (f *FachadaEstadisticas) ProcesarMensaje(body []byte) {
	log.Println("Eco [fachada]: ProcesarMensaje invocado")
	var evento dto.EventoReproduccionDTO
	if err := json.Unmarshal(body, &evento); err != nil {
		log.Println("Error decodificando evento:", err)
		return
	}
	total := f.repo.Registrar(evento)
	fmt.Printf(">> %s (%s) — reproducción #%d — %s\n",
		evento.Titulo, evento.TipoAudio, total, evento.FechaHora)
}
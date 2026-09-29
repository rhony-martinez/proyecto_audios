package fachada

import (
	"encoding/json"
	"fmt"
	"log"

	"servidor.local/estadisticas-servidor/capaAccesoDatos"
	"servidor.local/estadisticas-servidor/capaFachadaServices/dto"
)

type FachadaEstadisticas struct {
	// FachadaEstadisticas decodifica eventos y coordina el registro de reproducciones.
	repo *capaAccesoDatos.RepositorioEstadisticas
}

func NewFachadaEstadisticas(repo *capaAccesoDatos.RepositorioEstadisticas) *FachadaEstadisticas {
	// NewFachadaEstadisticas configura la fachada con el repositorio de contadores.
	return &FachadaEstadisticas{repo: repo}
}

func (f *FachadaEstadisticas) ProcesarMensaje(body []byte) {
	// ProcesarMensaje valida el evento JSON, incrementa su contador y muestra el total.
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
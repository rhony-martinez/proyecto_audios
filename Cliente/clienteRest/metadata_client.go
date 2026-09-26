package clienteRest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

const baseURL = "http://localhost:8081"

type TipoAudioDTO struct {
	IdTipo int    `json:"idTipo"`
	Nombre string `json:"nombre"`
}

type AudioResumenDTO struct {
	Titulo string `json:"titulo"`
}

func ObtenerTipos() ([]TipoAudioDTO, error) {
	log.Println("Eco [clienteRest]: ObtenerTipos invocado")
	resp, err := http.Get(baseURL + "/tipos")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tipos []TipoAudioDTO
	if err := json.NewDecoder(resp.Body).Decode(&tipos); err != nil {
		return nil, err
	}
	return tipos, nil
}

func ObtenerAudiosPorTipo(idTipo int) ([]AudioResumenDTO, error) {
	log.Printf("Eco [clienteRest]: ObtenerAudiosPorTipo llamado con idTipo=%d\n", idTipo)
	resp, err := http.Get(fmt.Sprintf("%s/tipos/%d/audios", baseURL, idTipo))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var audios []AudioResumenDTO
	if err := json.NewDecoder(resp.Body).Decode(&audios); err != nil {
		return nil, err
	}
	return audios, nil
}

// segmentoPorTipo mapea idTipo a la ruta REST correspondiente en el
// Servidor de Metadatos (definida en capaControladores del ese servidor).
func segmentoPorTipo(idTipo int) string {
	switch idTipo {
	case 1:
		return "musica"
	case 2:
		return "podcast"
	case 3:
		return "audiolibro"
	case 4:
		return "ruidoblanco"
	}
	return ""
}

// ObtenerDetalle retorna el detalle completo como mapa genérico, ya que cada
// tipo de audio tiene metadatos distintos (definidos en el DTO del servidor).
func ObtenerDetalle(idTipo int, titulo string) (map[string]interface{}, error) {
	log.Printf("Eco [clienteRest]: ObtenerDetalle llamado con titulo=%s\n", titulo)
	segmento := segmentoPorTipo(idTipo)
	urlConsulta := fmt.Sprintf("%s/audios/%s/%s", baseURL, segmento, url.PathEscape(titulo))

	resp, err := http.Get(urlConsulta)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("audio no encontrado (código %d)", resp.StatusCode)
	}

	var detalle map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&detalle); err != nil {
		return nil, err
	}
	return detalle, nil
}
package fachada

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"servidor.local/audio-servidor/capaAccesoDatos"
)

type FachadaAudios struct {
	repo *capaAccesoDatos.RepositorioAudios
}

func NewFachadaAudios(repo *capaAccesoDatos.RepositorioAudios) *FachadaAudios {
	return &FachadaAudios{repo: repo}
}

// ValidarYConstruirRuta valida que el archivo sea .mp3 y arma la ruta final de almacenamiento.
func (f *FachadaAudios) ValidarYConstruirRuta(nombreOriginal string) (string, error) {
	log.Printf("Eco [fachada]: ValidarYConstruirRuta llamado con nombreOriginal=%s\n", nombreOriginal)
	if strings.ToLower(filepath.Ext(nombreOriginal)) != ".mp3" {
		return "", fmt.Errorf("solo se permiten archivos .mp3")
	}
	return f.repo.RutaDestino(nombreOriginal), nil
}

func (f *FachadaAudios) ConfirmarAlmacenamiento(nombreArchivo string) {
	log.Printf("Eco [fachada]: audio almacenado correctamente: %s\n", nombreArchivo)
}
package capaAccesoDatos

import (
	"log"
	"os"
	"path/filepath"
	"sync"
)

// carpetaAudios es relativa a la raíz de ServidorDeAudios (donde corres "go run main/main.go").
const carpetaAudios = "audios"

type RepositorioAudios struct {
	mu sync.Mutex
}

var (
	instancia *RepositorioAudios
	once      sync.Once
)

// GetRepositorioAudios aplica patrón Singleton, igual que RepositorioCanciones en tu guía.
func GetRepositorioAudios() *RepositorioAudios {
	once.Do(func() {
		os.MkdirAll(carpetaAudios, os.ModePerm)
		instancia = &RepositorioAudios{}
	})
	return instancia
}

func (r *RepositorioAudios) RutaDestino(nombreArchivo string) string {
	return filepath.Join(carpetaAudios, nombreArchivo)
}

func (r *RepositorioAudios) ExisteArchivo(nombreArchivo string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, err := os.Stat(r.RutaDestino(nombreArchivo))
	existe := err == nil
	log.Printf("Eco [capaAccesoDatos]: ExisteArchivo(%s) = %v\n", nombreArchivo, existe)
	return existe
}
package capaAccesoDatos

import (
	"fmt"
	"os"
	"path/filepath"
)

// rutaAudios apunta a la carpeta compartida donde el Servidor de Audios
// guarda los mp3. Ajusta esta ruta según dónde quede tu carpeta compartida.
const rutaAudios = "../../ServidorDeAudios/audios"

// AbrirArchivoAudio abre el mp3 dado su nombre de archivo y lo devuelve como *os.File.
func AbrirArchivoAudio(nombreArchivo string) (*os.File, error) {
	ruta := filepath.Join(rutaAudios, nombreArchivo)
	file, err := os.Open(ruta)
	if err != nil {
		fmt.Println("Eco [capaAccesoDatos]: error al abrir audio:", nombreArchivo)
		return nil, fmt.Errorf("no se pudo abrir el archivo: %w", err)
	}
	fmt.Println("Eco [capaAccesoDatos]: audio abierto correctamente:", nombreArchivo)
	return file, nil
}
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var lector = bufio.NewReader(os.Stdin)

func leer(mensaje string) string {
	fmt.Print(mensaje)
	texto, _ := lector.ReadString('\n')
	return strings.TrimSpace(texto)
}

// subirAudio envía el mp3 al Servidor de Audios (REST) y retorna el nombre con
// el que quedó almacenado.
func subirAudio(rutaLocal string) (string, error) {
	file, err := os.Open(rutaLocal)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	nombreArchivo := filepath.Base(rutaLocal)
	part, _ := writer.CreateFormFile("archivo", nombreArchivo)
	io.Copy(part, file)
	writer.Close()

	resp, err := http.Post("http://localhost:8082/audios/upload", writer.FormDataContentType(), &buf)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		cuerpo, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("error subiendo audio: %s", cuerpo)
	}
	return nombreArchivo, nil
}

// registrarMetadata envía los metadatos al Servidor de Metadatos (REST).
func registrarMetadata(segmento string, body map[string]interface{}) error {
	data, _ := json.Marshal(body)
	resp, err := http.Post("http://localhost:8081/audios/"+segmento, "application/json", bytes.NewReader(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		cuerpo, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("error registrando metadata: %s", cuerpo)
	}
	return nil
}

func main() {
	fmt.Println("=== Panel de Administrador ===")
	rutaLocal := leer("Ruta del archivo .mp3 en tu disco: ")

	fmt.Println("Tipo de audio: 1.Música 2.Podcast 3.Audiolibro 4.Ruido Blanco")
	tipo := leer("Seleccione tipo: ")

	nombreArchivo, err := subirAudio(rutaLocal)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Audio almacenado en el servidor de audios como:", nombreArchivo)

	var segmento string
	var body map[string]interface{}

	switch tipo {
	case "1":
		segmento = "musica"
		anio, _ := strconv.Atoi(leer("Año de lanzamiento: "))
		body = map[string]interface{}{
			"titulo": leer("Título: "), "artistaPrincipal": leer("Artista Principal: "),
			"album": leer("Álbum: "), "generoMusical": leer("Género Musical: "),
			"selloDiscografico": leer("Sello Discográfico: "), "anioLanzamiento": anio,
			"archivo": nombreArchivo,
		}
	case "2":
		segmento = "podcast"
		body = map[string]interface{}{
			"nombrePodcast": leer("Nombre del Podcast: "), "tituloEpisodio": leer("Título del Episodio: "),
			"anfitrion": leer("Anfitrión: "), "temporadaEpisodio": leer("Temporada/Episodio: "),
			"notasShow": leer("Notas del Show: "), "clasificacion": leer("Clasificación: "),
			"archivo": nombreArchivo,
		}
	case "3":
		segmento = "audiolibro"
		body = map[string]interface{}{
			"tituloLibro": leer("Título del Libro: "), "autor": leer("Autor: "),
			"narrador": leer("Narrador: "), "editorial": leer("Editorial: "),
			"isbn": leer("ISBN: "), "capitulo": leer("Capítulo: "), "archivo": nombreArchivo,
		}
	case "4":
		segmento = "ruidoblanco"
		body = map[string]interface{}{
			"tipoSonido": leer("Tipo de Sonido: "), "fuenteAudio": leer("Fuente del Audio: "),
			"usoSugerido": leer("Uso Sugerido: "), "proveedorContenido": leer("Proveedor de Contenido: "),
			"duracionBucle": leer("Duración del Bucle: "), "frecuenciaDominante": leer("Frecuencia Dominante: "),
			"archivo": nombreArchivo,
		}
	default:
		fmt.Println("Tipo inválido.")
		return
	}

	if err := registrarMetadata(segmento, body); err != nil {
		fmt.Println("Error registrando metadata:", err)
		return
	}
	fmt.Println("Metadata registrada. El audio ya es visible en el catálogo del cliente.")
}
package vistas

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"cliente.local/cliente-audios/clienteCola"
	"cliente.local/cliente-audios/clienteRest"
	"cliente.local/cliente-audios/utilidades"
	pb "servidor.local/streaming-servidor/serviciosAudio"
)

var lector = bufio.NewReader(os.Stdin)

func leerOpcion() string {
	texto, _ := lector.ReadString('\n')
	return strings.TrimSpace(texto)
}

// MostrarMenuPrincipal - vista 1
func MostrarMenuPrincipal(grpcClient pb.AudioServiceClient, publisher *clienteCola.ClientePublisher) {
	for {
		fmt.Println("\n=== Spotify ===")
		fmt.Println("1. Ver tipo de audio")
		fmt.Println("2. Salir")
		fmt.Print("Seleccione una opción: ")

		switch leerOpcion() {
		case "1":
			mostrarTiposAudio(grpcClient, publisher)
		case "2":
			fmt.Println("Hasta pronto.")
			return
		default:
			fmt.Println("Opción inválida.")
		}
	}
}

// mostrarTiposAudio - vista 2
func mostrarTiposAudio(grpcClient pb.AudioServiceClient, publisher *clienteCola.ClientePublisher) {
	tipos, err := clienteRest.ObtenerTipos()
	if err != nil {
		log.Println("Error obteniendo tipos:", err)
		return
	}

	for {
		fmt.Println("\n=== Spotify ===")
		for i, t := range tipos {
			fmt.Printf("%d. %s\n", i+1, t.Nombre)
		}
		opcionAtras := len(tipos) + 1
		fmt.Printf("%d. Atrás\n", opcionAtras)
		fmt.Print("Seleccione una opción: ")

		opcion, err := strconv.Atoi(leerOpcion())
		if err != nil || opcion < 1 || opcion > opcionAtras {
			fmt.Println("Opción inválida.")
			continue
		}
		if opcion == opcionAtras {
			return
		}
		mostrarListaAudios(tipos[opcion-1], grpcClient, publisher)
	}
}

// mostrarListaAudios - vista 3
func mostrarListaAudios(tipo clienteRest.TipoAudioDTO, grpcClient pb.AudioServiceClient, publisher *clienteCola.ClientePublisher) {
	audios, err := clienteRest.ObtenerAudiosPorTipo(tipo.IdTipo)
	if err != nil {
		log.Println("Error obteniendo audios:", err)
		return
	}

	for {
		fmt.Printf("\n=== Spotify ===\nTipo: %s\n", tipo.Nombre)
		for i, a := range audios {
			fmt.Printf("%d. %s\n", i+1, a.Titulo)
		}
		opcionAtras := len(audios) + 1
		fmt.Printf("%d. Atrás\n", opcionAtras)
		fmt.Print("Seleccione una opción: ")

		opcion, err := strconv.Atoi(leerOpcion())
		if err != nil || opcion < 1 || opcion > opcionAtras {
			fmt.Println("Opción inválida.")
			continue
		}
		if opcion == opcionAtras {
			return
		}
		mostrarDetalleAudio(tipo, audios[opcion-1].Titulo, grpcClient, publisher)
	}
}

// mostrarDetalleAudio - vista 4
func mostrarDetalleAudio(tipo clienteRest.TipoAudioDTO, titulo string, grpcClient pb.AudioServiceClient, publisher *clienteCola.ClientePublisher) {
	imprimir, archivo, err := obtenerDetalle(tipo, titulo)
	if err != nil {
		log.Println("Error obteniendo detalle:", err)
		return
	}

	for {
		fmt.Println("\n=== Spotify ===")
		imprimir()
		fmt.Println("1. Reproducir")
		fmt.Println("2. Atrás")
		fmt.Print("Seleccione una opción: ")

		switch leerOpcion() {
		case "1":
			reproducirAudio(tipo, titulo, archivo, grpcClient, publisher)
		case "2":
			return
		default:
			fmt.Println("Opción inválida.")
		}
	}
}

// obtenerDetalle consulta el detalle según el tipo y retorna una función que
// imprime los campos en un ORDEN FIJO (evita el orden aleatorio de los mapas
// de Go), junto con el nombre del archivo mp3 para el streaming.
func obtenerDetalle(tipo clienteRest.TipoAudioDTO, titulo string) (func(), string, error) {
	switch tipo.IdTipo {
	case 1:
		d, err := clienteRest.ObtenerDetalleMusica(titulo)
		if err != nil {
			return nil, "", err
		}
		return func() {
			fmt.Printf("Música: %s\n", d.Titulo)
			fmt.Printf("- Artista Principal: %s\n", d.ArtistaPrincipal)
			fmt.Printf("- Álbum: %s\n", d.Album)
			fmt.Printf("- Género Musical: %s\n", d.GeneroMusical)
			fmt.Printf("- Sello Discográfico: %s\n", d.SelloDiscografico)
			fmt.Printf("- Año de Lanzamiento: %d\n", d.AnioLanzamiento)
		}, d.Archivo, nil

	case 2:
		d, err := clienteRest.ObtenerDetallePodcast(titulo)
		if err != nil {
			return nil, "", err
		}
		return func() {
			fmt.Printf("Podcast: %s - %s\n", d.NombrePodcast, d.TituloEpisodio)
			fmt.Printf("- Anfitrión: %s\n", d.Anfitrion)
			fmt.Printf("- Temporada/Episodio: %s\n", d.TemporadaEpisodio)
			fmt.Printf("- Notas del Show: %s\n", d.NotasShow)
			fmt.Printf("- Clasificación: %s\n", d.Clasificacion)
		}, d.Archivo, nil

	case 3:
		d, err := clienteRest.ObtenerDetalleAudiolibro(titulo)
		if err != nil {
			return nil, "", err
		}
		return func() {
			fmt.Printf("Audiolibro: %s\n", d.TituloLibro)
			fmt.Printf("- Autor: %s\n", d.Autor)
			fmt.Printf("- Narrador: %s\n", d.Narrador)
			fmt.Printf("- Editorial: %s\n", d.Editorial)
			fmt.Printf("- ISBN: %s\n", d.Isbn)
			fmt.Printf("- Capítulo: %s\n", d.Capitulo)
		}, d.Archivo, nil

	case 4:
		d, err := clienteRest.ObtenerDetalleRuidoBlanco(titulo)
		if err != nil {
			return nil, "", err
		}
		return func() {
			fmt.Printf("Ruido Blanco: %s\n", d.TipoSonido)
			fmt.Printf("- Fuente del Audio: %s\n", d.FuenteAudio)
			fmt.Printf("- Uso Sugerido: %s\n", d.UsoSugerido)
			fmt.Printf("- Proveedor de Contenido: %s\n", d.ProveedorContenido)
			fmt.Printf("- Duración del Bucle: %s\n", d.DuracionBucle)
			fmt.Printf("- Frecuencia Dominante: %s\n", d.FrecuenciaDominante)
		}, d.Archivo, nil
	}
	return nil, "", fmt.Errorf("tipo de audio desconocido: %d", tipo.IdTipo)
}

// reproducirAudio - vista 5
func reproducirAudio(tipo clienteRest.TipoAudioDTO, titulo, archivo string, grpcClient pb.AudioServiceClient, publisher *clienteCola.ClientePublisher) {
	// Envío asíncrono del evento de estadísticas (no bloquea la reproducción).
	publisher.PublicarEventoAsync(titulo, tipo.Nombre)

	detener, err := utilidades.IniciarReproduccion(grpcClient, archivo)
	if err != nil {
		log.Println("Error iniciando reproducción:", err)
		return
	}

	for {
		fmt.Printf("\n=== Spotify ===\nReproduciendo: %s\nReproduciendo audio...\n1. Salir\n", titulo)
		fmt.Print("Seleccione una opción: ")
		if leerOpcion() == "1" {
			detener()
			return
		}
	}
}
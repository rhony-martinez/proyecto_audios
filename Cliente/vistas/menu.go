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

// MostrarMenuPrincipal - vista1
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

// mostrarTiposAudio - vista2
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

// mostrarListaAudios - vista3
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

// mostrarDetalleAudio - vista4
func mostrarDetalleAudio(tipo clienteRest.TipoAudioDTO, titulo string, grpcClient pb.AudioServiceClient, publisher *clienteCola.ClientePublisher) {
	detalle, err := clienteRest.ObtenerDetalle(tipo.IdTipo, titulo)
	if err != nil {
		log.Println("Error obteniendo detalle:", err)
		return
	}

	for {
		fmt.Printf("\n=== Spotify ===\n%s: %s\n", tipo.Nombre, titulo)
		for clave, valor := range detalle {
			if clave == "archivo" {
				continue // detalle interno, no se muestra al usuario
			}
			fmt.Printf("- %s: %v\n", clave, valor)
		}
		fmt.Println("1. Reproducir")
		fmt.Println("2. Atrás")
		fmt.Print("Seleccione una opción: ")

		switch leerOpcion() {
		case "1":
			archivo, _ := detalle["archivo"].(string)
			reproducirAudio(tipo, titulo, archivo, grpcClient, publisher)
		case "2":
			return
		default:
			fmt.Println("Opción inválida.")
		}
	}
}

// reproducirAudio - vista5
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
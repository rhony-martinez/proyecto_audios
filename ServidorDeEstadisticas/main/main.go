package main

import (
	"log"

	"servidor.local/estadisticas-servidor/capaAccesoDatos"
	"servidor.local/estadisticas-servidor/capaControladores"
	"servidor.local/estadisticas-servidor/capaFachadaServices/fachada"
	"servidor.local/estadisticas-servidor/componenteConexionCola"
)

func main() {
	// TODO: reemplazar IP real que obtenida con ipconfig
	urlConexion := "amqp://admin:1234@<IP_WINDOWS>:5672/"

	conexion, err := componenteConexionCola.NewRabbitConsumer(urlConexion)
	if err != nil {
		log.Fatal(err)
	}
	defer conexion.Cerrar()

	repo := capaAccesoDatos.NewRepositorioEstadisticas()
	fach := fachada.NewFachadaEstadisticas(repo)
	consumidor := capaControladores.NewConsumidorEstadisticas(conexion, fach)

	if err := consumidor.EscucharCola(); err != nil {
		log.Fatal(err)
	}
}
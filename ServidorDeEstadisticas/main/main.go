package main

import (
	"log"

	"servidor.local/estadisticas-servidor/capaAccesoDatos"
	"servidor.local/estadisticas-servidor/capaControladores"
	"servidor.local/estadisticas-servidor/capaFachadaServices/fachada"
	"servidor.local/estadisticas-servidor/componenteConexionCola"
)

func main() {
	// main conecta con RabbitMQ, construye las capas y comienza el consumo de eventos.
	// TODO: reemplazar IP real que obtenida con ipconfig
	urlConexion := "amqp://admin:1234@192.168.0.108:5672/"

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

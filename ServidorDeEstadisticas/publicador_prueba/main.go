package main

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://admin:1234@192.168.0.108:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("estadisticas_reproduccion", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	body := `{"titulo":"La funa","tipoAudio":"Música","fechaHora":"2026-09-26T17:30:00"}`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx, "", q.Name, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        []byte(body),
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Eco [publicador de prueba]: mensaje publicado en", q.Name)
}

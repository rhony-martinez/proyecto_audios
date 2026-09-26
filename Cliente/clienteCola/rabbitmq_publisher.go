package clienteCola

import (
	"context"
	"encoding/json"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const NombreCola = "estadisticas_reproduccion"

type ClientePublisher struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewClientePublisher(urlConexion string) (*ClientePublisher, error) {
	conn, err := amqp.Dial(urlConexion)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	if _, err := ch.QueueDeclare(NombreCola, true, false, false, false, nil); err != nil {
		return nil, err
	}
	return &ClientePublisher{conn: conn, ch: ch}, nil
}

type eventoReproduccion struct {
	Titulo    string `json:"titulo"`
	TipoAudio string `json:"tipoAudio"`
	FechaHora string `json:"fechaHora"`
}

// PublicarEventoAsync envía el evento en una goroutine para NO bloquear al
// usuario mientras empieza la reproducción (requisito: envío asíncrono).
func (p *ClientePublisher) PublicarEventoAsync(titulo, tipoAudio string) {
	go func() {
		log.Println("Eco [clienteCola]: publicando evento de reproducción...")
		evento := eventoReproduccion{Titulo: titulo, TipoAudio: tipoAudio, FechaHora: time.Now().Format(time.RFC3339)}
		body, _ := json.Marshal(evento)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := p.ch.PublishWithContext(ctx, "", NombreCola, false, false, amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
		if err != nil {
			log.Println("Error publicando evento de estadísticas:", err)
			return
		}
		log.Println("Eco [clienteCola]: evento publicado correctamente")
	}()
}

func (p *ClientePublisher) Cerrar() {
	p.ch.Close()
	p.conn.Close()
}
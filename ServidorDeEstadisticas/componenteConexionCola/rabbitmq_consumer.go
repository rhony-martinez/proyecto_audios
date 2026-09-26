package componenteConexionCola

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitConsumer struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

// NewRabbitConsumer se conecta al bróker ubicado en el host de Windows,
// usando las credenciales del usuario creado con rabbitmqctl.
func NewRabbitConsumer(urlConexion string) (*RabbitConsumer, error) {
	conn, err := amqp.Dial(urlConexion)
	if err != nil {
		return nil, fmt.Errorf("error conectando a RabbitMQ: %w", err)
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("error abriendo canal: %w", err)
	}
	return &RabbitConsumer{conn: conn, ch: ch}, nil
}

func (r *RabbitConsumer) DeclararCola(nombreCola string) (amqp.Queue, error) {
	return r.ch.QueueDeclare(nombreCola, true, false, false, false, nil)
}

func (r *RabbitConsumer) Consumir(nombreCola string) (<-chan amqp.Delivery, error) {
	return r.ch.Consume(nombreCola, "", true, false, false, false, nil)
}

func (r *RabbitConsumer) Cerrar() {
	r.ch.Close()
	r.conn.Close()
}
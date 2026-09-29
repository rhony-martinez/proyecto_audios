package componenteConexionCola

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitConsumer struct {
	// RabbitConsumer conserva la conexión y el canal usados para consumir mensajes AMQP.
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
	// DeclararCola crea o recupera una cola durable con el nombre indicado.
	return r.ch.QueueDeclare(nombreCola, true, false, false, false, nil)
}

func (r *RabbitConsumer) Consumir(nombreCola string) (<-chan amqp.Delivery, error) {
	// Consumir suscribe el canal a la cola y devuelve el flujo de entregas.
	return r.ch.Consume(nombreCola, "", true, false, false, false, nil)
}

func (r *RabbitConsumer) Cerrar() {
	// Cerrar libera los recursos AMQP del consumidor.
	r.ch.Close()
	r.conn.Close()
}
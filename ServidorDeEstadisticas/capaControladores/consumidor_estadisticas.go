package capaControladores

import (
	"log"

	"servidor.local/estadisticas-servidor/capaFachadaServices/fachada"
	"servidor.local/estadisticas-servidor/componenteConexionCola"
)

const NombreCola = "estadisticas_reproduccion"

type ConsumidorEstadisticas struct {
	conexion *componenteConexionCola.RabbitConsumer
	fachada  *fachada.FachadaEstadisticas
}

func NewConsumidorEstadisticas(conexion *componenteConexionCola.RabbitConsumer, f *fachada.FachadaEstadisticas) *ConsumidorEstadisticas {
	return &ConsumidorEstadisticas{conexion: conexion, fachada: f}
}

func (c *ConsumidorEstadisticas) EscucharCola() error {
	if _, err := c.conexion.DeclararCola(NombreCola); err != nil {
		return err
	}
	mensajes, err := c.conexion.Consumir(NombreCola)
	if err != nil {
		return err
	}

	log.Println("Eco [cola]: esperando mensajes en la cola", NombreCola)
	for msg := range mensajes {
		log.Println("Eco [cola]: mensaje consumido desde", NombreCola)
		c.fachada.ProcesarMensaje(msg.Body)
	}
	return nil
}
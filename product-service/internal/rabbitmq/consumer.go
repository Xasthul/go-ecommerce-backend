package rabbitmq

import (
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
)

type OrderCreatedEvent struct {
	OrderID   string `json:"order_id"`
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

func ConsumeOrders(conn *amqp.Connection, handle func(*OrderCreatedEvent)) error {
	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	err = ch.ExchangeDeclare("orders", "fanout", true, false, false, false, nil)
	if err != nil {
		return err
	}

	q, err := ch.QueueDeclare("", false, true, true, false, nil)
	if err != nil {
		return err
	}

	err = ch.QueueBind(q.Name, "", "orders", false, nil)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			var event OrderCreatedEvent
			if err := json.Unmarshal(msg.Body, &event); err == nil {
				handle(&event)
			}
		}
	}()

	return nil
}

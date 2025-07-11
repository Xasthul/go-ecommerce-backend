package rabbitmq

import (
	"encoding/json"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	ch *amqp.Channel
}

func NewPublisher(conn *amqp.Connection) (*Publisher, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	err = ch.ExchangeDeclare(
		"orders",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)

	return &Publisher{ch: ch}, err
}

type OrderCreatedEvent struct {
	UserID     uuid.UUID `json:"user_id"`
	OrderID    uuid.UUID `json:"order_id"`
	ProductID  uuid.UUID `json:"product_id"`
	Quantity   int       `json:"quantity"`
	TotalCents int       `json:"total_cents"`
}

func (p *Publisher) PublishOrderCreated(payload *OrderCreatedEvent) error {
	body, _ := json.Marshal(&payload)
	return p.ch.Publish(
		"orders",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
			Type:        "order.created",
		},
	)
}

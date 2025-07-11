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
		"payments",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)

	return &Publisher{ch: ch}, err
}

type PaymentSucceededEvent struct {
	OrderID   uuid.UUID `json:"order_id"`
	PaymentID uuid.UUID `json:"payment_id"`
	Amount    int       `json:"amount"`
}

func (p *Publisher) PublishPaymentSucceeded(payload *PaymentSucceededEvent) error {
	body, _ := json.Marshal(&payload)
	return p.ch.Publish(
		"payments",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
			Type:        "payment.succeeded",
		},
	)
}

type PaymentFailedEvent struct {
	OrderID uuid.UUID `json:"order_id"`
	Reason  string    `json:"reason"`
}

func (p *Publisher) PublishPaymentFailed(payload *PaymentFailedEvent) error {
	body, _ := json.Marshal(&payload)
	return p.ch.Publish(
		"payments",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
			Type:        "payment.failed",
		},
	)
}

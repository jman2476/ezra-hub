package msgbroker

import (
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func SubscribeJSON[T any](
	conn *amqp.Connection,
	exchange,
	queueName,
	key string,
	queueType SimpleQueueType,
	handler func(T) AckType,
) error {
	channel, queue, err := DeclareAndBind(
		conn,
		exchange,
		queueName,
		key,
		queueType,
	)
	if err != nil {
		return fmt.Errorf("Subscibe error: DeclareAndBind fail: %w", err)
	}

	delivery, err := channel.Consume(
		queue.Name,
		"",
		false, false, false, false, nil,
	)
	if err != nil {
		return fmt.Errorf("could not get delivery channel: %w", err)
	}

	unmarshaller := func(data []byte) (T, error) {
		var target T
		err := json.Unmarshal(data, &target)
		return target, err
	}

	go func(delchan <-chan amqp.Delivery) {
		for msg := range delchan {
			data, err := unmarshaller(msg.Body)
			if err != nil {
				fmt.Println(
					fmt.Errorf("Failed to unmarshal: %w", err),
				)
				continue
			}

			ackType := handler(data)
			switch ackType {
			case Ack:
				msg.Ack(false)
			case NackRequeue:
				msg.Nack(false, true)
			case NackDiscard:
				msg.Nack(false, false)
			}
		}
	}(delivery)

	return nil
}

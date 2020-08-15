package mq

import (
	"encoding/json"
	"github.com/streadway/amqp"
)

//Publish is a
func Publish(queue amqp.Queue, body interface{}) error{
	payload, err := json.Marshal(body)
	if err != nil {
		return err
	}

	err = Channel.Publish(
		"",     // exchange
		queue.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        payload,
		})

	if err != nil {
		return err
	}

	return nil
}
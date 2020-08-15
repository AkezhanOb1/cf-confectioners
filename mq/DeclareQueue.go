package mq

import "github.com/streadway/amqp"

func DeclareQueue (queueName string) (amqp.Queue, error){
	queue, err := Channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return amqp.Queue{}, err
	}


	return queue, nil
}
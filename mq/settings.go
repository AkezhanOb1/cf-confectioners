package mq

import (
	"github.com/AkezhanOb1/cf-confectioners/config"
	"github.com/streadway/amqp"
	"log"
)

//Mq is a
var Channel *amqp.Channel

func init() {
	conn, err := amqp.Dial(config.RabbitConnection)
	if err != nil {
		log.Fatal(err)
	}

	Channel, err = conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("connected to mq")
}


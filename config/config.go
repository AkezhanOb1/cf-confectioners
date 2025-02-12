package config

import "fmt"

//postgresAddress is the address of the postgres
const postgresAddress = "127.0.0.1"
//const postgresAddress = "46.101.138.224"

//postgresPort is the port of the postgres
const postgresPort = "5432"

//postgresDataBase is the name of the database
const postgresDataBase = "postgres"

//postgresUsername is the name of the user inside DBA
const postgresUsername = "postgres"

//postgresPassword is the password of the user
const postgresPassword = "postgres"

//PostgresConnection is the connection string to the database
var PostgresConnection = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	postgresAddress, postgresPort, postgresUsername, postgresPassword, postgresDataBase)

//ServerAddress is
var ServerAddress = "localhost:50055"

//rabbitUserName is the name of the user in Rabbit MQ
const rabbitUserName = "guest"
//rabbitPassword is the password of the user
const rabbitPassword = "guest"
//RabbitConnection is the connection string to the Rabbit MQ
var RabbitConnection = fmt.Sprintf("amqp://%s:%s@127.0.0.1:5672/", rabbitUserName, rabbitPassword)
//ConfectionerGreetingQueue is a
var ConfectionerGreetingQueue = "confectioner-greeting"
//TokenServer is 46.101.138.224
//var TokenServer = "46.101.138.224:50052"
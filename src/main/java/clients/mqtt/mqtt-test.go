package main

import (
	"fmt"
	"os"
	"os/signal"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func onMessageReceived(client MQTT.Client, message MQTT.Message) {
	fmt.Printf("Received message on topic: %s\n", message.Topic())
	fmt.Printf("Message: %s\n", message.Payload())
}

func main() {
	// Создание нового клиента MQTT
	opts := MQTT.NewClientOptions().AddBroker("tcp://broker.hivemq.com:1883")
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	defer client.Disconnect(250)

	// Подписка на топики
	topics := []string{
		"ITMO/Student17/Value1",
		"ITMO/Student17/Value2",
		"ITMO/+/Value3", // "+" обозначает подписку на все топики ITMO с любым номером студента и любым значением Value3
	}
	for _, topic := range topics {
		if token := client.Subscribe(topic, 0, onMessageReceived); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}
	fmt.Println("Subscribed to topics")

	// Ожидание сигнала завершения (например, CTRL+C)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

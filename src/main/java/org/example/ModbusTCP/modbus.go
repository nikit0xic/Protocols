package main

import (
	"fmt"
	"github.com/goburrow/modbus"
	"log"
	"time"
)

func main() {
	tcpAddr := "109.167.241.225:601"
	slaveID := byte(1)

	studentNumber := 17
	startRegister := uint16(studentNumber * 100)

	handler := modbus.NewTCPClientHandler(tcpAddr)
	handler.SlaveId = slaveID
	defer handler.Close()
	err := handler.Connect()

	if err != nil {
		log.Fatalf("Connection error: %v", err)
	}

	client := modbus.NewClient(handler)
	// Чтение данных: 2 * uint16 (2 регистра) + строка ASCII (5 регистров, 10 символов)
	results, err := client.ReadHoldingRegisters(startRegister, 7)
	if err != nil {
		log.Fatalf("Input error: %v", err)
	}

	firstNumber := uint16(results[0])<<8 | uint16(results[1])
	secondNumber := uint16(results[2])<<8 | uint16(results[3])
	asciiString := string(results[4:14])

	fmt.Println("Current time:", time.Now())
	fmt.Printf("First number: %d\n", firstNumber)
	fmt.Printf("Second number: %d\n", secondNumber)
	fmt.Printf("ASCII String: %s\n", asciiString)
}

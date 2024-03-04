package UDP_FTP

import (
	"fmt"
	"net"
	"time"
)

func main() {
	serverAddr := "109.167.241.225"
	serverPort := 61557

	// Создание UDP клиента
	clientConn, err := net.DialUDP("udp", nil, &net.UDPAddr{IP: net.ParseIP(serverAddr), Port: serverPort})
	if err != nil {
		fmt.Printf("Ошибка при создании UDP клиента: %v\n", err)
		return
	}
	defer clientConn.Close()

	// Отправка данных
	for i := 1; i <= 100; i++ {
		data := fmt.Sprintf("<Фофанов>: %d", generateData(i))
		sendData := []byte(data)
		_, err := clientConn.Write(sendData)
		if err != nil {
			fmt.Printf("Ошибка при отправке данных: %v\n", err)
			return
		}
		time.Sleep(1 * time.Second) // Пауза между отправками
	}
}

// Генерация данных
func generateData(iteration int) int {
	// Пример функции: квадрат числа
	return iteration * iteration
}

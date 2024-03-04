package UDP_FTP

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	"io"
	"os"
	"strings"
)

func main() {
	server := "109.167.241.225"
	port := 21
	user := "Student"
	pass := "FksG5$%^rgtdSDFH"
	remoteFile := "UDP log.txt"

	// Подключение к FTP серверу
	ftpClient, err := ftp.Dial(fmt.Sprintf("%s:%d", server, port))
	if err != nil {
		fmt.Printf("Ошибка при подключении к FTP серверу: %v\n", err)
		return
	}
	defer ftpClient.Quit()

	err = ftpClient.Login(user, pass)
	if err != nil {
		fmt.Printf("Ошибка при аутентификации: %v\n", err)
		return
	}

	// Получение файла с FTP сервера
	file, err := os.Create(remoteFile)
	if err != nil {
		fmt.Printf("Ошибка при создании файла: %v\n", err)
		return
	}
	defer file.Close()

	err = ftpClient.Retrieve(remoteFile, file)
	if err != nil {
		fmt.Printf("Ошибка при получении файла: %v\n", err)
		return
	}

	fmt.Println("Файл успешно скачан.")
}

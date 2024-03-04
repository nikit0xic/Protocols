package main

import (
	"fmt"
	"net"
	"regexp"
	"time"
)

func main() {
	ip := "109.167.241.225"
	port := 6340
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	result := make([]byte, 0)
	buffer := make([]byte, 65536)
	for true {
		n, err := conn.Read(buffer)
		if err != nil {
			panic(err)
		}
		if n != 4 {
			result = append(result, buffer...)
			break
		}
		time.Sleep(1 * time.Second)
	}
	matched, _ := regexp.Compile("Student17\\s\\S+\\s\\S+")
	results := matched.FindAllString(string(result), 10)
	for _, res := range results {
		fmt.Println(res)
	}
}

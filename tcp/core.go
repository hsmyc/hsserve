package tcp

import (
	"bufio"
	"fmt"
	"net"
	"net/textproto"
	"strings"

	"github.com/hsmyc/hsserve/utils"
)

func handleRequest(conn net.Conn) {
	defer conn.Close()
	tp := textproto.NewReader(bufio.NewReader(conn))
	rl, err := tp.ReadLine()
	if err != nil {
		fmt.Println("Error reading", err)
		return
	}
	requestData := rl
	lines := strings.Split(requestData, "\n")
	if len(lines) == 0 {
		fmt.Println("No data received")
		conn.Write([]byte("HTTP/1.1 400 Bad Request\r\n\r\nInvalid request"))
		return
	}
	head := strings.TrimSpace(lines[0])
	if head == "GET / HTTP/1.1" {
		text := utils.HTML{Path: "static/index.html"}
		data, err := text.ReturnHTML()

		if err != nil {
			fmt.Println("Error reading file", err)
			return
		}

		conn.Write([]byte(data))
		return
	}
	conn.Write([]byte("HTTP/1.1 400 Bad Request\r\n\r\nInvalid request"))
}

func StartServer() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server", err)
		return
	}
	fmt.Println("Server started on port 8080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection", err)
			return
		}
		go handleRequest(conn)
	}

}

package main

import (
	"bufio"
	"fmt"
	"io"
	"mq-lite/internal/broker"
	"net"
	"os"
	"strings"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	b := broker.New()
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("failed to create listener, err:", err)
		os.Exit(1)
	}
	fmt.Println("listening on", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("failed to accept connection, err:", err)
			continue
		}
		go handleConnection(b, conn)
	}
}

func handleConnection(b *broker.Broker, conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		bytes, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Println("failed to read bytes from connection, err:", err)
			}
			return
		}
		fmt.Printf("received request: %s", bytes)

		line := strings.TrimSpace(bytes)
		if len(line) == 0 {
			continue
		}
		parts := strings.SplitN(line, " ", 3)
		if len(parts) < 2 {
			conn.Write([]byte("too few arguments"))
			continue
		}
		cmd := strings.ToUpper(parts[0])
		topic := parts[1]

		switch cmd {
		case "SUB":
			b.Subscribe(topic, conn)
		case "PUB":
			if len(parts) < 3 {
				conn.Write([]byte("too few arguments"))
			}
			b.Publish(topic, parts[2], conn)
			conn.Write([]byte("published"))
		default:
			conn.Write([]byte("unknown command: " + cmd))
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
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
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		bytes, err := reader.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Println("failed to read bytes from connection, err:", err)
			}
			return
		}
		fmt.Printf("received request: %s", bytes)

		line := fmt.Sprintf("%s\n", bytes)
		fmt.Println(line)
		conn.Write([]byte(line))
	}
}

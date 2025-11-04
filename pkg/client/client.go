package client

import (
	"bufio"
	"fmt"
	"net"
)

type Client struct {
	conn   net.Conn
	reader *bufio.Reader
}

func Connect(addr string) (*Client, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &Client{conn: conn, reader: bufio.NewReader(conn)}, nil
}

func (c *Client) Publish(topic, msg string) error {
	cmd := fmt.Sprintf("PUB %s %s\n", topic, msg)
	_, err := c.conn.Write([]byte(cmd))
	return err
}

func (c *Client) Subscribe(topic string) (<-chan string, error) {
	cmd := fmt.Sprintf("SUB %s\n", topic)
	_, err := c.conn.Write([]byte(cmd))
	if err != nil {
		return nil, err
	}
	msgChan := make(chan string, 10)
	go func() {
		defer close(msgChan)
		for {
			line, err := c.reader.ReadString('\n')
			if err != nil {
				return
			}
			msgChan <- line
		}
	}()
	return msgChan, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

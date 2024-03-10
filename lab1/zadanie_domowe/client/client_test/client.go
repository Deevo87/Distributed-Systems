package client_test

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
)

type Client struct {
	conn         net.Conn
	hostName     string
	serverAddr   string
	portAddr     string
	quitch       chan os.Signal
	serverClosed chan struct{}
}

func NewClient(hostname, portAddr, serverAddr string) *Client {
	return &Client{
		hostName:     hostname,
		portAddr:     portAddr,
		serverAddr:   serverAddr,
		quitch:       make(chan os.Signal, 1),
		serverClosed: make(chan struct{}),
	}
}

func (c *Client) Start() error {
	fmt.Println("Connected to the server...")
	conn, err := net.Dial("tcp", c.serverAddr)
	if err != nil {
		fmt.Println("Error occurred during client start: ", err)
	}
	c.conn = conn
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("Error during closing connection: ", conn.RemoteAddr(), err)
		}
	}(c.conn)
	c.acceptLoop()
	select {
	case <-c.quitch:
		fmt.Println("Connection with server closed...")
	case <-c.serverClosed:
		fmt.Println("Server closed connection...")
	}
	return nil
}

func (c *Client) acceptLoop() {
	go c.receiveMsg()
	go c.SendMsg()
}

func (c *Client) SendMsg() error {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		select {
		case <-c.serverClosed:
			return nil
		default:
			line := scanner.Text()
			if line == "" {
				continue
			} else if line == "U" {
				fmt.Println("Sending UDP")
			}
			msg := c.hostName + ": " + line
			_, err := c.conn.Write([]byte(msg))
			if err != nil {
				fmt.Println("Error during sending message from client: ", err)
				return err
			}
		}
	}
	signal.Notify(c.quitch, os.Interrupt) //captures ctrl + c and close connection
	return nil
}

func (c *Client) receiveMsg() error {
	for {
		select {
		case <-c.quitch:
			return nil
		default:
			buff := make([]byte, 2048)
			n, err := c.conn.Read(buff)
			if err != nil {
				if err == io.EOF {
					close(c.serverClosed)
					return err
				}
				fmt.Println("Error during read: ", err)
				return err
			}

			// buffer the thing we read not actually the whole buffer
			msg := buff[:n]
			fmt.Printf("> %s\n", string(msg))
		}
	}
}

func (c *Client) sendUDP() error {
	s, _ := net.ResolveUDPAddr("udp", c.serverAddr)
	conn, err := net.DialUDP("udp", nil, s)
	if err != nil {
		fmt.Println("Error during sending UDP from ", c.hostName)
		return err
	}

	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)
	return nil
}

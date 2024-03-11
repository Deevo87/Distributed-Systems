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
	connTCP      net.Conn
	connUDP      *net.UDPConn
	hostName     string
	serverAddr   string
	quitch       chan os.Signal
	serverClosed chan struct{}
}

func NewClient(hostname, serverAddr string) *Client {
	return &Client{
		hostName:     hostname,
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
	c.connTCP = conn
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("Error during closing connection: ", conn.RemoteAddr(), err)
		}
	}(c.connTCP)

	udpEndpoint, _ := net.ResolveUDPAddr("udp", c.connTCP.LocalAddr().String())

	udpLn, err := net.ListenUDP("udp", udpEndpoint)
	if err != nil {
		return err
	}

	defer func(udpLn *net.UDPConn) {
		err := udpLn.Close()
		if err != nil {
			fmt.Println("Error during closing UDP: ", err)
		}
	}(udpLn)
	c.connUDP = udpLn
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
	go c.receiveTCPMsg()
	go c.sendMsg()
	go c.receiveUDPMsg()
}

func (c *Client) sendMsg() error {
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
				c.sendUDP()
				continue
			}
			msg := c.hostName + ": " + line
			_, err := c.connTCP.Write([]byte(msg))
			if err != nil {
				fmt.Println("Error during sending message from client: ", err)
				return err
			}
		}
	}
	signal.Notify(c.quitch, os.Interrupt) //captures ctrl + c and close connection
	return nil
}

func (c *Client) receiveTCPMsg() error {
	for {
		select {
		case <-c.quitch:
			return nil
		default:

			buff := make([]byte, 2048)
			n, err := c.connTCP.Read(buff)
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

func (c *Client) receiveUDPMsg() error {
	for {
		select {
		case <-c.serverClosed:
			return nil
		default:
			buff := make([]byte, 2048)
			n, _, err := c.connUDP.ReadFromUDP(buff)
			if err != nil {
				if err == io.EOF {
					return err
				}
				continue
			}
			msg := buff[:n]
			fmt.Println(string(msg))
		}
	}
}

func (c *Client) createUDPConn(addr string) (*net.UDPConn, error) {
	udpEndpoint, _ := net.ResolveUDPAddr("udp", addr)
	conn, err := net.DialUDP("udp", nil, udpEndpoint)
	if err != nil {
		fmt.Println("Error while creating UDP connection ", addr)
		return nil, err
	}
	return conn, nil
}

func (c *Client) sendUDP() error {
	conn, err := c.createUDPConn(c.serverAddr)

	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	asciiArt := "    /\\_____/\\\n   /  o   o  \\\n  ( ==  ^  == )\n   )         (\n  (           )\n ( (  )   (  ) )\n(__(__)___(__)__)"
	_, err = conn.Write([]byte(asciiArt))
	if err != nil {
		fmt.Println("Error during sending message from client using UDP: ", err)
		return err
	}
	return nil
}

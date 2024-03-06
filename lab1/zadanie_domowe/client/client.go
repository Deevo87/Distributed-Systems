package main

import (
	"fmt"
	"net"
)

type Client struct {
	conn       net.Conn
	hostName   string
	serverAddr string
	portAddr   string
}

func NewClient(hostname, portAddr, serverAddr string) *Client {
	return &Client{
		hostName:   hostname,
		portAddr:   portAddr,
		serverAddr: serverAddr,
	}
}

func (c *Client) Start(msg string) error {
	conn, err := net.Dial("tcp", c.serverAddr)
	if err != nil {
		fmt.Println("Error occurred during client start: ", err)
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("Error occurred during closing connection on the clint side: ", err)
			// nie wiem czy tu ma być return czy nie, jak nie działa to potencjalnie tutaj to się może pierdolić
		}
	}(conn)
	c.conn = conn
	c.SendMsg(msg)
	return nil
}

func (c *Client) acceptLoop() {
	//for {
	//	conn, err := c.conn.
	//}
}

func (c *Client) SendMsg(msg string) error {
	data := []byte(msg)
	_, err := c.conn.Write(data)
	if err != nil {
		fmt.Println("Error during sending message from client: ", err)
		return err
	}
	return nil
}

func (c *Client) receiveMsg() {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("Error during closing connection: ", conn.RemoteAddr(), err)
		}
	}(c.conn)
	buff := make([]byte, 2048)
	for {
		n, err := c.conn.Read(buff)
		if err != nil {
			fmt.Println("Error during read: ", err)
			//TODO sprawdzaj EOFy
			continue
		}
		// buffer the thing we read not actually the whole buffer
		msg := buff[:n]
		fmt.Println(string(msg))
	}
}

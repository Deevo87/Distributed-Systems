package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

type Server struct {
	listenAddr string
	ln         net.Listener
	quitch     chan struct{}
	connMap    sync.Map
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitch:     make(chan struct{}),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		//fmt.Println("Error occurred: ", err)
		return err
	}

	defer func(ln net.Listener) {
		err := ln.Close()
		if err != nil {
			fmt.Println("Error during closing: ", err)
			return
		}
	}(ln)
	s.ln = ln

	go s.acceptLoop()

	<-s.quitch

	return nil
}

func (s *Server) acceptLoop() {
	// accepting connection
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("Error during accept: ", err)
			// there is no return to accept more connections
			continue
		}
		// for each connection there is new go routine (thread)
		s.connMap.Store(conn.RemoteAddr(), conn)
		fmt.Println("New connection to the server! Connection address: ", conn.RemoteAddr())
		go s.readLoop(conn)
	}
}

func (s *Server) readLoop(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("Error during closing connection: ", conn.RemoteAddr(), err)
		}
	}(conn)
	buff := make([]byte, 2048)
	for {
		n, err := conn.Read(buff)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("Connection with client: %s closed...\n", conn.RemoteAddr())
				s.connMap.Delete(conn.RemoteAddr())
				break
			}
			continue
		}
		//buffer the thing we read not actually the whole buffer
		msg := buff[:n]
		fmt.Printf("From %s\n", msg)
		s.sendAll(msg, conn.RemoteAddr())
	}
}

func (s *Server) sendAll(msg []byte, senderAddress net.Addr) {
	s.connMap.Range(func(key, value interface{}) bool {
		if conn, ok := value.(net.Conn); ok {
			if conn.RemoteAddr() != senderAddress {
				fmt.Println("Sending to: ", conn.RemoteAddr())
				conn.Write(msg)
			}
		}
		return true
	})
}

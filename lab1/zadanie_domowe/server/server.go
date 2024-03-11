package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

type Server struct {
	listenAddr  string
	tcpListener net.Listener
	udpListener *net.UDPConn
	quitch      chan struct{}
	connMap     sync.Map
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitch:     make(chan struct{}),
	}
}

func (s *Server) Start() error {
	tcpLn, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		//fmt.Println("Error occurred: ", err)
		return err
	}
	udpEndpoint, _ := net.ResolveUDPAddr("udp", s.listenAddr)

	udpLn, err := net.ListenUDP("udp", udpEndpoint)
	if err != nil {
		return err
	}

	defer func(ln net.Listener) {
		err := ln.Close()
		if err != nil {
			fmt.Println("Error during closing TCP: ", err)
			return
		}
	}(tcpLn)
	s.tcpListener = tcpLn

	defer func(udpLn *net.UDPConn) {
		err := udpLn.Close()
		if err != nil {
			fmt.Println("Error during closing UDP: ", err)
		}
	}(udpLn)
	s.udpListener = udpLn

	go s.acceptTCPLoop()
	go s.readUDPLoop()

	<-s.quitch

	return nil
}

func (s *Server) acceptTCPLoop() {
	for {
		conn, err := s.tcpListener.Accept()
		if err != nil {
			// there is no return to accept more connections
			continue
		}
		// for each connection there is new go routine (thread)
		s.connMap.Store(conn.RemoteAddr(), conn)
		fmt.Println("New connection to the server! Connection address: ", conn.RemoteAddr())
		go s.readTCPLoop(conn)
	}
}

func (s *Server) readUDPLoop() error {
	for {
		buff := make([]byte, 2048)
		n, addr, err := s.udpListener.ReadFromUDP(buff)
		if err != nil {
			fmt.Println("Error during UDP read: ", err)
			continue
		}
		fmt.Println("New UDP message to the server! Sender address: ", addr)
		msg := buff[:n]
		fmt.Println(string(msg))
		s.sendAllUDP(msg, addr.String())
	}
}

func (s *Server) readTCPLoop(conn net.Conn) {
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
		s.sendAllTCP(msg, conn.RemoteAddr())
	}
}

func (s *Server) sendAllTCP(msg []byte, senderAddress net.Addr) {
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

func (s *Server) sendAllUDP(msg []byte, senderAddress string) {
	s.connMap.Range(func(key, value interface{}) bool {
		addr := fmt.Sprintf("%v", key)
		conn, _ := s.createUDPConn(addr)
		_, err := conn.Write(msg)
		if err != nil {
			fmt.Println("Error during sending UDP to all clients: ", err)
			return false
		}
		fmt.Println("Sending to: ", addr)
		err = conn.Close()
		if err != nil {
			return false
		}
		return true
	})
}

func (s *Server) createUDPConn(addr string) (*net.UDPConn, error) {
	udpEndpoint, _ := net.ResolveUDPAddr("udp", addr)
	conn, err := net.DialUDP("udp", nil, udpEndpoint)
	if err != nil {
		fmt.Println("Error during sending UDP to ", addr)
		return nil, err
	}
	return conn, nil
}

package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

const defaultHostPort string = ":9000"

func main() {
	log.Printf("TCP Server Application")

	tcpAddr, err := net.ResolveTCPAddr("tcp", defaultHostPort)
	if err != nil {
		log.Printf("Erroe occured!")
		log.Println(err)
		log.Fatal(err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Printf("Erroe occured!")
		log.Println(err)
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()

		conn.(*net.TCPConn).SetKeepAlive(true)
		timeoutDuration := 5 * time.Second
		conn.(*net.TCPConn).SetDeadline(time.Now().Add(timeoutDuration))

		if err != nil {
			log.Printf("Erroe occured!")
			log.Println(err)
			continue
		}

		datetime := time.Now().String()
		go func() {
			conn.Write([]byte(fmt.Sprintf("Time is: %q", datetime)))
			defer conn.Close()
		}()
	}
}

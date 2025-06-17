package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	broadcastAddr := "255.255.255.255:8080"

	addr, err := net.ResolveUDPAddr("udp", broadcastAddr)
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Error dialing UDP:", err)
		return
	}
	defer conn.Close()

	message := []byte("Hello, broadcast!")

	for {
		_, err = conn.Write(message)
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}

		time.Sleep(10 * time.Microsecond)
	}
}

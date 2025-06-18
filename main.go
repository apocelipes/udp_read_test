package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"golang.org/x/net/ipv4"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "255.255.255.255:8080")
	if err != nil {
		panic(err)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	useBatchSend := len(os.Args) > 1 && os.Args[1] == "-batch"

	if useBatchSend {
		sendBatch(conn)
	} else {
		sendSingle(conn)
	}
}

func sendSingle(conn net.Conn) {
	message := []byte("Hello, udp!")
	for {
		_, err := conn.Write(message)
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}

		time.Sleep(10 * time.Microsecond)
	}
}

func sendBatch(conn net.PacketConn) {
	message := []byte("Hello, udp!")
	batchMsg := make([]ipv4.Message, 16)
	for i := range batchMsg {
		batchMsg[i] = ipv4.Message{
			Buffers: [][]byte{message},
			N:       len(message),
		}
	}
	pc := ipv4.NewPacketConn(conn)
	for {
		_, err := pc.WriteBatch(batchMsg, flags)
		if err != nil {
			panic(err)
		}
		time.Sleep(20 * time.Microsecond)
	}
}

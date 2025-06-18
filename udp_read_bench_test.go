package main

import (
	"net"
	"testing"

	"golang.org/x/net/ipv4"
)

func BenchmarkReadFrom(b *testing.B) {
	addr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		b.Fatal(err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		b.Fatal(err)
	}
	buff := make([]byte, 1024)
	for b.Loop() {
		count := 0
		for count <= 1000 {
			n, _, err := conn.ReadFrom(buff)
			if err != nil {
				b.Fatal(err)
			}
			if n <= 0 {
				b.Fatalf("length <= 0: %d", n)
			}
			count++
		}
	}
	conn.Close()
}

func BenchmarkReadFromUDP(b *testing.B) {
	addr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		b.Fatal(err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		b.Fatal(err)
	}
	buff := make([]byte, 1024)
	for b.Loop() {
		count := 0
		for count <= 1000 {
			n, _, err := conn.ReadFromUDP(buff)
			if err != nil {
				b.Fatal(err)
			}
			if n <= 0 {
				b.Fatalf("length <= 0: %d", n)
			}
			count++
		}
	}
	conn.Close()
}

func BenchmarkReadFromUPDAddrPort(b *testing.B) {
	addr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		b.Fatal(err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		b.Fatal(err)
	}
	buff := make([]byte, 1024)
	for b.Loop() {
		count := 0
		for count <= 1000 {
			n, _, err := conn.ReadFromUDPAddrPort(buff)
			if err != nil {
				b.Fatal(err)
			}
			if n <= 0 {
				b.Fatalf("length <= 0: %d", n)
			}
			count++
		}
	}
	conn.Close()
}

func BenchmarkRecvmmsg(b *testing.B) {
	addr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		b.Fatal(err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		b.Fatal(err)
	}
	c := ipv4.NewPacketConn(conn)
	batchBuffer := make([]ipv4.Message, 16)
	for i := range batchBuffer {
		batchBuffer[i].Buffers = [][]byte{make([]byte, 1024)}
		batchBuffer[i].N = 1024
	}
	for b.Loop() {
		count := 0
		for count <= 1000 {
			n, err := c.ReadBatch(batchBuffer, 0)
			if err != nil {
				b.Fatal(err)
			}
			if n <= 0 {
				b.Fatalf("length <= 0: %d", n)
			}
			count += n
		}
	}
	conn.Close()
}

/*
goos: darwin
goarch: arm64
pkg: udpreadtest
cpu: Apple M4
BenchmarkReadFrom-10               	      58	  24760452 ns/op	   64064 B/op	    2002 allocs/op
BenchmarkReadFromUDP-10            	      48	  25219753 ns/op	   64064 B/op	    2002 allocs/op
BenchmarkReadFromUPDAddrPort-10    	      50	  25412824 ns/op	       0 B/op	       0 allocs/op
BenchmarkRecvmmsg-10               	      45	  24835418 ns/op	  296098 B/op	   10998 allocs/op
PASS
ok  	udpreadtest	8.061s
*/

/*
goos: linux
goarch: arm64
pkg: udpreadtest
BenchmarkReadFrom-4              	       3	 484477973 ns/op	   64064 B/op	    2002 allocs/op
BenchmarkReadFromUDP-4           	       3	 491391101 ns/op	   64064 B/op	    2002 allocs/op
BenchmarkReadFromUPDAddrPort-4   	       3	 474905637 ns/op	       0 B/op	       0 allocs/op
BenchmarkRecvmmsg-4              	       3	 476800798 ns/op	   64920 B/op	    2005 allocs/op
PASS
ok  	udpreadtest	5.801s
*/

/*
goos: linux
goarch: amd64
pkg: udpreadtest
cpu: Intel(R) Core(TM) i7-14650HX
BenchmarkReadFrom-24                          27          43534583 ns/op           64064 B/op       2002 allocs/op
BenchmarkReadFromUDP-24                       26          44521251 ns/op           64064 B/op       2002 allocs/op
BenchmarkReadFromUPDAddrPort-24               26          44305431 ns/op               0 B/op          0 allocs/op
BenchmarkRecvmmsg-24                          26          44505049 ns/op           64268 B/op       2002 allocs/op
PASS
ok      udpreadtest     4.647s
*/

// batch send

/*
goos: linux
goarch: amd64
pkg: udpreadtest
cpu: Intel(R) Core(TM) i7-14650HX
BenchmarkReadFrom-24                         220           5568036 ns/op           64064 B/op       2002 allocs/op
BenchmarkReadFromUDP-24                      216           5478449 ns/op           64065 B/op       2002 allocs/op
BenchmarkReadFromUPDAddrPort-24              211           5512309 ns/op               0 B/op          0 allocs/op
BenchmarkRecvmmsg-24                         212           5561909 ns/op           64568 B/op       2015 allocs/op
PASS
ok      udpreadtest     4.799s
*/

/*
goos: linux
goarch: arm64
pkg: udpreadtest
BenchmarkReadFrom-4              	       4	 265467281 ns/op	   64064 B/op	    2002 allocs/op
BenchmarkReadFromUDP-4           	       4	 261408366 ns/op	   64064 B/op	    2002 allocs/op
BenchmarkReadFromUPDAddrPort-4   	       4	 252192508 ns/op	       0 B/op	       0 allocs/op
BenchmarkRecvmmsg-4              	       5	 245079165 ns/op	   64612 B/op	    2004 allocs/op
PASS
ok  	udpreadtest	4.194s
*/

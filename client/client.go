package main

import (
	"fmt"
	"net"
)

func main() {
	udp()
}
func tcp() {
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		panic(err)
		return
	}
	defer conn.Close()
	write, err := conn.Write([]byte("hello world"))
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(write)
	buf := [512]byte{}
	read, err := conn.Read(buf[:])
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(string(buf[:read]))
}
func udp() {
	udp, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 9001,
	})
	if err != nil {
		panic(err)
		return
	}
	defer udp.Close()
	bytes := []byte("Hello World")
	_, err = udp.Write(bytes)
	if err != nil {
		panic(err)
		return
	}
	b := make([]byte, 1024)
	n, addr, err := udp.ReadFromUDP(b)
	fmt.Println(string(b[:n]))
	fmt.Println(addr)
}

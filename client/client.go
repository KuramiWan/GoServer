package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		return
	}
	defer conn.Close()
	write, err := conn.Write([]byte("hello world"))
	if err != nil {
		return
	}
	fmt.Println(write)
	buf := [512]byte{}
	read, err := conn.Read(buf[:])
	if err != nil {
		return
	}
	fmt.Println(string(buf[:read]))
}

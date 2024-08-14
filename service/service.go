package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if n == 0 {
			break
		}
		if err != nil {
			fmt.Println("error:", err)
			break
		}
		s := string(buf[:n])
		fmt.Println("收到的数据", s)
		conn.Write([]byte(s))
	}

}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("error:", err)
			continue
		}
		go process(conn)
	}

}

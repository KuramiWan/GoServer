package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func tcpProcess(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("error:", err)
			break
		}
		s := string(buf[:n])
		fmt.Println("收到的数据", s)
		_, err = conn.Write([]byte(s))
		if err != nil {
			return
		}
	}

}
func tcp() {
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
		go tcpProcess(conn)
	}
}
func udp() {
	listenUDP, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4(0, 0, 0, 0), Port: 9001})
	if err != nil {
		panic(err)
		return
	}
	defer listenUDP.Close()
	for {
		var data [1024]byte
		n, addr, err := listenUDP.ReadFromUDP(data[:])
		if err != nil {
			continue

		}
		fmt.Println(string(data[:n]))
		_, err = listenUDP.WriteToUDP(data[:n], addr)
		if err != nil {
			continue
		}
	}
}
func main() {
	go tcp()
	go udp()
	select {}
}

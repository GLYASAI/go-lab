package main

import (
	"fmt"
	"net"
)

func main() {
	ip := net.ParseIP("0.0.0.0")
	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dstAddr := &net.UDPAddr{IP: ip, Port: 38080}
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	conn.Write([]byte("hello!!!\n"))
	fmt.Printf("<%s>\n", conn.RemoteAddr())
}

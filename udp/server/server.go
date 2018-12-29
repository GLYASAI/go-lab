package main

import (
	"fmt"
	"net"
	"os"
	"path"
)

func main() {
	output := "/run/output"
	if !isExists(output) {
		err := os.MkdirAll(path.Dir(output), 0777)
		if err != nil {
			panic(err)
		}
		if _, err := os.Create(output); err != nil {
			panic(err)
		}
	}
	f, err := os.OpenFile(output, os.O_WRONLY|os.O_APPEND, 0666)
	defer f.Close()
	if err != nil {
		panic(fmt.Sprintf("error opening file: %s: %v", output, err))
	}
	listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("0.0.0.0"), Port: 38080})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Local: <%s> \n", listener.LocalAddr().String())
	data := make([]byte, 1024)
	for {
		n, remoteAddr, err := listener.ReadFromUDP(data)
		if err != nil {
			fmt.Printf("error during read: %s", err)
		}
		fmt.Printf("<%s> %s\n", remoteAddr, data[:n])
		_, err = f.Write(data[:n])
		if err != nil {
			panic(err)
		}
	}
}

func isExists(f string) bool {
	_, err := os.Stat(f)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

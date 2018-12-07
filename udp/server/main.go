package main

import (
	"fmt"
	"io/ioutil"
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
		err = ioutil.WriteFile(output, data[:n], 0644)
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

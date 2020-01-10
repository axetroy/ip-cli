package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

func main() {
	fmt.Println("local", getLocalIP())
	fmt.Println("public", getPublicIP())
}

func getLocalIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")

	if err != nil {
		log.Panicln(err)
	}

	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func getPublicIP() string {
	url := "https://api.ipify.org?format=text" // we are using a public IP API, we're using ipify here, below are some others
	// https://www.ipify.org
	// http://myexternalip.com
	// http://api.ident.me
	// http://whatismyipaddress.com/api
	resp, err := http.Get(url)

	if err != nil {
		log.Panicln(err)
	}

	defer resp.Body.Close()

	ip, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Panicln(err)
	}

	return string(ip)
}

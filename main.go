package main

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/elazarl/goproxy"
	"golang.org/x/net/proxy"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s [http proxy listen port] [socks5 host] [socks5 port] [optional: listen host]\n", os.Args[0])
		os.Exit(1)
	}
	socks5 := os.Args[2] + ":" + os.Args[3]
	os.Setenv("HTTP_PROXY", "socks5://"+socks5)
	fmt.Println(os.Args)
	sock, err := proxy.SOCKS5("tcp", socks5, nil, proxy.Direct)
	if err != nil {
		fmt.Println("connect to socks5 proxy error", err)
		os.Exit(2)
	}

	proxy := goproxy.NewProxyHttpServer()
	proxy.ConnectDial = func(network string, address string) (net.Conn, error) {
		return sock.Dial(network, address)
	}
	host := "localhost"
	if len(os.Args) > 4 {
		host = os.Args[4]
	}
	http.ListenAndServe(host + ":"+os.Args[1], proxy)
}

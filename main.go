package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func customUsage() {
	fmt.Fprintf(os.Stderr,
		`Usage: %s [OPTION]...
A TCP Echo Server - sends data recived back to sender

Options:
`,
		os.Args[0])

	flag.PrintDefaults()
}

func echo(conn net.Conn) {
	defer conn.Close()

	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalln("Unable to read/write data")
	}

}

func main() {
	flag.Usage = customUsage

	var host string
	flag.StringVar(&host, "host", "0.0.0.0", "host to bind to")

	var port int
	flag.IntVar(&port, "port", 20080, "port to listen on")

	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))

	if err != nil {
		log.Fatalln("Unable to bind to port")
	}

	log.Printf("Listening on %s:%d\n", host, port)

	for {
		conn, err := listener.Accept()
		log.Printf("Received connection from %s\n", conn.RemoteAddr())

		if err != nil {
			log.Fatalln("Unable to accept connection")
		}

		go echo(conn)
	}
}

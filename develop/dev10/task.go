package main

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var timeout time.Duration
	flag.DurationVar(&timeout, "timeout", 10*time.Second, "timeout for connection")
	flag.Parse()

	if flag.NArg() < 2 {
		fmt.Println("Usage: go-telnet [--timeout=<timeout>] host port")
		os.Exit(1)
	}

	host := flag.Arg(0)
	port := flag.Arg(1)

	address := fmt.Sprintf("%s:%s", host, port)

	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Set up signal handling for Ctrl+C or Ctrl+D
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		for {
			// Read a single byte from stdin
			var b [1]byte
			_, err := os.Stdin.Read(b[:])
			if err != nil && err != io.EOF {
				fmt.Println("Error reading from stdin:", err)
				os.Exit(1)
			}

			// If Ctrl+D is pressed, close the connection
			if b[0] == 4 {
				fmt.Println("Ctrl+D pressed. Closing connection...")
				conn.Close()
				return
			}

			// Write the byte to the connection
			_, err = conn.Write(b[:])
			if err != nil {
				fmt.Println("Error writing to server:", err)
				return
			}
		}
	}()

	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if err != nil && err != io.EOF {
				fmt.Println("Error reading from server:", err)
				os.Exit(1)
			}
			fmt.Print(string(buf[:n]))
		}
	}()

	// Wait for Ctrl+C to close the connection
	<-sigCh
	fmt.Println("Closing connection...")
	conn.Close()
}

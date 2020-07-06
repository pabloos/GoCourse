package main

import (
	"bufio"
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("unable to start server: %s", err)
	}
	defer listener.Close()

	log.Printf("Chat server started on :8888")

	var users []net.Conn

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept connection: %s", err)
			continue
		}

		io.WriteString(conn, "Bienvenido al chat de GDG Marbella!\n")

		users = append(users, conn)

		go func() {
			for {
				msg, err := bufio.NewReader(conn).ReadString('\n')
				if err != nil {
					log.Println(err)
					continue
				}

				for _, user := range users {
					if user.RemoteAddr().String() != conn.RemoteAddr().String() {
						io.WriteString(user, msg)
					}
				}
			}
		}()
	}
}

package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	// names, err := net.LookupAddr("8.8.8.8") // google DNS
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// for _, name := range names {
	// 	fmt.Println(name)
	// }

	//DNS query
	// addrs, err := net.LookupHost("www.google.com")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// for _, addr := range addrs {
	// 	fmt.Println(addr)
	// }

	res := net.Resolver{
		PreferGo: true,
	}

	addrs, err := res.LookupIPAddr(context.Background(), "drive.google.com")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, addr := range addrs {
		fmt.Println(addr.String())
	}

	// interfaces, _ := net.Interfaces()
	// for _, iface := range interfaces {
	// 	addrs1 := iface.HardwareAddr.String()

	// 	if len(addrs1) > 0 {
	// 		fmt.Println(iface.Name)
	// 		fmt.Println(iface.Addrs)
	// 	}
	// }

	ip := net.ParseIP("10.0.1.1")

	// get the IP mask
	fmt.Println(ip.DefaultMask()) // ff000000

	_, ipv4Net, _ := net.ParseCIDR("10.0.1.1/8")

	ipv4Net.Contains(net.ParseIP("192.168.1.1")) //false

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer listener.Close()

	go func() {
		for {
			conn, err := net.Dial("tcp", "localhost:8080")
			if err != nil {
				fmt.Println(err)
				return
			}

			msg, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(string(msg))
			}
		}
	}()

	// go func() {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept connection: %s", err.Error())
			continue
		}

		fmt.Println("Served started")

		io.WriteString(conn, "Hello!")
	}
	// }()
}

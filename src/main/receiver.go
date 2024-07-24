package main

import (
	"fmt"
	"net"
	"net/http"
	"strconv"
)

func receiveToken(port int) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

	fmt.Printf("Receiver server started on %d", port)
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		go handleReceiverConn(conn)
	}
}

func handleReceiverConn(conn net.Conn) error {
	defer conn.Close()

	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)
	if err != nil {
		return err
	}

	amount, err := strconv.Atoi(string(buffer[:n]))
	if err != nil {
		fmt.Println("Error converting received data to integer:", err)
		return err
	}

	fmt.Printf("Received amount: %d\n", amount)
	UpdateWalletBalance(amount, "wallet-receiver.tok")
	return nil
}

package main

import (
	"fmt"
	"net"
	"strconv"
)

func sendToken(tx Transaction) error {

	amountBytes := []byte(strconv.Itoa(tx.Amount))

	conn, err := net.Dial("tcp", tx.Receiver)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write(amountBytes)
	if err != nil {
		fmt.Println("Error sending amount to the server:", err)
		return nil

	}

	UpdateWalletBalance(tx.Amount*-1, "wallet-sender.tok") // Not using a temp variable or importing math module

	return nil
}

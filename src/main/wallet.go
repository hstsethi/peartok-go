package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func UpdateWalletBalance(additionalAmount int, walletPath string) error {
	file, err := os.Open(walletPath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		initialAmount, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return err
		}

		totalAmount := initialAmount + additionalAmount

		file, err = os.OpenFile(walletPath, os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Println("Error opening file for writing:", err)
			return err
		}
		defer file.Close()

		_, err = fmt.Fprintf(file, "%d", totalAmount)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return err
		}
	}
	return nil
}

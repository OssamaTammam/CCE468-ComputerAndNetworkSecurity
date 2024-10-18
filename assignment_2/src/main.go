package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileName := "ciphertexts.txt"

	// Open the file and check for errors
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}
	defer file.Close()

	ciphertexts := []string{} // Holds all out ciphertexts

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cipher := scanner.Text()
		ciphertexts = append(ciphertexts, cipher)
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}

	
}

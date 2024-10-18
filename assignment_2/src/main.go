package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Position struct {
	Sentence int
	Letter   int
}

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

	// Make empty plaintexts to hold possible information while decrypting the ciphertexts
	plaintexts := make([]string, len(ciphertexts))
	placeholder := strings.Repeat("$", len(ciphertexts[0])/2)
	for i := range plaintexts {
		plaintexts[i] = placeholder
	}

	// Locate possible space positions
	possibleSpaces := []Position{}

}

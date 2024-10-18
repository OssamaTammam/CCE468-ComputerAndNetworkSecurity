package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Position struct {
	Sentence int
	Letter   int
}

func hexToBinary(hexString string) (byte, error) {
	// Parse the hexadecimal string as an unsigned integer (base 16)
	value, err := strconv.ParseUint(hexString, 16, 8)
	if err != nil {
		return 0, fmt.Errorf("invalid hexadecimal string: %w", err)
	}

	return byte(value), nil
}

func locateSpaces(ciphertexts []string, plaintexts *[]string) []Position {
	possibleSpaces := []Position{}

	return possibleSpaces
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
	plaintexts := make([][]rune, len(ciphertexts))
	for i := range plaintexts {
		plaintexts[i] = make([]rune, len(ciphertexts[0])/2)
		for j := 0; j < len(ciphertexts[0])/2; j++ {
			plaintexts[i][j] = '$'
		}
	}

	// Locate possible space positions
	possibleSpaces = locateSpaces()

}

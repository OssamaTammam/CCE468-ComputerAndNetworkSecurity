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

func hexToBinary(hexString string) byte {
	// Parse the hexadecimal string as an unsigned integer (base 16)
	value, _ := strconv.ParseUint(hexString, 16, 8)

	return byte(value)
}

func locateSpaces(ciphertexts []string, plaintexts *[][]rune) []Position {
	possibleSpaces := []Position{}

	ciphertextLen := len(ciphertexts[0])

	// Character position in ciphers
	for charPosition := 0; charPosition < ciphertextLen; charPosition += 2 {
		// Obtain a letter sentence position
		var letterSentence int
		for i := range ciphertexts {
			firstLetterBinary := hexToBinary(ciphertexts[i][charPosition : charPosition+2])
			secondLetterBinary := hexToBinary(ciphertexts[i+1][charPosition : charPosition+2])

			result := firstLetterBinary ^ secondLetterBinary

			// Same letter skip
			if result == 0 {
				continue
			} else if byte(1<<6)&result == 0 {
				letterSentence = i
				break
			}
		}

		// Get empty spaces using that character
		charBinary := hexToBinary(ciphertexts[letterSentence][charPosition : charPosition+2])
		for i := range ciphertexts {
			// Don't check with self
			if i == letterSentence {
				continue
			}

			currBinary := hexToBinary(ciphertexts[i][charPosition : charPosition+2])
			result := charBinary ^ currBinary

			// Space located
			if result&byte(1<<6) != 0 {
				(*plaintexts)[i][charPosition/2] = ' '
				possibleSpaces = append(possibleSpaces, Position{i, charPosition})
			}
		}
	}

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
	possibleSpaces := locateSpaces(ciphertexts, &plaintexts)

	// Use the spaces located to check for the rest of the letters

}

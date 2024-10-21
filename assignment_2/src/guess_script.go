package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func hexToBinary(hexString string) byte {
	// Parse the hexadecimal string as an unsigned integer (base 16)
	value, _ := strconv.ParseUint(hexString, 16, 8)

	return byte(value)
}

func main() {
	cipherFile, _ := os.Open("ciphertexts.txt")
	defer cipherFile.Close()

	ciphertexts := []string{} // Holds all out ciphertexts

	cipherScanner := bufio.NewScanner(cipherFile)
	for cipherScanner.Scan() {
		cipher := cipherScanner.Text()
		ciphertexts = append(ciphertexts, cipher)
	}

	decodedFile, _ := os.Open("decodedtexts.txt")
	defer decodedFile.Close()

	decodedtexts := []string{}

	decodedScanner := bufio.NewScanner(decodedFile)
	for decodedScanner.Scan() {
		decodedtext := decodedScanner.Text()
		decodedtexts = append(decodedtexts, decodedtext)
	}

	decryptedTexts := make([][]rune, len(ciphertexts))
	for i := range decryptedTexts {
		decryptedTexts[i] = make([]rune, len(ciphertexts[0])/2)
		for j := 0; j < len(ciphertexts[0])/2; j++ {
			decryptedTexts[i][j] = '$'
		}
	}

	decryptedText := "Secure key exchange is needed for symmetric key encryption"
	decryptedCipher := "E7BE2598988D4FC325D86F2CEA7193F117EC2588E19A2B859D67FA847426F230C10EAC3ECE55EAC170092D7FACAE8FDEF436B0A164EF3C267014"

	for cipherIndex := range ciphertexts {
		for charIndex := 0; charIndex < len(ciphertexts[0]); charIndex += 2 {
			cipherChar := hexToBinary(ciphertexts[cipherIndex][charIndex : charIndex+2])
			knownCharCipher := hexToBinary(decryptedCipher[charIndex : charIndex+2])
			knownCharBits := byte(decryptedText[charIndex/2])

			decryptedChar := cipherChar ^ knownCharCipher ^ knownCharBits
			decryptedTexts[cipherIndex][charIndex/2] = rune(decryptedChar)
		}
	}

	decodedSentences := make([]string, len(ciphertexts))
	for i := range ciphertexts {
		decodedSentences[i] = string(decryptedTexts[i])
	}

	// Print decoded sentences
	for i := range decodedSentences {
		fmt.Println(decodedSentences[i])
	}
}

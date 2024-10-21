package main

import (
	"fmt"
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

func decodeLetters(possibleSpaces []Position, ciphertexts []string, plaintexts *[][]rune) {
	// Store what cols are already decoded
	decodedPositions := make(map[int]bool)
	spaceBinary := byte(0b00100000)

	// Loop over possibleSpaces and decoded cols based on it
	for _, space := range possibleSpaces {
		// Check if position was already decoded
		if decodedPositions[space.Letter] {
			continue
		}

		firstCharBinary := hexToBinary(ciphertexts[space.Sentence][space.Letter : space.Letter+2])

		for i := range ciphertexts {
			// Don't decode self
			if i == space.Sentence {
				continue
			}

			secondCharBinary := hexToBinary(ciphertexts[i][space.Letter : space.Letter+2])
			result := firstCharBinary ^ secondCharBinary ^ spaceBinary

			// Update the plaintexts
			(*plaintexts)[i][space.Letter/2] = rune(result)

		}

		// Column/Letter position doesn't need to be checked again
		decodedPositions[space.Letter] = true
	}
}

func main() {
	// Holds all out ciphertexts
	ciphertexts := []string{
		"F9B4228898864FCB32D83F3DFD7589F109E33988FA8C7A9E9170FB923065F52DD648AA2B8359E1D122122738A8B9998BE278B2BD7CF3313C7609",
		"F5BF229F8F9B1C8832C0212DFD7F92EA18FF29C7E6C968848D6EFAC16074F129D640AB67CE59E3DC6109212AB4EB959FFD34F3B269EB292C7409",
		"FDAF668499C801C734813F3BF3718FF91AEA2C88FC862B999D6EE7C16369F83ADF57FF28CD18FCCC6F0D2B2BB5A295DEF436B0A164EF3C267014",
		"FDFB35858B8403882EC4392CE03289F50CF82588FC816ECB8B63F3843076F52CC059B035C718E0DB220D3B33B3A28692F478B2B07EF03D216B09",
		"E4BE239FCA9A0ADE29C43869FD74DBE31CE835DAE19D72CB9567FD897168FD2CDE5DFF35C65CFAD667136E29B2A7989BE339B1BA71F63C267A09",
		"F8BE279F848101CF60C9203EB26694B00EF929DCEDC9788E9B77EC843075FB39C759BE35C618E6C622016E31A2A8938DE239A1AA3DEC23267316",
		"E7BE2598988D4FC325D86F2CEA7193F117EC2588E19A2B859D67FA847426F230C10EAC3ECE55EAC170092D7FACAE8FDEF436B0A164EF3C267014",
		"E7BE259898811BD160C03B69E67A9EB01CF330CDE69A6ECB9764BE946367F636DF47AB3E835BE0C06E046E3BA6A69799F478A0B67EEA3A266B03",
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
	decodeLetters(possibleSpaces, ciphertexts, &plaintexts)

	// Make the final sentences
	decodedSentences := make([]string, len(ciphertexts))
	for i := range ciphertexts {
		decodedSentences[i] = string(plaintexts[i])
	}

	// Print decoded sentences
	fmt.Println("Decoded texts:")
	for i := range decodedSentences {
		fmt.Println(decodedSentences[i])
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

	decryptedSentences := make([]string, len(ciphertexts))
	for i := range ciphertexts {
		decryptedSentences[i] = string(decryptedTexts[i])
	}

	// Print decoded sentences
	println("\nDecrypted texts after guessing 7th sentence:")
	for i := range decryptedSentences {
		fmt.Println(decryptedSentences[i])
	}
}

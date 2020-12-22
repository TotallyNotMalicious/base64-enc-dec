package main

import (
	b64 "encoding/base64"
	"fmt"
)

var input string

func encrypt(input string) {
	encoded := b64.StdEncoding.EncodeToString([]byte(input))
	fmt.Println()
	fmt.Println(encoded)
}

func decrypt(input string) {
	decoded, _ := b64.StdEncoding.DecodeString(input)
	fmt.Println()
	fmt.Println(string(decoded))
}

func main() {

	fmt.Print("\nDo You Want To Encrypt A String Or Decrypt A String?: ")
	fmt.Scanln(&input)

	if input == "encrypt" || input == "ENCRYPT" {

		fmt.Print("\nWhat Is The String You Would Like To Encrypt?: ")
		fmt.Scanln(&input)
		encrypt(input)

	} else if input == "decrypt" || input == "DECRYPT" {

		fmt.Print("\nWhat Is The String You Would Like To Decrypt?: ")
		fmt.Scanln(&input)
		decrypt(input)
	}
}

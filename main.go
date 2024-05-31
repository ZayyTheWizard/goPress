package main

import (
	"fmt"
	"log"
	"os"

	EnFile "github.com/ZayyTheWizard/goPress/FileHandling"
	huffman "github.com/ZayyTheWizard/goPress/Huffman"
)

// Driver
func main() {
	if len(os.Args) == 1 {
		fmt.Println("Invalid Argument main.go <FileName> <Encode or Decode>")
		os.Exit(255) // Exit code 255 for invalid format
	}

	fileName := os.Args[1]

	if len(os.Args) == 2 {
		fmt.Println("Invalid Argument main.go <FileName> <Encode or Decode>")
		os.Exit(255) // Exit code 255 for invalid format
	}

	switch encodeOrDecode := os.Args[2]; encodeOrDecode {
	case "encode":
		fmt.Println("encode")
		fileContent, err := EnFile.FetchFileContent(fileName)
		if err != nil {
			log.Fatal(err)
		}

		huffman.Encode(fileContent)
	case "decode":
		fmt.Println("decode")
	default:
		fmt.Println("Encode or Decode")
		os.Exit(1)
	}
}

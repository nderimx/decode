package main

import (
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fileName := os.Args[1]
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("could not read file: %v\n", err)
	}
	inputEncoding := os.Args[2]
	outputEncoding := os.Args[3]
	var decodedInput []byte
	switch inputEncoding {
	case "bytes", "blob", "byte_array":
		decodedInput = data
	case "b64", "base64":
		decodedString, err := b64.StdEncoding.DecodeString(string(data))
		if err != nil {
			fmt.Printf("could not decode file data: %v\n", err)
		}
		decodedInput = []byte(decodedString)
	}
	switch outputEncoding {
	case "unicode", "string", "utf8", "utf-8":
		fmt.Println(string(decodedInput))
	case "b64", "base64":
		output := b64.StdEncoding.EncodeToString(decodedInput)
		fmt.Println(string(output))
	}
}

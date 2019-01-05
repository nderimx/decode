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
	case "bytes", "blob", "byte-array", "unknown":
		decodedInput = data
	case "b64", "base64":
		decodedString, err := b64.StdEncoding.DecodeString(string(data))
		if err != nil {
			fmt.Printf("could not decode file data: %v\n", err)
		}
		decodedInput = []byte(decodedString)
	}
	switch outputEncoding {
	case "string", "utf-8", "utf8":
		fmt.Println(string(decodedInput))
	case "bytes":
		fmt.Println(decodedInput)
	case "escaped-unprintables":
		fmt.Printf("%q\n", decodedInput)
	case "escaped-non-ascii":
		fmt.Printf("%+q\n", decodedInput)
	case "hex", "hexadecimal":
		fmt.Printf("% x\n", decodedInput)
	case "b64", "base64":
		output := b64.StdEncoding.EncodeToString(decodedInput)
		fmt.Println(string(output))
	case "utf-32", "utf32", "code-points", "runes":
		// 65533 means it didn't make it past utf-8 conversion
		s := string(decodedInput)
		var strs []rune
		for _, c := range s {
			strs = append(strs, c)
		}
		fmt.Printf("%v\n", strs)
	}

}

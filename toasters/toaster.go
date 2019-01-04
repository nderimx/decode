package main

import "fmt"

func main() {
	// the rune 65533 maps to 'unknown character'
	// when a byte value is utf-8 invalid it gets represented as the rune value 65533
	s := "\n\xbd½\xb2\x3d\xbc\x20\xe2\x8c\x98" //+"Ç"
	// s = s + string(byte(199)) + string(byte(169)) + "⺀"
	// plain old string printing
	fmt.Printf("\nString: %v\n %v \n\n", len(s), s)
	// if ascii is exceeded, utf-8 encoding kicks in
	// ( one char gets represented by more that one byte )
	fmt.Printf("Bytes: %v\n %v \n\n", len([]byte(s)), []byte(s))
	var strs []rune
	for _, c := range s {
		// unicode representation without U+ prefix,
		// more specifically utf-32, since a rune is an int32
		strs = append(strs, c)
	}
	fmt.Printf("Runes one by one: %v\n %v \n\n", len(strs), strs)

	var bts []byte
	for _, c := range s {
		// prints only the least significant utf-32 byte if it exeeds 256 bits
		bts = append(bts, byte(c))
	}
	fmt.Printf("Bytes one by one: %v\n %v \n\n", len(bts), bts)

	fmt.Printf("Hex bytes: % x\n\n", s)

	fmt.Printf("Hex bytes w/ escaped unprintables: %q\n\n", s)

	fmt.Printf("Hex bytes w/ escaped non ascii: %+q\n\n", s)
}

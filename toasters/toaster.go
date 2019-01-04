package main

import "fmt"

func main() {
	s := "!@#$%^&*()_+/~}|{:'\\[]Ç"
	s = s + string(byte(199)) + string(byte(169)) + "⺀"
	fmt.Printf("String: %v\n %v \n", len(s), s)
	// if ascii is exceeded, utf-8 encoding kicks in
	// ( one char gets represented by more that one byte )
	fmt.Printf("Bytes: %v\n %v \n", len([]byte(s)), []byte(s))
	fmt.Println("Chars one by one")
	var strs []rune
	for _, c := range s {
		// unicode representation without U+ prefix,
		// more specifically utf-32, since a rune is an int32
		strs = append(strs, c)
	}
	fmt.Printf("Bytes: %v\n %v \n", len(strs), strs)

	fmt.Println("Bytes one by one")
	var bts []byte
	for _, c := range s {
		// prints only the least significant byte if it exeeds bytes
		bts = append(bts, byte(c))
	}
	fmt.Printf("Bytes: %v\n %v \n", len(bts), bts)
}

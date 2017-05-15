package main

import (
	"fmt"

	"./unicode"
)

// Kodet for shift_jis kan alternativt bruke eucJP
const UNICODE3 = "\x22\x96\x6B\x82\xC6\x93\xEC\x20\x81\x68"

func main() {

	fmt.Println(unicode.Translate(UNICODE3, "jp"))
	fmt.Println(unicode.Translate(UNICODE3, "is"))
}

package main

import (
	"fmt"

	"./iso"
)

func main() {
	iso.IterateOverASCIIStringLiteral()
	fmt.Println(iso.GreetingExtendedASCII())
	fmt.Printf("%c", iso.GreetingExtendedASCII2())
}

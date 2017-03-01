package main

import (
	"fmt"

	"./iso"
)

func main() {
	iso.IterateOverASCIIStringLiteral()
	fmt.Println(iso.GreetingExtendedASCII())
}

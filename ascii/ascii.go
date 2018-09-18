package ascii

import "fmt"


func IterateOverASCIIStringLiteral(ASCII string) {

	for i := 0; i < len(ASCII); i++ {
		fmt.Printf("%X %c %b\n", ASCII[i], ASCII[i], ASCII[i])
	}
}

func GreetingASCII() string {
	a := "\x22\x48\x65\x6C\x6C\x6F\x20\x3A\x2D\x29\x22"
	for i := 0; i < len(a); i++ {
		fmt.Printf("%c", a[i])
	}
	fmt.Println()
	return a
}

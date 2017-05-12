package lang

import (
	"fmt"
)

func Lang() string {
	//var extascii []byte
	a := "\xEF\xDA\xA3\xD2\xD3\xCB\x0A\xEF\xDA\xA3\xD2\xD3\xCB\xC1\x0A\xEF"
	for i := 0; i < len(a); i++ {
		//fmt.Println(extascii)
		fmt.Printf("|%c|\n", a[i])

	}
	return a
}

func Lang2() string {
	//var extascii []byte
	b := "\xFE\xFD\x73\x6B\x61\x72\x0A\xFE\xFD\x73\x6B\x61\x72\x61\x6E\x61"
	for i := 0; i < len(b); i++ {
		//fmt.Println(extascii)
		fmt.Printf("|%c|\n", b[i])

	}
	return b
}

func Lang3() string {
	//var extascii []byte
	c := "\xF8\x79\x65\x73\x70\x65\x73\x69\x61\x6C\x69\x73\x74\x65\x6E\x0A"
	for i := 0; i < len(c); i++ {
		//fmt.Println(extascii)
		fmt.Printf("|%c|\n", c[i])

	}
	return c
}

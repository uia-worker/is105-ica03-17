package main

import (
	"fmt"
	"path/filepath"

	"./fileutils"
)

func main() {
	//KOI8-R
	absPathRu, _ := filepath.Abs("./files/lang01.wl")
	fmt.Printf("%s", fileutils.FileToByteslice(absPathRu))
	//WINDOWS1252
	absPathIce, _ := filepath.Abs("./files/lang02.wl")
	fmt.Printf("%s", fileutils.FileToByteslice(absPathIce))
	//WINDOWS1252
	absPathNor, _ := filepath.Abs("./files/lang03.wl")
	fmt.Printf("%s", fileutils.FileToByteslice(absPathNor))

	hexPath, _ := filepath.Abs("./treasure/treasure.txt")
	fmt.Printf("%s", fileutils.FileToByteslice(hexPath))
}

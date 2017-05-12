package main

import (
	"fmt"
	"path/filepath"

	"./fileutils"
)

func main() {
	//KOI8-R
	absPathRu, _ := filepath.Abs("/GitHub/org/is105-ica03/files/lang01.wl")
	fmt.Printf("%s", fileutils.FileToByteslice(absPathRu))
	//WINDOWS1252
	absPathIce, _ := filepath.Abs("/GitHub/org/is105-ica03/files/lang02.wl")
	fmt.Printf("%s", fileutils.FileToByteslice(absPathIce))
	//WINDOWS1252
	absPathNor, _ := filepath.Abs("/GitHub/org/is105-ica03/files/lang03.wl")
	fmt.Printf("%s", fileutils.FileToByteslice(absPathNor))

	hexPath, _ := filepath.Abs("/GitHub/org/is105-ica03/treasure/treasure.txt")
	fmt.Printf("%s", fileutils.FileToByteslice(hexPath))
}

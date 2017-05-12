package main

import (
	"path/filepath"

	"./treasure"
)

func main() {

	hexPath, _ := filepath.Abs(".//treasure/treasure.txt")
	//fmt.Printf("%s", fileutils.FileToByteslice(hexPath))
	treasure.PrintTreasureUTF8(hexPath)

}

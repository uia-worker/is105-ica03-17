package main

import (
	"path/filepath"

	"./treasure"
)

func main() {

	hexPath, _ := filepath.Abs("/GitHub/org/is105-ica03/treasure/treasure.txt")
	//fmt.Printf("%s", fileutils.FileToByteslice(hexPath))
	treasure.PrintTreasureUTF8(hexPath)

}

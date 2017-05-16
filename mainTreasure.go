package main

import (
	"path/filepath"

	"./treasure"
	"fmt"
)

func main() {

	hexPath, _ := filepath.Abs("./treasure/treasure.txt")

	treasure.PrintTreasureUTF8(hexPath)

}

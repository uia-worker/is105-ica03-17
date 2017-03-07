package treasure

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Kode for Oppgave 3c
// Bruk strengen fra filen treasure.txt som in-data for denne funksjonen
func PrintTreasureUTF8(treasureString string) []byte {

	file, err := os.Open(treasureString)

	if err != nil {
		log.Fatal(err)
	}
	finfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	sizeOfSlice := finfo.Size()

	// The file.Read() function can read a
	// tiny file into a large byte slice,
	// but io.ReadFull() will return an
	// error if the file is smaller than
	// the byte slice
	byteSlice := make([]byte, sizeOfSlice)

	_, err = io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(byteSlice)
	for i := 0; i < len(byteSlice); i++ {
		//fmt.Printf("%X %+q %b\n", byteSlice[i], byteSlice[i], byteSlice[i])
		//fmt.Printf("%+q", byteSlice[i])
	}
	return []byte{'\x00'} // returverdien er her kun en stedsholder
}

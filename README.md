# is105-ica03
Repository for IS-105 experiments with encoding of signs and symbols.


Oppgave 2
https://github.com/GB-Noname/is105-ica03/blob/master/iso/iso.go

a. kjøres i main.go, kode ligger i iso.go
b. kjøres i main.go, kode ligger i iso.go


Kode-kommentarer:

ascii.go 
(under mappen ascii)
Filen inneholder en konstant (const ASCII), og 2 funksjoner (func IterateOverASCIIStringLiteral, 
og func GreetingASCII)
Const ASCII er en String med karakterer representert med heksadesimal.
func IterateOverASCIIStringLiteral står for ittereringen av print funksjonen, og printer ut hver
karakter i const ASCII, og blir konvertert/oversatt på 3 forskjellige måter (%X, %+q og %b\n).
func GreetingASCII er av typen String. funksjonen printer ut teksten "hello"

ascii_test.go 
(under mappen ascii)
Filen inneholder 2 funksjoner; func TestGretingsASCII og func isASCII.
func TestGretingsASCII er en testing-funksjon sjekker om funksjonen GreetingASCII fra ascii.go er ascii.
func isASCII er en boolean, som basert på range ser om det er ASCII eller ikke.

fileutils.go
Filen har en funksjon i seg, som er FileToBytesslice. funksjonen henter en fil og konverterer den om til en 
byteslice, for så å returnere byteslicen. 
Dersom filen er mindre enn byteslicen, vil funksjonen returnere en feilmelding.

iso.go
Filen har 1 konstant (const ASCII), og 3 funksjoner (func IterateOverASCIIStringLiteral, func GreetingExtendedASCII,
og func GreetingExtendedASCII2. 
Const ASCII er en String med karakterer representert med heksadesimal.
func IterateOverASCIIStringLiteral er en iterator som printer ut hver byte i konstanten ASCII, og koverterer de på 3
forskjellige måter (%X, %+q og %b\n).
func GreetingExtendedASCII er a typen String. funksjonen printer ut teksten hello ved bruk av det utvidede ASCII-table.
func GreetingExtendedASCII2 er av typen String. funksjonen printer ut teksten hello ved bruk av det utvidede ASCII-table.

iso_test.go
(under mappen ascii)
Filen inneholder 2 funksjoner; func TestGretingsASCII og func isASCII.
func TestGretingsExtendedASCII er en testing-funksjon sjekker om funksjonen GreetingExtendedASCII fra iso.go er fra
det utvidede ascii-table.
func isASCII er en boolean, som basert på range ser om det er fra det utvidede ASCII-table eller ikke.

lang.go
Filen inneholder 3 funksjoner; func Lang, func Lang2, og func Lang3. alle er av typen String.
func Lang har en slice med bytes a som den koverterer til unicode-format, og returnerer.
func Lang2 har en slice med bytes b som den koverterer til unicode-format, og returnerer.
func Lang3 har en slice med bytes c som den koverterer til unicode-format, og returnerer.

treasure.go
Filen har 1 konstant (const tres), og en funksjon (func PrintTreasureUTF8).
const tres er en slice med karakterer representer med heksadesimal.
func PrintTreasureUTF8 bruker strengen fra filen treasure.txt som in-data. Funksjonen leser en liten fil om til en
større byteslice. Dersom filen er mindre en byteslicen, vil io.Readfull returnere en feilmelding.
return-verdien til funksjonen fungerer kun som en stedsholder.

unicode.go
filen har 1 funksjon (func Translate), og er av typen String.
funksjonen tar inn den norske strengen som inn-data, og returnerer strengen formatert med det nye språket.




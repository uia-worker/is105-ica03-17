# is105-ica03
Repository for IS-105 experiments with encoding of signs and symbols.

### Oppgave 1
https://github.com/GB-Noname/is105-ica03/tree/master/ascii

a. kjøres i mainASCII.go, kode ligger i ascii.go

<<<<<<< HEAD
a. kjøres i main.go, kode ligger i iso.go
b. kjøres i main.go, kode ligger i iso.go
=======
b. kjøres i mainASCII.go, kode ligger i ascii.go

c. ascii_test.go

### Oppgave 2
https://github.com/GB-Noname/is105-ica03/blob/master/iso/

a. kjøres i mainISO.go, kode ligger i iso.go

b. kjøres i mainISO.go, kode ligger i iso.go

c. iso_test.go

### Oppgave 3

https://github.com/GB-Noname/is105-ica03/tree/master/Riddle

a. kjøres i main_riddle.go, kode ligger i tre.go

https://github.com/GB-Noname/is105-ica03/tree/master/fileutils

b. kjøres i mainFileUtils.go, kode ligger i fileutils.go

https://github.com/GB-Noname/is105-ica03/tree/master/treasure

c. kjøres i mainTreasure.go, kode ligger i treasure.go

### Oppgave 4

https://github.com/GB-Noname/is105-ica03/tree/master/unicode

a. kjøres i mainUnicode.go, kode ligger i unicode.go

https://github.com/GB-Noname/is105-ica03/blob/master/server.go

b. server.go
>>>>>>> ed5029ecd5bab26b6c740d45a0580dfb1febd143


Kode-kommentarer:

<<<<<<< HEAD
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



=======
main_riddle.go

Importerer mappen riddle. Programmet kjøres gjennom func main som
henter func RiddleASCII fra tre.go i Riddle mappen.


mainASCII.go

Importerer mappen ascii slik at du kan hente kode fra andre .go 
filer. Programmet kjøres gjennom func main som henter
func IterateOverASCIIStringLiteral og func GreetingASCII fra
ascii.go i ascii mappen. "Constants" er deklarert som variabler,
bare med nøkkelord "const"


mainFileUtils.go

Importerer mappen "fileutils" slik at du kan hente kode fra fileutils.go.
Programmet kjøres gjennom func main som igjen henter func
FileToByteslice. Det hentes også filer fra mappen "files" hvor det hentes
karakter-sett til ulike språk. I denne oppgaven var det russisk, islandsk og norsk.
hexPath henter heksadesimalene fra en tekstfil fra mappen "treasure"


mainISO.go

Importerer mappen "iso" slik at du kan hente kode fra iso.go.
Programmet kjøres gjennom func main som har hentet 
func IterateOverExtendedASCIIStringLiteral og 
func GreetingExtendedASCII fra iso.go.
const deklarerer en konstant verdi.


mainTreasure.go

Importerer mappen "treasure" slik at du kan hente kode fra treasure.go.
Programmet kjøres gjennom func main som har hentet
func PrintTreasureUTF8 fra treasure.go. Det er også satt en hexPath
til tekstfilen i mappen treasure, slik at den kan bruke heksadesimalene
fra den filen.


mainUnicode.go

Importerer mappen "unicode" slik at du kan hente kode fra unicode.go.
Programmet kjøres gjennom func main som printer (fmt.Println) unicode
som er deklarert i "const". Henter også func Translate fra unicode.go
og bruker den i fmt.Println. UNICODE3 er jp = japansk.
Kodet for shift_jis kan alternativt bruke eucJP


server.go

Filen for å eksperimentere med webserver.
Starte server med "go run server.go".
Du kan ha tilgang til server fra nettleser med http://localhost:3000.
I func main har du ListenAndServe som starter HTTP server med en gitt
adresse og "handler". Adressen er 3000 og handler er nil, som betyr
at den skal bruke DefaultServeMux.

func foo skriver data som skal sendes til nettleser.
Standardinstilling er UTF-8.
w.Write er det som blir skrivet på webserveren.
w (parameter) = http.ResponseWriter.
>>>>>>> ed5029ecd5bab26b6c740d45a0580dfb1febd143

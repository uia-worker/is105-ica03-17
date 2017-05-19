# is105-ica03
Repository for IS-105 experiments with encoding of signs and symbols.

### Oppgave 1
https://github.com/GB-Noname/is105-ica03/tree/master/ascii

a. kjøres i mainASCII.go, kode ligger i ascii.go

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


Kode-kommentarer:

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
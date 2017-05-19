# is105-ica03
Repository for IS-105 experiments with encoding of signs and symbols.

#### Bidragsytere: 

På Github står det fem contributors. Alle har bidratt, vi har jobbet i team. 
Ampede15 har samarbeidet med CastleDev. 
Og adrianlo har samarbeidet med MortenSchibbye. 

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



#### Kode-kommentarer:

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

Filen har en funksjon i seg, som heter FileToBytesslice. funksjonen henter en fil og konverterer den om til en 
byteslice, for så å returnere byteslicen. 
Dersom filen er mindre enn byteslicen, vil funksjonen returnere en feilmelding.

iso.go

Filen har 2 funksjoner (func IterateOverExtendedASCIIStringLiteral og func GreetingExtendedASCII 
Const ASCII er en String med karakterer representert med heksadesimal.
func IterateOverASCIIStringLiteral er en iterator som printer ut hver byte i konstanten ASCII, og koverterer de på 3
forskjellige måter (%X, %+q og %b\n).
func GreetingExtendedASCII er a typen String. funksjonen printer ut teksten hello ved bruk av det utvidede ASCII-table.

iso_test.go

(under mappen ascii)
Filen inneholder 2 funksjoner; func TestGretingsASCII og func isASCII.
func TestGretingsExtendedASCII er en testing-funksjon sjekker om funksjonen GreetingExtendedASCII fra iso.go er fra
det utvidede ascii-table.
func isASCII er en boolean, som basert på range ser om det er fra det utvidede ASCII-table eller ikke.

tre.go

Filen har 1 funksjon som heter RiddleASCII. funksjonen printer ut byteslicen a.

treasure.go

Filen har 1 konstant (const tres), og en funksjon (func PrintTreasureUTF8).
const tres er en slice med karakterer representer med heksadesimal.
func PrintTreasureUTF8 bruker strengen fra filen treasure.txt som in-data. Funksjonen leser en liten fil om til en
større byteslice. Dersom filen er mindre en byteslicen, vil io.Readfull returnere en feilmelding.
return-verdien til funksjonen fungerer kun som en stedsholder.

unicode.go

filen har2 konstanter (const UNICODEIS, og const UNICODE3) og 1 funksjon (func Translate) av typen String.
const UNICODEIS er en byteslice kodet for 1253 ISLANDSK.
const UNICODE3 er en byteslice kodet for shift_jis kan alternativt bruke eucJP.
funksjonen func Translate har to parametere (expression string og language string). 
funksjonen vil returnere en av de 2 constantene, avhengig av hvilke input funksjonen får.
inputen er enten jp eller is.


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
Bruker KO8-R for det russiske, og WINDOWS1252 for det norske og islandske.


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

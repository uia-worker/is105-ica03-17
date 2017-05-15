package unicode

// Kodet for 1252 ISLANDSK
const UNICODEIS = "\x22\x6E\x6F\x72\xF0\x75\x72\x20\x6F\x67\x20\x73\x75\xF0\x75\x72\x94"

//const UNICODE2 = "\x81\x5f\x96\x5f\x82\xbe\x93\xde\x81\x5f"
// Kodet for shift_jis kan alternativt bruke eucJP
const UNICODE3 = "\x22\x96\x6B\x82\xC6\x93\xEC\x20\x81\x68"

// Kode for Oppgave 4a
func Translate(expression string, language string) string {

	if language == "jp" {

		language = UNICODE3
	}
	if language == "is" {
		language = UNICODEIS
	}
	return language
}

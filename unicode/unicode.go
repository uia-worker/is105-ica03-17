package unicode


const UNICODEIS = "\x22\x6E\x6F\x72\xF0\x75\x72\x20\x6F\x67\x20\x73\x75\xF0\x75\x72\x94"

const UNICODE3 = "\x22\x96\x6B\x82\xC6\x93\xEC\x20\x81\x68"

func Translate(expression string, language string) string {

	if language == "jp" {

		language = UNICODE3
	}
	if language == "is" {
		language = UNICODEIS
	}
	return language
}

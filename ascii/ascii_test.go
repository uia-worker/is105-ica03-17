package ascii

import "testing"

func TestGretingsASCII(t *testing.T) {
	ascii := GreetingASCII()
	if !(isASCII(ascii)) {
		t.Fail()
	}
}

func isASCII(s string) bool {
	for _, C := range s {
		if C > 127 {
			return false
		}
	}
	return true
}

// var istrue bool = false
//var s1 string = "诶"
//var ascii.GreetingASCII string

/*func TestisASCII(t *testing.T) (s1 string) {
	ascii: =GreetingASCII()
	for _, c := range ascii.GrettingASCII() {
		if c > 127 {
			istrue = false
			return "false"
		}
	}
	istrue = true
	return "true"
}
*/

/*func TestAscii(t *testing.T, s1 string) {
	for i := 0; i < len(s1); i++ {
		//fmt.Printf(" %+q \n", ascii[i])
		//s2 := isASCII()

		if TestisASCII(s1) == true {
			fmt.Print("har bare ascii?", istrue)
		}
	}
}

*/

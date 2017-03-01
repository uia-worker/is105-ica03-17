package iso

func isASCII(s string) bool {
	for _, C := range s {
		if C > 127 {
			return false
		}
	}
	return true
}

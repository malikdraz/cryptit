package decrypt

func Nimbus(str string) string {
	decryptedStr := ""
	for _, val := range str {
		ascii_code := int(val)
		decryptedStr += string(ascii_code - 3)
	}
	return decryptedStr
}

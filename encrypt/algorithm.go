package encrypt

func Nimbus(str string) string {
	encryptedStr := ""
	for _, val := range str {
		ascii_code := int(val)
		encryptedStr += string(ascii_code + 3)
	}
	return encryptedStr
}

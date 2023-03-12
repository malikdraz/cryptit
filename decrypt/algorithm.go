// Package decrypt consists of all the decryption algorithm
package decrypt

// Nimbus decrypts by reducing the ascii code by 3 for character
func Nimbus(str string) string {
	decryptedStr := ""
	for _, val := range str {
		ascii_code := int(val)
		decryptedStr += string(ascii_code - 3)
	}
	return decryptedStr
}

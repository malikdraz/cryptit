package main

import (
	"fmt"
	"github.com/malikdraz/cryptit/decrypt"
	"github.com/malikdraz/cryptit/encrypt"
)

func main() {
	message := encrypt.Nimbus("Malik Draz")
	fmt.Println(message)
	fmt.Println(decrypt.Nimbus(message))
}

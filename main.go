package main

import (
	"fmt"

	"github.com/Harshalyadav/crypit.git/encrypt"

	"github.com/Harshalyadav/crypit.git/dcrypt"
)


func main (){
     encryptVar :=encrypt.EncryptFunction("harshal")
	fmt.Println(encryptVar)

	fmt.Println(dcrypt.EncryptFunction(encryptVar))
}
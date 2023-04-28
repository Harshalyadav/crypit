package encrypt



func EncryptFunction(str string) string {

	encryptString := " "

	for _,c := range str {
		
			asciiCode := int(c)
			character := string(asciiCode + 3)
			encryptString += character

		

		}
		return encryptString
}
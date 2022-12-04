package libs

import (
	"math/rand"
	"time"
)

const alphanum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
const specialchars = "!@#$%&"

func GenerateRandomPassword(length int, option string) string {
	randomstring := alphanum
	if option == "all" {
		randomstring = randomstring + specialchars
	}

	randomstringLength := len(randomstring)
	var result string
	for i := 0; i < length; i++ {
		rand.Seed(time.Now().UnixNano())
		randomPos := rand.Intn(randomstringLength)
		result = result + string(randomstring[randomPos])
	}

	return result
}

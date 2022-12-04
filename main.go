package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const alphanum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
const specialchars = "!@#$%&"

func generateRandomPassword(length int, option string) string {
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

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Usage: randompass [number_of_chars] [alphanum|all]")
		fmt.Println("Example:")
		fmt.Println("randompass 25 alphanum")
		fmt.Println("randompass 32 all")
		return
	}

	option := os.Args[2]
	if option != "alphanum" && option != "all" {
		fmt.Println("Usage: randompass [number_of_chars] [alphanum|all]")
		fmt.Println("Example:")
		fmt.Println("randompass 25 alphanum")
		fmt.Println("randompass 32 all")
		return
	}

	passwordLength, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	fmt.Println(generateRandomPassword(passwordLength, option))
}

package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ksetyadi/randompass/libs"
)

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

	fmt.Println(libs.GenerateRandomPassword(passwordLength, option))
}

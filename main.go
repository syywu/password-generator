package main

import (
	"log"
	"math/rand"
	"strings"
)

func main() {
	log.Println("Your new generated password is:", generator())
}

func generator() string {
	var password strings.Builder

	for i := 0; i < 12; i++ {
		password.WriteString(string(rune(randomInt())))
	}

	return password.String()
}

func randomInt() int {
	min, max := 33, 126
	return rand.Intn(max-min) + min
}

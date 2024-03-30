package main

import (
	"fmt"
	"log"
	"math/rand"
	"os/exec"
	"strings"
)

func main() {
	password := generator()
	fmt.Println(password)

	cmd := exec.Command("./password.sh", password)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
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

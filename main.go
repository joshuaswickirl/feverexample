package main

import (
	"fmt"
	"log"

	"github.com/joshuaswickirl/feverexample/fever"
)

func main() {
	var input float64
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Fatal(err)
	}

	if fever.Determine(input) {
		fmt.Println("Fever")
	} else {
		fmt.Println("Allowed")
	}
}

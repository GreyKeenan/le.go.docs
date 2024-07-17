
package main

import (
	"fmt"

	"example.com/myLibMod"
)

var WORDS = [...]string {
	"bappledebonk",
	"bipplederbankle",
	"blappidydappideedodah",
}

func main() {
	var err error
	var message string

	message, err = myLibMod.Call(WORDS[0])
	if (err != nil) {
		panic(err)
	}

	fmt.Printf("\nString: %s\n", message)


	messages, err := myLibMod.Call_multiple(WORDS[:])
	if (err != nil) {
		panic(err)
	}

	fmt.Println(messages)
	fmt.Println("")
}

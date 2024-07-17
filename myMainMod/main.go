
package main

import (
	"fmt"

	"example.com/myLibMod"
)

func main() {
	var err error
	var message string

	message, err = myLibMod.Call("")
	if (err != nil) {
		panic(err)
	}

	fmt.Printf("String: %s\n", message)
}

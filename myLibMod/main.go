
package myLibMod

import (
	"fmt"
	"errors"
)

func Call(word string) (string, error) {

	if word == "" {
		return "", errors.New("invalid word given")
	}

	return fmt.Sprintf("blippidee blooh blah %v", word), nil
}


package myLibMod

import (
	"fmt"
	"errors"
	"math/rand"
)

var FORMATS = [...]string {
	"blippidee blooh blah %v",
	"boopidy bah bim %v",
	"bibbledee bip bap %v",
}

func Call(word string) (string, error) {

	if word == "" {
		return word, errors.New("invalid word given")
	}

	return fmt.Sprintf(randomFormat(), word), nil
}

func Call_multiple(words []string) (map[string]string, error) {
	var messages map[string]string = make(map[string]string)

	var err error

	var _ int
	var word string
	var message string
	for _, word = range words {
		message, err = Call(word)
		if (err != nil) {
			return nil, err
		}
		messages[word] = message
	}

	return messages, nil
}

func randomFormat() string {
	return FORMATS[rand.Intn(len(FORMATS))]
}

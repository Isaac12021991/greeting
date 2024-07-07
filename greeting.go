package greeting

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("Nombre vacio")
	}
	//message := fmt.Sprintf("Hola %v", name)

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {

	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil

}

func randomFormat() string {
	formats := []string{
		"Hola, %v ! Bienvenido",
		"Hola, %v ! 2",
		"Hola, %v ! 3",
	}

	return formats[rand.Intn(len(formats))]

}

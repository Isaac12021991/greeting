# Saludos en go

Inicializacion en Go prueba de repositorio de goolang

## Instalacion
```go get -u github.com/Isaac12021991/greeting```

## Uso

```
package main

import (
	"fmt"
	"log"

	"github.com/Isaac12021991/greeting"
)

func main() {

	log.SetPrefix("Greeting: ")
	log.SetFlags(2)

	message, err := greeting.Hello("Isaac")
	if err != nil {

		log.Fatal(err)

	}

	fmt.Println(message)

	names := []string{"Isaa", "MAma", "PE"}
	messages, err := greeting.Hellos(names)

	fmt.Println(messages)
}


```
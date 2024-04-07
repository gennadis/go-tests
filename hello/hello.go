package main

import (
	"fmt"
)

const (
	spanish = "Spanish"
	french  = "French"
	turkish = "Turkish"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchhHelloPrefix = "Bonjour, "
	turkishHelloPrefix = "Merhaba, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchhHelloPrefix
	case turkish:
		prefix = turkishHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", "English"))
}

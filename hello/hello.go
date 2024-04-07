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

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchhHelloPrefix
	case turkish:
		prefix = turkishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("world", "English"))
}

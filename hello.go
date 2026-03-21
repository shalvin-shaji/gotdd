package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrifix = "Hola, "
const frenchHelloPrifix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := englishHelloPrefix

	switch language {

	case spanish:
		prefix = spanishHelloPrifix
	case french:
		prefix = frenchHelloPrifix

	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("World", ""))
}

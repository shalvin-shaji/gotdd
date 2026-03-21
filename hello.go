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
	return greatingPrefix(language) + name
}

func greatingPrefix(language string) string {
	prefix := englishHelloPrefix

	switch language {

	case spanish:
		prefix = spanishHelloPrifix
	case french:
		prefix = frenchHelloPrifix

	}
	return prefix
}

func main() {
	fmt.Println(Hello("World", ""))
}

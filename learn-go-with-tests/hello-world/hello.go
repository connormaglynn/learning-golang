package main

const englishHello = "Hello, "
const spanishHello = "Hola, "

func Hello(name string, langauge string) string {
	if name == "" {
		name = "World"
	}
	if langauge == "Spanish" {
		return spanishHello + name
	}
	return englishHello + name
}

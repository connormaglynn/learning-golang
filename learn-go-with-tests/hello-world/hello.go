package main

const englishHello = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHello + name
}

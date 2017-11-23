package main

import (
	"fmt"
)

// type sample interface {
// 	demoFunc(string, int) (string, error)
// }

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "hello"
}

func (spanishBot) getGreeting() string {
	return "hola"
}

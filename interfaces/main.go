package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {

	e := englishBot{}
	s := spanishBot{}

	printGreeting(e)
	printGreeting(s)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "hello!!"
}

func (spanishBot) getGreeting() string {
	return "hola!!"
}

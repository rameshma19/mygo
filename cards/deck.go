package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of deck
// which is a slice of string

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

//receiver sample
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fn string) error {
	return ioutil.WriteFile(fn, []byte(d.toString()), 0666)
}

func newDeckFromFile(fn string) deck {
	bs, err := ioutil.ReadFile(fn)
	if err != nil {
		// Option#1 - log the error and return a call to newDeck()
		// Option#2 - log the error entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	//Ace of Spade, Two of Spades etc.
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffleMyDeck() {
	m1 := make(map[int]string)
	for i := range d {
		m1[i] = d[i]
	}

	//fmt.Println(m1)

	rand.Seed(time.Now().UnixNano())
	m2 := make(map[int]string)
	for j := range d {
		//fmt.Println(rand.Intn(16))
		m2[j] = m1[rand.Intn(16)]
	}

	fmt.Println(m2)
}

func (d deck) shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i := range d {
		j := rand.Intn(len(d) - 1)
		d[i], d[j] = d[j], d[i]
	}
}

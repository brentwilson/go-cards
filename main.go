package main

import "fmt"

func main() {
	cards := newDeckFromFile("cardsasbytes")
	fmt.Println(cards)

}

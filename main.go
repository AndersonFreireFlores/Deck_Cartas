package main

import "fmt"

func main() {
	cards := newDeck()
	cards.saveToFile("minhas_cartas")

	hand, remainigCards := deal(cards, 5)

	fmt.Println("Hand:", hand)
	fmt.Println("RemainigCards:", remainigCards)

}

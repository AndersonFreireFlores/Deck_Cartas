package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainigCards := deal(cards, 5)

	fmt.Println("Hand:", hand)
	fmt.Println("RemainigCards:", remainigCards)
}

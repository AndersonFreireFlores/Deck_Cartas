package main

func main() {

	cards := newDeckFromFile("minhas_cartas")
	cards.shuffle()
	cards.print()

}

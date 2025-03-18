package main


func main() {
	cards := newDeck()
	handOne, handTwo := cards.deal(5)
	
	handOne.print()
	println("----")
	handTwo.print()
}



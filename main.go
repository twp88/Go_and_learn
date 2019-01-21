package main

import "fmt"

func main() {
	deckOfCards := deck{"House of Cards"}

	deckOfCards.print()

	deckOfCards = deckOfCards.add("Bling a ling")
	deckOfCards = deckOfCards.add("Ting ting")

	deckOfCards.print()

	fmt.Println("=================")
	deckOfCards.shuffle()
	deckOfCards.print()

	grooms := groomList{"Desmond", "Kevin"}
	brides := brideList{"Raquel", "Elizabeth"}

	marriedCouples, _ := marryCouples(brides, grooms)

	saveResult := saveMarriedCouples("newlyWeds", marriedCouples.brideListToString())

	fmt.Println(saveResult)
}

func doSliceThing(s []string, lowerRange int, upperRange int) {
	fmt.Println("Here is the slice")
	fmt.Println(s)
	fmt.Println("This is the upper range")
	fmt.Println("I think it will print everthing from position 0 onwards upto but not including pos 4")
	fmt.Println(s[:upperRange])
	fmt.Println("This is the lower range")
	fmt.Println("I think it will print everthing from position 2 onwards")
	fmt.Println(s[lowerRange:])
}

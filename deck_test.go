package main

import (
	"fmt"
	"testing"
)

func TestDeckType(t *testing.T) {
	deckOfCards := deck{"Ace of Spades"}

	if deckOfCards[0] != "Ace of Spades" {
		t.Error("Houston we have an error")
	}
}

func TestAppendToDeck(t *testing.T) {
	deckOfCards := deck{"Ace of Spades,"}
	deckOfCards = deckOfCards.add("House of Cards")

	fmt.Println(deckOfCards)

	if deckOfCards[1] != "House of Cards" {
		t.Error("Houston wid da ting")
	}
}

func TestPrintDeckOfCards(t *testing.T) {
	deckOfCards := deck{"Ace of Spades"}

	if deckOfCards.print() != deckOfCards[0] {
		t.Error("Houston we have an error")
	}
}

func TestParkCarInGarage(t *testing.T) {
	firstCar := "first Example Car"
	secondCar := "second Example Car"

	expected := garage{firstCar, secondCar}

	actual := garage{}
	actual = actual.park(firstCar, secondCar)

	if expected[0] != actual[0] {
		t.Error("Houston we have a problem")
	}
}

func TestMarryCouples(t *testing.T) {
	grooms := groomList{"James", "Richard", "Sebastian"}
	brides := brideList{"Maria", "Siobahn", "Raquel"}

	expectedList, expectedPreviousNewlys := []string{"Maria and James", "Richard and Siobahn", "Sebastian and Raquel"}, "Raquel and Sebastian"

	actualList, actualPreviousNewlys := marryCouples(brides, grooms)

	if actualList[0] != expectedList[0] && actualList[1] != expectedList[1] && actualList[2] != expectedList[2] {
		t.Error("Houston we have a problem")
	}

	if actualPreviousNewlys != expectedPreviousNewlys {
		t.Error("Houston we have a problem")
	}
}

func TestToString(t *testing.T) {
	deckOfCards := deck{"Ace of Spades"}

	actual := deckOfCards.deckToString()
	expected := "Ace of Spades"

	if actual != expected {
		t.Error("This has gone wrong")
	}
}

func TestReadFromFile(t *testing.T) {
	deckOfCards := deck{"Ace of Spades"}
	deckOfCards.writeToFile("CoolFileName")

}

package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type deck []string
type garage []string
type groomList []string
type brideList []string

func (d deck) print() string {
	for i, card := range d {
		fmt.Println(i, card)
	}
	return d[0]
}

func (d deck) add(a string) deck {
	d = append(d, a)
	return d
}

func (g garage) park(firstCar string, secondCar string) garage {
	g = append(g, firstCar)
	g = append(g, secondCar)
	return g
}

func marryCouples(brides brideList, grooms groomList) (brideList, string) {
	marriedCouples := brideList{}
	newlyWeds := ""

	for _, bride := range brides { // this line has an error. It marries each bride twice
		for _, groom := range grooms {
			newlyWeds = bride + " and " + groom
			marriedCouples = append(marriedCouples, newlyWeds)
		}
	}

	return marriedCouples, newlyWeds
}

func saveMarriedCouples(fileName string, content string) error {
	return ioutil.WriteFile(fileName, []byte(content), 0666)
}

func (d deck) deckToString() string {
	return strings.Join(d, ", ")
}

func (s brideList) brideListToString() string {
	return strings.Join(s, ", ")
}

func (d deck) writeToFile(newFileName string) error {
	err := ioutil.WriteFile(newFileName, []byte(d.deckToString()), 0666)
	return err
}

func readFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("This has not gone well")
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() deck {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}

	return d
}

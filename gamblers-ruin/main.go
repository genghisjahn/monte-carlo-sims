package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

const heads = 1
const tails = 0

var s1 = rand.NewSource(time.Now().UnixNano())
var r1 = rand.New(s1)

var awins = 0
var bwins = 0

var numgames = 0
var am = 0
var bm = 0
var andy = gambler{}
var bart = gambler{}

func init() {
	flag.IntVar(&numgames, "games", 100, "number of games, default is 100")
	flag.IntVar(&am, "ma", 100, "Amount of money Andy starts with, default is 100")
	flag.IntVar(&bm, "mb", 100, "Amount of money Bart starts with, default is 100")
}

func main() {
	flag.Parse()
	andy = gambler{Name: "Andy", Money: am}
	bart = gambler{Name: "Bart", Money: bm}
	fmt.Println(andy.Status() + " and " + bart.Status())
	for i := 0; i < numgames; i++ {
		andy.Money = am
		bart.Money = bm
		for {
			flipCoin()
			if andy.Money == 0 || bart.Money == 0 {
				break
			}
		}
		if andy.Money == 0 {
			bwins++
		} else {
			awins++
		}
	}
	fmt.Printf("%v won %v times\n", andy.Name, awins)
	fmt.Printf("%v won %v times\n", bart.Name, bwins)
}

func flipCoin() {
	if r1.Float32() < .5 {
		andy.Money++
		bart.Money--
	} else {
		andy.Money--
		bart.Money++
	}

}

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"sort"
	"time"
)

var fighters = []*gunfighter{}

var s1 = rand.NewSource(time.Now().UnixNano())
var r1 = rand.New(s1)

var data []byte
var jErr error

var flagfile string
var flagfights int
var flaglog bool
var flagcontestants string

func init() {
	flag.IntVar(&flagfights, "r", 100, "Number of rounds, default is 100")
	flag.StringVar(&flagfile, "f", "default", "Name of file minus the .json suffix to pull fighter data from")
	flag.BoolVar(&flaglog, "log", false, "Default is false. Logs the output of each shot.")
	flag.StringVar(&flagcontestants, "c", "", "Specify which fighters file will fight, default is all (empty string)")
}

func main() {
	flag.Parse()
	data, jErr = ioutil.ReadFile(fmt.Sprintf("fighters/%v.json", flagfile))
	if jErr != nil {
		log.Fatal(jErr)
	}
	err := json.Unmarshal(data, &fighters)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < flagfights; i++ {
		fight()
	}

	sort.Sort(byVictories(fighters))
	if flagfights > 1 {
		for _, v := range fighters {
			fmt.Println(v.Name, v.Victories) //, v.Shots, len(v.Kills), v.ShotAt)
		}
	}

}

func fight() {
	for v, g := range fighters {
		g.Dead = false
		g.KilledBy = ""
		fighters[v] = g

	}
	sort.Sort(bySpeedScore(fighters)) //Sorted by speedscored (highest to lowest)

	var c = 0
	var totalShots = 0
	for {
		c++
		for k, v := range fighters {
			v.setspeed()
			fighters[k] = v
		}
		for k, g := range fighters {
			if !g.Dead {
				result, deceased := g.shoot()
				totalShots++
				fighters[k].Shots++
				if deceased != "" {
					fighters[k].Kills = append(fighters[k].Kills, deceased)
				}
				_ = result
				if flaglog {
					fmt.Println("Round:", c, " ", result)
				}
			}
		}
		acount := 0
		victor := ""
		for _, v := range fighters {
			if !v.Dead {
				acount++
				victor = v.Name

			}
		}
		if acount == 1 {
			if flaglog {
				fmt.Println(victor, "is victorious!")
				fmt.Println(c, " rounds")
				fmt.Println(totalShots, " total shots")
			}
			setVictory(victor)
			break
		}
	}

}

func setRndScores() {
	for k := range fighters {
		fighters[k].setrnd()
	}
}

func setVictory(name string) {
	for k, v := range fighters {
		if v.Name == name {
			fighters[k].Victories++
		}
	}
}

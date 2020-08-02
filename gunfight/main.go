package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"sort"
	"strings"
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
var flagtournament int

func init() {
	flag.IntVar(&flagfights, "r", 100, "Number of rounds, default is 100")
	flag.StringVar(&flagfile, "f", "default", "Name of file minus the .json suffix to pull fighter data from")
	flag.BoolVar(&flaglog, "log", false, "Default is false. Logs the output of each shot.")
	flag.StringVar(&flagcontestants, "c", "", "Comma delimited string species which fighters from the file selected will fight, default is all (empty string)")
	flag.IntVar(&flagtournament, "t", 2, "Tournament flag, number of fighters to be in each fight of a round robbin tournament, winner moves on.  Default is 2")
}

func main() {
	var rawfighters = []*gunfighter{}
	flag.Parse()
	data, jErr = ioutil.ReadFile(fmt.Sprintf("fighters/%v.json", flagfile))
	if jErr != nil {
		log.Fatal(jErr)
	}
	err := json.Unmarshal(data, &rawfighters)
	if err != nil {
		log.Fatal(err)
	}

	if flagcontestants != "" {
		names := strings.Split(flagcontestants, ",")
		for _, n := range names {
			for _, v := range rawfighters {
				if strings.ToLower(v.Name) == strings.TrimSpace(n) {
					fighters = append(fighters, v)
				}
			}
		}
		if len(fighters) < 2 {
			log.Fatal("Number of specified fighters must be at least 2")
		}
	} else {
		fighters = rawfighters
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

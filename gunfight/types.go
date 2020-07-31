package main

import (
	"fmt"
	"sort"
)

type gunfighter struct {
	Name       string  `json:"name"`
	Accuracy   float32 `json:"accuracy"`
	Speed      float32 `json:"speed"`
	SpeedScore float32 `json:"speed_score"`
	Confidence float32 `json:"confidence"`
	RndScore   float32 `json:"random_score"`
	Dead       bool
	Shots      int
	Kills      []string
	ShotAt     int
	KilledBy   string
	Victories  int
}

func (g *gunfighter) setspeed() {
	g.SpeedScore = r1.Float32() * g.Speed
}

func (g *gunfighter) setrnd() {
	g.RndScore = r1.Float32()
}

func (g *gunfighter) FinalResult() string {
	var r = fmt.Sprintf("killed by %v", g.KilledBy)
	if !g.Dead {
		r = " victorious"
	}
	return fmt.Sprintf("%v fired %v shot(s), killed (%v), was shot at %v time(s) and was %v", g.Name, g.Shots, len(g.Kills), g.ShotAt, r)
}

func (g *gunfighter) shoot() (string, string) {
	livingFighters := []*gunfighter{}
	setRndScores()
	for _, v := range fighters {
		if !v.Dead {
			livingFighters = append(livingFighters, v)
		}
	}
	e := "%v %v %v%v."
	r := ""

	t := r1.Float32() * 100
	var target = &gunfighter{}
	var comment = ""
	sort.Sort(byAccuracy(livingFighters))
	if g.Confidence*r1.Float32() > t {
		if livingFighters[0].Name != g.Name {
			target = livingFighters[0]
		} else {
			target = livingFighters[1]
		}
		comment = "(deadliest)"
	} else {
		sort.Sort(byRndScore(livingFighters))
		if livingFighters[0].Name != g.Name {
			target = livingFighters[0]
		} else {
			target = livingFighters[1]
		}
	}
	for sk, sa := range fighters {
		if sa.Name == target.Name {
			sa.ShotAt++
			fighters[sk] = sa
		}
	}
	r = "missed"
	d := ""
	shot := r1.Float32() * float32(100.0)
	if g.Accuracy > shot {
		target.Dead = true
		for k, v := range fighters {
			if v.Name == target.Name {
				fighters[k].Dead = true
				r = "killed"
				for sk, sa := range fighters {
					if sa.Name == target.Name {
						d = sa.Name
						sa.KilledBy = g.Name
					}
					fighters[sk] = sa
				}
			}
		}
	}
	return fmt.Sprintf(e, g.Name, r, target.Name, comment), d
}

type byAccuracy []*gunfighter

func (a byAccuracy) Len() int           { return len(a) }
func (a byAccuracy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAccuracy) Less(i, j int) bool { return a[i].Accuracy > a[j].Accuracy }

type bySpeedScore []*gunfighter

func (a bySpeedScore) Len() int           { return len(a) }
func (a bySpeedScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a bySpeedScore) Less(i, j int) bool { return a[i].SpeedScore > a[j].SpeedScore }

type byRndScore []*gunfighter

func (a byRndScore) Len() int           { return len(a) }
func (a byRndScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byRndScore) Less(i, j int) bool { return a[i].RndScore > a[j].RndScore }

type byVictories []*gunfighter

func (a byVictories) Len() int           { return len(a) }
func (a byVictories) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byVictories) Less(i, j int) bool { return a[i].Victories > a[j].Victories }

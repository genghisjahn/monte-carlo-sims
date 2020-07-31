package main

import "fmt"

type gambler struct {
	Name  string
	Money int
}

func (g *gambler) Change(r bool) {
	if r {
		g.Money++
		return
	}
	g.Money--
}

func (g *gambler) Status() string {
	return fmt.Sprintf("%v has %v dollars", g.Name, g.Money)
}

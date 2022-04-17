package main

import (
	eb "github.com/hajimehoshi/ebiten/v2"
	"math/rand"
	"time"
)

type deck struct {
	cards []*card
}

type card struct {
	number int
	symbol string
	image  *eb.Image
}

func (d *deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.cards), func(i, j int) { d.cards[i], d.cards[j] = d.cards[j], d.cards[i] })
}

func (d *deck) getCard() *card {
	c := d.cards[0]
	d.cards = d.cards[1:]
	return c
}

type players struct {
	players []player
}

type player struct {
	ID       string
	cards    []*card
	location location
}

type location struct {
	x, y  float64
	scale float64
}

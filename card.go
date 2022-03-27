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

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

type player struct {
	ID       string
	cards    []*card
	selected *card
	location location
}

type location struct {
	x, y  float64
	scale float64
}

func (p *player) nextSelection() *card {
	index := indexOf(p.cards, p.selected)
	index++
	index %= len(p.cards)
	return p.cards[index]
}

func indexOf(cards []*card, c *card) int {
	if c == nil {
		return -1
	}
	for i, v := range cards {
		if *v == *c {
			return i
		}
	}
	return -1
}

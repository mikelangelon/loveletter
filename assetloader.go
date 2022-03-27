package main

import (
	eb "github.com/hajimehoshi/ebiten/v2"
	"image"
)

const (
	witdh  = 184
	height = 286
)

func newDeck(img *eb.Image) *deck {
	var cards []*card
	for j := 0; j < 2; j++ {
		for i := 0; i < 5; i++ {
			cards = append(cards, &card{
				number: i,
				image:  img.SubImage(image.Rect(witdh*i, height*j, witdh*(i+1), height*(j+1))).(*eb.Image),
			})
		}
	}
	return &deck{
		cards: cards,
	}
}

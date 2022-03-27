package main

import (
	eb "github.com/hajimehoshi/ebiten/v2"
	"image"
)

const (
	cardWidth  = 184
	cardHeight = 286
)

func newDeck(img *eb.Image) *deck {
	var cards []*card
	for j := 0; j < 2; j++ {
		for i := 0; i < 5; i++ {
			im := img.SubImage(image.Rect(cardWidth*i, cardHeight*j, cardWidth*(i+1), cardHeight*(j+1))).(*eb.Image)

			cards = append(cards, &card{
				number: i,
				image:  im,
			})
		}
	}
	return &deck{
		cards: cards,
	}
}

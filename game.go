package main

import (
	"bytes"
	"fmt"
	eb "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"log"
)

const (
	screenW = 640
	screenH = 480
	scale   = 0.5
)

type Game struct {
	deck    *deck
	players players
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *eb.Image) {
	ebitenutil.DebugPrint(screen, "Hellow, World!")
	for i, v := range g.players.players[0].cards {
		op := &eb.DrawImageOptions{}
		op.GeoM.Scale(scale, scale)
		op.GeoM.Translate(float64(i*cardWidth)*scale+screenW/2-cardWidth*scale, screenH-cardHeight*scale)

		screen.DrawImage(v.image, op)

	}
	//for i, v := range g.deck.cards[7:] {
	//	op := &eb.DrawImageOptions{}
	//	op.GeoM.Translate(float64(i*2*200), 0)
	//	op.GeoM.Scale(0.5, 0.5)
	//	screen.DrawImage(v.image, op)
	//}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func newGame() *Game {
	img, _, err := image.Decode(bytes.NewReader(deckImage))
	if err != nil {
		log.Fatal(err)
	}

	d := newDeck(eb.NewImageFromImage(img))
	d.Shuffle()

	var ps []player
	for i := 0; i < 5; i++ {
		var player = player{
			ID:    fmt.Sprintf("%s", i),
			cards: []*card{d.getCard(), d.getCard()},
		}
		ps = append(ps, player)
	}

	return &Game{
		deck: d,
		players: players{
			players: ps,
		},
	}
}

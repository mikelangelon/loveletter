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

	for _, player := range g.players.players {
		for i, v := range player.cards {
			op := &eb.DrawImageOptions{}
			op.GeoM.Scale(player.location.scale, player.location.scale)
			op.GeoM.Translate(player.location.x+nextCard(i, player.location), player.location.y)

			screen.DrawImage(v.image, op)
		}
	}
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
	for i := 0; i < 4; i++ {
		var player = player{
			ID:       fmt.Sprintf("%s", i),
			cards:    []*card{d.getCard(), d.getCard()},
			location: getLocation(i, 3),
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

func getLocation(i, enemies int) location {
	switch i {
	case 0:
		return location{
			x:     float64(i*cardWidth)*scale + screenW/2 - cardWidth*scale,
			y:     screenH - cardHeight*scale,
			scale: scale,
		}
	case 1, 2, 3:
		fmt.Printf("%f", float64(i*screenW/3))
		//corrector := float64(i-1) * cardWidth * scale / 2
		return location{
			x:     float64(i*screenW/(enemies+1)) - cardWidth*scale/2,
			y:     0,
			scale: scale / 2,
		}
	}
	return location{}
}

func nextCard(i int, l location) float64 {
	return float64(i*cardWidth) * l.scale
}

package main

import (
	"bytes"
	_ "embed"
	eb "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	_ "image/png"
	"log"
)

//go:embed cards.png
var deckImage []byte

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *eb.Image) {
	ebitenutil.DebugPrint(screen, "Hellow, World!")
	for i, v := range d.cards[7:] {
		op := &eb.DrawImageOptions{}
		op.GeoM.Translate(float64(i*2*200), 0)
		op.GeoM.Scale(0.5, 0.5)
		screen.DrawImage(v.image, op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

var d *deck

func initGame() {
	img, _, err := image.Decode(bytes.NewReader(deckImage))
	if err != nil {
		log.Fatal(err)
	}

	d = newDeck(eb.NewImageFromImage(img))
	d.Shuffle()
}

func main() {
	eb.SetWindowSize(640, 480)
	eb.SetWindowTitle("Hello World!")
	initGame()
	if err := eb.RunGame(&Game{}); err != nil {
		panic(err)
	}
}

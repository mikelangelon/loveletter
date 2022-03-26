package main

import (
	_ "embed"
	eb "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *eb.Image) {
	ebitenutil.DebugPrint(screen, "Hellow, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	eb.SetWindowSize(640, 480)
	eb.SetWindowTitle("Hello World!")
	if err := eb.RunGame(&Game{}); err != nil {
		panic(err)
	}
}

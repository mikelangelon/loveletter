package main

import (
	_ "embed"
	eb "github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
)

//go:embed cards.png
var deckImage []byte

func main() {
	eb.SetWindowSize(screenW, screenH)
	eb.SetWindowTitle("Hello World!")

	if err := eb.RunGame(newGame()); err != nil {
		panic(err)
	}
}

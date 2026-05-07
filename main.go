package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const WIN_W = 800
const WIN_H = 300
const WIN_TITLE = "Digital Clock"

type ScreenState struct {
	currentBluePrint [6][7]bool
}

func (g *ScreenState) Update() error {
	g.currentBluePrint = getCurrentTimeAsBluePrint()
	return nil
}

func (g *ScreenState) Draw(screen *ebiten.Image) {
	DrawClock(screen, WIN_W, WIN_H, g.currentBluePrint)
}

func (g *ScreenState) Layout(outsideW, outsideH int) (int, int) {
	return WIN_W, WIN_H
}

func main() {
	ebiten.SetWindowSize(WIN_W, WIN_H)
	ebiten.SetWindowTitle(WIN_TITLE)

	if err := ebiten.RunGame(&ScreenState{}); err != nil {
		log.Fatal(err)
	}
}

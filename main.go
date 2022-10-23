package main

import (
	"app/internal/constant"
	"app/internal/game"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetTPS(constant.TPS)
	ebiten.SetWindowSize(constant.ScreenWidth, constant.ScreenHeight)
	ebiten.SetWindowTitle("Snake game")

	if err := ebiten.RunGame(&game.Game{}); err != nil {
		log.Fatal(err)
	}
}

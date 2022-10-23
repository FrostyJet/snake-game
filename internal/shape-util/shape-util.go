package shapeUtil

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func DrawGlow(image *ebiten.Image, x, y, scale float64, glowColor color.RGBA) {
	centerX := x + scale/2
	centerY := y + scale/2

	for i := 0; i < 5; i++ {
		glowColor.A += uint8(2)
		ebitenutil.DrawCircle(image, centerX, centerY, float64(int(scale*2)-i*5), glowColor)
	}

	glowColor.A = 75
	ebitenutil.DrawRect(image, x, y, scale, scale, glowColor)
}

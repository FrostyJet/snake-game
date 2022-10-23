package object

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Canvas struct {
	Width    int
	Height   int
	TileSize int
	Image    *ebiten.Image
}

func NewCanvas(w, h, tileSize int) *Canvas {
	c := &Canvas{
		Width:    w,
		Height:   h,
		TileSize: tileSize,
		Image:    ebiten.NewImage(w, h),
	}

	c.init()

	return c
}

func (c *Canvas) init() {
	// background
	bg := color.RGBA{33, 33, 33, 255}
	ebitenutil.DrawRect(c.Image, 0, 0, float64(c.Width), float64(c.Height), bg)

	// vertical lines on top of background
	lineColor := color.RGBA{38, 38, 38, 255}
	for x := 0; x < c.Width; x += c.TileSize {
		ebitenutil.DrawLine(c.Image, float64(x), 0, float64(x), float64(c.Height), lineColor)
	}

	// horizontal lines on top of background
	for y := 0; y < c.Height; y += c.TileSize {
		ebitenutil.DrawLine(c.Image, 0, float64(y), float64(c.Width), float64(y), lineColor)
	}
}

func (c *Canvas) Clear() {
	c.init()
}

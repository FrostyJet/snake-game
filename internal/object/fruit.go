package object

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	shapeUtil "app/internal/shape-util"
)

type Fruit struct {
	x     float64
	y     float64
	scale float64
}

func NewFruit(x, y, s float64) *Fruit {
	return &Fruit{
		x:     x,
		y:     y,
		scale: s,
	}
}

func (f *Fruit) GetX() float64 {
	return f.x
}

func (f *Fruit) GetY() float64 {
	return f.y
}

func (f *Fruit) SetPos(x, y float64) {
	f.y = y
	f.x = x
}

func (f *Fruit) Draw(image *ebiten.Image) {
	fruitColor := color.RGBA{5, 254, 8, 255}
	ebitenutil.DrawRect(image, f.x, f.y, f.scale, f.scale, fruitColor)

	shapeUtil.DrawGlow(image, f.x, f.y, f.scale, color.RGBA{255, 255, 255, 0})
}

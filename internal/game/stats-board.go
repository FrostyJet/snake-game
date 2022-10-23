package game

import (
	"app/internal/typography"
	"image/color"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type StatsBoard struct {
	Image *ebiten.Image
	data  *GameStats
}

func NewStatsBoard(w, h int, d *GameStats) *StatsBoard {
	s := &StatsBoard{
		Image: ebiten.NewImage(w, h),
		data:  d,
	}

	s.init()
	return s
}

func (s *StatsBoard) init() {
	padding := map[string]int{
		"left": 35,
		"top":  25,
	}
	w, h := s.Image.Size()

	// background
	ebitenutil.DrawRect(s.Image, 0, 0, float64(w), float64(h), color.RGBA{235, 235, 235, 255})

	tileImg, _, err := ebitenutil.NewImageFromFile("resources/images/tile.png")
	if err != nil {
		log.Fatal("images: could not open find tile image", err)
	}

	// draw Tile
	tileOps := &ebiten.DrawImageOptions{}
	tileOps.GeoM.Scale(0.4, 0.3)
	tileOps.GeoM.Translate(float64(padding["left"]), float64(padding["top"]))
	s.Image.DrawImage(tileImg, tileOps)

	font, err := typography.GetFront("Cocola", 24)
	if err != nil {
		log.Fatal(err)
	}

	title := "Your score"
	tBounds := text.BoundString(font, title)

	boundsX, boundsY := float64(-tBounds.Min.X), float64(-tBounds.Min.Y)
	boundsW, _ := float64(tBounds.Dx()), float64(tBounds.Dy())

	tileWidth := float64(tileImg.Bounds().Dx()) * 0.4
	tileLeftPadding := (tileWidth - boundsW) / 2

	op := &ebiten.DrawImageOptions{}

	tileStartX := boundsX + float64(padding["left"])
	tileStartY := boundsY + float64(padding["top"])

	op.GeoM.Translate(tileStartX+tileLeftPadding, tileStartY+20)
	op.ColorM.ChangeHSV(255, 0, 0)
	text.DrawWithOptions(s.Image, title, font, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(70, 60)
	op.ColorM.ChangeHSV(255, 0, 0)
	op.GeoM.Scale(2, 2)
	text.DrawWithOptions(s.Image, strconv.Itoa(s.data.Score), font, op)
}

func (s *StatsBoard) Update(data *GameStats) {
	s.data = data
}

func (s *StatsBoard) Clear() {
	s.init()
}

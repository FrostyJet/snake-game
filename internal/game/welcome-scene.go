package game

import (
	"app/internal/constant"
	"app/internal/sounds"
	"app/internal/typography"
	"image/color"
	_ "image/jpeg"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type WelcomeScene struct {
	image *ebiten.Image
}

func NewWelcomeScene() *WelcomeScene {
	return &WelcomeScene{
		image: ebiten.NewImage(constant.ScreenWidth, constant.ScreenHeight),
	}
}

func (w *WelcomeScene) Update(state *State) error {

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		sounds.GameStartPlayer.Rewind()
		sounds.GameStartPlayer.Play()

		state.SceneManager.GoTo(NewPlaygroundScene())
	}

	return nil
}

func (w *WelcomeScene) Draw(screen *ebiten.Image) {
	bg := w.getBackground(constant.ScreenWidth, constant.ScreenHeight)

	screen.DrawImage(&bg, &ebiten.DrawImageOptions{})
}

var (
	y  = 500
	dy = 2
)

func (w *WelcomeScene) getBackground(width, height int) ebiten.Image {
	img := ebiten.NewImage(width, height)
	ebitenutil.DrawRect(img, 0, 0, float64(width), float64(height), color.White)

	bgImg, _, err := ebitenutil.NewImageFromFile("resources/images/intro-bg.jpg")
	if err != nil {
		log.Fatal(err)
	}

	font, err := typography.GetFront("Cocola", 36)
	if err != nil {
		log.Fatal(err)
	}

	y += dy
	if y > 500 || y < 450 {
		dy *= -1
	}

	text.Draw(bgImg, "Press 'space' to start", font, 250, y, color.White)

	img.DrawImage(bgImg, &ebiten.DrawImageOptions{})

	return *img
}

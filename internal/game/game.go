package game

import (
	"app/internal/constant"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	SceneManager *SceneManager
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.SceneManager.Draw(screen)
}

func (g *Game) Update() error {
	if g.SceneManager == nil {
		g.SceneManager = NewSceneManager()
		g.SceneManager.GoTo(NewWelcomeScene())
	}

	err := g.SceneManager.Update()
	if err != nil {
		return err
	}

	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return constant.ScreenWidth, constant.ScreenHeight
}

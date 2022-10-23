package game

import (
	"app/internal/constant"
	"app/internal/object"
	"app/internal/sounds"
	mathUtil "app/pkg/math-util"

	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	canvasWidth  = 600
	canvasHeight = 600
	boxScale     = 20
)

var (
	s          *object.Snake
	f          *object.Fruit
	canvas     *object.Canvas
	statsBoard *StatsBoard
	gameStats  *GameStats
)

type PlaygroundScene struct {
	IsPaused bool
	stats    *GameStats
}

func _init() {
	// game status information
	gameStats = &GameStats{
		Score: 0,
	}

	// board object, where boundaries of game end
	canvas = object.NewCanvas(canvasWidth, canvasHeight, boxScale)

	// stats board, where score, time, etc. is written
	statsBoard = NewStatsBoard(300, 600, gameStats)

	// snake object
	s = object.NewSnake(100, 100, boxScale)
	s.BindToCanvas(canvas)

	// fruit object
	populateFruit()
}

func NewPlaygroundScene() *PlaygroundScene {
	_init()

	return &PlaygroundScene{
		stats: gameStats,
	}
}

func (p *PlaygroundScene) Update(state *State) error {
	if state.IsScenePaused {
		return nil
	}

	s.Update()
	statsBoard.Update(p.stats)

	p.processSnakeSelfCollision(state)
	return nil
}

func (p *PlaygroundScene) Draw(screen *ebiten.Image) {
	// player elements
	s.Draw(canvas.Image)
	f.Draw(canvas.Image)

	// Draw Canvas board
	screen.DrawImage(canvas.Image, &ebiten.DrawImageOptions{})
	canvas.Clear()

	// Draw stats board
	sbOp := &ebiten.DrawImageOptions{}
	sbOp.GeoM = ebiten.GeoM{}
	sbOp.GeoM.Translate(600, 0)
	screen.DrawImage(statsBoard.Image, sbOp)
	statsBoard.Clear()

	p.processSnakeToFruitCollision()
}

func (p *PlaygroundScene) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return constant.ScreenWidth, constant.ScreenHeight
}

func populateFruit() {
	centerX := math.Floor(float64(canvas.Width)/2) - boxScale
	centerY := math.Floor(float64(canvas.Height)/2) - boxScale

	f = object.NewFruit(centerX, centerY, boxScale)
}

func (p *PlaygroundScene) processSnakeToFruitCollision() {
	d := mathUtil.Distance(s.GetX(), s.GetY(), f.GetX(), f.GetY())

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	if d < boxScale {
		s.Grow()
		s.Boost()
		f.SetPos(float64(r1.Intn(canvas.Width)/boxScale)*boxScale, float64(r1.Intn(canvas.Height)/boxScale)*boxScale)

		sounds.BleepPlayer.Rewind()
		sounds.BleepPlayer.Play()

		p.stats.Score++
	}
}

func (p *PlaygroundScene) processSnakeSelfCollision(state *State) {
	if s.IsSelfColliding() {
		state.SceneManager.PauseScene()

		sounds.SnakeSelfBitePlayer.Rewind()
		sounds.SnakeSelfBitePlayer.Play()
	}
}

package object

import (
	"app/internal/constant"
	shapeUtil "app/internal/shape-util"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Snake struct {
	x      float64
	y      float64
	dX     int
	dY     int
	speed  int
	size   int
	scale  float64
	tail   []Point
	canvas *Canvas
}

func NewSnake(x, y, s float64) *Snake {
	return &Snake{
		x:     x,
		y:     y,
		dX:    1,
		dY:    0,
		speed: 12,
		size:  constant.SnakeInitialSize,
		scale: s,
	}
}

func (s *Snake) Draw(image *ebiten.Image) {
	bodyColor := color.RGBA{255, 255, 255, 255}
	dt := 125 / s.size

	for index, point := range s.tail {
		n := uint8(125 + dt*(s.size-index))

		bodyColor.R = n
		bodyColor.G = n
		bodyColor.B = n
		ebitenutil.DrawRect(image, point.x, point.y, s.scale-1, s.scale-1, bodyColor)

		// draw glow on head
		if index == 0 {
			shapeUtil.DrawGlow(image, point.x, point.y, s.scale, color.RGBA{252, 232, 3, 0})
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		s.setDir(0, 1)
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		s.setDir(1, 0)
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		s.setDir(0, -1)
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		s.setDir(-1, 0)
	}
}

var frameCount = 0

func (s *Snake) Update() {
	frameCount++

	if frameCount%s.speed == 0 {
		s.x += float64(s.dX) * s.scale
		s.y += float64(s.dY) * s.scale

		s.checkBoundaries()

		s.tail = append([]Point{{s.x, s.y}}, s.tail...)
		if len(s.tail) > s.size {
			s.tail = s.tail[:s.size]
		}
	}
}

func (s *Snake) IsSelfColliding() bool {
	for index, point := range s.tail {
		if index > 0 {
			if s.x == point.x && s.y == point.y {
				return true
			}
		}
	}

	return false
}

func (s *Snake) GetX() float64 {
	return s.x
}

func (s *Snake) GetY() float64 {
	return s.y
}

func (s *Snake) BindToCanvas(c *Canvas) {
	s.canvas = c
}

func (s *Snake) Grow() {
	s.size++
}

func (s *Snake) Boost() {
	if s.speed <= 5 {
		return
	}

	s.speed -= 3
}

func (s *Snake) setDir(dx, dy int) {
	if s.dX != -dx {
		s.dX = dx
	}

	if s.dY != -dy {
		s.dY = dy
	}
}

func (s *Snake) checkBoundaries() {
	if s.x >= float64(s.canvas.Width) {
		s.x = 0
	} else if s.x < 0 {
		s.x = float64(s.canvas.Width) - s.scale
	}

	if s.y >= float64(s.canvas.Height) {
		s.y = 0
	} else if s.y < 0 {
		s.y = float64(s.canvas.Height) - s.scale
	}
}

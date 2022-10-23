package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Update(state *State) error
	Draw(screen *ebiten.Image)
}

type SceneManager struct {
	isScenePaused   bool
	current         Scene
	next            Scene
	transitionCount int
}

func NewSceneManager() *SceneManager {
	return &SceneManager{
		current: nil,
		next:    nil,
	}
}

func (s *SceneManager) GoTo(scene Scene) {
	if s.current == nil {
		s.current = scene
	} else {
		s.next = scene
		s.transitionCount = 20
	}
}

func (s *SceneManager) PauseScene() {
	s.isScenePaused = true

	s.ResumeScene()
	s.GoTo(NewPlaygroundScene())
}

func (s *SceneManager) ResumeScene() {
	s.isScenePaused = false
}

func (s *SceneManager) Update() error {
	if s.transitionCount == 0 {
		return s.current.Update(&State{
			SceneManager:  s,
			IsScenePaused: s.isScenePaused,
		})
	}

	s.transitionCount--

	if s.transitionCount > 0 {
		return nil // wait for transition

	}

	// transitions just hit zero, time to switch scenes
	s.current = s.next
	s.next = nil
	return nil
}

func (s *SceneManager) Draw(screen *ebiten.Image) {
	if s.transitionCount == 0 {
		s.current.Draw(screen)
		return
	}
}

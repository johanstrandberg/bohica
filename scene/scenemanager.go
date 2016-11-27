package scene

import (
)

type Scene interface {
    Update(time string)

}

type SceneManager struct {
    currentScene *Scene
}

//Update updates the currentscene
func (s *SceneManager) Update(time string) {
    s.Update(time)
}
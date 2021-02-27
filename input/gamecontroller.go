package input

import (
	"github.com/go-pogo/errors"
	"github.com/veandco/go-sdl2/sdl"
)

func AnyGameController() (*sdl.GameController, error) {
	for i := 0; i < sdl.NumJoysticks(); i++ {
		if !sdl.IsGameController(i) {
			continue
		}

		return sdl.GameControllerOpen(i), nil
	}

	return nil, errors.New("input: no game controllers found")
}

package sdlkit

import (
	"github.com/veandco/go-sdl2/sdl"
	sdlttf "github.com/veandco/go-sdl2/ttf"
)

type FontsMap map[string]*sdlttf.Font

func OpenFontFromMem(data []byte, size int) (*sdlttf.Font, error) {
	src, err := sdl.RWFromMem(data)
	if err != nil {
		return nil, err
	}

	return sdlttf.OpenFontRW(src, 1, size)
}

func OpenFontIndexFromMem(data []byte, size, index int) (*sdlttf.Font, error) {
	src, err := sdl.RWFromMem(data)
	if err != nil {
		return nil, err
	}

	return sdlttf.OpenFontIndexRW(src, 1, size, index)
}

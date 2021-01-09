package sdlkit

import (
	"github.com/go-pogo/errors"
	"github.com/veandco/go-sdl2/sdl"
)

type TexturesMap struct {
	ren  *sdl.Renderer
	list map[string]*sdl.Texture
}

func NewTexturesMap(ren *sdl.Renderer, size uint) *TexturesMap {
	return &TexturesMap{
		ren:  ren,
		list: make(map[string]*sdl.Texture, size),
	}
}

func (t *TexturesMap) Find(name string) (*sdl.Texture, bool) {
	tx, ok := t.list[name]
	if !ok || tx == nil {
		return nil, false
	}

	return tx, true
}

func (t *TexturesMap) Get(name string) (*sdl.Texture, error) {
	tx, ok := t.list[name]
	if !ok || tx == nil {
		return nil, errors.Newf("sdlkit: texture %s does not exist", name)
	}

	return tx, nil
}

func (t *TexturesMap) Add(name string, tx *sdl.Texture) *sdl.Texture {
	t.list[name] = tx
	return tx
}

func (t *TexturesMap) Remove(name string, destroy bool) error {
	tx, err := t.Get(name)
	if err != nil {
		return errors.Trace(err)
	}

	t.list[name] = nil

	if destroy {
		return errors.Trace(tx.Destroy())
	}
	return nil
}

// Destroy destroys all sdl.Textures within the TexturesMap.
func (t *TexturesMap) Destroy() (err error) {
	for n, tx := range t.list {
		if tx == nil {
			continue
		}

		errors.Append(&err, tx.Destroy())
		t.list[n] = nil
	}
	return err
}

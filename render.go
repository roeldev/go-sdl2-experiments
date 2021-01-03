package sdlkit

import (
	"github.com/go-pogo/errors"
	"github.com/veandco/go-sdl2/sdl"
)

type TextureDisplay struct {
	Alpha     uint8
	BlendMode sdl.BlendMode
	Flip      sdl.RendererFlip
}

type RenderTarget interface {
	Renderer() *sdl.Renderer
	Clear()
	Draw(...Drawable)
	Err() error
}

type StageRenderTarget struct {
	stage *Stage
	err   error
}

func (s *StageRenderTarget) Stage() *Stage { return s.stage }

func (s *StageRenderTarget) Renderer() *sdl.Renderer { return s.stage.Renderer() }

func (s *StageRenderTarget) Err() error { return s.err }

func (s *StageRenderTarget) Clear() {
	s.err = nil

	errors.Append(&s.err,
		s.stage.renderer.SetDrawColor(
			s.stage.BgColor.R,
			s.stage.BgColor.G,
			s.stage.BgColor.B,
			0xFF,
		),
		s.stage.renderer.Clear(),
	)
}

func (s *StageRenderTarget) Draw(draw ...Drawable) {
	for _, d := range draw {
		errors.Append(&s.err, d.Draw(s.stage.renderer))
	}
}

func (s *StageRenderTarget) ClearAndDraw(draw ...Drawable) {
	s.Clear()
	s.Draw(draw...)
}

type TextureRenderTarget struct {
	renderer *sdl.Renderer
}

package draw

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/geom"
)

func DrawTexture(r *sdl.Renderer, tx *sdl.Texture, src, dst *sdl.Rect, tr geom.Transform, td sdlkit.TextureDisplay) error {
	// doAlpha, doBlendMode := td.Alpha < 255, td.BlendMode != sdl.BLENDMODE_NONE

	// var alpha uint8
	// if doAlpha {
	// 	alpha, _ = tx.GetAlphaMod()
	_ = tx.SetAlphaMod(td.Alpha)
	// }
	// var bm sdl.BlendMode
	// if doBlendMode {
	// bm, _ = tx.GetBlendMode()
	// _ = tx.SetBlendMode(td.BlendMode)
	// }

	err := r.CopyEx(tx, src, dst, float64(tr.Rotation), nil, td.Flip)

	// if doAlpha {
	// 	_ = tx.SetAlphaMod(alpha)
	// }
	// if doBlendMode {
	// 	_ = tx.SetBlendMode(bm)
	// }

	return err
}

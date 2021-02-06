package sdlkit

import (
	"github.com/veandco/go-sdl2/sdl"
)

func GetClosestDisplayModeRatio(displayIndex int, mode sdl.DisplayMode) (*sdl.DisplayMode, error) {
	dm, err := sdl.GetDisplayMode(displayIndex, 0)
	if err != nil {
		return nil, err
	}

	// ratio of largest display mode, assume this is also the screen's ratio
	ratio := displayModeRatio(dm.W, dm.H)

	closest, err := sdl.GetClosestDisplayMode(displayIndex, &mode, &dm)
	if err == nil && ratio == displayModeRatio(dm.W, dm.H) {
		return closest, nil
	}

	n, err := sdl.GetNumDisplayModes(displayIndex)
	if err != nil {
		if closest != nil {
			return closest, nil
		} else {
			return nil, err
		}
	}

	result := *closest
	for i := 0; i < n; i++ {
		dm, err = sdl.GetDisplayMode(displayIndex, i)
		if err != nil || ratio != displayModeRatio(dm.W, dm.H) {
			continue
		}
		if dm.W >= mode.W && dm.H >= mode.H {
			result = dm
		}
	}
	return &result, nil
}

func displayModeRatio(w, h int32) int32 {
	return int32((float32(w) / float32(h)) * 100)
}

package sdlkit

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/veandco/go-sdl2/sdl"
)

func TestTime_ConvTicks(t *testing.T) {
	gt := NewTime(60)
	assert.Equal(t, gt.startTime, gt.ConvTicks(gt.startTick))

	time.Sleep(time.Second * 3)
	assert.Equal(t, int64(0), time.Now().Sub(gt.ConvTicks(sdl.GetTicks())).Milliseconds())
}

package ui

import (
	"github.com/veandco/go-sdl2/sdl"
)

type KeyAction uint8

const (
	KANone KeyAction = iota
	KAMoveForward
	KAMovebackwards
	KATurnLeft
	KATurnRight
	KAStrafeLeft
	KAStrafeRight
	KAJump
	KADuck
	KAShootPrimary
	KAShootSecundary
)

type KeyInput struct {
	Code sdl.Keycode
	Mod  sdl.Keymod
}

type KeyMap map[KeyInput]KeyAction

type Keyboard struct {
}

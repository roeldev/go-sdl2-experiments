package input

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Action uint8

const (
	NoAction Action = iota
	ActionMoveForwards
	ActionMoveBackwards
	ActionTurnLeft
	ActionTurnRight
	ActionStrafeLeft
	ActionStrafeRight
	ActionJump
	ActionDuck
	ActionShootPrimary
	ActionShootSecundary
	ActionCustom
)

type KeyInput struct {
	Code sdl.Keycode
	Mod  sdl.Keymod
}

func KeyInputFromEvent(e *sdl.KeyboardEvent) KeyInput {
	return KeyInput{Code: e.Keysym.Sym, Mod: sdl.Keymod(e.Keysym.Mod)}
}

type KeyMap map[KeyInput]Action

type KeyboardState struct {
	keyMap  KeyMap
	actions map[Action]KeyActionState
}

func NewKeyboardState(keyMap KeyMap) *KeyboardState {
	return &KeyboardState{
		keyMap:  keyMap,
		actions: make(map[Action]KeyActionState, len(keyMap)),
	}
}

func (ks *KeyboardState) HandleKeyboardEvent(e *sdl.KeyboardEvent) error {
	input := KeyInputFromEvent(e)
	action := ks.keyMap[input]
	if action != NoAction {
		state := ks.actions[action]
		state.action = action
		state.Pressed = e.State == sdl.PRESSED
		state.Released = !state.Pressed
		ks.actions[action] = state
	}
	return nil
}

func (ks *KeyboardState) KeyActionState(action Action) KeyActionState {
	return ks.actions[action]
}

type KeyActionState struct {
	Pressed  bool
	Released bool

	action Action
}

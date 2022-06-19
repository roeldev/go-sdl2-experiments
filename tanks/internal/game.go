// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"io/fs"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/display"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/ecs"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/event"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/input"
	"github.com/roeldev/go-sdl2-experiments/tanks/internal/tank"
)

const (
	tankComponent ecs.ComponentTag = 1 << iota
	colliderComponent
)

type tanksGame struct {
	event.Manager

	stage  *sdlkit.Stage
	camera *sdlkit.Camera
	mouse  *input.MouseState

	assets  fs.ReadFileFS
	ground  *display.Tile
	world   *display.TileMap
	objects *sdlkit.TextureAtlas
	players []*tank.Tank
	ecs     *ecs.Manager
}

func NewGame(stage *sdlkit.Stage, assets fs.ReadFileFS) (sdlkit.Scene, error) {
	load := sdlkit.NewAssetsLoader(assets, stage.Renderer())
	objects, err := load.TextureAtlasXml("assets/onlyObjects_default.xml")
	if err != nil {
		return nil, err
	}

	terrain, err := load.UniformTextureAtlas("assets/terrainTiles_retina.png", 128, 128, 40)
	if err != nil {
		return nil, err
	}

	world, err := display.NewTileMap(terrain)
	if err != nil {
		return nil, err
	}

	game := &tanksGame{
		stage:   stage,
		camera:  sdlkit.NewCamera(0, 0, stage.Width(), stage.Height()),
		mouse:   input.NewMouseState(input.TrackMouseBtnLeft),
		assets:  assets,
		ground:  display.MustNewTile(terrain.GetFomIndex(20)),
		world:   world,
		objects: objects,
		ecs:     ecs.NewManager(),
	}
	game.ground.StretchMode = display.StretchTile
	game.ground.W = 10000
	game.ground.H = 3000
	game.ground.X = 5005
	game.ground.Y = 1505

	stage.Canvas().SetCamera(game.camera)
	game.MustRegisterHandler(
		stage,
		game,
		game.mouse,
	)
	return game, nil
}

func (game *tanksGame) SceneName() string { return "" }

func (game *tanksGame) Activate() (err error) {
	game.addPlayer(tank.Green, &tank.KeyboardMouse{})

	gc, err := input.AnyGameController()
	if err == nil {
		game.addPlayer(tank.Blue, &tank.GameController{
			Device: gc,
		})

		// demo.players[0].SetX(demo.players[0].GetX() + 100)
		// demo.players[1].SetX(demo.players[1].GetX() - 100)
		// demo.players[1].SetHeading(math.Pi)
	}

	return nil
}

func (game *tanksGame) HandleWindowSizeChangedEvent(e *sdl.WindowEvent) error {
	game.camera.Resize(e.Data1, e.Data2)
	return nil
}

func (game *tanksGame) HandleRenderEvent(_ *sdl.RenderEvent) error {
	// return demo.player.Prerender(demo.canvas)
	return nil
}

func (game *tanksGame) Update(dt float64) {
	// demo.player1.SetTurretRotation(demo.player1.TurretPosition().SubXY(demo.mouse).Angle()+math.Pi)

	for _, tc := range game.ecs.Components(tankComponent) {
		tc.(*tank.Tank).Update(dt)
	}
}

func (game *tanksGame) Render(_ *sdl.Renderer) error {
	canvas := game.stage.Canvas()
	game.camera.Follow(game.players[0])
	game.ground.Draw(canvas)
	// errors.Append(&err, game.world.Draw(ren))

	for _, tc := range game.ecs.Components(tankComponent) {
		tc.(*tank.Tank).Draw(canvas)
	}

	return canvas.Done()
}

func (game *tanksGame) Destroy() error {
	return game.objects.Destroy()
}

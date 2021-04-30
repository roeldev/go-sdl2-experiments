package internal

import (
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit-examples/tanks/internal/tank"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/geom"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/geom/align"
	"github.com/roeldev/go-sdl2-experiments/pkg/sdlkit/physics"
)

func (game *tanksGame) addPlayer(color tank.Color, control tank.Control) *tank.Tank {
	player := tank.NewPlayer(game.objects, color, control)
	align.XYInSdlRect(align.ToCenter, player, game.stage.Size())

	game.RegisterHandler(player, game.stage.Size())
	game.players = append(game.players, player)

	entity := game.ecs.Create(nil)
	entity.AddComponent(tankComponent, player)
	entity.AddComponent(colliderComponent, player.Collider)
	return player
}

func (game *tanksGame) addTree(x, y float64) {
	tree := physics.NewStaticBody(&geom.Circle{X: x, Y: y, Radius: 10}, nil)

	entity := game.ecs.Create(nil)
	entity.AddComponent(colliderComponent, tree.Collider)
}

func (game *tanksGame) addBarrel() {
}

func (game *tanksGame) addSandbag() {

}

func (game *tanksGame) addSandbags() {

}

func (game *tanksGame) addBarricade() {

}

func (game *tanksGame) addCrate() {

}

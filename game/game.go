package game

import (
	"fmt"
	"log"
	"time"

	"github.com/petros-an/server-test/common/utils"
	"github.com/petros-an/server-test/common/vector"
	"github.com/petros-an/server-test/game/character"
	"github.com/petros-an/server-test/game/config"
	"github.com/petros-an/server-test/game/player"
	"github.com/petros-an/server-test/game/projectile"
	"github.com/petros-an/server-test/game/world"
)

type GameState struct {
	Characters  []*character.Character
	Players     map[player.PlayerId]*player.Player
	Projectiles []*projectile.Projectile
}

type Game struct {
	State  GameState
	Input  chan GameStateInput
	Output chan GameState
}

func New() *Game {

	outputChannel := make(chan GameState, 100)
	inputChannel := make(chan GameStateInput, 100)

	return &Game{
		Input:  inputChannel,
		Output: outputChannel,
		State: GameState{
			Characters: []*character.Character{},
			Players:    map[player.PlayerId]*player.Player{},
		},
	}
}

func (g *Game) String() string {
	str := ""
	for _, c := range g.State.Characters {
		str += fmt.Sprintf("[x:%f, y:%f, vx: %f, vy: %f, id: %s],", c.RigidBody.LocalPosition.X, c.RigidBody.LocalPosition.Y, c.RigidBody.Velocity.X, c.RigidBody.Velocity.Y, c.Tag)
	}
	return str
}

func (g *Game) Run() {

	log.Println("Starting game run")

	outputTicker := time.NewTicker(time.Duration(int64(config.SendTickerSeconds*1000)) * time.Millisecond)
	gameLoopTicker := time.NewTicker(time.Duration(int64(config.DT*1000)) * time.Millisecond)
	evictorTicker := time.NewTicker(time.Second)

	go func() {
		for {
			<-g.Output
		}
	}()

	for {
		select {
		case <-outputTicker.C:
			g.Output <- g.State
		case input := <-g.Input:
			input.UpdateState(&g.State)
		case <-evictorTicker.C:
			removeInactivePlayers(&g.State)
		case <-gameLoopTicker.C:
			applyGameLoopUpdate(&g.State)
		}
	}

}

func (g *Game) AddPlayer(playerId player.PlayerId) {
	g.Input <- NewPlayerUpdate{
		PlayerId: playerId,
	}
}

func (g *Game) UpdateCharacterDirection(playerId player.PlayerId, direction vector.Vector2D) {
	g.Input <- DirectionUpdate{
		PlayerId:  playerId,
		Direction: direction,
	}
}

func (g *Game) FireProjectile(playerId player.PlayerId, direction vector.Vector2D) {

	player, exists := g.GetPlayer(playerId)
	if !exists {
		return
	}

	g.Input <- ProjectileFiredUpdate{
		Position:  player.Character.Position(),
		Direction: direction,
		FiredBy:   player,
	}
}

func (g *Game) ReadState() chan GameState {
	return g.Output
}

func (g *Game) GetPlayer(playerId player.PlayerId) (*player.Player, bool) {
	player, exists := g.State.Players[playerId]
	return player, exists
}

func applyVitalsUpdate(state *GameState, id player.PlayerId) {
	if _, exists := state.Players[id]; exists {
		state.Players[id].RefreshVitals()
	}
}

func applyGameLoopUpdate(state *GameState) {
	for _, c := range state.Characters {
		c.Update(config.DT)
		c.SetPosition(
			world.RestrictPositionWithinBorder(c.Position()),
		)
	}

	for _, p := range state.Projectiles {
		p.Update(config.DT)
	}

	/*
		TODO: collision check
	*/
}

func removeInactivePlayers(state *GameState) {
	for id, p := range state.Players {
		if time.Since(p.LastVital) > config.EVICTION_INTERVAL {
			utils.RemoveElementFromSlice(&state.Characters, p.Character)
			delete(state.Players, id)
			log.Println("kicking due to connection timeout: " + id)
		}
	}
}

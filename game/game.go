package game

import (
	"fmt"
	"log"
	"time"

	"github.com/petros-an/server-test/common/color"
	"github.com/petros-an/server-test/common/vector"
	"github.com/petros-an/server-test/game/config"
	"github.com/petros-an/server-test/game/player"
	"github.com/petros-an/server-test/game/state"
	"github.com/petros-an/server-test/game/world"
)

type Game struct {
	State  state.GameState
	Input  chan GameStateInput
	Output chan GameStateOutput
}

type Character struct {
	vector.PSR2D
	Tag    string
	Color  color.RGBColor
	Health float64
}
type Projectile struct {
	vector.PSR2D
	Color color.RGBColor
}

type GameStateOutput struct {
	Characters  []Character
	Projectiles []Projectile
}

func StateToOutput(gameState state.GameState) GameStateOutput {
	outState := GameStateOutput{
		Characters:  make([]Character, len(gameState.Characters)),
		Projectiles: make([]Projectile, len(gameState.Projectiles)),
	}
	for i, character := range gameState.Characters {
		outState.Characters[i] = Character{
			PSR2D:  character.RigidBody.PSR2D,
			Tag:    character.Tag,
			Color:  character.Color,
			Health: character.Health,
		}
	}
	for i, projectile := range gameState.Projectiles {
		outState.Projectiles[i] = Projectile{
			PSR2D: projectile.RigidBody.PSR2D,
			Color: projectile.Color,
		}
	}
	return outState
}

func New() *Game {

	outputChannel := make(chan GameStateOutput, 100)
	inputChannel := make(chan GameStateInput, 100)

	return &Game{
		Input:  inputChannel,
		Output: outputChannel,
		State:  state.New(),
	}
}

func (g *Game) String() string {
	str := ""
	for _, c := range g.State.Characters {
		str += fmt.Sprintf("[x:%f, y:%f, vx: %f, vy: %f, id: %s],", c.RigidBody.Position.X, c.RigidBody.Position.Y, c.RigidBody.Velocity.X, c.RigidBody.Velocity.Y, c.Tag)
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
			g.Output <- StateToOutput(g.State)
		case input := <-g.Input:
			input.UpdateState(&g.State)
			g.State.RefreshPlayerVitals(input.GetPlayerId())
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

func (g *Game) UpdateCharacterMoveDirection(playerId player.PlayerId, direction vector.Vector2D) {

	_, exists := g.GetPlayer(playerId)
	if !exists {
		return
	}

	g.Input <- CharacterMoveDirectionUpdate{
		PlayerId:  playerId,
		Direction: direction,
	}
}

func (g *Game) UpdateCharacterRotation(playerId player.PlayerId, target vector.Vector2D) {

	player, exists := g.GetPlayer(playerId)
	if !exists {
		return
	}

	g.Input <- CharacterRotationUpdate{
		PlayerId:  playerId,
		Direction: target.Sub(player.Character.Position()).Normalized(),
	}
}

func (g *Game) FireProjectile(playerId player.PlayerId, target vector.Vector2D) {

	player, exists := g.GetPlayer(playerId)
	if !exists {
		return
	}
	g.Input <- ProjectileFiredUpdate{
		Position:  player.Character.Position(),
		Direction: target.Sub(player.Character.Position()).Normalized(),
		FiredBy:   player,
	}
}

func (g *Game) ReadStateOutpu() chan GameStateOutput {
	return g.Output
}

func (g *Game) GetPlayer(playerId player.PlayerId) (*player.Player, bool) {
	return g.State.GetPlayer(playerId)
}

func applyVitalsUpdate(state *state.GameState, id player.PlayerId) {
	state.RefreshPlayerVitals(id)
}

func applyGameLoopUpdate(s *state.GameState) {
	s.Update()
}

func removeInactivePlayers(state *state.GameState) {
	for id, p := range state.Players {
		if time.Since(p.LastVital) > config.EVICTION_INTERVAL {
			state.RemovePlayer(p.PlayerId)
			log.Println("kicking due to connection timeout: " + id)
		}
	}
}

func (g *Game) GetScores() state.Scores {
	return g.State.GetScores()
}

func (g *Game) GetWorldBorders() world.Borders {
	return world.WorldBorders
}

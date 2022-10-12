package state

import (
	"github.com/petros-an/server-test/common/color"
	"github.com/petros-an/server-test/common/vector"
	"github.com/petros-an/server-test/game/state"
)

type GameStateSchema struct {
	Characters  []CharacterSchema
	Projectiles []ProjectileSchema
}

type CharacterSchema struct {
	vector.PSR2D
	Tag    string
	Color  color.RGBColor
	Health float64
}
type ProjectileSchema struct {
	vector.PSR2D
	Color color.RGBColor
}

func StateToOutput(gameState state.GameState) GameStateSchema {
	outState := GameStateSchema{
		Characters:  make([]CharacterSchema, len(gameState.Characters)),
		Projectiles: make([]ProjectileSchema, len(gameState.Projectiles)),
	}
	for i, character := range gameState.Characters {
		outState.Characters[i] = CharacterSchema{
			PSR2D:  character.RigidBody.PSR2D,
			Tag:    character.Tag,
			Color:  character.Color,
			Health: character.Health,
		}
	}
	for i, projectile := range gameState.Projectiles {
		outState.Projectiles[i] = ProjectileSchema{
			PSR2D: projectile.RigidBody.PSR2D,
			Color: projectile.Color,
		}
	}
	return outState
}

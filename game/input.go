package game

import (
	"github.com/petros-an/server-test/common/vector"
	"github.com/petros-an/server-test/game/player"
	"github.com/petros-an/server-test/game/projectile"
)

type GameStateInput interface {
	UpdateState(*GameState)
	GetPlayerId() player.PlayerId
}

//

type DirectionUpdate struct {
	PlayerId  player.PlayerId
	Direction vector.Vector2D
}

func (u DirectionUpdate) GetPlayerId() player.PlayerId {
	return u.PlayerId
}

func (u DirectionUpdate) UpdateState(s *GameState) {
	if _, exists := s.Players[u.PlayerId]; !exists {
		return
	}

	s.Players[u.PlayerId].Character.MoveDirection = u.Direction
}

//

type NewPlayerUpdate struct {
	PlayerId player.PlayerId
}

func (u NewPlayerUpdate) GetPlayerId() player.PlayerId {
	return u.PlayerId
}

func (u NewPlayerUpdate) UpdateState(state *GameState) {
	playerId := u.PlayerId
	if _, exists := state.Players[playerId]; exists {
		return
	}
	newPlayer := player.New(playerId)
	state.Players[playerId] = newPlayer
	state.Characters = append(state.Characters, newPlayer.Character)
}

//

type ProjectileFiredUpdate struct {
	Position  vector.Vector2D
	Direction vector.Vector2D
	FiredBy   *player.Player
}

func (u ProjectileFiredUpdate) GetPlayerId() player.PlayerId {
	return u.FiredBy.PlayerId
}

func (u ProjectileFiredUpdate) UpdateState(state *GameState) {
	state.Projectiles = append(state.Projectiles, projectile.New(u.FiredBy.Character, u.Position, u.Direction))
}

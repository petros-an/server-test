package game

import (
	"github.com/petros-an/server-test/common/vector"
	"github.com/petros-an/server-test/game/player"
	"github.com/petros-an/server-test/game/state"
)

type GameStateInput interface {
	UpdateState(*state.GameState)
	GetPlayerId() player.PlayerId
}

//

type CharacterMoveDirectionUpdate struct {
	PlayerId  player.PlayerId
	Direction vector.Vector2D
}

func (u CharacterMoveDirectionUpdate) GetPlayerId() player.PlayerId {
	return u.PlayerId
}

func (u CharacterMoveDirectionUpdate) UpdateState(s *state.GameState) {
	s.UpdatePlayerMoveDirection(u.PlayerId, u.Direction)
}

//

type CharacterRotationUpdate struct {
	PlayerId  player.PlayerId
	Direction vector.Vector2D
}

func (u CharacterRotationUpdate) GetPlayerId() player.PlayerId {
	return u.PlayerId
}

func (u CharacterRotationUpdate) UpdateState(s *state.GameState) {
	if _, exists := s.Players[u.PlayerId]; !exists {
		return
	}

	s.Players[u.PlayerId].Character.RigidBody.SetRotation(u.Direction)
}

//

type NewPlayerUpdate struct {
	PlayerId player.PlayerId
}

func (u NewPlayerUpdate) GetPlayerId() player.PlayerId {
	return u.PlayerId
}

func (u NewPlayerUpdate) UpdateState(state *state.GameState) {
	state.AddPlayerIfNotExists(u.PlayerId)
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

func (u ProjectileFiredUpdate) UpdateState(s *state.GameState) {
	s.AddProjectile(u.FiredBy, u.Position, u.Direction)
}

package state

import (
	"github.com/petros-an/server-test/common/utils"
	"github.com/petros-an/server-test/common/vector"
	"github.com/petros-an/server-test/game/character"
	"github.com/petros-an/server-test/game/config"
	gameobject "github.com/petros-an/server-test/game/gameObject"
	"github.com/petros-an/server-test/game/player"
	"github.com/petros-an/server-test/game/projectile"
)

type GameState struct {
	Characters  []*character.Character
	Players     map[player.PlayerId]*player.Player
	Projectiles []*projectile.Projectile
}

func New() GameState {
	s := GameState{
		Characters:  []*character.Character{},
		Players:     map[player.PlayerId]*player.Player{},
		Projectiles: []*projectile.Projectile{},
	}
	return s
}

func (s *GameState) GetPlayer(playerId player.PlayerId) (*player.Player, bool) {
	p, exists := s.Players[playerId]
	return p, exists
}

func (s *GameState) GetGameObjects() []gameobject.GameObject {
	var res []gameobject.GameObject = []gameobject.GameObject{}
	for _, o := range append(s.Characters) {
		res = append(res, o)
	}
	for _, o := range s.Projectiles {
		res = append(res, o)
	}
	return res
}

func (s *GameState) RefreshVitals(playerId player.PlayerId) {
	if player, exists := s.GetPlayer(playerId); exists {
		player.RefreshVitals()
	}
}

func (s *GameState) RemovePlayer(playerId player.PlayerId) {
	if p, exists := s.GetPlayer(playerId); exists {
		utils.RemoveElementFromSlice(&s.Characters, p.Character)
		delete(s.Players, playerId)
	}
}

func (s *GameState) UpdatePlayerMoveDirection(playerId player.PlayerId, direction vector.Vector2D) {
	if p, exists := s.GetPlayer(playerId); exists {
		p.Character.SetMoveDirection(direction)
	}
}

func (s *GameState) AddPlayerIfNotExists(playerId player.PlayerId) {
	if _, exists := s.GetPlayer(playerId); exists {
		return
	}
	newPlayer := player.New(playerId)
	s.Players[playerId] = newPlayer
	s.Characters = append(s.Characters, newPlayer.Character)
}

func (s *GameState) AddProjectile(firedBy *player.Player, position vector.Vector2D, direction vector.Vector2D) {

	s.Projectiles = append(s.Projectiles, projectile.New(firedBy.Character, position, direction))
}

func (s *GameState) RemoveProjectile(proj *projectile.Projectile) {
	for _, p := range s.Projectiles {
		if p == proj {
			utils.RemoveElementFromSlice(&s.Projectiles, p)
			return
		}
	}
}

func (s *GameState) Update() {
	for _, obj := range s.GetGameObjects() {
		obj.Update(config.DT)
	}

	for _, p := range s.Projectiles {
		if p.IsOutsideWorld() {
			s.RemoveProjectile(p)
		}
	}

	for _, p := range s.Projectiles {
		for _, c := range s.Characters {
			if p.CollidesWithCharacter(c) && p.FiredBy != c {
				s.RemoveProjectile(p)
				c.GetDamaged(p.Damage)
				// log.Printf("Projectile %d hit character %s", p.Id, c.Tag)
				break
			}
		}
	}
}

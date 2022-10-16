package state

import (
	"github.com/petros-an/server-test/common/collider"
	"github.com/petros-an/server-test/common/utils"
	"github.com/petros-an/server-test/common/vector"
	"github.com/petros-an/server-test/game/character"
	"github.com/petros-an/server-test/game/config"
	gameobject "github.com/petros-an/server-test/game/gameObject"
	"github.com/petros-an/server-test/game/player"
	"github.com/petros-an/server-test/game/projectile"
	"github.com/petros-an/server-test/game/world"
)

type GameState struct {
	Characters  []*character.Character
	Players     map[player.PlayerId]*player.Player
	Projectiles []*projectile.Projectile
	Colliders   []collider.Collider2D
}

func New() GameState {
	s := GameState{
		Characters:  []*character.Character{},
		Players:     map[player.PlayerId]*player.Player{},
		Projectiles: []*projectile.Projectile{},
		Colliders:   []collider.Collider2D{},
	}
	return s
}

func (s *GameState) SetUp() {
	borders := world.SetUpWorldBorders()
	for _, border := range borders {
		s.AddCollider(border.Collider)
	}
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

func (s *GameState) RefreshPlayerVitals(playerId player.PlayerId) {
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
	s.AddCollider(newPlayer.Character.Collider)
}

func (s *GameState) AddProjectile(firedBy *player.Player, position vector.Vector2D, direction vector.Vector2D) {

	newProj := projectile.New(firedBy.Character, position, direction)
	s.Projectiles = append(s.Projectiles, newProj)
	s.AddCollider(newProj.Collider)
}

func (s *GameState) AddCollider(collider collider.Collider2D) {
	s.Colliders = append(s.Colliders, collider)
}

func (s *GameState) RemoveProjectile(proj *projectile.Projectile) {
	s.RemoveCollider(proj.Collider)
	for _, p := range s.Projectiles {
		if p == proj {
			utils.RemoveElementFromSlice(&s.Projectiles, p)
			return
		}
	}
}
func (s *GameState) RemoveCollider(coll collider.Collider2D) {
	for i, c := range s.Colliders {
		if c == coll {
			utils.RemoveElementFromSliceAtIndex(&s.Colliders, i)
			return
		}
	}
}

func (s *GameState) Update() {
	gameobjects := s.GetGameObjects()
	for _, obj := range gameobjects {
		obj.Update(config.DT)
	}

	collider.CheckCollidersCollisions(&s.Colliders)

	// delete ToDestroy objects
	len := len(gameobjects)
	for i := len - 1; i >= 0; i-- {
		obj := gameobjects[i]
		if obj.ToDestroy() {
			switch obj.GetType() {
			case gameobject.Projectile:
				s.RemoveProjectile(obj.(*projectile.Projectile))
				break
			}
		}
	}

}

func (s *GameState) GetPlayerFromCharacter(c *character.Character) *player.Player {
	for _, p := range s.Players {
		if p.Character == c {
			return p
		}
	}
	return nil
}

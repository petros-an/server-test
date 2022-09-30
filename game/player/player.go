package player

import (
	"time"

	"github.com/petros-an/server-test/game/character"
)

type PlayerId string

type Player struct {
	Character *character.Character
	LastVital time.Time
	PlayerId
}

func (p *Player) RefreshVitals() {
	p.LastVital = time.Now()
}

func New(id PlayerId) *Player {
	newCharacter := character.RandomNew()
	newCharacter.Tag = string(id)
	return &Player{
		Character: newCharacter,
		LastVital: time.Now(),
		PlayerId:  id,
	}
}

package state

import "github.com/petros-an/server-test/game/player"

type PlayerScore struct {
	KillCount uint
}

type Scores map[player.PlayerId]PlayerScore

func (s *GameState) GetScores() Scores {
	res := Scores{}
	for _, p := range s.Players {
		res[p.PlayerId] = PlayerScore{
			KillCount: p.KillCount(),
		}
	}
	return res
}

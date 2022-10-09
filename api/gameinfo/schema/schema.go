package schema

import (
	"github.com/petros-an/server-test/game/player"
	"github.com/petros-an/server-test/game/state"
	"github.com/petros-an/server-test/game/world"
)

type BordersSchema struct {
	Xmin float64
	Ymin float64
	Xmax float64
	Ymax float64
}

type PlayerScoreSchema struct {
	KillCount uint
}

type GameInfoSchema struct {
	Scores  map[player.PlayerId]PlayerScoreSchema
	Borders BordersSchema
}

func Build(
	scores state.Scores,
	borders world.Borders,
) GameInfoSchema {
	res := GameInfoSchema{Scores: map[player.PlayerId]PlayerScoreSchema{}}
	for playerTag, playerScores := range scores {
		res.Scores[playerTag] = PlayerScoreSchema{KillCount: playerScores.KillCount}
	}
	res.Borders = BordersSchema{
		Xmin: borders.Xmin,
		Xmax: borders.Xmax,
		Ymin: borders.Ymin,
		Ymax: borders.Ymax,
	}
	// log.Println(res)
	return res
}

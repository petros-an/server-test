package state

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	api_common "github.com/petros-an/server-test/api/common"
	"github.com/petros-an/server-test/common/utils"
	"github.com/petros-an/server-test/common/vector"
	"github.com/petros-an/server-test/game"
	"github.com/petros-an/server-test/game/player"
)

func GetEndpoint(
	g *game.Game,
) api_common.Endpoint {
	endpoint := func(w http.ResponseWriter, r *http.Request) {

		if conn, err := api_common.Upgrade(w, r); err != nil {
			log.Print("upgrade:", err)
			return
		} else {
			defer conn.Close()

			log.Println("Starting connetion")

			playerId := readNewPlayerFromConnection(g, conn)

			log.Println(playerId)

			go readInputsFromConnection(g, conn, playerId)

			sendStateToConnection(g, conn)
		}

	}

	return endpoint
}

func sendStateToConnection(g *game.Game, conn *websocket.Conn) {

	for newState := range g.ReadState() {

		message := utils.SerializeJson(newState)
		err := conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("write:", err)
			return
		}
		// log.Println(string(message))
	}

}

func readNewPlayerFromConnection(g *game.Game, conn *websocket.Conn) player.PlayerId {
	_, message, err := conn.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return ""
	}

	playerId := player.PlayerId(message)

	g.AddPlayer(playerId)

	return playerId

}

func readInputsFromConnection(g *game.Game, conn *websocket.Conn, charId player.PlayerId) {
	for {
		if _, data, err := conn.ReadMessage(); err != nil {
			log.Println("read:", err)
			return
		} else {
			var parsedInput PlayerInputSchema
			err = utils.ParseJson(data, &parsedInput)
			if err != nil {
				log.Println(err)
				continue
			}
			PropagatePlayerInput(parsedInput, g, charId)
		}
	}
}

func PropagatePlayerInput(input PlayerInputSchema, g *game.Game, playerId player.PlayerId) {
	// log.Println(PlayerInputSchema)
	if input.MoveDirection != nil {
		g.UpdateCharacterMoveDirection(
			playerId,
			vector.New(
				input.MoveDirection.X, input.MoveDirection.Y,
			),
		)
	}

	if input.CharacterRotation != nil {
		g.UpdateCharacterRotation(
			playerId,
			vector.New(
				input.CharacterRotation.X, input.CharacterRotation.Y,
			),
		)
	}

	if input.ProjectileInput != nil {
		g.FireProjectile(
			playerId,
			vector.New(
				input.ProjectileInput.X, input.ProjectileInput.Y,
			),
		)
	}
}

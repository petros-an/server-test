package scores

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/petros-an/server-test/api/common"
	"github.com/petros-an/server-test/common/utils"
	"github.com/petros-an/server-test/game"
)

func GetEndpoint(
	g *game.Game,
) common.Endpoint {
	endpoint := func(w http.ResponseWriter, r *http.Request) {

		if conn, err := common.Upgrade(w, r); err != nil {
			log.Print("upgrade:", err)
			return
		} else {
			defer conn.Close()

			for {
				scores := g.GetScores()
				str := utils.SerializeJson(scores)
				err = conn.WriteMessage(websocket.TextMessage, str)
				if err != nil {
					log.Println("error during scores sending:", err)
					return
				}
				time.Sleep(time.Second)
			}

		}

	}

	return endpoint
}

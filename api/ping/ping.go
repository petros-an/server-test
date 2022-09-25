package ping

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/petros-an/server-test/api/common"
	"github.com/petros-an/server-test/common/utils"
)

func Endpoint(w http.ResponseWriter, r *http.Request) {
	conn, err := common.Upgrade(w, r)
	if err != nil {
		log.Print("(ping)upgrade:", err)
		return
	}
	defer conn.Close()

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		message := utils.SerializeJson(map[string]interface{}{})
		err = conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("error during ping sending:", err)
			return
		}
	}

}

package route

import (
	"github.com/gorilla/websocket"
	log "unknwon.dev/clog/v2"

	"cli2ws/internal/cmd"
	"cli2ws/internal/context"
)

var upgrader = websocket.Upgrader{}

func HandleWs(ctx context.Context) {
	ws, err := upgrader.Upgrade(ctx.ResponseWriter(), ctx.Request().Request, nil)
	if err != nil {
		log.Error("upgrade:", err)
		return
	}
	command := ctx.Query("cmd")
	log.Trace("command: %v", command)

	err = cmd.Execute(command, ws)
	if err != nil {
		log.Error("Failed to execute command: %v", err)
		ctx.ServerError()
	}
	defer func() { _ = ws.Close() }()

}

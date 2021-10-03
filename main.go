package main

import (
	"github.com/flamego/flamego"
	log "unknwon.dev/clog/v2"

	"cli2ws/internal/context"
	"cli2ws/internal/route"
)

func main() {
	// Use clog in project.
	defer log.Stop()
	err := log.NewConsole()
	if err != nil {
		panic("failed to init console logger: " + err.Error())
	}

	f := flamego.Classic()
	f.Use(context.Contexter())
	f.Any("/ws", route.HandleWs)
	f.Run()
}

package main

import (
	"github.com/cockroachdb/errors"
	log "unknwon.dev/clog/v2"
)

func main() {
	// Use clog in project.
	defer log.Stop()
	err := log.NewConsole()
	if err != nil {
		panic("failed to init console logger: " + err.Error())
	}

	log.Info("Hello World")
}

func err() error {
	return errors.New("Use cockroachdb/errors as error package")
}

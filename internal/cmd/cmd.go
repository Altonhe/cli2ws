package cmd

import (
	"bufio"
	"os/exec"
	"strings"
	"time"

	"github.com/creack/pty"
	log "unknwon.dev/clog/v2"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second
)

func Execute(cmd string, ws *websocket.Conn) error {
	arr := strings.Split(cmd, " ")
	c := exec.Command(arr[0], arr[1:]...)
	f, err := pty.Start(c)
	if err != nil {
		return err
	}

	outScanner := bufio.NewScanner(f)
	for outScanner.Scan() {
		err := ws.SetWriteDeadline(time.Now().Add(writeWait))
		if err != nil {
			log.Error("Failed to set write deadline: %v", err)

		}
		if err := ws.WriteMessage(websocket.BinaryMessage, outScanner.Bytes()); err != nil {
			log.Error("Failed to write message: %v", err)
			ws.Close()
			break
		}
	}
	c.Wait()

	return nil
}

package websocket_test

import (
	"testing"
	"time"

	"github.com/zjxpcyc/websocket"
)

func TestSend(t *testing.T) {
	h := websocket.NewHub()

	go func() {
		time.Sleep(2 * time.Second)

		data := websocket.Message{
			Data: []byte(`this string sended by test case`),
		}

		h.Send(data)
	}()

	go h.Run()

	time.Sleep(5 * time.Second)
}

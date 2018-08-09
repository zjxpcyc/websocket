package main

// 主要用来测试
import (
	"fmt"
	"log"
	"net/http"

	"github.com/zjxpcyc/websocket"
)

var c = make(chan int)

func main() {
	websocket.Run()

	// http server
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		cli, err := websocket.NewClient(websocket.ClientID{"ID-1", nil}, w, r, websocket.GHub)
		if err != nil {
			panic(err)
		}

		cli.Run()
		c <- 1
	})

	go func() {
		for {
			select {
			case <-c:
				websocket.Send(websocket.Message{
					To: websocket.ClientID{
						ID: "ID-1",
					},
					Data: []string{"This", "slice", "come", "from", "server"},
				})
			}
		}
	}()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	fmt.Println("finished ...")
}

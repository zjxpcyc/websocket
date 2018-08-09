package websocket

// GHub 默认提供一个 Hub
var GHub = NewHub()

// Run 执行默认 HUB
func Run() {
	go GHub.Run()
}

// Send 发送内容
func Send(message Message) {
	GHub.Send(message)
}

package websocket

// Message 消息
type Message struct {
	From ClientID    `json:"from"`
	To   ClientID    `json:"to"`
	Data interface{} `json:"data"`
}

// JSONMessage message json 格式
type JSONMessage []byte

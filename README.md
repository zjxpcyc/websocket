# websocket

Websocket 库使用 [gorilla/websocket](https://github.com/gorilla/websocket)

对 [官方案例](https://github.com/gorilla/websocket/tree/master/examples/chat) 进行了部分改造，以适应我们的业务场景

## Install
```bash
go get github.com/zjxpcyc/websocket
```


## Useage

> 基础使用可以参考 test 目录



**中控端定义**
```golang
// 1、初始化 Hub
hub := websocket.NewHub()

// 2、启动 Hub
go hub.Run()

// 另外, 可以使用系统默认提供的 GHub 来代替 new Hub 操作
// 以上两个步骤, 可以使用以下方式代替
websocket.Run()  // 前面没有 go 关键字
```


**初始化客户端**
```golang
// 客户端的初始化需要 http 的 Response 与 Request 对象
func NewClient(id ClientID, w http.ResponseWriter, r *http.Request, hub *Hub) (*Client, error)

// 初始化完成之后, 需要注册 client
client.Run()
```


**消息体**
```golang

// 消息体需要符合 websocket.Message 的定义
// Message 消息
type Message struct {
	From ClientID    `json:"from"`
	To   ClientID    `json:"to"`
	Data interface{} `json:"data"`
}

// ClientID 客户端ID
type ClientID struct {
	ID     string   `json:"id"`
	Groups []string `json:"groups"`
}
```

**发送消息**
```golang
// 1、如果是自定义的 Hub
hub.Send(message)

// 2、也可以使用系统内置的 Send
websocket.Send(message)
```
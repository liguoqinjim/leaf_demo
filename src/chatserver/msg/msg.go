package msg

import (
	"github.com/name5566/leaf/network/json"
	"github.com/name5566/leaf/network/protobuf"
)

var (
	JSONProcessor     = json.NewProcessor()
	ProtobufProcessor = protobuf.NewProcessor()
)

func init() {
	JSONProcessor.Register(&C2S_AddUser{})
	JSONProcessor.Register(&C2S_Message{})
	JSONProcessor.Register(&C2S_Action{})

	JSONProcessor.Register(&S2C_Login{})
	JSONProcessor.Register(&S2C_Joined{})
	JSONProcessor.Register(&S2C_Left{})
	JSONProcessor.Register(&S2C_Typing{})
	JSONProcessor.Register(&S2C_StopTyping{})
	JSONProcessor.Register(&S2C_Message{})
}

//客户端发送给服务器
type C2S_AddUser struct {
	UserName string
}

type C2S_Message struct {
	Message string
}

type C2S_Action struct {
	Op string
}

//服务器传给客户端
type S2C_Login struct {
	NumUsers int
}

type S2C_Joined struct {
	NumUsers int
	UserName string
}

type S2C_Left struct {
	NumUsers int
	UserName string
}

type S2C_Typing struct {
	UserName string
}

type S2C_StopTyping struct {
	UserName string
}

type S2C_Message struct {
	UserName string
	Message  string
}

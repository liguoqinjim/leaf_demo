package internal

import (
	"chatserver/conf"
	"chatserver/game"
	"chatserver/msg"
	"github.com/name5566/leaf/gate"
	"log"
)

type Module struct {
	*gate.Gate
}

func (m *Module) OnInit() {
	m.Gate = &gate.Gate{
		MaxConnNum:      conf.Server.MaxConnNum,
		PendingWriteNum: conf.PendingWriteNum,
		MaxMsgLen:       conf.MaxMsgLen,
		WSAddr:          conf.Server.WSAddr,
		HTTPTimeout:     conf.HTTPTimeout,
		TCPAddr:         conf.Server.TCPAddr,
		LenMsgLen:       conf.LenMsgLen,
		LittleEndian:    conf.LittleEndian,
		AgentChanRPC:    game.ChanRPC,
	}

	switch conf.Encoding {
	case "json":
		m.Gate.Processor = msg.JSONProcessor
	case "protobuf":
		m.Gate.Processor = msg.ProtobufProcessor
	default:
		log.Fatalf("unknown encoding:%v", conf.Encoding)
	}
}

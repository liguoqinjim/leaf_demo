package internal

import (
	"fmt"
	"github.com/name5566/leaf/module"
	"github.com/name5566/leaf/recordfile"
	"log"
	"math/rand"
	"chatserver/base"
	"chatserver/msg"
	"time"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton

	readRobots()
	robotMessage()
}

func (m *Module) OnDestroy() {

}

type Robot struct {
	Name    string
	Message string
}

var robots *recordfile.RecordFile

func readRobots() {
	rf, err := recordfile.New(Robot{})
	if err != nil {
		log.Fatalf("recordfile New error:%v", err)
	}
	err = rf.Read("../../bin/leaf_demo/gamedata/robots.txt")
	if err != nil {
		log.Fatalf("rf read error:%v", err)
	}

	robots = rf

	//每小时读一次robots.txt
	skeleton.AfterFunc(time.Hour, readRobots)
}

func robotMessage() {
	if robots == nil || robots.NumRecord() == 0 {
		log.Fatalf("robots empty")
	}

	robot := robots.Record(rand.Intn(robots.NumRecord())).(*Robot)

	now := time.Now().In(loc)
	message := fmt.Sprintf("@%02d:%02d %s", now.Hour(), now.Minute(), robot.Message)

	pm := &messages[messageIndex]
	pm.userName = robot.Name
	pm.message = message
	messageIndex++
	if messageIndex >= maxMessages {
		messageIndex = 0
	}

	broadcastMsg(&msg.S2C_Message{
		UserName: robot.Name,
		Message:  message,
	}, nil)

	n := time.Duration(rand.Intn(11) + 5)
	skeleton.AfterFunc(n*time.Minute, robotMessage)
}

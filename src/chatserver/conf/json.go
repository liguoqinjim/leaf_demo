package conf

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var Server struct {
	LogLevel   string
	LogPath    string
	WSAddr     string
	TCPAddr    string
	MaxConnNum int
}

func init() {
	//读取运行路径
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd error:%v", err)
	}
	log.Println("当前运行路径:", dir)

	data, err := ioutil.ReadFile("../../bin/leaf_demo/conf/chatserver.json")
	if err != nil {
		log.Fatalf("readFile error:%v", err)
	}
	err = json.Unmarshal(data, &Server)
	if err != nil {
		log.Fatalf("json unmarshal error:%v", err)
	}
}

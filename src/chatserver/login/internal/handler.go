package internal

import "reflect"

func handleMsg(m, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {

}

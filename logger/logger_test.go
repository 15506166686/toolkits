package logger

import (
	"testing"
	"log"
)

func Test_output(t *testing.T) {
	InitLog("file", "debug", "gateway.log")
	Println("msg", "test")
	Infoln("msg", "test")
	Warnln("msg", "test")
	Errorln("msg", "test")
}

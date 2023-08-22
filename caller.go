package module

import (
	"github.com/aiteung/atmessage"
	"github.com/aiteung/module/handler"
	"github.com/aiteung/module/model"
)

func Caller(Modulename string, Pesan model.IteungMessage) (reply string) {
	switch Modulename {
	case "tokengrup":
		reply = handler.TokenGroup(Pesan)
	}
	return
}

func CallAndSend(Modulename string, Pesan model.IteungMessage, WAIface model.IteungWhatsMeowConfig) {
	reply := Caller(Modulename, Pesan)
	atmessage.SendMessage(reply, WAIface.Info.Sender, WAIface.Waclient)
}

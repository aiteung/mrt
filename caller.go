package module

import (
	"github.com/aiteung/atmessage"
	"github.com/aiteung/module/handler"
	"github.com/aiteung/module/model"
	"go.mau.fi/whatsmeow"
)

func Caller(Modulename string, Pesan model.IteungMessage) (reply string) {
	switch Modulename {
	case "tokengrup":
		reply = handler.TokenGroup(Pesan)
	}
	return
}

func CallAndSend(Modulename string, Pesan model.IteungMessage, WAIface model.IteungWhatsMeowConfig) (resp whatsmeow.SendResponse, err error) {
	reply := Caller(Modulename, Pesan)
	resp, err = atmessage.SendMessage(reply, WAIface.Info.Chat, WAIface.Waclient)
	return
}

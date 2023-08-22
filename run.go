package module

import (
	"os"

	"github.com/aiteung/module/model"
)

func Run(Pesan model.IteungMessage, DBIface model.IteungDBConfig) (Modulename string) {
	NormalizeAndTypoCorrection(&Pesan.Message, DBIface.MongoConn, DBIface.TypoCollection)
	if IsIteungCall(Pesan) {
		Modulename = GetModuleName(Pesan, DBIface.MongoConn, DBIface.ModuleCollection)
	}
	reply := Caller(Modulename, Pesan)
	var msg = model.GowaNotif{
		User:     Pesan.Chat_number,
		Server:   Pesan.Chat_server,
		Messages: reply,
	}
	var ApiWa string = os.Getenv("URLAPIWA")
	SendToGoWAAPI(msg, ApiWa)
	return
}

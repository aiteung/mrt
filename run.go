package module

import (
	"github.com/aiteung/module/model"
)

func Run(Pesan model.IteungMessage, DBIface model.IteungDBConfig, URLApiWa string) (Modulename string) {
	NormalizeAndTypoCorrection(&Pesan.Message, DBIface.MongoConn, DBIface.TypoCollection)
	if IsIteungCall(Pesan) {
		Modulename = GetModuleName(Pesan, DBIface.MongoConn, DBIface.ModuleCollection)
	}
	go CallAndSend(Modulename, Pesan, URLApiWa)
	return
}

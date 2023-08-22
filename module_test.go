package module

import (
	"fmt"
	"os"
	"testing"

	"github.com/aiteung/atdb"
	"github.com/aiteung/module/model"
)

var MongoInfo = atdb.DBInfo{
	DBString: os.Getenv("MONGOSTRING"),
	DBName:   "iteung",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

var Pesan = model.IteungMessage{
	Phone_number: "6281312000300",
	Chat_server:  "g.us",
	Group_name:   "NN257S",
	Group_id:     "1234566954",
	Message:      "teung minta token group",
	Group:        "",
	Alias_name:   "Rolly Maulana Awangga",
	Is_group:     true,
}

var NewModule = model.Module{
	Name:    "pomodoro",
	Keyword: []string{"pomodoro", "cek", "status"},
}
var ModuleCollection = "module"

var NewTypo = model.Typo{
	From: "grub",
	To:   "grup",
}

var TypoCollection = "typo"

func TestInsertCollection(t *testing.T) {
	//atdb.InsertOneDoc(MongoConn, ModuleCollection, NewModule)
	//atdb.InsertOneDoc(MongoConn, TypoCollection, NewTypo)

}

func TestModuleCall(t *testing.T) {
	var modulename string
	if modulename != "" {
		fmt.Println(modulename)
	} else {
		fmt.Println("inisiasi module empty string defaultnya")
	}
	NormalizeAndTypoCorrection(&Pesan.Message, MongoConn, TypoCollection)
	if IsIteungCall(Pesan) {
		modulename = GetModuleName(Pesan, MongoConn, ModuleCollection)
		if modulename != "" {
			fmt.Println(modulename)
		} else {
			fmt.Println("tidak ada module")
		}

	} else {
		fmt.Println("Pesan tidak memanggil iteung")
	}
	if modulename != "" {
		reply := Caller(modulename, Pesan)
		fmt.Println(reply)
	}
}

func TestAPI(t *testing.T) {
	var msg = model.GowaNotif{
		User:     "6281312000300",
		Server:   "s.whatsapp.net",
		Messages: "alo",
	}
	var ApiWa string = os.Getenv("URLAPIWA")
	resp := SendToGoWAAPI(msg, ApiWa)
	fmt.Println(resp)
}

func TestAPIBEV2(t *testing.T) {
	var urlv2iteung string = os.Getenv("URLITEUNGBEV2")
	resp := SendToIteungAPI(Pesan, urlv2iteung)
	fmt.Println(resp)
}

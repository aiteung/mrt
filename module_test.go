package module

import (
	"fmt"
	"os"
	"testing"

	"github.com/aiteung/atdb"
)

var MongoInfo = atdb.DBInfo{
	DBString: os.Getenv("MONGOSTRING"),
	DBName:   "iteung",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

var Pesan = IteungMessage{
	Phone_number: "6281312000300",
	Chat_server:  "g.us",
	Group_name:   "NN257S",
	Group_id:     "",
	Message:      "teung minta token group",
	Group:        "",
	Alias_name:   "Rolly Maulana Awangga",
	Is_group:     "true",
}

var NewModule = Module{
	Name:    "pomodoro",
	Keyword: []string{"pomodoro", "cek", "status"},
}
var ModuleCollection = "module"

var NewTypo = Typo{
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
		Caller(modulename)
	}

}

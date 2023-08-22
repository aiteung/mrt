package module

import (
	"os"
	"testing"

	"github.com/aiteung/atdb"
)

var MongoInfo = atdb.DBInfo{
	DBString: os.Getenv("MONGOSTRING"),
	DBName:   "hris",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

var Pesan = IteungMessage{
	Phone_number: "6281312000300",
	Chat_server:  "g.us",
	Group_name:   "NN257S",
	Group_id:     "",
	Messages:     "cuk",
	Group:        "",
	Alias_name:   "Rolly Maulana Awangga",
	Is_group:     "true",
}

var NewModule = Module{
	Name:    "pomodoro",
	Keyword: []string{"minta", "token", "grup"},
}

func TestInsertModule(t *testing.T) {

	atdb.InsertOneDoc(MongoConn, "module", NewModule)
}

package module

import (
	"fmt"
	"strings"

	"github.com/aiteung/atdb"
	"github.com/aiteung/musik"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"go.mongodb.org/mongo-driver/mongo"
)

func IteungModuleCall(Info *types.MessageInfo, Message *waProto.Message, waclient *whatsmeow.Client, MongoConn *mongo.Database, TypoCollection string, ModuleCollection string) (Pesan IteungMessage, modulename string) {
	Pesan = Whatsmeow2Struct(Info, Message, waclient)
	NormalizeAndTypoCorrection(&Pesan.Message, MongoConn, TypoCollection)
	if IsIteungCall(Pesan) {
		modulename = GetModuleName(Pesan, MongoConn, ModuleCollection)
	}
	return
}

func Whatsmeow2Struct(Info *types.MessageInfo, Message *waProto.Message, waclient *whatsmeow.Client) (im IteungMessage) {
	im.Phone_number = Info.Sender.User
	im.Chat_server = Info.Chat.Server
	im.Group_name = ""
	im.Alias_name = Info.PushName
	m := Message.GetConversation()
	im.Message = m
	im.Is_group = "false"
	im.Filename = ""
	im.Filedata = ""
	im.Latitude = 0.0
	im.Longitude = 0.0
	if Info.Chat.Server == "g.us" {
		groupInfo, err := waclient.GetGroupInfo(Info.Chat)
		fmt.Println("cek err : ", err)
		if groupInfo != nil {
			im.Group = groupInfo.GroupName.Name + "@" + Info.Chat.User
			im.Group_name = groupInfo.GroupName.Name
			im.Group_id = Info.Chat.User
		} else {
			fmt.Println("groupInfo : ", groupInfo)
		}
		im.Is_group = "true"
		if strings.Contains(Message.GetConversation(), "teung") || strings.Contains(Message.GetConversation(), "Teung") {
			go waclient.SendChatPresence(Info.Chat, "composing", "")
		}
	}
	return
}

func IsIteungCall(im IteungMessage) bool {
	if (strings.Contains(im.Message, "teung") && im.Chat_server == "g.us") || (im.Chat_server == "s.whatsapp.net") {
		return true
	} else {
		return false
	}
}

func GetModuleName(im IteungMessage, MongoConn *mongo.Database, ModuleCollection string) (modulename string) {
	modules := atdb.GetAllDoc[[]Module](MongoConn, ModuleCollection)
	for _, mod := range modules {
		complete, _ := musik.IsMatch(im.Message, mod.Keyword...)
		if complete {
			modulename = mod.Name
		}
	}
	return
}

package module

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/aiteung/musik"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
)

func Whatmeow2Struct(Info *types.MessageInfo, Message *waProto.Message, waclient *whatsmeow.Client) (im IteungMessage) {
	im.Phone_number = Info.Sender.User
	im.Group_name = Info.Sender.User
	im.Alias_name = Info.PushName
	im.Messages = Message.GetConversation()
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

func IsMultiKey(Info *types.MessageInfo, Message *waProto.Message, db *sql.DB) bool {
	m := musik.NormalizeString(Message.GetConversation())
	if (strings.Contains(m, "teung") && Info.Chat.Server == "g.us") || (Info.Chat.Server == "s.whatsapp.net") {
		complete, match := musik.IsMatch(m, "jadwal", "kuliah", "pertemuan", "jumlah", "ngajar")
		fmt.Println(complete)
		if match >= 2 && IsTerdaftar() {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func IsTerdaftar() bool {
	return true
}

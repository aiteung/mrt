package module

import (
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type PhoneList struct {
	Phones []string `json:"phones,omitempty"`
}

type Response struct {
	Response string `json:"response"`
}

type IteungV1Message struct {
	Phone_number string  `json:"phone_number"`
	Group_name   string  `json:"group_name"`
	Alias_name   string  `json:"alias_name"`
	Messages     string  `json:"messages"`
	Is_group     string  `json:"is_group"`
	Filename     string  `json:"filename"`
	Filedata     string  `json:"filedata"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Api_key      string  `json:"api_key"`
}

type IteungMessage struct {
	Phone_number string  `json:"phone_number,omitempty" bson:"phone_number,omitempty"`
	Chat_server  string  `json:"chat_server,omitempty" bson:"chat_server,omitempty"`
	Group_name   string  `json:"group_name,omitempty" bson:"group_name,omitempty"`
	Group_id     string  `json:"group_id,omitempty" bson:"group_id,omitempty"`
	Group        string  `json:"group,omitempty" bson:"group,omitempty"`
	Alias_name   string  `json:"alias_name,omitempty" bson:"alias_name,omitempty"`
	Message      string  `json:"messages,omitempty" bson:"messages,omitempty"`
	Is_group     string  `json:"is_group,omitempty" bson:"is_group,omitempty"`
	Filename     string  `json:"filename,omitempty" bson:"filename,omitempty"`
	Filedata     string  `json:"filedata,omitempty" bson:"filedata,omitempty"`
	Latitude     float64 `json:"latitude,omitempty" bson:"latitude,omitempty"`
	Longitude    float64 `json:"longitude,omitempty" bson:"longitude,omitempty"`
}

type Module struct {
	Name    string   `json:"name,omitempty" bson:"name,omitempty"`
	Keyword []string `json:"keyword,omitempty" bson:"keyword,omitempty"`
}

type Typo struct {
	From string `json:"from,omitempty" bson:"from,omitempty"`
	To   string `json:"to,omitempty" bson:"to,omitempty"`
}

type IteungWhatsMeowConfig struct {
	Info     *types.MessageInfo
	Message  *waProto.Message
	Waclient *whatsmeow.Client
}

type IteungDBConfig struct {
	MongoConn        *mongo.Database
	TypoCollection   string
	ModuleCollection string
}

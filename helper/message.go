package helper

import (
	"strings"

	"github.com/aiteung/atmessage/mediadecrypt"
	"github.com/aiteung/module/model"
	waProto "go.mau.fi/whatsmeow/binary/proto"
)

func GetPhoneNumber(WAIface model.IteungWhatsMeowConfig) (phonenumber string) {
	if WAIface.Message.ExtendedTextMessage != nil {
		if WAIface.Message.ExtendedTextMessage.ContextInfo != nil {
			if WAIface.Message.ExtendedTextMessage.ContextInfo.Participant != nil {
				phonenumber = strings.Split(*WAIface.Message.ExtendedTextMessage.ContextInfo.Participant, "@")[0]
			}
		}
	} else {
		phonenumber = WAIface.Info.Sender.User
	}

	return

}

func GetMessage(Message *waProto.Message) (message string) {
	if Message.ExtendedTextMessage != nil {
		message = *Message.ExtendedTextMessage.Text
	} else if Message.DocumentMessage != nil {
		if Message.DocumentMessage.Caption != nil {
			message = *Message.DocumentMessage.Caption
		}
	} else {
		message = Message.GetConversation()
	}
	return

}

func GetLongLat(Message *waProto.Message) (long, lat float64) {
	if Message.ExtendedTextMessage != nil {
		if Message.ExtendedTextMessage.ContextInfo != nil {
			if Message.ExtendedTextMessage.ContextInfo.Participant != nil {
				if Message.ExtendedTextMessage.ContextInfo.QuotedMessage.LiveLocationMessage != nil {
					lat = *Message.ExtendedTextMessage.ContextInfo.QuotedMessage.LiveLocationMessage.DegreesLatitude
					long = *Message.ExtendedTextMessage.ContextInfo.QuotedMessage.LiveLocationMessage.DegreesLongitude
				}
			}

		}
	} else if Message.LiveLocationMessage != nil {
		long = *Message.LiveLocationMessage.DegreesLongitude
		lat = *Message.LiveLocationMessage.DegreesLatitude
	}
	return
}

func GetFile(Message *waProto.Message) (filename, filedata string) {
	if Message.ExtendedTextMessage != nil {
		if Message.ExtendedTextMessage.ContextInfo != nil {
			if Message.ExtendedTextMessage.ContextInfo.Participant != nil {
				if Message.ExtendedTextMessage.ContextInfo.QuotedMessage.DocumentMessage != nil {
					filename = *Message.ExtendedTextMessage.ContextInfo.QuotedMessage.DocumentMessage.DirectPath
					filedata = mediadecrypt.GetBase64Filedata(Message.ExtendedTextMessage.ContextInfo.QuotedMessage.DocumentMessage.Url, Message.ExtendedTextMessage.ContextInfo.QuotedMessage.DocumentMessage.MediaKey)
				}
				if Message.ExtendedTextMessage.ContextInfo.QuotedMessage.DocumentWithCaptionMessage != nil {
					filename = *Message.ExtendedTextMessage.ContextInfo.QuotedMessage.DocumentWithCaptionMessage.Message.DocumentMessage.DirectPath
					filedata = mediadecrypt.GetBase64Filedata(Message.ExtendedTextMessage.ContextInfo.QuotedMessage.DocumentWithCaptionMessage.Message.DocumentMessage.Url, Message.ExtendedTextMessage.ContextInfo.QuotedMessage.DocumentWithCaptionMessage.Message.DocumentMessage.MediaKey)
				}
			}
		}
	} else if Message.DocumentMessage != nil {
		switch {
		case Message.DocumentMessage.Title != nil:
			filename = *Message.DocumentMessage.Title
		case Message.DocumentMessage.FileName != nil:
			filename = *Message.DocumentMessage.FileName
		}
		filedata = mediadecrypt.GetBase64Filedata(Message.DocumentMessage.Url, Message.DocumentMessage.MediaKey)
	}
	return

}

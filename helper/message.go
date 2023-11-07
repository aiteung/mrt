package helper

import (
	"strings"

	"github.com/aiteung/atmessage/mediadecrypt"
	"github.com/aiteung/module/model"
	waProto "go.mau.fi/whatsmeow/binary/proto"
)

func GetPhoneNumber(WAIface model.IteungWhatsMeowConfig) (phonenumber string) {
	phonenumber = WAIface.Info.Sender.User
	if WAIface.Message.ExtendedTextMessage != nil {
		if WAIface.Message.ExtendedTextMessage.ContextInfo != nil {
			//Kalo pake whatsapp Desktop masuk module ExtendedTextMessage ContextInfo expiration:0
			if WAIface.Message.ExtendedTextMessage.ContextInfo.Participant != nil {
				phonenumber = strings.Split(*WAIface.Message.ExtendedTextMessage.ContextInfo.Participant, "@")[0]
			}
		}
	}

	return

}

func GetMessage(Message *waProto.Message) (message string) {
	switch {
	case Message.ExtendedTextMessage != nil:
		message = *Message.ExtendedTextMessage.Text
	case Message.DocumentMessage != nil:
		if Message.DocumentMessage.Caption != nil {
			message = *Message.DocumentMessage.Caption
		}
	case Message.ImageMessage != nil:
		if Message.ImageMessage.Caption != nil {
			message = *Message.ImageMessage.Caption
		}
	case Message.LiveLocationMessage != nil:
		message = Message.LiveLocationMessage.GetCaption()
	default:
		message = Message.GetConversation()
	}
	return

}

func GetLongLat(Message *waProto.Message) (long, lat float64, liveloc bool) {
	if Message.ExtendedTextMessage != nil {
		if Message.ExtendedTextMessage.ContextInfo != nil {
			if Message.ExtendedTextMessage.ContextInfo.Participant != nil {
				if Message.ExtendedTextMessage.ContextInfo.QuotedMessage.LiveLocationMessage != nil {
					lat = *Message.ExtendedTextMessage.ContextInfo.QuotedMessage.LiveLocationMessage.DegreesLatitude
					long = *Message.ExtendedTextMessage.ContextInfo.QuotedMessage.LiveLocationMessage.DegreesLongitude
					liveloc = true
				}
			}

		}
	} else if Message.LiveLocationMessage != nil {
		long = *Message.LiveLocationMessage.DegreesLongitude
		lat = *Message.LiveLocationMessage.DegreesLatitude
		liveloc = true
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
	} else if Message.ImageMessage != nil {
		filename = strings.ReplaceAll(*Message.ImageMessage.Mimetype, "/", ".")
		filedata = mediadecrypt.GetBase64Filedata(Message.ImageMessage.Url, Message.ImageMessage.MediaKey)
	}
	return

}

func GetStatusFromLink(WAIface model.IteungWhatsMeowConfig) (whmsg bool) {
	if WAIface.Message.ExtendedTextMessage != nil && WAIface.Info.Chat.Server == "s.whatsapp.net" {
		if WAIface.Message.ExtendedTextMessage.ContextInfo != nil {
			if WAIface.Message.ExtendedTextMessage.ContextInfo.EntryPointConversionSource != nil {
				msg := *WAIface.Message.ExtendedTextMessage.ContextInfo.EntryPointConversionSource
				if msg == "click_to_chat_link" {
					whmsg = true
				}
			}
		}
	}
	return
}

func GetFromLinkDelay(Message *waProto.Message) uint32 {
	return *Message.ExtendedTextMessage.ContextInfo.EntryPointConversionDelaySeconds
}

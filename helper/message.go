package helper

import (
	"encoding/base64"
	"go.mau.fi/whatsmeow"
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

func GetFile(client *whatsmeow.Client, Message *waProto.Message) (filename, filedata string) {
	if extMsg := Message.GetExtendedTextMessage(); extMsg != nil {
		if extMsg.ContextInfo == nil {
			return
		}
		if extMsg.ContextInfo.Participant == nil {
			return
		}
		if extMsg.ContextInfo.QuotedMessage.DocumentMessage != nil {
			filename = *extMsg.ContextInfo.QuotedMessage.DocumentMessage.DirectPath
			payload, err := client.Download(extMsg.ContextInfo.QuotedMessage.DocumentMessage)
			if err != nil {
				return
			}
			filedata = base64.StdEncoding.EncodeToString(payload)
		}
		if extMsg.ContextInfo.QuotedMessage.DocumentWithCaptionMessage != nil {
			filename = *extMsg.ContextInfo.QuotedMessage.DocumentWithCaptionMessage.Message.DocumentMessage.DirectPath
			payload, err := client.Download(extMsg.ContextInfo.QuotedMessage.DocumentWithCaptionMessage.Message.DocumentMessage)
			if err != nil {
				return
			}
			filedata = base64.StdEncoding.EncodeToString(payload)
		}
	} else if doc := Message.GetDocumentMessage(); doc != nil {
		switch {
		case doc.Title != nil:
			filename = *doc.Title
		case doc.FileName != nil:
			filename = *doc.FileName
		}
		payload, err := client.Download(doc)
		if err != nil {
			return
		}
		filedata = base64.StdEncoding.EncodeToString(payload)
	} else if img := Message.GetImageMessage(); img != nil {
		filename = strings.ReplaceAll(*img.Mimetype, "/", ".")
		filedata = mediadecrypt.GetBase64Filedata(img.URL, img.GetMediaKey())
		payload, err := client.Download(img)
		if err != nil {
			return
		}
		filedata = base64.StdEncoding.EncodeToString(payload)
	} else if docCap := Message.GetDocumentWithCaptionMessage(); docCap != nil {
		if docCap.GetMessage() == nil {
			return
		}
		switch {
		case docCap.GetMessage().GetDocumentMessage().Title != nil:
			filename = docCap.GetMessage().GetDocumentMessage().GetTitle()
		case docCap.GetMessage().GetDocumentMessage().FileName != nil:
			filename = docCap.GetMessage().GetDocumentMessage().GetFileName()
		}
		payload, err := client.Download(docCap.Message.DocumentMessage)
		if err != nil {
			return
		}

		filedata = base64.StdEncoding.EncodeToString(payload)
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

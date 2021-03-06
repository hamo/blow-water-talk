package weixin

import (
	"encoding/xml"
	"errors"
	"strings"
	"time"
)

var (
	ErrMessageFormat = errors.New("")
)

type MessageBase struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
}

type MessageSendBase MessageBase

type MessageReceiveBase struct {
	MessageBase

	MsgId uint64
}

type MessageReceiveText struct {
	MessageReceiveBase

	Content string
}

type MessageReceiveImage struct {
	MessageReceiveBase

	PicUrl  string
	MediaId string
}

type MessageReceiveVoice struct {
	MessageReceiveBase

	MediaId string
	Format  string
}

type MessageReceiveVideo struct {
	MessageReceiveBase

	MediaId      string
	ThumbMediaId string
}

type MessageReceiveLocal struct {
	MessageReceiveBase

	Location_X float32
	Location_Y float32
	Scale      int
	Label      string
}

type MessageReceiveLink struct {
	MessageReceiveBase

	Title       string
	Description string
	Url         string
}

type MessageSendText struct {
	MessageSendBase

	Content string

	XMLName struct{} `xml:"xml"`
}

type MessageSendImage struct {
	MessageSendBase

	MediaId string

	XMLName struct{} `xml:"xml"`
}

type MessageSendAudio struct {
	MessageSendBase

	MediaId string

	XMLName struct{} `xml:"xml"`
}

type MessageSendVideo struct {
	MessageSendBase

	MediaId     string
	Title       string
	Description string

	XMLName struct{} `xml:"xml"`
}

type MessageSendMusic struct {
	MessageSendBase

	Title        string
	Description  string
	MusicURL     string
	HQMusicUrl   string
	ThumbMediaId string

	XMLName struct{} `xml:"xml"`
}

//type MessageSendTextAndImage struct

func MessageDecodeReceive(msg string) (MsgType string, p interface{}, err error) {
	d := xml.NewDecoder(strings.NewReader(msg))

	mrb := new(MessageReceiveBase)
	var t string
	var re interface{}

	if err := d.Decode(mrb); err != nil {
		goto errDecode
	}

	t = mrb.MsgType

	// reset Decoder
	d = xml.NewDecoder(strings.NewReader(msg))

	switch t {
	case "text":
		re = new(MessageReceiveText)
	case "image":
		re = new(MessageReceiveImage)
	case "voice":
		re = new(MessageReceiveVoice)
	case "video":
		re = new(MessageReceiveVideo)
	case "location":
		re = new(MessageReceiveLocal)
	case "link":
		re = new(MessageReceiveLink)
	default:
		goto errDecode
	}

	if err := d.Decode(re); err != nil {
		goto errDecode
	}

	return t, re, nil

errDecode:
	return "", nil, ErrMessageFormat
}

func MessageCreateText(from, to, content string) *MessageSendText {
	re := new(MessageSendText)

	re.FromUserName = from
	re.ToUserName = to
	re.MsgType = "text"

	// FIXME: location
	re.CreateTime = time.Now().Unix()

	re.Content = content

	return re
}

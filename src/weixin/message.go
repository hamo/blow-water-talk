package weixin

import ()

type MessageBase struct {
	ToUserName   string
	FromUserName string
	CreateTime   int
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

type MessageReceiveAudio struct {
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
}

type MessageSendImage struct {
	MessageSendBase

	MediaId string
}

type MessageSendAudio struct {
	MessageSendBase

	MediaId string
}

type MessageSendVideo struct {
	MessageSendBase

	MediaId     string
	Title       string
	Description string
}

type MessageSendMusic struct {
	MessageSendBase

	Title        string
	Description  string
	MusicURL     string
	HQMusicUrl   string
	ThumbMediaId string
}

//type MessageSendTextAndImage struct

package main

import (
	"encoding/xml"
	"fmt"

	"weixin"
)

func msgHandle(answer chan []byte, t string, p interface{}) {
	switch t {
	case "text":
		msgHandleText(answer, p.(*weixin.MessageReceiveText))

	// FIXME: Just for Maimeng
	case "image":
		msgHandleImage(answer, p.(*weixin.MessageReceiveImage))
	case "voice":
		msgHandleVoice(answer, p.(*weixin.MessageReceiveVoice))
	case "video":
		msgHandleVideo(answer, p.(*weixin.MessageReceiveVideo))
	case "location":
		msgHandleLocal(answer, p.(*weixin.MessageReceiveLocal))
	case "link":
		msgHandleLink(answer, p.(*weixin.MessageReceiveLink))
	default:
		return
	}
}

func msgHandleText(answer chan []byte, msg *weixin.MessageReceiveText) {
	from := msg.ToUserName
	to := msg.FromUserName

	a := weixin.MessageCreateText(from, to, "Hello 世界")

	x, _ := xml.Marshal(a)

	fmt.Printf("%s\n", string(x))
	answer <- x
	return
}

func msgHandleImage(answer chan []byte, msg *weixin.MessageReceiveImage) {
	from := msg.ToUserName
	to := msg.FromUserName

	a := weixin.MessageCreateText(from, to, "讨厌啦，人家还没学会呢～")
	x, _ := xml.Marshal(a)
	answer <- x
	return
}

func msgHandleVoice(answer chan []byte, msg *weixin.MessageReceiveVoice) {

	from := msg.ToUserName
	to := msg.FromUserName

	a := weixin.MessageCreateText(from, to, "讨厌啦，人家还没学会呢～")
	x, _ := xml.Marshal(a)
	answer <- x
	return
}

func msgHandleVideo(answer chan []byte, msg *weixin.MessageReceiveVideo) {
	from := msg.ToUserName
	to := msg.FromUserName

	a := weixin.MessageCreateText(from, to, "讨厌啦，人家还没学会呢～")
	x, _ := xml.Marshal(a)
	answer <- x
	return
}

func msgHandleLocal(answer chan []byte, msg *weixin.MessageReceiveLocal) {
	from := msg.ToUserName
	to := msg.FromUserName

	a := weixin.MessageCreateText(from, to, "讨厌啦，人家还没学会呢～")
	x, _ := xml.Marshal(a)
	answer <- x
	return
}

func msgHandleLink(answer chan []byte, msg *weixin.MessageReceiveLink) {
	from := msg.ToUserName
	to := msg.FromUserName

	a := weixin.MessageCreateText(from, to, "讨厌啦，人家还没学会呢～")
	x, _ := xml.Marshal(a)
	answer <- x
	return
}

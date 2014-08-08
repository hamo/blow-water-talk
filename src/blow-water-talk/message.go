package main

import (
	"encoding/xml"
	"fmt"

	"weixin"
)

func msgHandle(answer chan []byte, t string, p interface{}) {
	var from string
	var to string

	switch t {
	case "text":
		msgHandleText(answer, p.(*weixin.MessageReceiveText))

	// FIXME: Just for Maimeng
	case "image":
		from = p.(*weixin.MessageReceiveImage).ToUserName
		to = p.(*weixin.MessageReceiveImage).FromUserName
		goto maimeng
	case "voice":
		from = p.(*weixin.MessageReceiveVoice).ToUserName
		to = p.(*weixin.MessageReceiveVoice).FromUserName
		goto maimeng
	case "video":
		from = p.(*weixin.MessageReceiveVideo).ToUserName
		to = p.(*weixin.MessageReceiveVideo).FromUserName
		goto maimeng
	case "location":
		from = p.(*weixin.MessageReceiveLocal).ToUserName
		to = p.(*weixin.MessageReceiveLocal).FromUserName
		goto maimeng
	case "link":
		from = p.(*weixin.MessageReceiveLink).ToUserName
		to = p.(*weixin.MessageReceiveLink).FromUserName
		goto maimeng
	default:
		return
	}

maimeng:
	a := weixin.MessageCreateText(from, to, "讨厌啦，人家还没学会呢～")
	x, _ := xml.Marshal(a)
	answer <- x
	return
}

func msgHandleText(answer chan []byte, text *weixin.MessageReceiveText) {
	from := text.ToUserName
	to := text.FromUserName

	a := weixin.MessageCreateText(from, to, "Hello 世界")

	x, _ := xml.Marshal(a)

	fmt.Printf("%s\n", string(x))
	answer <- x

}

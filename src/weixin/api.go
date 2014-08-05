package weixin

import (
	"time"
)

const (
	ApiBaseURL = "https://api.weixin.qq.com/cgi-bin/"
)

type Api struct {
	Appid  string
	Secret string

	BaseURL string

	access_token string
	expires_in   time.Duration
}

func NewApi(appid string, secret string) (re *Api) {
	re = new(Api)

	re.Appid = appid
	re.Secret = secret

	re.BaseURL = ApiBaseURL
}

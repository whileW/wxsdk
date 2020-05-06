package xcx

import (
	"fmt"
	"wxsdk"
)

const (
	code2Session = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
)

type LoginStruct struct {
	Openid				string			`json:"openid"`
	SessionKey			string			`json:"session_key"`
	Unionid				string			`json:"unionid"`
	wxsdk.WXError
}

func Login(js_code string) (*LoginStruct,error) {
	url := fmt.Sprintf(code2Session, wxsdk.AppId, wxsdk.AppSecret, js_code)
	resp := &LoginStruct{}
	err := wxsdk.Get(url, resp)
	return resp, err
}
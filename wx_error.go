package wxsdk

import "fmt"

type wxResp interface {
	// 似有的 error 方法，保证外部(其他包)定义的 struct 只能内嵌
	// WXError 的才能实现这个方法，才能作为当前包 http 方法的参数
	error() error
}

type WXError struct {
	ErrCode int    `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}

func (e *WXError) error() error {
	if e.ErrCode == 0 {
		return nil
	}
	return e
}

func (e *WXError) Error() string {
	return fmt.Sprintf("weixin: (%d)%s", e.ErrCode, e.Errmsg)
}

type WXXmlError struct {
	ReturnCode string    	`xml:"return_code"`
	ReturnMsg  string 		`xml:"return_msg"`
}

func (e *WXXmlError) error() error {
	if e.ReturnCode == "SUCCESS" {
		return nil
	}
	return e
}

func (e *WXXmlError) Error() string {
	return fmt.Sprintf("weixin: (%s)%s", e.ReturnCode, e.ReturnMsg)
}

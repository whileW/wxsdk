package xcx

import (
	"encoding/xml"
	"fmt"
	"github.com/go-errors/errors"
	"github.com/satori/go.uuid"
	"github.com/whileW/wxsdk"
	"strings"
)

const (
	unifiedorder = "https://api.mch.weixin.qq.com/pay/unifiedorder"
)

type UnifiedorderReqStruct struct {
	Appid			string				`json:"appid"`			//小程序id
	MchId			string				`json:"mch_id"`			//商户号
	nonce_str		string				`json:"nonce_str"`		//随机字符串--32位
	Sign			string				`json:"sign"`			//签名
	Body			string 				`json:"body"`			//商品描述
	OutTradeNo		string				`json:"out_trade_no"`	//订单号
	TotalFee		int					`json:"total_fee"`		//订单金额	单位：分
	SpbillCreateIp	string				`json:"spbill_create_ip"`	//终端ip
	NotifyUrl		string				`json:"notify_url"`		//通知地址
	TradeType		string				`json:"trade_type"`		//交易类型
}
type UnifiedorderRespStruct struct {
	ResultCode			string			`json:"result_code"`
	ErrCode				string			`json:"err_code"`
	ErrCodeDes			string			`json:"err_code_des"`
	wxsdk.WXXmlError
}

func Unifiedorder(body string,total int,ip string,notify_url string,trade_type string) (string,error) {
	reqStruct := &UnifiedorderReqStruct{
		Appid:wxsdk.AppId,
		MchId:wxsdk.MchId,
		nonce_str:strings.Replace(uuid.NewV1().String(),"-","",0),
		Body:body,
		OutTradeNo:strings.Replace(uuid.NewV1().String(),"-","",0),
		TotalFee:total,
		SpbillCreateIp:ip,
		NotifyUrl:notify_url,
		TradeType:trade_type,
	}
	req,err := xml.Marshal(reqStruct)
	if err != nil {
		return "",err
	}
	resp := &UnifiedorderRespStruct{}
	err = wxsdk.PostXml(unifiedorder,req, resp)
	if err != nil {
		if resp.ResultCode != "SUCCESS" {
			return "",errors.New(fmt.Sprintf("weixin: (%s)%s", resp.ErrCode, resp.ErrCodeDes))
		}
	}
	return reqStruct.OutTradeNo, err
}
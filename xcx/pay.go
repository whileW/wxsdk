package xcx

import (
	"code.aliyun.com/sxs/utils/encryption"
	"errors"
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/whileW/wxsdk"
	"strings"
	xml2 "encoding/xml"
)

const (
	unifiedorder = "https://api.mch.weixin.qq.com/pay/unifiedorder"
)

type xml struct {
	Appid			string			`xml:"appid"`			//小程序id
	Body			string 			`xml:"body"`			//商品描述
	MchId			string			`xml:"mch_id"`			//商户号
	NonceStr		string			`xml:"nonce_str"`			//随机字符串--32位
	NotifyUrl		string			`xml:"notify_url"`			//通知地址
	OutTradeNo		string			`xml:"out_trade_no"`		//订单号
	Openid			string			`xml:"openid"`				//openid
	Sign			string			`xml:"sign"`			//签名
	SpbillCreateIp	string			`xml:"spbill_create_ip"`		//终端ip
	TotalFee		int				`xml:"total_fee"`			//订单金额	单位：分
	TradeType		string			`xml:"trade_type"`			//交易类型
}
type UnifiedorderRespStruct struct {
	ResultCode			string			`xml:"result_code"`
	ErrCode				string			`xml:"err_code"`
	ErrCodeDes			string			`xml:"err_code_des"`
	PrepayId			string			`xml:"prepay_id"`			//统一下单返回id
	wxsdk.WXXmlError
}

func Unifiedorder(order_id string,body string,total int,ip string,notify_url string,
	trade_type string,openid string) (string,error) {
	reqStruct := &xml{
		Appid:wxsdk.AppId,
		MchId:wxsdk.MchId,
		NonceStr:strings.Replace(uuid.NewV1().String(),"-","",-1),
		Body:body,
		OutTradeNo:order_id,
		TotalFee:total,
		SpbillCreateIp:ip,
		NotifyUrl:notify_url,
		TradeType:trade_type,
		Openid:openid,
	}
	reqStruct.Sign = getSign(reqStruct)
	req,err := xml2.Marshal(reqStruct)
	if err != nil {
		return "",err
	}
	resp := &UnifiedorderRespStruct{}
	err = wxsdk.PostXml(unifiedorder,req, resp)
	if err != nil {
		return "",err
	}else {
		if resp.ResultCode != "SUCCESS" {
			return "",errors.New(fmt.Sprintf("weixin: (%s)%s", resp.ErrCode, resp.ErrCodeDes))
		}
	}
	return resp.PrepayId, err
}
func getSign(req *xml) string {
	str := fmt.Sprintf("appid=%s&body=%s&mch_id=%s&nonce_str=%s&notify_url=%s&" +
		"openid=%s&out_trade_no=%s&spbill_create_ip=%s&total_fee=%d&trade_type=%s&key=%s",
		req.Appid,req.Body,req.MchId,req.NonceStr,
		req.NotifyUrl,req.Openid,req.OutTradeNo,req.SpbillCreateIp,req.TotalFee,req.TradeType,wxsdk.MchKey)
	md5str := encryption.Md5V(str)
	//wxsdk.ComputeHmacSha256(md5str,wxsdk.MchKey)
	return strings.ToLower(md5str)
}
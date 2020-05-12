package xcx

import (
	"github.com/whileW/wxsdk"
	"testing"
)

func TestUnifiedorder(t *testing.T) {
	wxsdk.AppId = "wx14973ca54b510f35"
	wxsdk.MchId = "1529206941"
	wxsdk.MchKey = "79GQAJwZigkyvgqgEMmqqawV3YHvN9ZC"
	str,err := Unifiedorder("jjdskaskf","弘途在线-测试商品",100,"117.89.241.15","https://api.uxsw.cn/local/v1/zyjy/order/wxpay/pay","JSAPI","o1XVI43q0xTQ-pCoNZ5ITNMIteSY")
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(str)
}
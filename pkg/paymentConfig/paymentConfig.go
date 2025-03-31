package paymentConfig

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

type Config struct {
	pid    string // 商户号
	secret string //密钥
}

type Payment struct {
	pid          string               // 商户号
	method       string               // 接口调用类型 web、hump、jsapi、app、scan、applet
	device       string               // 设备号
	paymentType  string               // wxpay，qqpay，alipay
	out_trade_no string               // 商户订单号
	notify_url   string               // 通知地址
	return_url   string               // 返回地址
	name         string               // 商品名称
	money        string               // 商品金额
	client_ip    string               // 客户端ip
	param        string               // 附加参数
	timestamp    *timestamp.Timestamp // 时间戳
	sign         string               // 签名
	sign_type    string               // 签名类型
}

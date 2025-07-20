# 彩虹易支付(V2) Go语言版本SDK

本项目是彩虹易支付（V2）平台的 Go 语言 SDK，支持订单创建、签名、支付请求等常用功能，便于在 Go 项目中集成彩虹易支付。

## 目录结构

```
.
├── .env
├── export.go
├── go.mod
├── go.sum
├── README.md
├── internal/
│   ├── models/
│   │   ├── Order.go
│   │   └── Response.go
│   └── utils/
│       ├── Fetches.go
│       ├── MD5.go
│       └── RSA.go
└── pkg/
    └── PayService/
        └── makePayment.go
```

## 功能特性

- 订单结构体与参数封装
- MD5、RSA 签名工具
- 支持多种 POST/GET 请求方式
- 支付下单接口封装
- 便于扩展和二次开发

## 快速开始

1. **安装依赖**

   ```bash
   go mod tidy
   ```

2. **配置环境变量**

   在 `.env` 文件中配置你的支付参数，例如：

   ```
   PAY_PID=你的商户ID
   PAY_NOTIFY_URL=你的回调地址
   PAY_RETURN_URL=你的前端回跳地址
   PAY_SIGN_TYPE=MD5
   PRIVATE_KEY=你的密钥
   ```

3. **创建订单并发起支付**

   示例代码：

   ```go
   import (
       "github.com/JeffreyZhu0201/rainbow-pay-sdk-go/internal/models"
       "github.com/JeffreyZhu0201/rainbow-pay-sdk-go/pkg/PayService"
   )

   order := models.Order{
       OutTradeNo:   "20240510001",
       CommodityName: "VIP会员",
       Amount:       9.9,
   }
   resp := PayService.CreateOrder(order)
   fmt.Println(resp)
   ```

## 主要模块说明

- [`internal/models/Order.go`](internal/models/Order.go)：订单结构体定义
- [`internal/utils/MD5.go`](internal/utils/MD5.go)：MD5 签名相关
- [`internal/utils/RSA.go`](internal/utils/RSA.go)：RSA 签名相关
- [`internal/utils/Fetches.go`](internal/utils/Fetches.go)：HTTP 请求工具
- [`pkg/PayService/makePayment.go`](pkg/PayService/makePayment.go)：支付下单逻辑

## 贡献

欢迎 issue 和 PR！

##
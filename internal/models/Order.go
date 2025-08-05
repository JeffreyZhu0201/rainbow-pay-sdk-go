package models

// import (
// )

type Order struct {
	OutTradeNo string //自定义订单号
	// PaidUser    string   // 支付用户
	CommodityName string //商品ID
	// PaidStatus  string   	//支付状态
	Amount float32
}

// // BeforeCreate 在创建记录之前生成 UUID
// func (u *Order) BeforeCreate(tx *gorm.DB) (err error) {
// 	u.ID = uuid.New().String()
// 	return
// }

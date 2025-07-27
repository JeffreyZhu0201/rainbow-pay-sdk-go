/*
 * @Author: JeffreyZhu 1624410543@qq.com
 * @Date: 2025-07-27 17:13:53
 * @LastEditors: JeffreyZhu 1624410543@qq.com
 * @LastEditTime: 2025-07-27 17:46:50
 * @FilePath: /workspace/rainbow-pay-sdk-go/internal/models/Order.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
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

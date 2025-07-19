/*
 * @Author: Jeffrey Zhu 1624410543@qq.com
 * @Date: 2025-04-24 19:31:57
 * @LastEditors: Jeffrey Zhu 1624410543@qq.com
 * @LastEditTime: 2025-04-24 19:31:57
 * @FilePath: \RocketVPN\go-backend\models\Response.go
 * @Description: File Description Here...
 *
 * Copyright (c) 2025 by JeffreyZhu, All Rights Reserved.
 */
package models

// Response represents the standard API response structure
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

package utils

import (
	"fmt"
	"sort"
	"strings"
)

/**
 * @description: 返回排序后字符串
 * @param {map[string]interface{}} m
 * @return {*}
 */
func SortMap(m map[string]interface{}) string {
	signParams := make(map[string]string)

	for k, v := range m {
		if v != "" && k != "sign" && k != "sign_type" {
			signParams[k] = fmt.Sprint(v)
		}
	}

	// 按ASCII码排序参数名
	var keys []string
	for k := range signParams {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 拼接参数
	var signStr strings.Builder
	for i, k := range keys {
		if i > 0 {
			signStr.WriteString("&")
		}
		signStr.WriteString(string(k))
		signStr.WriteString("=")
		signStr.WriteString(string(signParams[k]))
	}
	return signStr.String()
}

/*
 * @Author: JeffreyZhu 1624410543@qq.com
 * @Date: 2025-07-27 17:13:53
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2025-07-28 14:43:54
 * @FilePath: /workspace/rainbow-pay-sdk-go/internal/utils/MD5.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package utils

// Package utils provides common utility functions including MD5 hashing operations.
// The imported packages support MD5 hash generation, hex encoding, string manipulation,
// file operations, and sorting functionality required by the utilities.
import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"sort"
	"strings"
)

/**
 * @description:
 * @param {map[string]interface{}}
 * @return {*}
 */
func SortMapAndSignMD5(m map[string]interface{}) (string, string) {

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

	// 添加商户密钥并计算MD5
	finalStr := signStr.String() + os.Getenv("PRIVATE_KEY")
	h := md5.New()
	h.Write([]byte(finalStr))
	//paymentParams["sign"] = hex.EncodeToString(h.Sum(nil))
	signStr.WriteString("&sign_type=MD5&sign=" + hex.EncodeToString(h.Sum(nil)))

	return signStr.String(), hex.EncodeToString(h.Sum(nil))
}

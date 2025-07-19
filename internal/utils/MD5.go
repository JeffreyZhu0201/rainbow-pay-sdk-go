package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"sort"
	"strings"
)

func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func SortMapAndSign(m map[string]interface{}) (strings.Builder, string) {

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

	return signStr, hex.EncodeToString(h.Sum(nil))
}

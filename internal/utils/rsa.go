/*
 * @Author: Jeffrey Zhu 1624410543@qq.com
 * @Date: 2025-05-08 00:06:43
 * @LastEditors: Jeffrey Zhu 1624410543@qq.com
 * @LastEditTime: 2025-05-08 00:06:52
 * @FilePath: \RocketVPN\go-backend\utils\RSA.go
 * @Description: File Description Here...
 *
 * Copyright (c) 2025 by JeffreyZhu, All Rights Reserved.
 */
package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

func LoadPrivateKeyFromPEM(privateKeyPEM string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil {
		return nil, errors.New("failed to parse PEM block")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

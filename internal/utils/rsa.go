/*
 * @Author: Jeffrey Zhu 1624410543@qq.com
 * @Date: 2025-05-08 00:06:43
 * @LastEditors: JeffreyZhu 1624410543@qq.com
 * @LastEditTime: 2025-07-28 15:04:47
 * @FilePath: \RocketVPN\go-backend\utils\RSA.go
 * @Description: File Description Here...
 *
 * Copyright (c) 2025 by JeffreyZhu, All Rights Reserved.
 */
package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"os"
)

// func main() {
// 	// 1. 准备要签名的字符串数据
// 	data := "这是一条需要签名的敏感交易数据: {转账金额: 15000, 收款人: 张三, 交易时间: 2023-07-15}"

// 	// 2. 生成或加载密钥
// 	privateKey, publicKey := getKeyPair()

// 	// 3. 创建签名
// 	signature, err := SignString(privateKey, data)
// 	if err != nil {
// 		log.Fatalf("签名失败: %v", err)
// 	}

// 	fmt.Printf("签名结果 (Base64):\n%s\n", signature)

// 	// 4. 验证签名
// 	if VerifyString(publicKey, data, signature) {
// 		fmt.Println("✅ 签名验证成功 - 数据完整且来源可信")
// 	} else {
// 		fmt.Println("❌ 签名验证失败 - 数据可能被篡改")
// 	}
// }

// 获取密钥对（生成新密钥或从文件加载）
func getKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
	privKeyFile := "private.pem"
	pubKeyFile := "public.pem"

	// 如果存在密钥文件则加载
	if fileExists(privKeyFile) {
		privateKey, err := LoadPrivateKey(privKeyFile)
		if err != nil {
			log.Fatalf("私钥加载失败: %v", err)
		}
		return privateKey, &privateKey.PublicKey
	}

	// 生成新密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("密钥生成失败: %v", err)
	}

	// 保存密钥对
	if err := SavePrivateKey(privateKey, privKeyFile); err != nil {
		log.Fatalf("私钥保存失败: %v", err)
	}

	if err := SavePublicKey(&privateKey.PublicKey, pubKeyFile); err != nil {
		log.Fatalf("公钥保存失败: %v", err)
	}

	fmt.Println("✅ 已生成新的RSA密钥对")
	return privateKey, &privateKey.PublicKey
}

// 检查文件是否存在
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

// 保存私钥到文件
func SavePrivateKey(privateKey *rsa.PrivateKey, filename string) error {
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return pem.Encode(file, privateKeyPEM)
}

// 保存公钥到文件
func SavePublicKey(publicKey *rsa.PublicKey, filename string) error {
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}

	publicKeyPEM := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return pem.Encode(file, publicKeyPEM)
}

// 从文件加载私钥
func LoadPrivateKey(filename string) (*rsa.PrivateKey, error) {
	keyBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, errors.New("无效的PEM格式")
	}

	if block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("不是RSA私钥")
	}

	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

// 对字符串进行SHA256WithRSA签名
func SignStringRSA(privateKey *rsa.PrivateKey, data string) (string, error) {
	// 计算字符串的SHA256哈希
	hashed := sha256.Sum256([]byte(data))

	// 使用私钥进行PKCS1v15签名
	signature, err := rsa.SignPKCS1v15(
		rand.Reader,
		privateKey,
		crypto.SHA256,
		hashed[:],
	)
	if err != nil {
		return "", fmt.Errorf("签名创建失败: %w", err)
	}

	// 返回Base64编码的签名
	return base64.StdEncoding.EncodeToString(signature), nil
}

// 验证字符串的SHA256WithRSA签名
func VerifyString(publicKey *rsa.PublicKey, data string, base64Signature string) bool {
	// 计算字符串的SHA256哈希
	hashed := sha256.Sum256([]byte(data))

	// 解码Base64签名
	signature, err := base64.StdEncoding.DecodeString(base64Signature)
	if err != nil {
		log.Printf("签名解码失败: %v", err)
		return false
	}

	// 使用公钥验证签名
	err = rsa.VerifyPKCS1v15(
		publicKey,
		crypto.SHA256,
		hashed[:],
		signature,
	)

	return err == nil
}

func SortMapAndSignRSA(m map[string]interface{}) (string, string) {

	privateKey, _ := getKeyPair()
	// 3. 创建签名
	signature, err := SignStringRSA(privateKey, SortMap(m))
	if err != nil {
		log.Fatalf("签名失败: %v", err)
	}

	fmt.Printf("签名结果 (Base64):\n%s\n", signature)
	return signature, ""
}

func VerifyStringRSA(m map[string]interface{}, signature string) bool {

	_, publicKey := getKeyPair()

	originalStr := SortMap(m)

	// signature, err := SignStringRSA(privateKey, originalStr)

	if VerifyString(publicKey, originalStr, signature) {
		fmt.Println("✅ 签名验证成功 - 数据完整且来源可信")
		return true
	} else {
		fmt.Println("❌ 签名验证失败 - 数据可能被篡改")
		return false
	}
}

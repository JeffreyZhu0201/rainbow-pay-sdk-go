/*
 * @Author: Jeffrey Zhu 1624410543@qq.com
 * @Date: 2025-03-31 14:28:12
 * @LastEditors: Jeffrey Zhu 1624410543@qq.com
 * @LastEditTime: 2025-03-31 14:30:11
 * @FilePath: \rainbow-pay-sdk-go\internal\utils\rsa.go
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
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// GenerateSignContent generates content for signing by sorting and formatting struct fields
func GenerateSignContent(params interface{}) string {
	// Get struct fields and values
	v := reflect.ValueOf(params)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()

	// Create slice to store field-value pairs
	pairs := make([]string, 0)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		// Skip sign and sign_type fields
		if field.Tag.Get("json") == "sign" || field.Tag.Get("json") == "sign_type" {
			continue
		}

		// Get json tag name or field name
		name := field.Tag.Get("json")
		if name == "" {
			name = field.Name
		}

		// Add non-empty values to pairs
		if value.String() != "" {
			pairs = append(pairs, fmt.Sprintf("%s=%v", name, value.Interface()))
		}
	}

	// Sort by ASCII code
	sort.Strings(pairs)

	// Join with &
	return strings.Join(pairs, "&")
}

// SignWithRSA signs content using RSA-SHA256
func SignWithRSA(content string, privateKey string) (string, error) {
	// Decode private key
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return "", fmt.Errorf("failed to decode private key")
	}

	// Parse private key
	priKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	// Calculate hash
	h := sha256.New()
	h.Write([]byte(content))
	hash := h.Sum(nil)

	// Sign
	signature, err := rsa.SignPKCS1v15(rand.Reader, priKey, crypto.SHA256, hash)
	if err != nil {
		return "", err
	}

	// Return base64 encoded signature
	return base64.StdEncoding.EncodeToString(signature), nil
}

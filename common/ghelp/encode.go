// json常用编码、解码操作

package ghelp

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
)

// JsonEncode json编码操作
func JsonEncode(data interface{}) string {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(jsonStr)
}

// JsonDecode json解码操作
// 成功返回interface, 失败返回nil
func JsonDecode(jsonStr string) interface{} {
	var v interface{}
	err := json.Unmarshal([]byte(jsonStr), &v)
	if err != nil {
		return nil
	}
	return v
}

// JsonDecodeTo json解码操作
// 失败返回错误，成功返回nil
func JsonDecodeTo(jsonStr string, dst interface{}) error {
	err := json.Unmarshal([]byte(jsonStr), &dst)
	return err
}

// Base64Encode base64编码操作
func Base64Encode(data string) string {
	bData := []byte(data)
	return base64.StdEncoding.EncodeToString(bData)
}

// Base64Decode base64解码操作
func Base64Decode(encodeStr string) string {
	bData, err := base64.StdEncoding.DecodeString(encodeStr)
	if err != nil {
		return ""
	}
	return string(bData)
}

// Md5 生成md5字符串， binary为是否设置原始16 字符二进制格式，默认为false
func Md5(data string, binary ...bool) string {
	h := md5.New()
	_, err := h.Write([]byte(data))
	if err != nil {
		return ""
	}
	d := h.Sum(nil)
	if len(binary) > 0 && binary[0] {
		return string(d)
	} else {
		return hex.EncodeToString(d)
	}
}

// Sha1
func Sha1(data string) string {
	h := sha1.New()
	_, err := h.Write([]byte(data))
	if err != nil {
		return ""
	}
	return hex.EncodeToString(h.Sum(nil))
}

// HmacSha256
func HmacSha256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

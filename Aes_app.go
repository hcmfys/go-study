package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
)

func main() {
	orig := "hello world12345499998866vv666"
	key := "123456781234567812345678"
	iv := "xx456781234567812xxx678v"
	fmt.Println("原文：", orig)

	encryptCode := AesEncrypt(orig, key, iv)
	fmt.Println("aes密文：", encryptCode)

	decryptCode := AesDecrypt(encryptCode, key, iv)
	fmt.Println("aes解密结果：", decryptCode)

	key = "2fa6c1e9"
	decEncrypt, _ := DesEncrypt(orig, []byte(key))
	fmt.Println("des密文：", decEncrypt)

	desDecrypt, _ := DesDecrypt(decEncrypt, []byte(key))
	fmt.Println("dess解密结果：", desDecrypt)

}

func AesEncrypt(orig string, key string, iv string) string {

	origData := []byte(orig)
	k := []byte(key)

	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	//补全码
	origData = PKCS7Padding(origData, blockSize)
	//加密模式
	blockMode := cipher.NewCBCEncrypter(block, []byte(iv)[:blockSize])
	//创建数组
	cryted := make([]byte, len(origData))
	blockMode.CryptBlocks(cryted, origData)
	return base64.StdEncoding.EncodeToString(cryted)

}

func AesDecrypt(crypted string, key string, iv string) string {

	cryptedByte, _ := base64.StdEncoding.DecodeString(crypted)
	k := []byte(key)
	//分组秘钥
	block, _ := aes.NewCipher(k)
	// 解密
	blockSize := block.BlockSize()
	// 加密模式
	blockeMode := cipher.NewCBCDecrypter(block, []byte(iv)[:blockSize])
	// 解密
	orig := make([]byte, len(cryptedByte))
	// 解密
	blockeMode.CryptBlocks(orig, cryptedByte)
	// 解密
	orig = PKCS7UnPadding(orig)
	return string(orig)

}

//补码
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func DesEncrypt(text string, key []byte) (string, error) {
	src := []byte(text)
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()
	src = ZeroPadding(src, bs)
	if len(src)%bs != 0 {
		return "", errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	return hex.EncodeToString(out), nil
}

func DesDecrypt(decrypted string, key []byte) (string, error) {
	src, err := hex.DecodeString(decrypted)
	if err != nil {
		return "", err
	}
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	out := make([]byte, len(src))
	dst := out
	bs := block.BlockSize()
	if len(src)%bs != 0 {
		return "", errors.New("crypto/cipher: input not full blocks")
	}
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	out = ZeroUnPadding(out)
	return string(out), nil
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}

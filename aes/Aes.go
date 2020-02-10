package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"fmt"
	"strings"
)

type Aes struct {
}

//PKCS7填充，
func (a *Aes) PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//PKCS7反填充，
func (a *Aes) PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//PKCS5填充，
func (a *Aes) PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//PKCS5反填充，
func (a *Aes) PKCS5UnPadding(ciphertext []byte) []byte {
	length := len(ciphertext)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(ciphertext[length-1])
	return ciphertext[:(length - unpadding)]
}

//0填充
func (a *Aes) ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding) //用0去填充
	return append(ciphertext, padtext...)
}

//反0填充
func (a *Aes) ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}

/**
 * AesCBC加密
 * @param str []byte 加密的字符串
 * @param key []byte 密钥
 * @param iv []byte 偏移量
 * @param padding string 填充方式 默认pkcs5填充 如：zero领填充 pkcs7填充
 */
func (a *Aes) AesCbcEncode(str, key, iv []byte, padding string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	if padding == "zero" {
		str = a.ZeroPadding(str, blockSize)
	} else if padding == "pkcs7" {
		str = a.PKCS7Padding(str, blockSize)
	} else {
		str = PKCS5Padding(str, blockSize)
	}
	var blockMode cipher.BlockMode
	blockMode = cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(str))
	blockMode.CryptBlocks(crypted, str)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

/**
 * AesCBC解密
 * @param str []byte 解密的字符串
 * @param key []byte 密钥
 * @param iv []byte 偏移量
 * @param padding string 填充方式 默认pkcs5填充 如：zero领填充 pkcs7填充
 */
func (a *Aes) AesCbcDecode(str, key, iv []byte, padding string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	var (
		blockMode cipher.BlockMode
	)

	blockMode = cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(str))
	blockMode.CryptBlocks(origData, str)
	if padding == "zero" {
		origData = a.ZeroUnPadding(origData)
	} else if padding == "pkcs7" {
		origData = a.PKCS7UnPadding(origData)
	} else {
		origData = a.PKCS5UnPadding(origData)
	}
	return string(origData), nil
}

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbEncrypter ecb

// NewECBEncrypter returns a BlockMode which encrypts in electronic code book
// mode, using the given Block.
func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypter)(newECB(b))
}
func (x *ecbEncrypter) BlockSize() int { return x.blockSize }
func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ecbDecrypter ecb

// NewECBDecrypter returns a BlockMode which decrypts in electronic code book
// mode, using the given Block.
func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}
func (x *ecbDecrypter) BlockSize() int { return x.blockSize }
func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

/*
 * aesecb 加密
 * @param origData []byte 加密的内容
 * @param key []byte 密钥
 * @param padding string 填充 （空pkcs5填充 可选 ZeroPadding）
 */
func (a *Aes) AesECBEncrypt(origData, key []byte, padding string, mode string) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	if padding == "zero" {
		origData = a.ZeroPadding(origData, blockSize)
	} else if padding == "pkcs7" {
		origData = a.PKCS7Padding(origData, blockSize)
	} else {
		origData = a.PKCS5Padding(origData, blockSize)
	}
	var blockMode cipher.BlockMode
	blockMode = NewECBEncrypter(block)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func (a *Aes) genKey(key string) string {
	if len(key) < 32 {
		key += strings.Repeat("0", 32)
	}
	return key[:32]
}

// 3DES解密
func (a *Aes) tripleDesDecrypt(crypted, key []byte, IV string) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, []byte(IV))
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = a.PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

// 3DES加密
func (a *Aes) TripleDesEncrypt(origData, key []byte, IV string) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	origData = a.PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, []byte(IV))
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

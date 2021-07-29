package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
	"golang.org/x/crypto/bcrypt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type CaptchaResponse struct {
	CaptchaId  string
	CaptchaUrl string
}

//获取验证码
func GetCaptcha() *CaptchaResponse {
	captchaId := captcha.NewLen(4)
	return &CaptchaResponse{
		CaptchaId:  captchaId,
		CaptchaUrl: fmt.Sprintf("/admin/auth/captcha/%s.png", captchaId),
	}
}

//模仿php的array_key_exists,判断是否存在map中
func KeyInMap(key string, m map[string]interface{}) bool {
	_, ok := m[key]
	if ok {
		return true
	} else {
		return false
	}
}

//模仿php的in_array,判断是否存在string数组中
func InArrayForString(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

//模仿php的in_array,判断是否存在int数组中
func InArrayForInt(items []int, item int) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

//php的函数password_hash
func PasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

//php的函数password_verify
func PasswordVerify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//int数组转string数组
func IntArrToStringArr(arr []int) []string {
	var stringArr []string
	for _, v := range arr {
		stringArr = append(stringArr, strconv.Itoa(v))
	}
	return stringArr
}

//对字符串进行MD5哈希
func GetMd5String(str string) string {
	t := md5.New()
	io.WriteString(t, str)
	return fmt.Sprintf("%x", t.Sum(nil))
}

//对字符串进行SHA1哈希
func GetSha1String(str string) string {
	t := sha1.New()
	io.WriteString(t, str)
	return fmt.Sprintf("%x", t.Sum(nil))
}

//字符串命名风格转换
func ParseName(name string, ptype int, ucfirst bool) string {
	if ptype > 0 {
		//解释正则表达式
		reg := regexp.MustCompile(`_([a-zA-Z])`)
		if reg == nil {
			beego.Error("MustCompile err")
			return ""
		}
		//提取关键信息
		result := reg.FindAllStringSubmatch(name, -1)
		for _, v := range result {
			name = strings.ReplaceAll(name, v[0], strings.ToUpper(v[1]))
		}

		if ucfirst {
			return Ucfirst(name)
		} else {
			return Lcfirst(name)
		}
	} else {
		//解释正则表达式
		reg := regexp.MustCompile(`[A-Z]`)
		if reg == nil {
			beego.Error("MustCompile err")
			return ""
		}
		//提取关键信息
		result := reg.FindAllStringSubmatch(name, -1)

		for _, v := range result {
			name = strings.ReplaceAll(name, v[0], "_"+v[0])
		}
		return strings.ToLower(name)
	}
}

//首字母大写
func Ucfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

//首字母小写
func Lcfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

//加密
func Encrypt(byteData []byte, aeskey []byte) string {
	xpass, err := AesEncrypt(byteData, aeskey)
	if err != nil {
		beego.Error(err)
		return ""
	}
	return base64.StdEncoding.EncodeToString(xpass)
}

//解密
func Decrypt(encryptData string, aeskey []byte) string {
	bytesPass, err := base64.StdEncoding.DecodeString(encryptData)
	if err != nil {
		beego.Error(err)
		return ""
	}

	decryptData, err := AesDecrypt(bytesPass, aeskey)
	if err != nil {
		beego.Error(err)
		return ""
	}
	return string(decryptData)
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

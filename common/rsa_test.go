package common

import (
	"encoding/hex"
	"fmt"
	"testing"
)

const prk = "-----BEGIN RSA PRIVATE KEY-----\nMIICXQIBAAKBgQC3emEnCtAUvuD3c/+KqrIRTGEFrvzssq+pamBhrv0ozVx14eu2\nvT/29ORPWzox5h7ceN7JsWmIJ9PZUbDNdvnL6Skw8+x6FtUbcQZpfX+1QiUpYSwZ\n1MuQ1ylzWAP9DTE++u2vSO+j0OCUIKWY8/gYftG4hkZjZBbczzWo/McpSwIDAQAB\nAoGAEUG9aYKm14ysdBnA6zXq0Z2xcmtm9oxH4VNUBVwEC5ZlH+FD3kgmf//AiYY3\nDwJp3KqxqZ66Ikg8sK/yRSDvlYlqW247SDajh7k5nXLDMhGYN30+HJXEofUJtY7s\nEwOjU2QAYrotJm4/oo3I7Q5tfd2gTfIa8xLp+rpwRVFC+oECQQDR1hv61BvEUreB\nCyzzFZtkpSS1MoPOeOoMyA35SCF8pqP7eogLum9PcCamHgGbhVxH+KNyq32R7WPc\ngmqc7u1RAkEA39fICEQfx3j9PYGv+4T6mcq62ezicpYcLBOKjlgaCjbrGTVoQQio\nfIklpz03BPJ5TVnZAUtGQz//U5fXz6KV2wJAR43HhMUHou7B/JMfBNV9Y9icp91N\n7P52cV1Wxoa+RI9eo8ao1bcBdgk8ZLEewzW6viAfPF8WNsjIoM0oJdOjwQJBAI+X\ng5lR4jT6tzESlYq6pmursiuEG0u4YcAglPx1JdcxnaThLsyxOiwRapca3MWOqiPl\npCCBYkRXtHmyaV2oBYsCQQDFsFbez4dabHMWRq0dFsZsMGqQnHVg+jkg52EokMQP\nXPlOtpQQAsrcnuAFWJL42IStGCPpYoedAEV7lit80Yv4\n-----END RSA PRIVATE KEY-----"
const puk = "-----BEGIN PUBLIC KEY-----\nMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC3emEnCtAUvuD3c/+KqrIRTGEF\nrvzssq+pamBhrv0ozVx14eu2vT/29ORPWzox5h7ceN7JsWmIJ9PZUbDNdvnL6Skw\n8+x6FtUbcQZpfX+1QiUpYSwZ1MuQ1ylzWAP9DTE++u2vSO+j0OCUIKWY8/gYftG4\nhkZjZBbczzWo/McpSwIDAQAB\n-----END PUBLIC KEY-----"

func TestRsaEncrypt(t *testing.T) {
	//prvKey, pubKey := GenRsaKey()
	//fmt.Println(string(prvKey))
	//fmt.Println(string(pubKey))
	//var data = "123456"
	//ciphertext := RsaEncrypt([]byte(data), []byte(puk))
	sss := "61d063fc4f95f03251392441d5c289fceda1136052df6d61d690b4bce58c4afdf4d4053fb83478a03f495465823011cc2b6f31d8fcaf6f06f05a9c9b7b5f0bcbbb4bf3dd62fd7ce21b464a5b00406748539ef12fe7aa99aa8bb681cebc82e8606ed98e196234d092cfa6f8a41f96e1f568e0e262f8744b4092c689c2642774ec"
	a, _ := hex.DecodeString(sss)
	sourceData := RsaDecrypt(a, []byte(prk))
	fmt.Println("私钥解密后的数据：", string(sourceData))
}


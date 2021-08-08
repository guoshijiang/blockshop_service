package common

import (
	"fmt"
	"testing"
)

func TestRsaEncrypt(t *testing.T) {
	sk, pk := GenerateRSAKey(1024)
	fmt.Println(string(sk), string(pk))
	a := RsaEncrypt([]byte("123456"), pk)
	fmt.Println(string(a))
	b := RsaDecrypt(a, sk)
	fmt.Println(string(b))
}


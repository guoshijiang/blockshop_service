package common

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

func GenerateRSAKey(bits int) ([]byte, []byte){
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		panic(err)
	}
	X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	privateBlock := pem.Block{Type: "RSA Private Key", Bytes: X509PrivateKey}
	publicKey := privateKey.PublicKey
	X509PublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}
	publicBlock := pem.Block{Type: "RSA Public Key", Bytes: X509PublicKey}
	return pem.EncodeToMemory(&privateBlock), pem.EncodeToMemory(&publicBlock)
}

func RsaEncrypt(plainText []byte, pk []byte) []byte {
	block, _ := pem.Decode(pk)
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		panic(err)
	}
	return cipherText
}

func RsaDecrypt(cipherText []byte, sk []byte) []byte{
	block, _ := pem.Decode(sk)
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil{
		panic(err)
	}
	plainText,_:=rsa.DecryptPKCS1v15(rand.Reader,privateKey,cipherText)
	return plainText
}


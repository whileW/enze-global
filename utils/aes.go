package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func AesDecrypt(Input string,Iv,Key []byte) ([]byte,error) {
	block, err := aes.NewCipher(Key)
	if err != nil {
		return nil,err
	}
	msg,_ := base64.StdEncoding.DecodeString(Input)
	cipherText := make([]byte, len(msg))
	mode := cipher.NewCBCDecrypter(block, Iv)
	mode.CryptBlocks(cipherText,msg)
	return decode2(cipherText),nil
}
func  decode2(decrypted []byte)[]byte {
	pad := int(decrypted[len(decrypted)-1])
	if pad < 1 || pad > 32 {
		pad = 0
	}
	res := decrypted[:len(decrypted)-pad]
	return res
}
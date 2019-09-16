// package today

// import (
// 	"bytes"
// 	"encoding/base64"
// 	"fmt"
// 	"goEncrypt"
// 	"log"

// 	"github.com/tjfoc/gmsm/sm2"
// )

// func today() {
// 	priv, err := sm2.GenerateKey() // 生成密钥对
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	msg := []byte("Tongji Fintech Research Institute")
// 	pub := &priv.PublicKey
// 	ciphertxt, err := pub.Encrypt(msg)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("加密结果:%x\n", ciphertxt)
// 	plaintxt, err := priv.Decrypt(ciphertxt)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if !bytes.Equal(msg, plaintxt) {
// 		log.Fatal("原文不匹配")
// 	}

// 	r, s, err := sm2.Sign(priv, msg)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	isok := sm2.Verify(pub, msg, r, s)
// 	fmt.Printf("Verified: %v\n", isok)
// 	plaintext := []byte("床前明月光，疑是地上霜，举头望明月，学习go语言") //明文
// 	fmt.Println("明文为：", string(plaintext))

// 	//传入明文和自己定义的密钥，密钥为8字节，如果不足8字节函数内部自动补全，超过8字节函数内部截取
// 	cryptText, err1 := goEncrypt.DesCbcEncrypt(plaintext, []byte("asd12345")) //得到密文
// 	if err1 != nil {
// 		return
// 	}
// 	fmt.Println("DES的CBC模式加密后的密文为:", base64.StdEncoding.EncodeToString(cryptText))

// 	//传入密文和自己定义的密钥，需要和加密的密钥一样，不一样会报错，8字节，如果不足8字节函数内部自动补全，超过8字节函数内部截取
// 	newplaintext, err2 := goEncrypt.DesCbcDecrypt(cryptText, []byte("asd12345")) //解密得到密文
// 	if err2 != nil {
// 		return
// 	}
// 	fmt.Println("DES的CBC模式解密完：", string(newplaintext))
// 	cryptText, err3 := goEncrypt.TripleDesEncrypt(plaintext, []byte("wumansgy12345678asdfghjk"))
// 	if err3 != nil {
// 		return
// 	}
// 	fmt.Println("三重DES的CBC模式加密后的密文为:", base64.StdEncoding.EncodeToString(cryptText))

// 	//传入密文和自己定义的密钥，需要和加密的密钥一样，不一样会报错，24字节，如果不足24字节函数内部自动补全，超过24字节函数内部截取
// 	newplaintext, err4 := goEncrypt.TripleDesDecrypt(cryptText, []byte("wumansgy12345678asdfghjk"))
// 	if err4 != nil {
// 		return
// 	}
// 	fmt.Println("三重DES的CBC模式解密：", string(newplaintext))
// 	//传入明文和自己定义的密钥，密钥为16字节，如果不足16字节函数内部自动补全，超过16字节函数内部截取
// 	cryptText, err5 := goEncrypt.AesCtrEncrypt(plaintext, []byte("626262"))
// 	if err5 != nil {
// 		return
// 	}
// 	fmt.Println("AES的CTR模式加密后的密文为:", base64.StdEncoding.EncodeToString(cryptText))

// 	//传入密文和自己定义的密钥，需要和加密的密钥一样，不一样会报错，16字节，如果不足16字节函数内部自动补全，超过16字节函数内部截取
// 	newplaintext, err6 := goEncrypt.AesCtrDecrypt(cryptText, []byte("626262"))
// 	if err6 != nil {
// 		return
// 	}
// 	fmt.Println("AES的CTR模式解密完：", string(newplaintext))
// }

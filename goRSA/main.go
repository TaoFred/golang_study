package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {

	qrMap := map[string]any{
		"manufacturer_code": "000015",
		"device_type_code":  "0211",
		"app_name":          "PRIDE-ControlValveDiag",
		"app_version":       "1.0.beta",
		"protocol":          "HART",
	}
	plaintext, _ := json.Marshal(qrMap)
	ciphertext := EncryptRSA(plaintext)
	fmt.Printf("ciphertext: %v\n", ciphertext)
	decryptedByte := DecryptRSA(ciphertext)
	fmt.Printf("decryptedStr: %v\n", string(decryptedByte))

	var resultMap map[string]any
	if err := json.Unmarshal(decryptedByte, &resultMap); err != nil {
		fmt.Println("Failed to unmarshal data")
		return
	}
	fmt.Printf("resultMap: %v\n", resultMap)

}

func EncryptRSA(plaintext []byte) string {
	pubKey := GetPubKey()
	rsaSize := pubKey.Size() - 11
	var ciphertext []byte
	idx := 0
	for {
		var chunk []byte
		var err error
		if idx+rsaSize < len(plaintext) {
			temp := plaintext[idx : idx+rsaSize]
			chunk, err = rsa.EncryptPKCS1v15(rand.Reader, pubKey, temp)
		} else {
			chunk, err = rsa.EncryptPKCS1v15(rand.Reader, pubKey, plaintext[idx:])

		}
		if err != nil {
			fmt.Println("Failed to encode plaintext")
			return ""
		}
		ciphertext = append(ciphertext, chunk...)
		idx += rsaSize
		if idx >= len(plaintext) {
			break
		}
	}

	return base64.StdEncoding.EncodeToString(ciphertext)
}

func DecryptRSA(cipherString string) []byte {
	rsaPrivateKey := GetPrivateKey()

	cipherBytes, err := base64.StdEncoding.DecodeString(cipherString)
	if err != nil {
		fmt.Println("Failed to Decode base64: ", err.Error())
		return nil
	}

	// 整体解密
	plainBytes, err := rsa.DecryptPKCS1v15(rand.Reader, rsaPrivateKey, cipherBytes)
	fmt.Println("整体解密", string(plainBytes), err)

	// 分段解密
	defaultLength := 128
	var decryptedData []byte
	for offset := 0; offset < len(cipherBytes); offset += defaultLength {
		var temp []byte
		var err error
		if offset+defaultLength <= len(cipherBytes) {
			temp, err = rsa.DecryptPKCS1v15(rand.Reader, rsaPrivateKey, cipherBytes[offset:offset+defaultLength])
		} else {
			temp, err = rsa.DecryptPKCS1v15(rand.Reader, rsaPrivateKey, cipherBytes[offset:])
		}
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		decryptedData = append(decryptedData, temp...)
	}
	fmt.Println("分段解密：", string(decryptedData))

	return decryptedData
}

func GetPubKey() *rsa.PublicKey {
	pubKeyBytes, err := os.ReadFile("./public_key.pem")
	if err != nil {
		fmt.Println("Failed to read public key file")
		return nil
	}

	// 解析公钥
	pubBlock, _ := pem.Decode(pubKeyBytes)
	if pubBlock == nil {
		fmt.Println("Failed to decode public key")
		return nil
	}
	pubKey, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if err == nil {
		if rsaPubKey, ok := pubKey.(*rsa.PublicKey); ok {
			return rsaPubKey
		}
	}
	rasPubKey, err := x509.ParsePKCS1PublicKey(pubBlock.Bytes)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return rasPubKey
}

func GetPrivateKey() *rsa.PrivateKey {
	privateBytes, err := os.ReadFile("./private_key.pem")
	if err != nil {
		fmt.Println("Failed to read private key file")
		return nil
	}
	// 解析私钥
	privateBlock, _ := pem.Decode(privateBytes)
	if privateBlock == nil {
		fmt.Println("Failed to decode private key")
		return nil
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(privateBlock.Bytes)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	if rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey); ok {
		// pubKeyBytes, err := x509.MarshalPKIXPublicKey(&rsaPrivateKey.PublicKey)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		// 	Type:  "RSA PUBLIC KEY",
		// 	Bytes: pubKeyBytes,
		// })
		// fmt.Println("Public Key:\n", string(publicKeyPEM))
		return rsaPrivateKey
	}
	return nil
}

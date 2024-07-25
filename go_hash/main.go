package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	HashFile()
	HashStr()
}

func HashFile() {
	filePath := "D:\\Integrity\\业务理解\\文档仓库\\disk.zip"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		log.Fatal(err)
	}

	hashValue := hash.Sum(nil)
	fmt.Println(hashValue)
	fmt.Println(len(hashValue))
	fmt.Printf("File Hash (SHA256): %x\n", hashValue)
	fmt.Printf("hex.EncodeToString(hashValue): %v\n", hex.EncodeToString(hashValue))
}

func HashStr() {
	str := "hello, world"
	hashValue := sha256.Sum256([]byte(str))
	fmt.Println(hashValue)
	fmt.Println(len(hashValue))
	fmt.Printf("File Hash (SHA256): %x\n", hashValue)
}

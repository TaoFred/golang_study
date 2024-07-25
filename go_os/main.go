package main

import (
	"fmt"
	"os"
)

func main() {
	// GetEnv()
	fmt.Println("MINIO_ROOT_USER: ", os.Getenv("MINIO_ROOT_USER"))
	fmt.Println("SUPOS_APP_TENANT_ID", os.Getenv("SUPOS_APP_TENANT_ID"))
}

func GetEnv() {
	fmt.Println(os.Getenv("GOPATH"))
	fmt.Println(os.Getenv("GOROOT"))
	fmt.Println(os.Getenv("GOBIN"))
	fmt.Println(os.Getenv("PATH"))
	fmt.Println(os.Getenv("HOME"))
	fmt.Println(os.Getenv("USER"))
	fmt.Println(os.Getenv("PWD"))
	fmt.Println(os.Getenv("SHELL"))
	fmt.Println(os.Getenv("TMPDIR"))
}

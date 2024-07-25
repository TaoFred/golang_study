package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

// Service 结构体用于解析XML中的service元素
type Service struct {
	ID          string `xml:"id"`
	Name        string `xml:"name"`
	Description string `xml:"description"`
	Env         []Env  `xml:"env"`
	Executable  string `xml:"executable"`
	Arguments   string `xml:"arguments"`
	Log         Log    `xml:"log"`
}

// Env 结构体用于解析XML中的env元素
type Env struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

// Log 结构体用于解析XML中的log元素
type Log struct {
	Mode                string `xml:"mode,attr"`
	SizeThreshold       int    `xml:"sizeThreshold"`
	Pattern             string `xml:"pattern"`
	AutoRollAtTime      string `xml:"autoRollAtTime"`
	ZipOlderThanNumDays int    `xml:"zipOlderThanNumDays"`
	ZipDateFormat       string `xml:"zipDateFormat"`
}

func main() {
	// 读取XML文件
	xmlFile, err := os.Open("minio-server.xml")
	if err != nil {
		fmt.Println("无法打开XML文件:", err)
		return
	}
	defer xmlFile.Close()

	// 读取文件内容
	xmlData, err := io.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("无法读取XML内容:", err)
		return
	}

	// 创建Service对象
	var service Service

	// 解析XML内容
	err = xml.Unmarshal(xmlData, &service)
	if err != nil {
		fmt.Println("XML解析失败:", err)
		return
	}

	// 打印解析得到的结果
	fmt.Println("ID:", service.ID)
	fmt.Println("Name:", service.Name)
	fmt.Println("Description:", service.Description)

	fmt.Println("Environment Variables:")
	for _, env := range service.Env {
		fmt.Printf("  Name: %s, Value: %s\n", env.Name, env.Value)
	}

	fmt.Println("Executable:", service.Executable)
	fmt.Println("Arguments:", service.Arguments)

	fmt.Println("Log Mode:", service.Log.Mode)
	fmt.Println("Size Threshold:", service.Log.SizeThreshold)
	fmt.Println("Pattern:", service.Log.Pattern)
	fmt.Println("Auto Roll At Time:", service.Log.AutoRollAtTime)
	fmt.Println("Zip Older Than Num Days:", service.Log.ZipOlderThanNumDays)
	fmt.Println("Zip Date Format:", service.Log.ZipDateFormat)
}

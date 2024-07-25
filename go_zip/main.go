package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	GetFiles()
}

type FileInfo struct {
	Name           string
	Path           string
	DestinationDir string
}

func GetFiles() {

	files := []FileInfo{
		{Name: "备用资源表格模块化显示.png", Path: "in/备用资源表格模块化显示.png", DestinationDir: "aa/bc"},
		{Name: "备用资源模块化显示初步设计.pdf", Path: "in/备用资源模块化显示初步设计.pdf", DestinationDir: "aa"},
		{Name: "中心主题.xmind", Path: "in/中心主题.xmind", DestinationDir: ""},
		{Name: "中心主题.xmind", Path: "in/中心主题.xmind", DestinationDir: "1"},
	}

	outputPath := "out/output.zip"

	err := createZip(files, outputPath)
	if err != nil {
		fmt.Println("Failed to create zip file:", err)
		return
	}

	fmt.Println("Zip file created successfully:", outputPath)

}

func createZip(files []FileInfo, outputPath string) error {
	zipFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}

	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()
	for _, file := range files {
		srcFile, err := os.Open(file.Path)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		// info, err := srcFile.Stat()
		// if err != nil {
		// 	return err
		// }

		// header, err := zip.FileInfoHeader(info)
		// if err != nil {
		// 	return err
		// }
		// // Set desired path inside the zip
		// header.Name = filepath.Join(file.DestinationDir, file.Name)

		writer, err := zipWriter.Create(filepath.Join(file.DestinationDir, file.Name))
		if err != nil {
			return err
		}

		_, err = io.Copy(writer, srcFile)
		if err != nil {
			return err
		}
	}
	return nil
}

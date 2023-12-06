package main

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
    // "fmt"
)

// 壓縮資料夾
func ZipFolder(folderPath string, zipFilePath string) error {
	zipFile, err := os.Create(zipFilePath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// 遍歷資料夾
	filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
        // fmt.Println(info)
		// 如果是資料夾，則跳過 // 因為遍歷會包含他本身這個資料夾，所以需要排除
		if info.IsDir() {
			return nil
		}

		// 打開文件
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// 建立 zip 內的文件結構
		relPath := strings.TrimPrefix(path, folderPath)
		zipFileWriter, err := zipWriter.Create(relPath)
		if err != nil {
			return err
		}

		// 複製文件內容到 zip 檔案中
		_, err = io.Copy(zipFileWriter, file)
		return err
	})

	return nil
}

func main() {
	folderPath := "./folder_to_zip" // 要壓縮的資料夾路徑
	zipFilePath := "./result.zip"   // 壓縮後的 zip 檔案路徑

	err := ZipFolder(folderPath, zipFilePath)
	if err != nil {
		panic(err)
	}
}

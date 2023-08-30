package util

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func AddFileToZip(zipWriter *zip.Writer, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 获取文件信息
	info, err := file.Stat()
	if err != nil {
		return err
	}

	// 创建一个新的zip文件条目
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}
	header.Name = filepath.Base(filePath)
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	// 将文件内容复制到zip文件中
	_, err = io.Copy(writer, file)
	return err
}

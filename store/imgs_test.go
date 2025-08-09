package store

import (
	"encoding/base64"
	"io"
	"log"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

// 将本地图片转换为base64
func Base64Img(imgPath string) (string, error) {
	img, err := os.Open(imgPath)
	if err != nil {
		return "", err
	}
	defer img.Close()
	imgBytes, err := io.ReadAll(img)
	if err != nil {
		return "", err
	}
	imgBase64 := base64.StdEncoding.EncodeToString(imgBytes)
	return imgBase64, nil
}
func TestBase64Img(t *testing.T) {
	imgDir := "./imgs"
	files, err := os.ReadDir(imgDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		imgPath := imgDir + "/" + file.Name()
		logrus.Info(imgPath)
		imgBase64, err := Base64Img(imgPath)
		if err != nil {
			log.Fatal(err)
			return
		}

		img := NewImg()
		// base64 前缀
		img.Base64 = "data:image/png;base64," + imgBase64
		err = img.CreateImg()
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}

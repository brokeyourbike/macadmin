package helpers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/groob/plist"
	"github.com/schollz/progressbar/v3"
)

type Ipsw struct {
	ProductTypes map[string]ProductType `plist:"MobileDeviceProductTypes"`
}

type ProductType struct {
	Name string
	Id   string
}

// type ProductTypes map[string]ProductType

// GetIpswData return data from com_apple_macOSIPSW.xml (which is actually a plist)
func FetchIpswData(feed string) (Ipsw, error) {
	filePath, err := DownloadFile(feed, "/tmp")
	if err != nil {
		return Ipsw{}, err
	}

	ipsw, err := ParseIpswFile(filePath)
	return ipsw, err
}

func ParseIpswFile(path string) (Ipsw, error) {
	content, err := os.Open(path)
	if err != nil {
		return Ipsw{}, err
	}

	var data Ipsw
	err = plist.NewXMLDecoder(content).Decode(&data)

	return data, err
}

func DownloadFile(url, outputDir string) (string, error) {
	savePath := path.Join(outputDir, path.Base(url))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return savePath, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return savePath, err
	}

	defer resp.Body.Close()

	f, err := os.OpenFile(savePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return savePath, err
	}

	defer f.Close()

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		fmt.Sprintf("Downloading: %s", url),
	)
	_, err = io.Copy(io.MultiWriter(f, bar), resp.Body)

	return savePath, err
}

package utils

import (
	"fmt"
	"github.com/fzxiao233/Go-Emoji-Utils"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var client *http.Client

func init() {
	client = createClient()
}

func createClient() *http.Client {
	if Config.EnableProxy == true {
		client = createSOCKS5Proxy()
	} else {
		client = http.DefaultClient
	}
	return client
}

func HttpGet(url string, header map[string]string) ([]byte, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64; rv:60.0) Gecko/20100101 Firefox/60.0")
	req.Header.Set("Accept-Language", "en-US,en;q=0.8")
	for k, v := range header {
		req.Header.Set(k, v)
	}
	res, err := client.Do(req)
	if res != nil {
		defer func() {
			err := res.Body.Close()
			if err != nil {
				return
			}
		}()
	}
	if err != nil {
		err = fmt.Errorf("HttpGet error %w", err)
		log.Print(err)
		return []byte{}, err
	}
	htmlBody, _ := ioutil.ReadAll(res.Body)
	return htmlBody, nil
}
func createSOCKS5Proxy() *http.Client {
	proxyUrl, _ := url.Parse("socks5://" + Config.Proxy)
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}

	//adding the Transport object to the http Client
	client := &http.Client{
		Transport: transport,
		Timeout:   60 * time.Second,
	}
	return client
}
func IsFileExist(aFilepath string) bool {
	if _, err := os.Stat(aFilepath); err == nil {
		return true
	} else {
		return false
	}
}
func GenerateFilepath(UserName string, VideoTitle string) string {
	pathSlice := []string{GenerateDownloadDir(UserName), RemoveIllegalChar(VideoTitle)}
	aFilepath := strings.Join(pathSlice, "/")
	if IsFileExist(aFilepath) {
		return ChangeName(aFilepath)
	} else {
		return aFilepath
	}
}
func GenerateDownloadDir(UserName string) string {
	dirPath := Config.DownloadDir + "/" + UserName
	if !IsFileExist(dirPath) {
		err := os.Mkdir(dirPath, 0775)
		if err != nil {
			log.Fatalf("mkdir error: %s", dirPath)
		}
	}
	return dirPath
}
func ChangeName(aFilepath string) string {
	dir, file := filepath.Split(aFilepath)
	ext := path.Ext(file)
	filename := strings.TrimSuffix(path.Base(file), ext)
	filename += strconv.FormatInt(time.Now().Unix(), 10)
	return dir + filename + ext
}
func GetTimeNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
func RemoveIllegalChar(Title string) string {
	illegalChars := []string{"|", "/", "\\", ":", "?"}
	Title = emoji.RemoveAll(Title)
	for _, char := range illegalChars {
		Title = strings.ReplaceAll(Title, char, "#")
	}
	return Title
}

func I2b(i int) bool {
	if i != 0 {
		return true
	} else {
		return false
	}
}

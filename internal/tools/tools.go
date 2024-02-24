package tools

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/ShinonomeSetsuna/Pancake50/internal/types"
)

var authorization string

func GetResource(url string) []byte {
	split := strings.Split(url, "/")
	fn := split[len(split)-1]
	path := "./temp/score/"
	if strings.HasSuffix(fn, ".webp") {
		path = "./temp/cover/"
	}
	var result []byte
	if fb, err := os.Open(path + fn); err == nil {
		// 读取本地文件
		result, _ = io.ReadAll(fb)
	} else {
		// 从网络获取
		if strings.HasSuffix(url, ".webp") { // webp补充后缀
			url += "-thumbnail.large"
		}
		client := &http.Client{}
		if req, err := http.NewRequest("GET", url, nil); err != nil {
			log.Fatalln("创建请求失败: ", err)
		} else {
			fmt.Println("开始获取文件：", url)
			req.Header.Set("Authorization", GetAuthorization())
			if resp, err := client.Do(req); err != nil {
				log.Fatalln("请求失败: ", err)
			} else {
				if res, err := io.ReadAll(resp.Body); err != nil {
					log.Fatalln("读取失败: ", err)
				} else {
					result = res
				}
			}
		}
		var check types.NetRespone
		if err := json.Unmarshal(result, &check); err == nil {
			if check.Code != "ok" {
				log.Fatalln("出现错误：", check.Message)
			}
		}
		os.Mkdir("./temp", fs.ModeDir)
		os.Mkdir(path, fs.ModeDir)
		file, _ := os.Create(path + fn)
		defer file.Close()
		file.Write(result)
	}
	return result
}

func GetAuthorization() string {
	if authorization == "" {
		if info, err := os.ReadFile("./Authorization.txt"); err != nil {
			if file, err := os.Create("./Authorization.txt"); err != nil {
				log.Fatalln("创建文件失败：", err)
			} else {
				fmt.Println("成功创建文件，请编辑后再次启动")
				defer file.Close()
				os.Exit(0)
			}

		} else {
			authorization = string(info)
		}
	}
	return authorization
}

package newWorld

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

// CheckIn 执行签到操作的函数
func (l *NewWorld) CheckIn() error {
	// 创建签到请求
	url := "https://neworld.space/user/checkin"

	// 创建一个空的请求体
	payload := []byte{}

	// 创建POST请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	// 设置请求头
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Length", "0")
	req.Header.Set("Cookie", l.Coookie)
	req.Header.Set("Origin", "https://neworld.space")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Sec-Ch-Ua", "\"Google Chrome\";v=\"117\", \"Not;A=Brand\";v=\"8\", \"Chromium\";v=\"117\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"macOS\"")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	// 创建HTTP客户端
	client := &http.Client{
		Timeout: time.Second * 10, // 设置超时时间
	}

	// 发送请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	// 打印响应内容
	println("签到响应内容:", string(body))
	return nil
}

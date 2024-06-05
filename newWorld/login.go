package newWorld

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// PerformLogin 执行登录操作并返回cookie
func (l *NewWorld) PerformLogin() error {
	// 构建请求体
	payload := url.Values{}
	payload.Set("email", l.Email)
	payload.Set("passwd", l.Passwd)

	// 创建POST请求
	url := "https://neworld.space/auth/login"
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(payload.Encode()))
	if err != nil {
		return err
	}

	// 设置请求头
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Origin", "https://neworld.space")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Referer", "https://neworld.space/auth/login")
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
		Timeout: time.Second * 10,
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to login, status code: %d", resp.StatusCode)
	}
	// 获取cookie
	cookies := resp.Cookies()

	cookiesStr := createCookieHeader(cookies)
	l.Coookie = cookiesStr
	return nil
}

// createCookieHeader 将cookie切片转换为请求头所需的字符串格式
func createCookieHeader(cookies []*http.Cookie) string {
	cookieStrings := make([]string, len(cookies))
	for i, cookie := range cookies {
		cookieStrings[i] = cookie.Name + "=" + cookie.Value
	}
	return strings.Join(cookieStrings, "; ")
}

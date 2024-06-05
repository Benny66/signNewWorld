package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// AccountConfig 存储单个账号的配置
type AccountConfig struct {
	Cron  string `json:"cron"`  // 定时任务表达式 比如：0 55 8 * * *
	Delay int    `json:"delay"` // 随机延迟时间，单位秒，防止被ban， 默认240s
	User  string `json:"user"`  // 账号用户名
	Pass  string `json:"pass"`  // 账号密码
}

// LoadDatabaseConfigs 从环境变量中加载多个数据库配置
func LoadAccountConfigs() ([]AccountConfig, error) {
	accountConfigs := make([]AccountConfig, 0)

	// 加载.env文件
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return nil, err
	}

	// 获取所有数据库前缀的环境变量
	prefixesStr := os.Getenv("ACCOUNT_PREFIXES")
	if prefixesStr == "" {
		fmt.Println("ACCOUNT_PREFIXES not set in environment variables")
		return nil, errors.New("没有账号")
	}

	// 将字符串分割成前缀数组
	prefixes := strings.Split(prefixesStr, ",")

	for _, prefix := range prefixes {
		accountConfig := AccountConfig{
			User: os.Getenv(prefix + "_USER"),
			Pass: os.Getenv(prefix + "_PASS"),
		}
		cronStr := os.Getenv(prefix + "_CRON")
		delay := os.Getenv(prefix + "_DELAY")
		delayInt, err := strconv.Atoi(delay)
		if err != nil {
			fmt.Println("Invalid delay value: ", delay, err)
			return nil, err
		}
		if delayInt <= 0 {
			delayInt = 240 // 默认延迟时间为240秒
		} else if delayInt > 600 {
			delayInt = 600 // 最大延迟时间为240秒
		}
		accountConfig.Cron = cronStr
		accountConfig.Delay = delayInt
		accountConfigs = append(accountConfigs, accountConfig)
	}

	return accountConfigs, nil
}

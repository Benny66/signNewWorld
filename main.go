package main

import (
	"fmt"
	"log"
	"math/rand"
	"signNewWorld/config"
	"signNewWorld/newWorld"
	"time"

	"github.com/robfig/cron/v3"
)

var accounts []config.AccountConfig = make([]config.AccountConfig, 0)
var err error

// CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
func main() {
	fmt.Println("程序启动...")
	c := cron.New(cron.WithSeconds()) // 允许秒级精度
	// 加载账号配置文件
	accounts, err = config.LoadAccountConfigs()
	if err != nil {
		log.Fatal(err)
	}
	for _, account := range accounts {
		fmt.Printf("账号：%s, 定时 %s \n", account.User, account.Cron)
		// 在每次迭代中初始化一个新变量
		localAccount := account // 创建一个新的变量，复制当前迭代的 account 值
		// 添加任务，使用cron的语法
		_, err := c.AddFunc(account.Cron, func() {
			performTask(localAccount)
		})
		if err != nil {
			log.Fatal(err)
		}
	}

	// 启动cron调度器
	c.Start()

	// 阻塞主线程，直到程序被中断
	select {}
}

var n int = 1

// 执行任务
func performTask(account config.AccountConfig) {
	fmt.Printf("第%d次签到，准备登录账号：%s，时间：%s\n", n, account.User, time.Now().Format("2006-01-02 15:04:05"))
	// 添加一个随机延迟，防止被ban
	rand.New(rand.NewSource(time.Now().UnixNano()))
	delay := rand.Intn(account.Delay) + 1
	fmt.Printf("等待 %d 秒...\n", delay)
	time.Sleep(time.Duration(delay) * time.Second) // 暂停相应的秒数
	fmt.Println("开始登录...")
	// 创建登录实例
	newWorld := newWorld.NewNewWorld(account.User, account.Pass)
	// 执行登录并获取cookie
	err := newWorld.PerformLogin()
	if err != nil {
		fmt.Println("Login failed:", err)
		return
	}
	fmt.Println("登录成功...")
	// 执行签到
	// 添加一个随机延迟，防止被ban
	fmt.Println("开始签到...")
	rand.New(rand.NewSource(time.Now().UnixNano()))
	delay = rand.Intn(10) + 1                      // 生成1到5之间的随机整数
	time.Sleep(time.Duration(delay) * time.Second) // 暂停相应的秒数
	err = newWorld.CheckIn()
	if err != nil {
		fmt.Println("CheckIn failed:", err)
		return
	}
	fmt.Println("签到完成")
	n++
}

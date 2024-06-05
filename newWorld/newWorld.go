package newWorld

// Login 结构体用于存储登录所需的信息
type NewWorld struct {
	Email   string
	Passwd  string
	Coookie string
}

// NewLogin 创建并返回一个新的Login实例
func NewNewWorld(email, passwd string) *NewWorld {
	return &NewWorld{
		Email:  email,
		Passwd: passwd,
	}
}

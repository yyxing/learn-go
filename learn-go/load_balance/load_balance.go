package load_balance

type LoadBalance interface {
	// 添加服务
	Add(...string) error
	// 获取服务
	Get(string) (string, error)
	// 动态修改
	Update()
}

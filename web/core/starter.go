package core

// 启动器接口 类似Spring的基础Bean在系统启动时
type Starter interface {
	// init上下文所需资源 配置文件的加载
	Init(ApplicationContext)
	// 根据初始化的资源，将一些全局的连接，连接池启动运行 类似mysql redis mq
	Start(ApplicationContext)
	// 销毁 资源的释放
	Finalize(ApplicationContext)
	// 获取Starter启动顺序
	GetOrder() int
}

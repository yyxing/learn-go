package types

// 每个请求的地址，以及对url的处理函数
type Request struct {
	// 请求路径
	Url string
	// 请求方法
	Method string
	// 请求体 非get请求使用
	Body []byte
	// 对每个url的处理接口 输入数据为url获取到的页面数据
	ParseFunc func([]byte) ParseResult
}

// 每个请求的处理请求，包含请求得到的数据以及从url中获取到的下一个请求
type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

// 空处理
func NilParseFunc(content []byte) ParseResult {
	return ParseResult{}
}

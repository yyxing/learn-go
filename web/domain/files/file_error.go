package files

// 账户相关业务异常
type fileError struct {
	s string
}

func (e fileError) Error() string {
	return e.s
}

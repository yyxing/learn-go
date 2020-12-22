package accounts

// 账户相关业务异常
type accountError struct {
	s string
}

func (e accountError) Error() string {
	return e.s
}

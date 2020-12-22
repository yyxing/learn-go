package envelopes

// 账户相关业务异常
type envelopeError struct {
	s string
}

func (e envelopeError) Error() string {
	return e.s
}

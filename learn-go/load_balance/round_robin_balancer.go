package load_balance

import (
	"errors"
)

type RoundRobinBalancer struct {
	curIndex int
	rs       []string
}

func (lb *RoundRobinBalancer) Add(param ...string) error {
	if len(param) == 0 {
		return errors.New("param len 1 at least")
	}
	addr := param[0]
	// TODO http校验
	lb.rs = append(lb.rs, addr)
	return nil
}

func (lb *RoundRobinBalancer) Next() string {
	next := lb.rs[lb.curIndex]
	lb.curIndex = (lb.curIndex + 1) % len(lb.rs)
	return next
}
func (lb *RoundRobinBalancer) Get(url string) (string, error) {
	return lb.Next(), nil
}

func (lb *RoundRobinBalancer) Update() {

}

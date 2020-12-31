package load_balance

import (
	"errors"
	"math/rand"
	"time"
)

var (
	seed = rand.NewSource(time.Now().UnixNano())
)

// 随机负载均衡
type RandomLoadBalancer struct {
	rs       []string
	curIndex int
}

func (lb *RandomLoadBalancer) Add(param ...string) error {
	if len(param) == 0 {
		return errors.New("param len 1 at least")
	}
	addr := param[0]
	// TODO http校验
	lb.rs = append(lb.rs, addr)
	return nil
}

func (lb *RandomLoadBalancer) Next() string {
	r := rand.New(seed)
	return lb.rs[r.Intn(len(lb.rs))]
}
func (lb *RandomLoadBalancer) Get(url string) (string, error) {
	return lb.Next(), nil
}

func (lb *RandomLoadBalancer) Update() {

}

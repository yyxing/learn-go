package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Test(t *testing.T) {
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, i)
	}

	ss := s[:10]
	ss[2] = 111
	t.Log(s, ss)
}

func returnMultiValue() (int, int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	time.Sleep(1)
	return r.Intn(10), r.Intn(20)
}

func TestMulti(t *testing.T) {
	for i := 0; i < 10; i++ {
		value, i2 := returnMultiValue()
		t.Log(value, i2)
	}
}
func timer(inner func(op int) int) func(op int) int {
	return func(op int) int {
		start := time.Now()
		ret := inner(op)
		fmt.Printf("time cost %fs:", time.Since(start).Seconds())
		return ret
	}
}
func slowFunc(op int) int {
	time.Sleep(time.Second)
	return op + 1
}
func TestTimerFunc(t *testing.T) {
	timerFunc := timer(slowFunc)
	t.Log(timerFunc(100))
}

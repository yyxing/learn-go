package load_balance

import (
	"fmt"
	"strconv"
	"testing"
)

var (
	services = []string{
		"127.0.0.1:2000",
		"127.0.0.1:2001",
		"127.0.0.1:2002",
		"127.0.0.1:2003",
	}
)

func TestRandomLoadBalance(t *testing.T) {
	randomLoadBalance := RandomLoadBalancer{}
	for _, service := range services {
		randomLoadBalance.Add(service)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(randomLoadBalance.Get(""))
	}
}

func TestRoundRobinLoadBalance(t *testing.T) {
	roundRobinLoadBalance := RoundRobinBalancer{}
	for _, service := range services {
		roundRobinLoadBalance.Add(service)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(roundRobinLoadBalance.Get(""))
	}
}

func TestWeightRobinBalancer(t *testing.T) {
	weightRobinBalancer := WeightRobinBalancer{}
	weight := 4
	for _, service := range services {
		weightRobinBalancer.Add(service, strconv.Itoa(weight))
		weight--
	}
	for i := 0; i < 10; i++ {
		fmt.Println(weightRobinBalancer.Get(""))
	}
}
func TestConsistentHashBalancer(t *testing.T) {
	consistentHashBalancer := NewConsistentHashBalancer(10, nil)
	for _, service := range services {
		consistentHashBalancer.Add(service)
	}
	fmt.Println(consistentHashBalancer.Get("http://127.0.0.1:2002/base/getInfo"))
	fmt.Println(consistentHashBalancer.Get("http://127.0.0.1:2002/base/error"))
	fmt.Println(consistentHashBalancer.Get("http://127.0.0.1:2002/base/error"))
	fmt.Println(consistentHashBalancer.Get("http://127.0.0.1:2002/base/getInfo"))

	fmt.Println(consistentHashBalancer.Get("127.0.0.1"))
	fmt.Println(consistentHashBalancer.Get("192.168.0.65"))
	fmt.Println(consistentHashBalancer.Get("127.0.0.1"))
}

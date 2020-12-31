package load_balance

type LoadBalancerType uint8

const (
	RandomLB LoadBalancerType = iota
	WeightRobinLB
	RoundRobinLB
	ConsistentHashLB
)

func LoadBalanceFactory(lbType LoadBalancerType) LoadBalance {
	switch lbType {
	case RandomLB:
		return &RandomLoadBalancer{}
	case WeightRobinLB:
		return &WeightRobinBalancer{}
	case RoundRobinLB:
		return &RoundRobinBalancer{}
	case ConsistentHashLB:
		return NewConsistentHashBalancer(4, nil)
	default:
		return &RoundRobinBalancer{}
	}
}

package load_balance

import "strconv"

type WeightRobinBalancer struct {
	curIndex int
	rs       []*WeightNode
	rsw      []int64
}

type WeightNode struct {
	addr            string
	currentWeight   int64
	weight          int64
	effectiveWeight int64
}

func (lb *WeightRobinBalancer) Add(param ...string) error {
	addr := param[0]
	weight, err := strconv.ParseInt(param[1], 10, 64)
	if err != nil {
		return err
	}
	// TODO http校验
	lb.rsw = append(lb.rsw, weight)
	lb.rs = append(lb.rs, &WeightNode{
		addr:            addr,
		currentWeight:   0,
		weight:          weight,
		effectiveWeight: weight,
	})
	return nil
}

func (lb *WeightRobinBalancer) Next() string {
	var total int64
	var best *WeightNode
	for _, node := range lb.rs {
		// step 1. 累加所有服务器的total
		total += node.effectiveWeight
		// step 2. 变更节点的临时权重为当前的权重加上分配的权重
		node.currentWeight += node.effectiveWeight
		// step 3. 若当前通讯异常将节点的权重变小，若成功通讯则逐步恢复权重大小至weight
		if node.effectiveWeight < node.weight {
			node.effectiveWeight++
		}
		// step 4. 获取服务列表中权重最大的节点
		if best == nil || node.currentWeight > best.currentWeight {
			best = node
		}
	}
	if best == nil {
		return ""
	}
	// step 5. 选中节点减去节点权重和 使其被选中的几率降低
	best.currentWeight -= total
	return best.addr
}
func (lb *WeightRobinBalancer) Get(url string) (string, error) {
	return lb.Next(), nil
}

func (lb *WeightRobinBalancer) Update() {

}

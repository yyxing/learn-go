package load_balance

import (
	"errors"
	"fmt"
	"hash/crc32"
	"sort"
	"sync"
)

type Hash func(data []byte) uint32

type Uint32Slice []uint32

func (s Uint32Slice) Len() int {
	return len(s)
}

func (s Uint32Slice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Uint32Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type ConsistentHashBalancer struct {
	hash     Hash              // 哈希算法
	replicas int               // 虚拟节点复制因子
	keys     Uint32Slice       // 排序的节点切片
	hashMap  map[uint32]string // key节点hash value:节点key(节点地址)

	lock sync.RWMutex
}

func NewConsistentHashBalancer(replicas int, fn Hash) *ConsistentHashBalancer {
	if fn == nil {
		fn = crc32.ChecksumIEEE
	}
	return &ConsistentHashBalancer{
		hash:     fn,
		replicas: replicas,
		hashMap:  make(map[uint32]string),
	}
}
func (lb *ConsistentHashBalancer) Add(param ...string) error {
	if len(param) == 0 {
		return errors.New("param len 1 at least")
	}
	// 服务名称或者ip
	addr := param[0]
	lb.lock.Lock()
	defer lb.lock.Unlock()
	// TODO http校验
	// 添加虚拟节点
	for i := 0; i < lb.replicas; i++ {
		// 获取虚拟节点的hash
		hash := lb.hash([]byte(fmt.Sprintf("%d#%s", i+1, addr)))
		// 添加到hash keys节点中
		lb.keys = append(lb.keys, hash)
		// 将hash与服务地址对应
		lb.hashMap[hash] = addr
	}
	sort.Sort(lb.keys)
	return nil
}

func (lb *ConsistentHashBalancer) isEmpty() bool {
	return len(lb.keys) == 0
}

func (lb *ConsistentHashBalancer) Get(key string) (string, error) {
	if lb.isEmpty() {
		return "", errors.New("node is empty")
	}
	hash := lb.hash([]byte(key))

	idx := sort.Search(len(lb.keys), func(i int) bool {
		return lb.keys[i] >= hash
	})
	if idx == len(lb.keys) {
		idx = 0
	}
	lb.lock.RLock()
	defer lb.lock.RUnlock()
	return lb.hashMap[lb.keys[idx]], nil
}

func (lb *ConsistentHashBalancer) Update() {

}

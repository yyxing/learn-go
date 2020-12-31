package router

// Trie树节点
type node struct {
	pattern  string
	part     string
	children []*node
	isWild   bool
}

// /test/:name/hello
func (n *node) match(part string) {

}

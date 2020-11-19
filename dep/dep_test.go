package dep

import (
	"fmt"
	"testing"
)

func TestBlock(t *testing.T) {
	s := "中华人民共和国!@"
	for i, str := range []rune(s) {
		fmt.Printf("%d %c\n", i, str)
	}
}

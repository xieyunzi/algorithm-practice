package golang

import (
	// "fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsert(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	tree := &Tree{}
	for _, v := range values {
		tree.Insert(v)
	}

	assert.Equal(t, true, tree.Find(1))
	assert.Equal(t, false, tree.Find(0))
	// tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Key, " ") })
	// fmt.Println()
	assert.Equal(t, true, tree.Delete(5))
	// tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Key, " ") })
	// fmt.Println()
}

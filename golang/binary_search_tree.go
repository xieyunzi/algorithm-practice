package golang

type Tree struct {
	Root *Node
}

type Node struct {
	Left, Right *Node
	Key         int
}

func (n *Node) Insert(key int) bool {
	if n == nil {
		return false
	}

	switch {
	case key == n.Key:
		return false
	case key < n.Key:
		if n.Left == nil {
			n.Left = &Node{Key: key}
			return true
		}

		return n.Left.Insert(key)
	case key > n.Key:
		if n.Right == nil {
			n.Right = &Node{Key: key}
			return true
		}

		return n.Right.Insert(key)
	}

	return false
}

func (n *Node) findMax(parent *Node) (*Node, *Node) {
	if n.Right == nil {
		return parent, n
	}
	return n.Right.findMax(n)
}

func (n *Node) replaceNode(parent, replacement *Node) bool {
	if n == nil {
		return false
	}

	if n == parent.Left {
		parent.Left = replacement
		return true
	}
	if n == parent.Right {
		parent.Right = replacement
		return true
	}

	return false
}

func (n *Node) Delete(parent *Node, key int) bool {
	if n == nil {
		return false
	}

	switch {
	case key == n.Key:
		if n.Left == nil && n.Right == nil {
			return n.replaceNode(parent, nil)
		}
		if n.Left == nil {
			return n.replaceNode(parent, n.Right)
		}
		if n.Right == nil {
			return n.replaceNode(parent, n.Left)
		}

		replParent, replacement := n.Left.findMax(n)
		n.Key = replacement.Key
		return replacement.Delete(replParent, replacement.Key)
	case key < n.Key:
		return n.Left.Delete(n, key)
	case key > n.Key:
		return n.Right.Delete(n, key)
	}

	return false
}

func (n *Node) Find(key int) bool {
	if n == nil {
		return false
	}

	switch {
	case key == n.Key:
		return true
	case key < n.Key:
		return n.Left.Find(key)
	case key > n.Key:
		return n.Right.Find(key)
	}

	return false
}

func (t *Tree) Insert(key int) bool {
	if t.Root == nil {
		t.Root = &Node{Key: key}
	}

	return t.Root.Insert(key)
}

func (t *Tree) Find(key int) bool {
	if t.Root == nil {
		return false
	}

	return t.Root.Find(key)
}

func (t *Tree) Delete(key int) bool {
	if t.Root == nil {
		return false
	}

	fakeParent := &Node{Right: t.Root}
	if t.Root.Delete(fakeParent, key) {
		if fakeParent.Right == nil {
			t.Root = nil
		}

		return true
	}

	return false
}

func (t *Tree) Traverse(n *Node, f func(*Node)) {
	if n == nil {
		return
	}

	t.Traverse(n.Left, f)
	f(n)
	t.Traverse(n.Right, f)
}

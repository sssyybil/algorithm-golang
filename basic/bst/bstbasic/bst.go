package bstbasic

type Key any
type Value any

type Node struct {
	Key         Key
	Value       Value
	Left, Right *Node
}

type BST struct {
	Root  *Node
	Count int
}

func (bst BST) size() (count int) {
	return bst.Count
}

func (bst BST) isEmpty() bool {
	return bst.Count == 0
}

// Insert 向以 Node 为根的二分搜索树中，插入节点（key,value）
func (bst BST) Insert(key Key, value Value) {
}

//func insert(node *Node, key Key, value Value) {
//	if key == node.Key {
//		node.Value = value
//	} else if key < node.Key {
//		node.Left = insert(node.Left, key, value)
//	}
//}

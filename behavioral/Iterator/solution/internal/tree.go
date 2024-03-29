package internal

// Tree Iterator

type TreeNode struct {
	val      TransferInterface
	children []TreeNode
}

type TreePointer struct {
	node *TreeNode
}

func (t TreePointer) toArray(node *TreeNode) []TransferInterface {
	if node == nil {
		return nil
	}

	arr := []TransferInterface{node.val}

	for i := range node.children {
		arr = append(arr, t.toArray(&node.children[i])...)
	}

	return arr
}

func (t TreePointer) toLinkedList(node *TreeNode, lNode *LinkedNode) *LinkedNode {
	if node == nil {
		return nil
	}

	lNode = &LinkedNode{
		val:  node.val,
		next: lNode,
	}

	for i := range node.children {
		lNode = t.toLinkedList(&node.children[i], lNode)
	}

	return lNode
}

func (t TreePointer) Iterator() TransferIterator {
	return NewImplementLinkedList(t.toLinkedList(t.node, nil))
}

func NewImplementTree(node *TreeNode) TreePointer {
	return TreePointer{
		node: node,
	}
}

var TreeMockupData = &TreeNode{
	val: Profile{name: "Peter", balance: 25},
	children: []TreeNode{
		{
			val: Profile{name: "Tom", balance: 25},
			children: []TreeNode{
				{val: Profile{name: "Mary", balance: 25}},
				{val: Profile{name: "Vincent", balance: 25}},
				{val: Profile{name: "Vicky", balance: 25}},
			},
		},
		{
			val: Profile{name: "bob", balance: 25},
			children: []TreeNode{
				{val: Profile{name: "Alice", balance: 25}},
			},
		},
	},
}

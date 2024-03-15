package internal

// Linked-list iterator

type LinkedListIterator struct {
	node *LinkedNode
}

func (l *LinkedListIterator) HasNext() bool {
	return l.node != nil
}

func (l *LinkedListIterator) Next() TransferInterface {
	node := l.node
	l.node = node.next
	return node.val
}

type LinkedNode struct {
	val  TransferInterface
	next *LinkedNode
}

func NewImplementLinkedList(node *LinkedNode) TransferIterator {
	return &LinkedListIterator{node: node}
}

var LinkedListMockupData = &LinkedNode{
	val: Profile{name: "Peter", balance: 10},
	next: &LinkedNode{
		val: Profile{name: "Mary", balance: 10},
		next: &LinkedNode{
			val: Profile{name: "Tom", balance: 10},
			next: &LinkedNode{
				val:  Profile{name: "Harry", balance: 10},
				next: nil,
			},
		},
	},
}

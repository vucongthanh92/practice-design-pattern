package internal

// Linked-list iterator

type LinkedNode struct {
	val  TransferInterface
	next *LinkedNode
}

type LinkedPointer struct {
	node *LinkedNode
}

func (l *LinkedPointer) HasNext() bool {
	return l.node != nil
}

func (l *LinkedPointer) Next() TransferInterface {
	node := l.node
	l.node = node.next
	return node.val
}

func NewImplementLinkedList(node *LinkedNode) TransferIterator {
	return &LinkedPointer{node: node}
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

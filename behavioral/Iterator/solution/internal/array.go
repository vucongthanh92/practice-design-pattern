package internal

// Array iterator

type ArrayIterator struct {
	index int
	arr   []TransferInterface
}

func (a *ArrayIterator) HasNext() bool {
	return len(a.arr) > 0 && a.index < len(a.arr)
}

func (a *ArrayIterator) Next() TransferInterface {
	item := a.arr[a.index]
	a.index++
	return item
}

func NewImplementArray(arr []TransferInterface) TransferIterator {
	return &ArrayIterator{index: 0, arr: arr}
}

var ArrayMockupData = []TransferInterface{
	Profile{name: "Peter", balance: 0},
	Profile{name: "Mary", balance: 0},
	Profile{name: "Tom", balance: 0},
	Profile{name: "Harry", balance: 0},
}

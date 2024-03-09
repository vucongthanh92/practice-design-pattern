package main

import "fmt"

type Follower interface {
	Receive(msg string)
}

type Profile struct {
	name string
}

func (p Profile) Receive(msg string) {
	fmt.Printf("%s has received message %s \n", p.name, msg)
}

type FollowerIterator interface {
	Next() Follower
	HasNext() bool
}

func sendMessage(iterator FollowerIterator, msg string) {
	for iterator.HasNext() {
		iterator.Next().Receive(msg)
	}
}

// Array iterator

type FollowerArrayIterator struct {
	currentIdx int
	arr        []Follower
}

func (f *FollowerArrayIterator) HasNext() bool {
	return len(f.arr) > 0 && f.currentIdx < len(f.arr)
}

func (f *FollowerArrayIterator) Next() Follower {
	currentItem := f.arr[f.currentIdx]
	f.currentIdx++

	return currentItem
}

func NewFollowerArrayIterator(arr []Follower) FollowerIterator {
	return &FollowerArrayIterator{currentIdx: 0, arr: arr}
}

var arrayFollowers = []Follower{
	Profile{name: "Peter"},
	Profile{name: "Mary"},
	Profile{name: "Tom"},
	Profile{name: "Harry"},
}

// Linked-list iterator

type FollowerLinkedListIterator struct {
	node *LinkedNode
}

func (f *FollowerLinkedListIterator) HasNext() bool {
	return f.node != nil
}

func (f *FollowerLinkedListIterator) Next() Follower {
	node := f.node
	f.node = node.next

	return node.val
}

func NewFollowerLinkedListIterator(node *LinkedNode) FollowerIterator {
	return &FollowerLinkedListIterator{node: node}
}

var linkedListOfFollowers = &LinkedNode{
	val: Profile{name: "Peter"},
	next: &LinkedNode{
		val: Profile{name: "Mary"},
		next: &LinkedNode{
			val: Profile{name: "Tom"},
			next: &LinkedNode{
				val:  Profile{name: "Harry"},
				next: nil,
			},
		},
	},
}

// Tree Iterator

type FollowerTreeStorage struct {
	node *TreeNode
}

func NewFollowerTreeStorage(node *TreeNode) FollowerTreeStorage {
	return FollowerTreeStorage{
		node: node,
	}
}

func (f FollowerTreeStorage) toArray(node *TreeNode) []Follower {
	if node == nil {
		return nil
	}
	follower := []Follower{node.val}

	for i := range node.children {
		follower = append(follower, f.toArray(&node.children[i])...)
	}

	return follower
}

func (f FollowerTreeStorage) toLinkedList(node *TreeNode, lNode *LinkedNode) *LinkedNode {
	if node == nil {
		return nil
	}

	lNode = &LinkedNode{
		val:  node.val,
		next: lNode,
	}

	for i := range node.children {
		lNode = f.toLinkedList(&node.children[i], lNode)
	}

	return lNode
}

type TreeNode struct {
	val      Follower
	children []TreeNode
}

type LinkedNode struct {
	val  Follower
	next *LinkedNode
}

var treeOfFollower = &TreeNode{
	val: Profile{name: "Peter"},
	children: []TreeNode{
		{
			val: Profile{name: "Tom"},
			children: []TreeNode{
				{val: Profile{name: "Mary"}},
				{val: Profile{name: "Vincent"}},
				{val: Profile{name: "Vicky"}},
			},
		},
		{
			val: Profile{name: "bob"},
			children: []TreeNode{
				{val: Profile{name: "Alice"}},
			},
		},
	},
}

func (f FollowerTreeStorage) Iterator() FollowerIterator {
	return NewFollowerLinkedListIterator(f.toLinkedList(f.node, nil))
}

func main() {
	msg := "hello world"

	fmt.Println("*** [a,b,c] Array iterator")
	iterator := NewFollowerArrayIterator(arrayFollowers)
	sendMessage(iterator, msg)

	fmt.Println("*** a->b->c Linked-list iterator")
	iterator = NewFollowerLinkedListIterator(linkedListOfFollowers)
	sendMessage(iterator, msg)

	fmt.Println("*** a->[b->[e,f],c] Tree iterator")
	iterator = NewFollowerTreeStorage(treeOfFollower).Iterator()
	sendMessage(iterator, msg)
}

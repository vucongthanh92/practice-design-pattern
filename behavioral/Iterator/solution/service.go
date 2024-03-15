package main

import (
	"design-pattern/behavioral/Iterator/solution/internal"
	"fmt"
)

// main function
func main() {
	var amount float64 = 15

	fmt.Println("\n****** [a,b,c] Array iterator ******")
	iterator := internal.NewImplementArray(internal.ArrayMockupData)
	internal.Deposit(iterator, amount)

	fmt.Println("\n****** a->b->c Linked-list iterator ******")
	iterator = internal.NewImplementLinkedList(internal.LinkedListMockupData)
	internal.Deposit(iterator, amount)

	fmt.Println("\n****** a->[b->[e,f],c] Tree iterator ******")
	iterator = internal.NewImplementTree(internal.TreeMockupData).Iterator()
	internal.Deposit(iterator, amount)
}

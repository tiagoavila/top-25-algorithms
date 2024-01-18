package floydscycledetection

import (
	"testing"
)

func TestFindCycleInLinkedListNoCycle(t *testing.T) {
	// Create a linked list without a cycle
	linkedList := &LinkedList{}
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)

	// Call the function and expect the result to be false
	result := FindCycleInLinkedList(linkedList)
	if result {
		t.Errorf("Expected false, got true. No cycle in the linked list.")
	}
}

func TestFindCycleInLinkedListWithCycle(t *testing.T) {
	// Create a linked list with a cycle
	linkedList := &LinkedList{}
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)

	// Create a cycle by connecting the last node to the second node
	lastNode := linkedList.Head
	for lastNode.Next != nil {
		lastNode = lastNode.Next
	}
	lastNode.Next = linkedList.Head.Next

	// Call the function and expect the result to be true
	result := FindCycleInLinkedList(linkedList)
	if !result {
		t.Errorf("Expected true, got false. Cycle exists in the linked list.")
	}
}

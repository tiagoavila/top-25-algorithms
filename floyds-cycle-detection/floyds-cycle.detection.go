package floydscycledetection

// Node represents a node in the linked list
type Node struct {
	Data int
	Next *Node
}

// LinkedList represents a linked list
type LinkedList struct {
	Head *Node
}

// Append adds a new node with the given data to the end of the linked list
func (ll *LinkedList) Append(data int) {
	newNode := &Node{Data: data, Next: nil}

	if ll.Head == nil {
		ll.Head = newNode
		return
	}

	lastNode := ll.Head
	for lastNode.Next != nil {
		lastNode = lastNode.Next
	}

	lastNode.Next = newNode
}

// Floyd’s cycle finding algorithm or Hare-Tortoise algorithm is a pointer algorithm that uses only two pointers,
// moving through the sequence at different speeds. This algorithm is used to find a loop in a linked list.
// It uses two pointers one moving twice as fast as the other one. The faster one is called the fast pointer and the other one is called the slow pointer.
// While traversing the linked list one of these things will occur-
// The Fast pointer may reach the end (NULL) this shows that there is no loop in the linked list.
// The Fast pointer again catches the slow pointer at some time therefore a loop exists in the linked list.
func FindCycleInLinkedList(linkedList *LinkedList) bool {
	slow, fast := linkedList.Head, linkedList.Head

	for fast != nil && fast.Next != nil {
		// move slow by one
		slow = slow.Next

		// move fast by two
		fast = fast.Next.Next

		// if they meet any node, the linked list contains a cycle
		if slow == fast {
			return true
		}
	}

	return false
}

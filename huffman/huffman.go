package huffman

import (
	queue "github.com/Jcowwell/go-algorithm-club/PriorityQueue"
)

type HuffmanTreeNode struct {
	Value     string
	Frequency int
	Left      *HuffmanTreeNode
	Right     *HuffmanTreeNode
}

func NewHuffmanTreeNode(value string, frequency int) *HuffmanTreeNode {
	return &HuffmanTreeNode{
		Value:     value,
		Frequency: frequency,
	}
}

func lessThan(h1, h2 HuffmanTreeNode) bool {
	return h1.Frequency < h2.Frequency
}

func HuffmanCompression(value string) string {
	codes := GetCodes(value)
	compressedString := ""

	for _, char := range value {
		compressedString += codes[string(char)]
	}

	return compressedString
}

func DecompressHuffman(value string, codes map[string]string) string {
	decompressedString, tempCode := "", ""

	for _, char := range value {
		tempCode += string(char)

		if code, exists := codes[tempCode]; exists {
			decompressedString += code
			tempCode = ""
		}
	}

	return decompressedString
}

func GetCodes(value string) map[string]string {
	charsFrequency := make(map[string]int)

	// Count the frequency of each character and store it in a Map
	for _, char := range value {
		frequency, hasHey := charsFrequency[string(char)]
		if hasHey {
			charsFrequency[string(char)] = frequency + 1
		} else {
			charsFrequency[string(char)] = 1
		}
	}

	// Iterate through the map and create a Huffman tree node for each character and add it to the priority queue
	leafNodes := make(map[string]*HuffmanTreeNode)
	priorityQueue := queue.PriorityQueueInit(lessThan)

	for key, value := range charsFrequency {
		node := NewHuffmanTreeNode(key, value)
		leafNodes[key] = node
		priorityQueue.Enqueue(*node)
	}

	rootNode := NewHuffmanTreeNode("Root", 0)

	for !priorityQueue.IsEmpty() {
		leftNode, _ := priorityQueue.Dequeue()
		rightNode, hasRightNode := priorityQueue.Dequeue()

		// If there is a right node, create an internal node and add it to the priority queue
		if hasRightNode {
			internalNode := NewHuffmanTreeNode("Internal Node", leftNode.Frequency+rightNode.Frequency)
			internalNode.Left = &leftNode
			internalNode.Right = &rightNode
			priorityQueue.Enqueue(*internalNode)
		} else {
			// If there is no right node, the left node is the root node
			rootNode = &leftNode
		}
	}

	codes := make(map[string]string)

	// Traverse the tree and get the codes for each character
	traverseTree(rootNode, &codes, "")

	return codes
}

func traverseTree(node *HuffmanTreeNode, codes *map[string]string, code string) {
	if node == nil {
		return
	}

	// If the node is a leaf node, add the code to the map
	if node.Left == nil && node.Right == nil {
		(*codes)[node.Value] = code
	}

	// Recursively traverse the left subtree.
	traverseTree(node.Left, codes, code+"0")

	// Recursively traverse the right subtree.
	traverseTree(node.Right, codes, code+"1")
}

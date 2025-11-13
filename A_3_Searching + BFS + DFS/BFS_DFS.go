package main

import "fmt"

type Node struct {
	Left  *Node
	Right *Node
	Value int
}

func NewNode(value int) *Node {
	return &Node{
		Left:  nil,
		Right: nil,
		Value: value,
	}
}

type BinarySearchTree struct {
	Root *Node
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{
		Root: nil,
	}
}

func (bst *BinarySearchTree) Insert(value int) *BinarySearchTree {
	newNode := NewNode(value)
	if bst.Root == nil {
		bst.Root = newNode
	} else {
		currentNode := bst.Root
		for {
			if value < currentNode.Value {
				// Left
				if currentNode.Left == nil {
					currentNode.Left = newNode
					return bst
				}
				currentNode = currentNode.Left
			} else {
				// Right
				if currentNode.Right == nil {
					currentNode.Right = newNode
					return bst
				}
				currentNode = currentNode.Right
			}
		}
	}
	return bst
}

func (bst *BinarySearchTree) Lookup(value int) *Node {
	if bst.Root == nil {
		return nil
	}
	currentNode := bst.Root
	for currentNode != nil {
		if value < currentNode.Value {
			currentNode = currentNode.Left
		} else if value > currentNode.Value {
			currentNode = currentNode.Right
		} else if currentNode.Value == value {
			return currentNode
		}
	}
	return nil
}

func (bst *BinarySearchTree) Remove(value int) bool {
	if bst.Root == nil {
		return false
	}
	currentNode := bst.Root
	var parentNode *Node = nil

	for currentNode != nil {
		if value < currentNode.Value {
			parentNode = currentNode
			currentNode = currentNode.Left
		} else if value > currentNode.Value {
			parentNode = currentNode
			currentNode = currentNode.Right
		} else if currentNode.Value == value {
			// We have a match, get to work!

			// Option 1: No right child
			if currentNode.Right == nil {
				if parentNode == nil {
					bst.Root = currentNode.Left
				} else {
					// if parent > current value, make current left child a child of parent
					if currentNode.Value < parentNode.Value {
						parentNode.Left = currentNode.Left
					} else if currentNode.Value > parentNode.Value {
						// if parent < current value, make left child a right child of parent
						parentNode.Right = currentNode.Left
					}
				}
			} else if currentNode.Right.Left == nil {
				// Option 2: Right child which doesn't have a left child
				if parentNode == nil {
					bst.Root = currentNode.Right
				} else {
					currentNode.Right.Left = currentNode.Left

					// if parent > current, make right child of the left the parent
					if currentNode.Value < parentNode.Value {
						parentNode.Left = currentNode.Right
					} else if currentNode.Value > parentNode.Value {
						// if parent < current, make right child a right child of the parent
						parentNode.Right = currentNode.Right
					}
				}
			} else {
				// Option 3: Right child that has a left child
				// find the Right child's left most child
				leftmost := currentNode.Right.Left
				leftmostParent := currentNode.Right

				for leftmost.Left != nil {
					leftmostParent = leftmost
					leftmost = leftmost.Left
				}

				// Parent's left subtree is now leftmost's right subtree
				leftmostParent.Left = leftmost.Right
				leftmost.Left = currentNode.Left
				leftmost.Right = currentNode.Right

				if parentNode == nil {
					bst.Root = leftmost
				} else {
					if currentNode.Value < parentNode.Value {
						parentNode.Left = leftmost
					} else if currentNode.Value > parentNode.Value {
						parentNode.Right = leftmost
					}
				}
			}
			return true
		}
	}
	return false
}

func (bst *BinarySearchTree) BreadthFirstSearch() []int {
	currentNode := bst.Root
	list := []int{}
	queue := []*Node{}
	queue = append(queue, currentNode)

	for len(queue) > 0 {
		currentNode = queue[0]
		queue = queue[1:]
		list = append(list, currentNode.Value)

		if currentNode.Left != nil {
			queue = append(queue, currentNode.Left)
		}
		if currentNode.Right != nil {
			queue = append(queue, currentNode.Right)
		}
	}
	return list
}

func (bst *BinarySearchTree) BreadthFirstSearchR(queue []*Node, list []int) []int {
	if len(queue) == 0 {
		return list
	}
	currentNode := queue[0]
	queue = queue[1:]
	list = append(list, currentNode.Value)

	if currentNode.Left != nil {
		queue = append(queue, currentNode.Left)
	}
	if currentNode.Right != nil {
		queue = append(queue, currentNode.Right)
	}

	return bst.BreadthFirstSearchR(queue, list)
}

func (bst *BinarySearchTree) DFSInOrder() []int {
	return traverseInOrder(bst.Root, []int{})
}

func (bst *BinarySearchTree) DFSPostOrder() []int {
	return traversePostOrder(bst.Root, []int{})
}

func (bst *BinarySearchTree) DFSPreOrder() []int {
	return traversePreOrder(bst.Root, []int{})
}

func traverseInOrder(node *Node, list []int) []int {
	if node == nil {
		return list
	}
	fmt.Println(node.Value)
	if node.Left != nil {
		list = traverseInOrder(node.Left, list)
	}
	list = append(list, node.Value)
	if node.Right != nil {
		list = traverseInOrder(node.Right, list)
	}
	return list
}

func traversePreOrder(node *Node, list []int) []int {
	if node == nil {
		return list
	}
	fmt.Println(node.Value)
	list = append(list, node.Value)
	if node.Left != nil {
		list = traversePreOrder(node.Left, list)
	}
	if node.Right != nil {
		list = traversePreOrder(node.Right, list)
	}
	return list
}

func traversePostOrder(node *Node, list []int) []int {
	if node == nil {
		return list
	}
	fmt.Println(node.Value)
	if node.Left != nil {
		list = traversePostOrder(node.Left, list)
	}
	if node.Right != nil {
		list = traversePostOrder(node.Right, list)
	}
	list = append(list, node.Value)
	return list
}

func traverse(node *Node) map[string]interface{} {
	if node == nil {
		return nil
	}
	tree := make(map[string]interface{})
	tree["value"] = node.Value
	if node.Left == nil {
		tree["left"] = nil
	} else {
		tree["left"] = traverse(node.Left)
	}
	if node.Right == nil {
		tree["right"] = nil
	} else {
		tree["right"] = traverse(node.Right)
	}
	return tree
}

func main() {
	tree := NewBinarySearchTree()
	tree.Insert(9)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(20)
	tree.Insert(170)
	tree.Insert(15)
	tree.Insert(1)

	// fmt.Println("BFS:", tree.BreadthFirstSearch())
	// fmt.Println("BFS Recursive:", tree.BreadthFirstSearchR([]*Node{tree.Root}, []int{}))
	fmt.Println("DFS Post Order:", tree.DFSPostOrder())

	//     9
	//  4     20
	//1  6  15  170
}

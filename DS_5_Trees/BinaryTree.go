package main

import (
	"encoding/json"
	"fmt"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func NewNode(value int) *Node {
	return &Node{
		value: value,
		left:  nil,
		right: nil,
	}
}

type BinarySearchTree struct {
	root *Node
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{
		root: nil,
	}
}

func (bst *BinarySearchTree) Insert(value int) *BinarySearchTree {
	newNode := NewNode(value)
	if bst.root == nil {
		bst.root = newNode
	} else {
		currentNode := bst.root
		for {
			if currentNode.value > value {
				if currentNode.left == nil {
					currentNode.left = newNode
					return bst
				}
				currentNode = currentNode.left
			} else {
				if currentNode.right == nil {
					currentNode.right = newNode
					return bst
				}
				currentNode = currentNode.right
			}
		}
	}
	return bst
}

func (bst *BinarySearchTree) Lookup(value int) interface{} {
	if bst.root == nil {
		return false
	}
	currentNode := bst.root
	for currentNode != nil {
		if currentNode.value > value {
			currentNode = currentNode.left
		} else if currentNode.value < value {
			currentNode = currentNode.right
		} else if currentNode.value == value {
			return currentNode
		}
	}
	return false
}

func (bst *BinarySearchTree) Remove(value int) bool {
	if bst.root == nil {
		return false
	}
	currentNode := bst.root
	var parentNode *Node = nil

	for currentNode != nil {
		if value < currentNode.value {
			parentNode = currentNode
			currentNode = currentNode.left
		} else if value > currentNode.value {
			parentNode = currentNode
			currentNode = currentNode.right
		} else if currentNode.value == value {
			if currentNode.right == nil {
				if parentNode == nil {
					bst.root = currentNode.left
				} else {
					if currentNode.value < parentNode.value {
						parentNode.left = currentNode.left
					} else if currentNode.value > parentNode.value {
						parentNode.right = currentNode.left
					}
				}
			} else if currentNode.right.left == nil {
				currentNode.right.left = currentNode.left
				if parentNode == nil {
					bst.root = currentNode.right
				} else {
					if currentNode.value < parentNode.value {
						parentNode.left = currentNode.right
					} else if currentNode.value > parentNode.value {
						parentNode.right = currentNode.right
					}
				}
			} else {
				leftmost := currentNode.right.left
				leftmostParent := currentNode.right
				for leftmost.left != nil {
					leftmostParent = leftmost
					leftmost = leftmost.left
				}
				leftmostParent.left = leftmost.right
				leftmost.left = currentNode.left
				leftmost.right = currentNode.right
				if parentNode == nil {
					bst.root = leftmost
				} else {
					if currentNode.value < parentNode.value {
						parentNode.left = leftmost
					} else if currentNode.value > parentNode.value {
						parentNode.right = leftmost
					}
				}
			}
			return true
		}
	}
	return false
}

type TreeNode struct {
	Value int       `json:"value"`
	Left  *TreeNode `json:"left"`
	Right *TreeNode `json:"right"`
}

func traverse(node *Node) *TreeNode {
	if node == nil {
		return nil
	}
	tree := &TreeNode{Value: node.value}
	tree.Left = traverse(node.left)
	tree.Right = traverse(node.right)
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
	tree.Remove(170)

	treeData := traverse(tree.root)
	jsonData, _ := json.Marshal(treeData)
	fmt.Println(string(jsonData))
	fmt.Println(tree.Lookup(20))
}

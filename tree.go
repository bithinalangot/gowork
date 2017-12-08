package main

import "fmt"

type TreeNode struct {
  value int
  key int
  leftChild *TreeNode
  rightChild *TreeNode
  parent *TreeNode
}

type BinaryTree struct{
  root *TreeNode
}

// If the tree node don't have a left child then return 'false'
func (T *TreeNode) haslC() bool {
  if T.leftChild == nil {
    return false
  }
  return true
}

// If the tree node don't have a right child then return 'false'
func (T *TreeNode) hasrC() bool {
  if T.rightChild == nil {
      return false
  }
  return true
}

func (T *TreeNode) isRightChild() bool {
  if T.parent.leftChild == T {
    return true
  }
  return false
}

func (T *TreeNode) isLeftChild() bool {
  if T.parent.leftChild == T {
    return true
  }
  return false
}

//create a new tree node with values
func newTreeNode(parent *TreeNode, value, key int) (*TreeNode) {
  tree := &TreeNode {
    parent:     parent,
    leftChild:  nil,
    rightChild: nil,
    value:      value,
    key:        key,
  }
  return tree
}

func newBinaryTree()(*BinaryTree){
  bt := &BinaryTree {
    root: nil,
  }
  return bt
}

// recursive function to construct a tree.
func buildTree(parent *TreeNode, value, key int){
  if key <= parent.key {
    if parent.haslC() {
      buildTree(parent.leftChild, value, key)
    } else {
      parent.leftChild = newTreeNode(parent, value, key)
    }
  } else {
    if parent.hasrC() {
      buildTree(parent.rightChild, value, key)
    } else {
      parent.rightChild = newTreeNode(parent, value, key)
    }
  }
  return
}

func (bt *BinaryTree) createTree(a map[int]int) {
  for key, value := range a {
    if bt.root == nil {
      bt.root = newTreeNode(nil, value, key)
    } else {
      buildTree(bt.root, value, key)
    }
  }
}

func printInorder(node *TreeNode) {
  if node == nil {
    return
  }
  fmt.Printf("%d\n", node.key)
  printInorder(node.leftChild)
  printInorder(node.rightChild)
}

func main() {
  bt := newBinaryTree()
  a := map[int]int{
    1 : 45,
    22 : 23,
    3 : 53,
    4 : 34,
    5 : 23,
  }
  bt.createTree(a)
  printInorder(bt.root)
}

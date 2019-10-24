package main

import "fmt"

type BinarySearchTree struct {
	root *BinarySearchTreeNode
}

type BinarySearchTreeNode struct {
	data byte
	left *BinarySearchTreeNode
	right *BinarySearchTreeNode
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

func NewBinarySearchTreeNode(data byte) *BinarySearchTreeNode {
	return &BinarySearchTreeNode{
		data: data,
	}
}

func (t *BinarySearchTree) insert(data byte) {
	if t.root == nil {
		t.root = NewBinarySearchTreeNode(data)
	} else {
		t.root.insert(data)
	}
}

func (t *BinarySearchTreeNode) insert(data byte) {
	if data <= t.data {
		if t.left == nil {
			t.left = NewBinarySearchTreeNode(data)
		} else {
			t.left.insert(data)
		}
	} else {
		if t.right == nil {
			t.right = NewBinarySearchTreeNode(data)
		} else {
			t.right.insert(data)
		}
	}
}

func (t *BinarySearchTree) print() {
	if t.root != nil {
		t.root.print()
	}
}

func (t *BinarySearchTreeNode) print() {
	if t == nil {
		return
	}

	t.left.print()
	fmt.Printf("%d ", t.data)
	t.right.print()
}
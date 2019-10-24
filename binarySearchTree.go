package main

import (
	"fmt"
	"sync"
)

type BinarySearchTree struct {
	root *BinarySearchTreeNode
	mtx sync.Mutex
}

type BinarySearchTreeNode struct {
	data  byte
	left  *BinarySearchTreeNode
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
		t.mtx.Lock()
		t.root.insert(data)
		t.mtx.Unlock()
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

func (t *BinarySearchTree) printInOrder() {
	if t.root != nil {
		t.mtx.Lock()
		t.root.printInOrder()
		t.mtx.Unlock()
	}
}

func (t *BinarySearchTreeNode) printInOrder() {
	if t == nil {
		return
	}

	t.left.printInOrder()
	fmt.Printf("%d ", t.data)
	t.right.printInOrder()
}

func (t *BinarySearchTree) printPreOrder() {
	if t.root != nil {
		t.mtx.Lock()
		t.root.printPreOrder()
		t.mtx.Unlock()
	}
}

func (t *BinarySearchTreeNode) printPreOrder() {
	if t == nil {
		return
	}

	fmt.Printf("%d ", t.data)
	t.left.printPreOrder()
	t.right.printPreOrder()
}

func (t *BinarySearchTree) printPostOrder() {
	if t.root != nil {
		t.mtx.Lock()
		t.root.printPostOrder()
		t.mtx.Unlock()
	}
}

func (t *BinarySearchTreeNode) printPostOrder() {
	if t == nil {
		return
	}

	t.left.printPostOrder()
	t.right.printPostOrder()
	fmt.Printf("%d ", t.data)
}

package main

func main() {
	tree := NewBinarySearchTree()
	tree.insert(12)
	tree.insert(1)
	tree.insert(3)
	tree.insert(2)
	tree.insert(14)
	tree.insert(6)
	tree.insert(0)
	tree.insert(2)
	tree.insert(5)

	tree.print()
}

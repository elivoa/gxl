package collections

type Tree struct {
	Root *Node
}

type Node struct {
	Parent   *Node
	Children []*Node
}

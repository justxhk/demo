package math

import "fmt"

type TreeNode struct {
	Value     int
	LeftNode  *TreeNode
	RightNode *TreeNode
}

func (tree *TreeNode) Add(value int) {
	insert(tree, value)
}

func insert(tree *TreeNode, value int) {
	if tree.Value == value {
		return
	}
	if tree.Value > value {
		if tree.LeftNode == nil {
			tree.LeftNode = &TreeNode{Value: value}
			return
		}
		insert(tree.LeftNode, value)
	} else {
		if tree.RightNode == nil {
			tree.RightNode = &TreeNode{Value: value}
			return
		}
		insert(tree.RightNode, value)
	}
}

func (tree *TreeNode) PrintTree() {
	fmt.Print(tree.Value, " ")
	if tree.LeftNode != nil {
		tree.LeftNode.PrintTree()
	}
	if tree.RightNode != nil {
		tree.RightNode.PrintTree()
	}
}

func (tree *TreeNode) MiddlePrintTree() {
	if tree.LeftNode != nil {
		tree.LeftNode.MiddlePrintTree()
	}
	fmt.Print(tree.Value, " ")
	if tree.RightNode != nil {
		tree.RightNode.MiddlePrintTree()
	}
}

func (tree *TreeNode) PrintTreeNoRe() {
	var stack []*TreeNode
	for tree != nil || len(stack) != 0 {
		if tree != nil {
			fmt.Print(tree.Value, " ")
			stack = append(stack, tree)
			tree = tree.LeftNode
		} else {
			tree = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			tree = tree.RightNode
		}
	}
}

func (tree *TreeNode) PrintTreeByLevel() {
	var stack []*TreeNode
	stack = append(stack, tree)
	i := 0
	for len(stack) != 0 {
		i++
		fmt.Print("第", i, "层，元素: ")
		var tmp []*TreeNode
		for _, t := range stack {
			if t.LeftNode != nil {
				tmp = append(tmp, t.LeftNode)
			}
			if t.RightNode != nil {
				tmp = append(tmp, t.RightNode)
			}
			fmt.Print(t.Value, " ")
		}
		fmt.Println()
		stack = tmp
	}
}

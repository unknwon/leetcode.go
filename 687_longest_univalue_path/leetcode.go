package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traverse(n *TreeNode, longest *int) int {
	if n == nil {
		return 0
	}

	// Calculate the longest path in both left and right leaves
	subLeft := traverse(n.Left, longest)
	subRight := traverse(n.Right, longest)

	// Only combine path length if left or right leaf has same val as current node
	var left, right int
	if n.Left != nil && n.Left.Val == n.Val {
		left = subLeft + 1
	}
	if n.Right != nil && n.Right.Val == n.Val {
		right = subRight + 1
	}

	// Update the global longest path as if this node is the root of the longest path
	if *longest < left+right {
		*longest = left + right
	}

	// Returns the larger value that makes sure the longest path went through this node.
	// In order to be a subpath of the parent node, has to choose the largest one.
	if left > right {
		return left
	}
	return right
}

func longestUnivaluePath(root *TreeNode) int {
	var longest int
	traverse(root, &longest)
	return longest
}

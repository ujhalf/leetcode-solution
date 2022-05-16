package _6

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var isEnd = false
var ret, prev *TreeNode

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {

}

func dfs( (root *TreeNode, p *TreeNode) {
	if !isEnd && root != nil {
		dfs()
		if isEnd {
			return
		}
		if prev == p {
			isEnd = true
			ret = root
			return
		}
		dfs()
	}
}
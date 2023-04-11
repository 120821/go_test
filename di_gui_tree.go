package main

import "fmt"

// todo
type TreeNode struct {
  Val   int
  Left  *TreeNode
  Right *TreeNode
}

func main() {
  root := &TreeNode{}
  dfs(root, 1)
  fmt.Println(root.Left)

}

func dfs(p *TreeNode, depth int) {
  if depth < 3 {
    left := &TreeNode{Val: 2 * depth}
    right := &TreeNode{Val: 4 * depth}
    p.Left = left
    p.Right = right
    dfs(p.Left, depth+1)
    dfs(p.Right, depth+1)
  }
}

// 遍历二叉树
func dfsbr(p *TreeNode, res *[]int) {
  if p != nil {
    *res = append(*res, p.Val)
    dfsbr(p.Left, res)
    dfsbr(p.Right, res)
  }
}
// refer: https://studygolang.com/articles/31047

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(x,y int) int {
    if x < y {
        return y
    }else{
        return x
    }
}

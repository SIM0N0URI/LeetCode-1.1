func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return 1 + min(minDepth(root.Left), minDepth(root.Right))
}

func min(x, y int) int{
    if x == 0 {
        return y
    }
    if y == 0 {
        return x
    }
    if x > y {
        return y
    }else {
        return x
    }
}
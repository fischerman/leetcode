/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
 func kthSmallest(root *TreeNode, k int) int {
    
    // we can assume we found the element because k is valid
    _, res := kthSmallestCore(root, &k)
    return res
    
}

func kthSmallestCore(node *TreeNode, c *int) (bool, int) {
    if node.Left != nil {
        found, res := kthSmallestCore(node.Left, c)
        if found {
            return true, res
        }
    }
    if *c == 1 {
        return true, node.Val
    }
    *c--
    if node.Right != nil {
        found, res := kthSmallestCore(node.Right, c)
        if found {
            return true, res
        }
    }
    return false, 0
}

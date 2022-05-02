/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    curr := root
    var inorder []int
    for curr != nil {
        if curr.Left == nil {
            inorder = append(inorder, curr.Val)
            curr = curr.Right
        } else {
            prev := curr.Left
            for prev.Right != nil && prev != curr {
                prev = prev.Right
            }
            if prev.Right == nil {
                prev.Right = curr
                curr = curr.Left
            } else {
                prev.Right = nil
                inorder = append(inorder, curr.Val)
                curr = curr.Right
            }
        }
    }
    return inorder
}
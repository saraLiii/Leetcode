/*
给你二叉搜索树的根节点 root ，该树中的 恰好 两个节点的值被错误地交换。请在不改变其结构的情况下，恢复这棵树 
输入：root = [1,3,null,null,2]
输出：[3,1,null,null,2]
解释：3 不能是 1 的左孩子，因为 3 > 1 。交换 1 和 3 使二叉搜索树有效。

输入：root = [3,1,4,null,null,2]
输出：[2,1,4,null,null,3]
解释：2 不能在 3 的右子树中，因为 2 < 3 。交换 2 和 3 使二叉搜索树有效。
*/

/*
做题感想：如果发现结构比较复杂，需要想到用额外的空间来存储中间结果，然后再处理
然后需要想清楚换成另一个形态的本质是什么，比如说这个本质就是逆序对，但是放在树里面，换位置这个就很奇特，一时想不明白


官方还有两个优化的版本，就是给予变成数组之后，其实是找到两个逆序对。
所以后面就变成非递归的中跟遍历，寻找逆序对
以及 Morris 遍历，空间复杂度O(1)的解法
可以记住，但是morris遍历确实比较复杂，不是很常用，也容易用错，暂时不考虑在面试中使用。
*/


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getOrder(root *TreeNode) []*TreeNode{
    if root==nil{
        return []*TreeNode{}
    }
    order := append(getOrder(root.Left),root)
    return append(order,getOrder(root.Right)...)
}
func recoverTree(root *TreeNode)  {
    order:=getOrder(root)
    n:=len(order)
    i,j:=0,n-1
    for ;i<j&&order[i].Val <= order[i+1].Val;i++{
        continue
    }
    for ;i<j&&order[j].Val >= order[j-1].Val;j--{
        continue
    }
    order[i].Val,order[j].Val = order[j].Val,order[i].Val 
}
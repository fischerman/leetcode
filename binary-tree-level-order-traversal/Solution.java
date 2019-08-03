/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
class Solution {
    public List<List<Integer>> levelOrder(TreeNode root) {
        List<List<Integer>> lo = new ArrayList<>();
        List<TreeNode> curLevel = new ArrayList<TreeNode>();
        if(root == null) {
            return lo;
        }
        curLevel.add(root);
        while (curLevel.size() != 0) {
            ArrayList<TreeNode> nextLevel = new ArrayList<TreeNode>();
            List<Integer> o = new ArrayList<>();
            
            for(int i = 0; i < curLevel.size(); i++) {
                TreeNode node = curLevel.get(i);
                if(node.left != null) nextLevel.add(node.left);
                if(node.right != null) nextLevel.add(node.right);
                o.add(node.val);
            }
            
            lo.add(o);
            curLevel = nextLevel;
        }
        return lo;
    }
}
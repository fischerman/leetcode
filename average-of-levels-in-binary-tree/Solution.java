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
    public List<Double> averageOfLevels(TreeNode root) {
        List<Double> averages = new ArrayList<>();
        List<TreeNode> curLevel = new ArrayList<TreeNode>();
        curLevel.add(root);
        while (curLevel.size() != 0) {
            ArrayList<TreeNode> nextLevel = new ArrayList<TreeNode>();
            double sum = 0;
            
            for(int i = 0; i < curLevel.size(); i++) {
                TreeNode node = curLevel.get(i);
                if(node.left != null) nextLevel.add(node.left);
                if(node.right != null) nextLevel.add(node.right);
                sum += node.val;
            }
            
            averages.add(sum / curLevel.size());
            curLevel = nextLevel;
        }
        return averages;
    }
}
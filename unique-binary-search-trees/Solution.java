class Solution {
    
    private Map<Integer, Integer> cache = new HashMap<>();
    
    public int numTrees(int n) {
        if(n < 2) {
            return 1;
        }
        int sum = 0;
        for(int root = 0; root < n; root++) {
            sum += cachedNumTrees(root) * cachedNumTrees(n - root - 1);
        }
        return sum;
    }
    
    private int cachedNumTrees(int n) {
        if(this.cache.containsKey(n)) {
            return this.cache.get(n);
        } else {
            int size = this.numTrees(n);
            this.cache.put(n, size);
            return size;
        }
    }
}
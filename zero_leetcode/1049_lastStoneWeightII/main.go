package _049_lastStoneWeightII

//本题其实就是尽量让石头分成重量相同的两堆，相撞之后剩下的石头最小，这样就化解成01背包问题了。
//本题物品的重量为stone[i]，物品的价值也为stone[i]。
//dp[j]表示容量（这里说容量更形象，其实就是重量）为j的背包，最多可以背dp[j]这么重的石头。
//01背包的递推公式为：dp[j] = max(dp[j], dp[j - weight[i]] + value[i]);
//本题则是：dp[j] = max(dp[j], dp[j - stones[i]] + stones[i]);

func lastStoneWeightII(stones []int) int {
	//因为提示中给出1 <= stones.length <= 30，1 <= stones[i] <= 100，所以最大重量就是30 * 100 。
	//而我们要求的target其实只是最大重量的一半，所以dp数组开到1500大小就可以了。
	dp := make([]int, 1501)
	sum := 0
	for _, v := range stones {
		sum += v
	}
	target := sum / 2
	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}
	//最后dp[target]里是容量为target的背包所能背的最大重量。
	//那么分成两堆石头，一堆石头的总重量是dp[target]，另一堆就是sum - dp[target]。
	//在计算target的时候，target = sum / 2 因为是向下取整，所以sum - dp[target] 一定是大于等于dp[target]的。
	//那么相撞之后剩下的最小石头重量就是 (sum - dp[target]) - dp[target]。
	return sum - 2*dp[target]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

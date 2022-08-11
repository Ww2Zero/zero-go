package _22_coinChange

import (
	"math"
)

//leetcode submit region begin(Prohibit modification and deletion)

// 输入 [1,2,5] 11
// 输出 3
func coinChangeBrute(coins []int, amount int) (res int) {
	helpN(amount, ">amount:%v-res:%v", amount, res)
	defer func() {
		helpN(amount, " < amount:%v-res:%v", amount, res)
	}()
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	res = math.MaxInt
	for i := 0; i < len(coins); i++ {
		cur := coins[i]
		subProblem := coinChangeBrute(coins, amount-cur)
		if subProblem == -1 {
			continue
		}
		res = min(res, subProblem+1)
	}
	if res == math.MinInt {
		return -1
	}

	return res
}
func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}
func coinChange(coins []int, amount int) (res int) {
	memo := make([]int, amount+1)
	for i := 0; i < amount+1; i++ {
		memo[i] = -666
	}
	return coinChangeMemo(coins, amount, memo)
}
func coinChangeMemo(coins []int, amount int, memo []int) (res int) {
	helpN(amount, " > amount:%v-res:%v", amount, res)
	defer func() {
		helpN(amount, " < amount:%v-res:%v", amount, res)
	}()
	if amount == 0 {
		res = 0
		return
	}
	if amount < 0 {
		res = -1
		return
	}
	if memo[amount] != -666 {
		res = memo[amount]
		return
	}
	res = math.MaxInt
	for i := 0; i < len(coins); i++ {
		cur := coins[i]
		subProblem := coinChangeMemo(coins, amount-cur, memo)
		if subProblem == -1 {
			continue
		}
		res = min(res, subProblem+1)
	}
	if res == math.MaxInt {
		memo[amount] = -1
	} else {
		memo[amount] = res
	}
	res = memo[amount]
	return
}

//=== RUN   Test_coinChange
//=== RUN   Test_coinChange/test
//=========== > amount:11-res:0
//========== > amount:10-res:0
//========= > amount:9-res:0
//======== > amount:8-res:0
//======= > amount:7-res:0
//====== > amount:6-res:0
//===== > amount:5-res:0
//==== > amount:4-res:0
//=== > amount:3-res:0
//== > amount:2-res:0
//= > amount:1-res:0
//> amount:0-res:0
//< amount:0-res:0
//> amount:-1-res:0
//< amount:-1-res:-1
//> amount:-4-res:0
//< amount:-4-res:-1
//= < amount:1-res:1
//> amount:0-res:0
//< amount:0-res:0
//> amount:-3-res:0
//< amount:-3-res:-1
//== < amount:2-res:1
//= > amount:1-res:0
//= < amount:1-res:1
//> amount:-2-res:0
//< amount:-2-res:-1
//=== < amount:3-res:2
//== > amount:2-res:0
//== < amount:2-res:1
//> amount:-1-res:0
//< amount:-1-res:-1
//==== < amount:4-res:2
//=== > amount:3-res:0
//=== < amount:3-res:2
//> amount:0-res:0
//< amount:0-res:0
//===== < amount:5-res:1
//==== > amount:4-res:0
//==== < amount:4-res:2
//= > amount:1-res:0
//= < amount:1-res:1
//====== < amount:6-res:2
//===== > amount:5-res:0
//===== < amount:5-res:1
//== > amount:2-res:0
//== < amount:2-res:1
//======= < amount:7-res:2
//====== > amount:6-res:0
//====== < amount:6-res:2
//=== > amount:3-res:0
//=== < amount:3-res:2
//======== < amount:8-res:3
//======= > amount:7-res:0
//======= < amount:7-res:2
//==== > amount:4-res:0
//==== < amount:4-res:2
//========= < amount:9-res:3
//======== > amount:8-res:0
//======== < amount:8-res:3
//===== > amount:5-res:0
//===== < amount:5-res:1
//========== < amount:10-res:2
//========= > amount:9-res:0
//========= < amount:9-res:3
//====== > amount:6-res:0
//====== < amount:6-res:2
//=========== < amount:11-res:3
//--- PASS: Test_coinChange (0.00s)
//--- PASS: Test_coinChange/test (0.00s)
//PASS
//
//Process finished with the exit code 0

//leetcode submit region end(Prohibit modification and deletion)

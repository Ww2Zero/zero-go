package _273_deleteTreeNodes

import "fmt"

//parent: []int{-1, 0, 0, 1, 2, 2, 2},
//value:  []int{1, -2, 4, 0, -2, -1, -1},
func deleteTreeNodes(nodes int, parent []int, value []int) int {
	edge := make([][]int, nodes)

	for i := 0; i < len(parent); i++ {
		if parent[i] != -1 {
			edge[parent[i]] = append(edge[parent[i]], i)
		}
	}
	cnt := make([]int, nodes)
	for i := 0; i < len(cnt); i++ {
		cnt[i] = 1
	}
	var dfs func(now int)

	dfs = func(now int) {
		cur := edge[now]
		p(now, "cur:  %d: %v\n", now, cur)
		for i := 0; i < len(cur); i++ {
			subNode := cur[i]
			p(now, "subNode: %d\n", subNode)
			dfs(subNode)
			value[now] += value[subNode]
			p(now, "value[now=%v]: %d, %v \n", now, value[now], value)
			cnt[now] += cnt[subNode]
			p(now, "cnt[now=%v]: %d, %v \n", now, cnt[now], cnt)
		}
		if value[now] == 0 {
			// 设置为0表示子节点被删除
			cnt[now] = 0
		}
	}

	dfs(0)

	return cnt[0]
}

func p(now int, format string, args ...interface{}) {
	for n := 0; n < now; n++ {
		fmt.Printf("=")
	}
	fmt.Printf(format, args...)
}

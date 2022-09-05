package _84_findRedundantConnection

var (
	n      = 1005
	father = make([]int, 1005)
)

func initialize() {
	for i := 0; i < n; i++ {
		father[i] = i
	}
}

func find(u int) int {
	if u == father[u] {
		return u
	}
	tmp := find(father[u])
	father[u] = tmp
	return father[u]
}

func join(u, v int) {
	u = find(u)
	v = find(v)
	if u == v {
		return
	}
	father[v] = u
}
func same(u, v int) bool {
	u = find(u)
	v = find(v)
	return u == v
}
func findRedundantConnection(edges [][]int) []int {
	initialize()
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		if same(u, v) {
			return edge
		} else {
			join(u, v)
		}
	}
	return []int{}
}

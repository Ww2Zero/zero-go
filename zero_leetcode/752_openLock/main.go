package _52_openLock

func openLock(deadends []string, target string) int {
	const start = "0000"
	if target == start {
		return 0
	}

	dead := map[string]bool{}
	for _, s := range deadends {
		dead[s] = true
	}
	if dead[start] {
		return -1
	}

	// 枚举 status 通过一次旋转得到的数字
	get := func(status string) (ret []string) {
		s := []byte(status)
		for i, b := range s {
			s[i] = b - 1
			if s[i] < '0' {
				s[i] = '9'
			}
			ret = append(ret, string(s))
			s[i] = b + 1
			if s[i] > '9' {
				s[i] = '0'
			}
			ret = append(ret, string(s))
			s[i] = b
		}
		return
	}

	type pair struct {
		status string
		step   int
	}
	step := 0
	q := []string{start}
	seen := map[string]bool{start: true}
	for len(q) > 0 {
		n := len(q)
		for i := 0; i < n; i++ {
			status := q[i]
			for _, nxt := range get(status) {
				if dead[nxt] {
					continue
				}
				if seen[nxt] {
					continue
				}
				if nxt == target {
					// 之前的步骤次数+当前的扭动的步骤次数
					return step + 1
				}
				seen[nxt] = true
				q = append(q, nxt)
			}
		}
		step++
	}
	return -1
}

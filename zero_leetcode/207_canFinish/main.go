package _07_canFinish

//leetcode submit region begin(Prohibit modification and deletion)

// https://leetcode.cn/problems/course-schedule/solution/bao-mu-shi-ti-jie-shou-ba-shou-da-tong-tuo-bu-pai-/

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 1.记录课程入度table
	inDegreeTable := make(map[int]int, 0)
	for i := 0; i < numCourses; i++ {
		inDegreeTable[i] = 0
	}
	// 2.记录课程之间依赖关系,更新课程的入度
	relationTable := make(map[int][]int, 0)
	for i := 0; i < len(prerequisites); i++ {
		r := prerequisites[i]
		cur := r[1]
		nxt := r[0]
		// update inDegree table
		inDegreeTable[nxt] = inDegreeTable[nxt] + 1
		// update relation table
		relationTable[cur] = append(relationTable[cur], nxt)
	}

	// 3.bfs ，将入度为0的课程加入队列，
	var zeroQueue []int
	for k, v := range inDegreeTable {
		if v == 0 {
			zeroQueue = append(zeroQueue, k)
		}
	}
	// 4. 从队列中取课程，学习完该课程之后，修改依赖该课程的入度信息, 将入度为0的课程 ，加入到队列
	for len(zeroQueue) > 0 {
		q := zeroQueue[0]
		// remove first node
		zeroQueue = zeroQueue[1:]
		if subArr, ok := relationTable[q]; ok {
			for i := 0; i < len(subArr); i++ {
				//
				inDegreeTable[subArr[i]] = inDegreeTable[subArr[i]] - 1
				if inDegreeTable[subArr[i]] == 0 {
					zeroQueue = append(zeroQueue, subArr[i])
				}
			}
		} else {
			continue
		}
	}
	// 5. 遍历课程的入度table，若还有课程入度>0，则返回false
	for _, v := range inDegreeTable {
		if v != 0 {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

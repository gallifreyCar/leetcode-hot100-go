package courseschedule

// 207. 课程表
// 难度：中等
// 标签：拓扑排序、图
// 链接：https://leetcode.cn/problems/course_schedule/

// CanFinish DFS + 拓扑排序
// 时间复杂度: O(V+E)
// 空间复杂度: O(V+E)
func CanFinish(numCourses int, prerequisites [][]int) bool {
	// 构建邻接表
	graph := make([][]int, numCourses)
	for _, pre := range prerequisites {
		graph[pre[1]] = append(graph[pre[1]], pre[0])
	}

	// 0: 未访问, 1: 访问中, 2: 已访问
	visited := make([]int, numCourses)

	var hasCycle func(int) bool
	hasCycle = func(course int) bool {
		if visited[course] == 1 {
			return true // 检测到环
		}
		if visited[course] == 2 {
			return false // 已经访问过
		}

		visited[course] = 1
		for _, next := range graph[course] {
			if hasCycle(next) {
				return true
			}
		}
		visited[course] = 2
		return false
	}

	for i := 0; i < numCourses; i++ {
		if hasCycle(i) {
			return false
		}
	}
	return true
}

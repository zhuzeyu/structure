package algorithm

import "fmt"

//将jobs平均分配到m个节点上

func split(jobs []int, m int) [][]int {
	var res [][]int
	lenJobs := len(jobs)
	if m <= 0 || lenJobs == 0 {
		return res
	}

	num := lenJobs / m
	if lenJobs%m != 0 {
		num = num + 1
	}
	res = make([][]int, num)

	for x := 0; x < num; x++ {
		res[x] = make([]int, m)
	}

	for i := 0; i < lenJobs; i++ {
		res[i/m][i%m] = jobs[i]
	}
	return res
}

func splitM(jobs []int, m int) [][]int {
	var res [][]int
	lenJobs := len(jobs)
	if m <= 0 || lenJobs == 0 {
		return res
	}

	res = make([][]int, m)

	len := lenJobs / m
	if lenJobs%m != 0 {
		len = len + 1
	}
	for i := 0; i < m; i++ {
		res[i] = make([]int, len)
	}

	for i := 0; i < lenJobs; i++ {
		res[i%m][i/m] = jobs[i]
	}
	return res
}

func main() {
	fmt.Println(split([]int{1, 2, 3, 4, 5, 6, 7}, 3))
	fmt.Println(splitM([]int{1, 2, 3, 4, 5, 6, 7}, 3))
}

package main

import "sort"

type Task struct {
	Time     int
	Deadline int
}

type Tasks []Task

// 以下インタフェースを満たす
// https://qiita.com/Jxck_/items/fb829b818aac5b5f54f7
func (t Tasks) Len() int {
	return len(t)
}

func (t Tasks) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t Tasks) Less(i, j int) bool {
	// Sort by the deadline
	return t[i].Deadline < t[j].Deadline
}

func Megalomania(array1, array2 []int) bool {
	n := len(array1)
	if n != len(array2) {
		panic("Wrong input argument")
	}

	tasks := make(Tasks, n)

	for i := 0; i < n; i++ {
		tasks[i] = Task{array1[i], array2[i]}
	}
	sort.Sort(tasks) //O(NlogN)

	current_end_time := 0
	for i := 0; i < n; i++ {
		// pickup task
		task := tasks[i]
		//do task
		current_end_time += task.Time

		if current_end_time > task.Deadline {
			return false
		}
	}
	return true

}

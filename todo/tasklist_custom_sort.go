package todo

import (
	"sort"
)

// ----------------------------------------------------------------------------
//  TaskList.Sort()
// ----------------------------------------------------------------------------

// CustomSort allows a TaskList to be sorted by a custom function.
//
// The providing function must return true if taskA is less than taskB.
func (tasklist *TaskList) CustomSort(isALessThanB func(taskA, taskB Task) bool) {
	sort.Slice(
		*tasklist,
		func(i int, j int) bool {
			return isALessThanB((*tasklist)[i], (*tasklist)[j])
		},
	)
}

package todo

// ----------------------------------------------------------------------------
//  Type: tasklistSort
// ----------------------------------------------------------------------------

// tasklistSort is an implementation of sort.Interface.
type tasklistSort struct {
	by        func(task1, task2 *Task) bool
	tasklists TaskList
}

// ----------------------------------------------------------------------------
//  Sort Methods: tasklistSort
// ----------------------------------------------------------------------------

// Len returns the length of the tasklist. It is required to satisfy the
// sort.Interface.
func (ts *tasklistSort) Len() int {
	return len(ts.tasklists)
}

// Swap swaps the position of two tasks in the tasklist. It is required to satisfy
// the sort.Interface.
func (ts *tasklistSort) Swap(l, r int) {
	ts.tasklists[l], ts.tasklists[r] = ts.tasklists[r], ts.tasklists[l]
}

// Less compares two tasks and returns true if the first task is less than the
// second. It is required to satisfy the sort.Interface.
func (ts *tasklistSort) Less(l, r int) bool {
	return ts.by(&ts.tasklists[l], &ts.tasklists[r])
}

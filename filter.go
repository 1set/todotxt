package todotxt

// Filter filters the current TaskList for the given predicate (a function that takes a task as input and returns a bool),
// and returns a new TaskList. The original TaskList is not modified.
func (tasklist *TaskList) Filter(predicate func(Task) bool) *TaskList {
	var newList TaskList
	for _, t := range *tasklist {
		if predicate(t) {
			newList = append(newList, t)
		}
	}
	return &newList
}

// FilterCompleted filters completed tasks.
func FilterCompleted(t Task) bool {
	return t.Completed
}

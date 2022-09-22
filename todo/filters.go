package todo

// ----------------------------------------------------------------------------
//  Filter functions
// ----------------------------------------------------------------------------
//  These functions are used to filter a TaskList for specific tasks.

// FilterCompleted filters completed tasks.
func FilterCompleted(t Task) bool {
	return t.Completed
}

// FilterNotCompleted filters tasks that are not completed.
func FilterNotCompleted(t Task) bool {
	return !t.Completed
}

// FilterDueToday filters tasks that are due today.
func FilterDueToday(t Task) bool {
	return t.IsDueToday()
}

// FilterOverdue filters tasks that are overdue.
func FilterOverdue(t Task) bool {
	return t.IsOverdue()
}

// FilterHasDueDate filters tasks that have due date.
func FilterHasDueDate(t Task) bool {
	return t.HasDueDate()
}

// FilterHasPriority filters tasks that have priority.
func FilterHasPriority(t Task) bool {
	return t.HasPriority()
}

package todotxt

import "strings"

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

// FilterReverse returns a reversed filter for existing predicate.
func FilterReverse(predicate func(Task) bool) func(Task) bool {
	return func(t Task) bool {
		return !predicate(t)
	}
}

// FilterCompleted filters completed tasks.
func FilterCompleted(t Task) bool {
	return t.Completed
}

// FilterDueToday filters tasks that are due today.
func FilterDueToday(t Task) bool {
	return t.IsDueToday()
}

// FilterOverdue filters tasks that are overdue.
func FilterOverdue(t Task) bool {
	return t.IsOverdue()
}

func FilterHasDueDate(t Task) bool {
	return t.HasDueDate()
}

func FilterHasPriority(t Task) bool {
	return t.HasPriority()
}

func FilterByPriority(priority string) func(Task) bool {
	priority = strings.ToUpper(priority)
	return func(t Task) bool {
		return t.Priority == priority
	}
}

func FilterByProject(project string) func(Task) bool {
	return func(t Task) bool {
		for _, p := range t.Projects {
			if strings.EqualFold(p, project) {
				return true
			}
		}
		return false
	}
}

func FilterByContext(context string) func(Task) bool {
	return func(t Task) bool {
		for _, c := range t.Contexts {
			if strings.EqualFold(c, context) {
				return true
			}
		}
		return false
	}
}

package todotxt

import "strings"

// Predicate is a function that takes a task as input and returns a bool.
type Predicate func(Task) bool

// Filter filters the current TaskList for the given predicate,
// and returns a new TaskList. The original TaskList is not modified.
func (tasklist *TaskList) Filter(predicate Predicate) *TaskList {
	var newList TaskList
	for _, t := range *tasklist {
		if predicate(t) {
			newList = append(newList, t)
		}
	}
	return &newList
}

// FilterReverse returns a reversed filter for existing predicate.
func FilterReverse(predicate Predicate) Predicate {
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

// FilterHasDueDate filters tasks that have due date.
func FilterHasDueDate(t Task) bool {
	return t.HasDueDate()
}

// FilterHasPriority filters tasks that have priority.
func FilterHasPriority(t Task) bool {
	return t.HasPriority()
}

// FilterByPriority returns a filter for tasks that have the given priority.
// String comparison in the filters is case-insensitive.
func FilterByPriority(priority string) Predicate {
	priority = strings.ToUpper(priority)
	return func(t Task) bool {
		return t.Priority == priority
	}
}

// FilterByProject returns a filter for tasks that have the given project.
// String comparison in the filters is case-insensitive.
func FilterByProject(project string) Predicate {
	return func(t Task) bool {
		for _, p := range t.Projects {
			if strings.EqualFold(p, project) {
				return true
			}
		}
		return false
	}
}

// FilterByContext returns a filter for tasks that have the given context.
// String comparison in the filters is case-insensitive.
func FilterByContext(context string) Predicate {
	return func(t Task) bool {
		for _, c := range t.Contexts {
			if strings.EqualFold(c, context) {
				return true
			}
		}
		return false
	}
}

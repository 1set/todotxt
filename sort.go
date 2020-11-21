package todotxt

import (
	"errors"
	"sort"
	"time"
)

// TaskSortByType represents type of sorting element and order.
//go:generate stringer -type TaskSortByType -trimprefix Sort -output sort_type.go
type TaskSortByType uint8

// Flags for defining sorting element and order.
const (
	SortTaskIDAsc TaskSortByType = iota + 1
	SortTaskIDDesc
	SortTodoTextAsc
	SortTodoTextDesc
	SortPriorityAsc
	SortPriorityDesc
	SortCreatedDateAsc
	SortCreatedDateDesc
	SortCompletedDateAsc
	SortCompletedDateDesc
	SortDueDateAsc
	SortDueDateDesc
	SortContextAsc
	SortContextDesc
	SortProjectAsc
	SortProjectDesc
)

// Sort allows a TaskList to be sorted by certain predefined fields. Multiple-key sorting is supported.
// See constants Sort* for fields and sort order.
func (tasklist *TaskList) Sort(flag TaskSortByType, flags ...TaskSortByType) error {
	combined := make([]TaskSortByType, len(flags)+1)
	j := 0
	for i := len(flags) - 1; i >= 0; i-- {
		combined[j] = flags[i]
		j++
	}
	combined[j] = flag

	for _, flag := range combined {
		switch flag {
		case SortTaskIDAsc, SortTaskIDDesc:
			tasklist.sortByTaskID(flag)
		case SortTodoTextAsc, SortTodoTextDesc:
			tasklist.sortByTodoText(flag)
		case SortPriorityAsc, SortPriorityDesc:
			tasklist.sortByPriority(flag)
		case SortCreatedDateAsc, SortCreatedDateDesc:
			tasklist.sortByCreatedDate(flag)
		case SortCompletedDateAsc, SortCompletedDateDesc:
			tasklist.sortByCompletedDate(flag)
		case SortDueDateAsc, SortDueDateDesc:
			tasklist.sortByDueDate(flag)
		case SortContextAsc, SortContextDesc:
			tasklist.sortByContext(flag)
		case SortProjectAsc, SortProjectDesc:
			tasklist.sortByProject(flag)
		default:
			return errors.New("unrecognized sort option")
		}
	}
	return nil
}

type tasklistSort struct {
	tasklists TaskList
	by        func(t1, t2 *Task) bool
}

func (ts *tasklistSort) Len() int {
	return len(ts.tasklists)
}

func (ts *tasklistSort) Swap(l, r int) {
	ts.tasklists[l], ts.tasklists[r] = ts.tasklists[r], ts.tasklists[l]
}

func (ts *tasklistSort) Less(l, r int) bool {
	return ts.by(&ts.tasklists[l], &ts.tasklists[r])
}

func (tasklist *TaskList) sortBy(by func(t1, t2 *Task) bool) *TaskList {
	ts := &tasklistSort{
		tasklists: *tasklist,
		by:        by,
	}
	sort.Stable(ts)
	return tasklist
}

func (tasklist *TaskList) sortByTaskID(order TaskSortByType) *TaskList {
	tasklist.sortBy(func(t1, t2 *Task) bool {
		if t1.ID < t2.ID {
			return order == SortTaskIDAsc
		}
		return order == SortTaskIDDesc
	})
	return tasklist
}

func (tasklist *TaskList) sortByTodoText(order TaskSortByType) *TaskList {
	tasklist.sortBy(func(t1, t2 *Task) bool {
		if t1.Todo < t2.Todo {
			return order == SortTodoTextAsc
		}
		return order == SortTodoTextDesc
	})
	return tasklist
}

func (tasklist *TaskList) sortByPriority(order TaskSortByType) *TaskList {
	tasklist.sortBy(func(t1, t2 *Task) bool {
		if order == SortPriorityAsc { // ASC
			if t1.HasPriority() && t2.HasPriority() {
				return t1.Priority < t2.Priority
			}
			return t1.HasPriority()
		}
		// DESC
		if t1.HasPriority() && t2.HasPriority() {
			return t1.Priority > t2.Priority
		}
		return !t1.HasPriority()
	})
	return tasklist
}

func sortByDate(asc bool, hasDate1, hasDate2 bool, date1, date2 time.Time) bool {
	if asc { // ASC
		if hasDate1 && hasDate2 {
			return date1.Before(date2)
		}
		return hasDate2
	}
	// DESC
	if hasDate1 && hasDate2 {
		return date1.After(date2)
	}
	return !hasDate2
}

func (tasklist *TaskList) sortByCreatedDate(order TaskSortByType) *TaskList {
	tasklist.sortBy(func(t1, t2 *Task) bool {
		return sortByDate(order == SortCreatedDateAsc, t1.HasCreatedDate(), t2.HasCreatedDate(), t1.CreatedDate, t2.CreatedDate)
	})
	return tasklist
}

func (tasklist *TaskList) sortByCompletedDate(order TaskSortByType) *TaskList {
	tasklist.sortBy(func(t1, t2 *Task) bool {
		return sortByDate(order == SortCompletedDateAsc, t1.HasCompletedDate(), t2.HasCompletedDate(), t1.CompletedDate, t2.CompletedDate)
	})
	return tasklist
}

func (tasklist *TaskList) sortByDueDate(order TaskSortByType) *TaskList {
	tasklist.sortBy(func(t1, t2 *Task) bool {
		return sortByDate(order == SortDueDateAsc, t1.HasDueDate(), t2.HasDueDate(), t1.DueDate, t2.DueDate)
	})
	return tasklist
}

// lessStrings checks if the string slices a is exactly less than b in lexicographical order.
func lessStrings(a, b []string) bool {
	la, lb, min := len(a), len(b), 0
	if la == 0 && lb == 0 {
		return false
	} else if la == 0 && lb > 0 {
		return false
	} else if la > 0 && lb == 0 {
		return true
	}

	if min = la; la > lb {
		min = lb
	}
	for i := 0; i < min; i++ {
		if a[i] < b[i] {
			return true
		} else if a[i] > b[i] {
			return false
		}
	}

	if la == lb {
		return false
	}
	return la < lb
}

func (tasklist *TaskList) sortByContext(order TaskSortByType) *TaskList {
	tasklist.sortBy(func(t1, t2 *Task) bool {
		if order == SortContextAsc {
			return lessStrings(t1.Contexts, t2.Contexts)
		}
		return lessStrings(t2.Contexts, t1.Contexts)
	})
	return tasklist
}

func (tasklist *TaskList) sortByProject(order TaskSortByType) *TaskList {
	tasklist.sortBy(func(t1, t2 *Task) bool {
		if order == SortProjectAsc {
			return lessStrings(t1.Projects, t2.Projects)
		}
		return lessStrings(t2.Projects, t1.Projects)
	})
	return tasklist
}

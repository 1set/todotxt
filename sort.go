package todotxt

import (
	"errors"
	"sort"
	"time"
)

// Flags for defining sort element and order.
const (
	SortPriorityAsc = iota + 1
	SortPriorityDesc
	SortCreatedDateAsc
	SortCreatedDateDesc
	SortCompletedDateAsc
	SortCompletedDateDesc
	SortDueDateAsc
	SortDueDateDesc
)

// Sort allows a TaskList to be sorted by certain predefined fields.
// See constants Sort* for fields and sort order.
func (tasklist *TaskList) Sort(sortFlag int) error {
	switch sortFlag {
	case SortPriorityAsc, SortPriorityDesc:
		tasklist.sortByPriority(sortFlag)
	case SortCreatedDateAsc, SortCreatedDateDesc:
		tasklist.sortByCreatedDate(sortFlag)
	case SortCompletedDateAsc, SortCompletedDateDesc:
		tasklist.sortByCompletedDate(sortFlag)
	case SortDueDateAsc, SortDueDateDesc:
		tasklist.sortByDueDate(sortFlag)
	default:
		return errors.New("unrecognized sort option")
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

func (tasklist *TaskList) sortByPriority(order int) *TaskList {
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

func (tasklist *TaskList) sortByCreatedDate(order int) *TaskList {
	tasklist.sortBy(func(t1, t2 *Task) bool {
		return sortByDate(order == SortCreatedDateAsc, t1.HasCreatedDate(), t2.HasCreatedDate(), t1.CreatedDate, t2.CreatedDate)
	})
	return tasklist
}

func (tasklist *TaskList) sortByCompletedDate(order int) *TaskList {
	tasklist.sortBy(func(t1, t2 *Task) bool {
		return sortByDate(order == SortCompletedDateAsc, t1.HasCompletedDate(), t2.HasCompletedDate(), t1.CompletedDate, t2.CompletedDate)
	})
	return tasklist
}

func (tasklist *TaskList) sortByDueDate(order int) *TaskList {
	tasklist.sortBy(func(t1, t2 *Task) bool {
		return sortByDate(order == SortDueDateAsc, t1.HasDueDate(), t2.HasDueDate(), t1.DueDate, t2.DueDate)
	})
	return tasklist
}

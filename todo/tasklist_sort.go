package todo

import (
	"sort"

	"github.com/pkg/errors"
)

// ----------------------------------------------------------------------------
//  TaskList.Sort()
// ----------------------------------------------------------------------------

// Sort allows a TaskList to be sorted by certain predefined fields. Multiple-key
// sorting is supported. See constants Sort* for fields and sort order.
//
//nolint:cyclop // complexity is 12 but leave it as is
func (tasklist *TaskList) Sort(flag TaskSortByType, flags ...TaskSortByType) error {
	lenFlags := len(flags)
	combined := make([]TaskSortByType, lenFlags+1)
	index := 0

	for i := lenFlags - 1; i >= 0; i-- {
		combined[index] = flags[i]
		index++
	}

	combined[index] = flag

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

// ----------------------------------------------------------------------------
//  TaskList.Sort() helpers
// ----------------------------------------------------------------------------

//nolint:unparam // false positive
func (tasklist *TaskList) sortBy(by func(task1, task2 *Task) bool) *TaskList {
	ts := &tasklistSort{
		tasklists: *tasklist,
		by:        by,
	}
	sort.Stable(ts)

	return tasklist
}

func (tasklist *TaskList) sortByCompletedDate(order TaskSortByType) *TaskList {
	tasklist.sortBy(func(task1, task2 *Task) bool {
		return sortByDate(
			order == SortCompletedDateAsc, // is asc
			task1.HasCompletedDate(),      // hasDate1
			task2.HasCompletedDate(),      // hasDate2
			task1.CompletedDate,           // date1
			task2.CompletedDate,           // date2
		)
	})

	return tasklist
}

func (tasklist *TaskList) sortByContext(order TaskSortByType) *TaskList {
	tasklist.sortBy(func(task1, task2 *Task) bool {
		if order == SortContextAsc {
			return lessStrings(task1.Contexts, task2.Contexts)
		}

		return lessStrings(task2.Contexts, task1.Contexts)
	})

	return tasklist
}

func (tasklist *TaskList) sortByCreatedDate(order TaskSortByType) *TaskList {
	tasklist.sortBy(func(task1, task2 *Task) bool {
		return sortByDate(
			order == SortCreatedDateAsc, // is asc
			task1.HasCreatedDate(),      // hasDate1
			task2.HasCreatedDate(),      // hasDate2
			task1.CreatedDate,           // date1
			task2.CreatedDate,           // date2
		)
	})

	return tasklist
}

func (tasklist *TaskList) sortByDueDate(order TaskSortByType) *TaskList {
	tasklist.sortBy(func(task1, task2 *Task) bool {
		return sortByDate(
			order == SortDueDateAsc, // is asc
			task1.HasDueDate(),      // hasDate1
			task2.HasDueDate(),      // hasDate2
			task1.DueDate,           // date1
			task2.DueDate,           // date2
		)
	})

	return tasklist
}

func (tasklist *TaskList) sortByPriority(order TaskSortByType) *TaskList {
	tasklist.sortBy(func(task1, task2 *Task) bool {
		// ASC
		if order == SortPriorityAsc {
			if task1.HasPriority() && task2.HasPriority() {
				return task1.Priority < task2.Priority
			}

			return task1.HasPriority()
		}

		// DESC
		if task1.HasPriority() && task2.HasPriority() {
			return task1.Priority > task2.Priority
		}

		return !task1.HasPriority()
	})

	return tasklist
}

func (tasklist *TaskList) sortByProject(order TaskSortByType) *TaskList {
	tasklist.sortBy(func(task1, task2 *Task) bool {
		if order == SortProjectAsc {
			return lessStrings(task1.Projects, task2.Projects)
		}

		return lessStrings(task2.Projects, task1.Projects)
	})

	return tasklist
}

func (tasklist *TaskList) sortByTaskID(order TaskSortByType) *TaskList {
	tasklist.sortBy(func(task1, task2 *Task) bool {
		if task1.ID < task2.ID {
			return order == SortTaskIDAsc
		}

		return order == SortTaskIDDesc
	})

	return tasklist
}

func (tasklist *TaskList) sortByTodoText(order TaskSortByType) *TaskList {
	tasklist.sortBy(func(task1, task2 *Task) bool {
		if task1.Todo < task2.Todo {
			return order == SortTodoTextAsc
		}

		return order == SortTodoTextDesc
	})

	return tasklist
}

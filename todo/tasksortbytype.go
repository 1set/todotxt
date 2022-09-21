package todo

// ----------------------------------------------------------------------------
//  Type: TaskSortByType
// ----------------------------------------------------------------------------

// TaskSortByType represents type of sorting element and order.
//
// The stringer implementation `String()` is defined in sort_type.go.
// See doc.go as well.
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

package todo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

//nolint:paralleltest // do not parallel to avoid race conditions
func TestTaskList_Filter(t *testing.T) {
	// Load test data
	testTasklist := testLoadFromPath(t, testInputTasklist)

	// Filter list to get only completed tasks
	{
		completedList := testTasklist.Filter(FilterCompleted)

		expectLen := 33
		actualLen := len(completedList)
		require.Equal(t, expectLen, actualLen,
			"filtering by completed tasks did not work as expected:\n%s", completedList.String())
	}

	// Filter list to get only tasks with a due date
	{
		dueDateList := testTasklist.Filter(FilterHasDueDate)

		expectLen := 26
		actualLen := len(dueDateList)
		require.Equal(t, expectLen, actualLen,
			"filtering by tasks with a due date did not work as expected:\n%s", dueDateList.String())
	}

	// Filter list to get only tasks with "B" priority
	{
		prioBList := testTasklist.Filter(FilterByPriority("b"))

		expectLen := 17
		actualLen := len(prioBList)
		require.Equal(t, expectLen, actualLen,
			"filtering by tasks with 'B' priority did not work as expected:\n%s", prioBList.String())
	}
}

//nolint:paralleltest // do not parallel to avoid race conditions
func TestTaskList_Filter_filter_by_multiple_predicates(t *testing.T) {
	// Load test data
	testTasklist := testLoadFromPath(t, testInputFilter)

	now := time.Now()
	testTasklist[0].DueDate = now.AddDate(0, 0, -2)
	testTasklist[1].DueDate = now
	testTasklist[2].DueDate = now.AddDate(0, 0, 1)

	// and -- filters tasks with priority and is overdue and is completed
	{
		filteredList := testTasklist.Filter(FilterHasPriority).Filter(FilterOverdue).Filter(FilterCompleted)

		expectLen := 2
		actualLen := len(filteredList)
		require.Equal(t, expectLen, actualLen,
			"'and' filter failed. filtering by multiple predicates did not work as expected:\n%s", filteredList.String())
	}

	// or -- filters tasks with priority A or B
	{
		filteredList := testTasklist.Filter(FilterByPriority("a"), FilterByPriority("b"))

		expectLen := 7
		actualLen := len(filteredList)
		require.Equal(t, expectLen, actualLen,
			"'or' filter failed. filtering by multiple predicates did not work as expected:\n%s", filteredList.String())
	}

	// or -- filters incompleted tasks that are overdue or have set priority
	{
		filteredList := testTasklist.Filter(FilterNot(FilterCompleted)).Filter(FilterOverdue, FilterHasPriority)

		expectLen := 10
		actualLen := len(filteredList)
		require.Equal(t, expectLen, actualLen,
			"'or' filter failed. filtering by multiple predicates did not work as expected:\n%s", filteredList.String())
	}
}

func TestTaskList_Filter_Helpers(t *testing.T) {
	t.Parallel()

	// Load test data
	testTasklist := testLoadFromPath(t, testInputFilter)

	now := time.Now()
	testTasklist[0].DueDate = now.AddDate(0, 0, -2)
	testTasklist[1].DueDate = now
	testTasklist[2].DueDate = now.AddDate(0, 0, 1)

	for testNum, test := range []struct {
		predicate Predicate
		expectLen int
	}{
		{FilterCompleted, 9},
		{FilterNotCompleted, 17},
		{FilterNot(FilterCompleted), 17},
		{FilterNot(FilterNot(FilterCompleted)), 9},
		{FilterDueToday, 1},
		{FilterOverdue, 9},
		{FilterHasDueDate, 12},
		{FilterHasPriority, 9},
		{FilterByPriority("a"), 2},
		{FilterByPriority("B"), 5},
		{FilterByPriority("c"), 2},
		{FilterByPriority("e"), 0},
		{FilterByProject("unknown"), 0},
		{FilterByProject("Family"), 2},
		{FilterByProject("planning"), 2},
		{FilterByContext("unknown"), 0},
		{FilterByContext("call"), 2},
		{FilterByContext("go"), 9},
	} {
		filteredList := testTasklist.Filter(test.predicate)

		expectLen := test.expectLen
		actualLen := len(filteredList)
		require.Equal(t, expectLen, actualLen,
			"test case #%d failed. It did not filter as expected:\n%s", testNum+1, filteredList.String())
	}
}

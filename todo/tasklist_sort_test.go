package todo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTaskSortByType(t *testing.T) {
	t.Parallel()

	names := map[TaskSortByType]string{
		SortTaskIDAsc:         "TaskIDAsc",
		SortTaskIDDesc:        "TaskIDDesc",
		SortTodoTextAsc:       "TodoTextAsc",
		SortTodoTextDesc:      "TodoTextDesc",
		SortPriorityAsc:       "PriorityAsc",
		SortPriorityDesc:      "PriorityDesc",
		SortCreatedDateAsc:    "CreatedDateAsc",
		SortCreatedDateDesc:   "CreatedDateDesc",
		SortCompletedDateAsc:  "CompletedDateAsc",
		SortCompletedDateDesc: "CompletedDateDesc",
		SortDueDateAsc:        "DueDateAsc",
		SortDueDateDesc:       "DueDateDesc",
		SortContextAsc:        "ContextAsc",
		SortContextDesc:       "ContextDesc",
		SortProjectAsc:        "ProjectAsc",
		SortProjectDesc:       "ProjectDesc",
		0:                     "TaskSortByType(0)",
	}

	for nameType, expect := range names {
		actual := nameType.String()

		require.Equal(t, expect, actual, "the String method returned an unexpected name for TaskSortByType")
	}
}

func TestTaskList_Sort_sort_by_priority(t *testing.T) {
	t.Parallel()

	// Load the test tasklist
	testTasklist := testLoadFromPath(t, testInputSort)

	const taskID = 0

	// Get 6 tasks from the test tasklist
	actualTasklist := testTasklist[taskID : taskID+6]

	// SortPriorityAsc
	{
		require.NoError(t, actualTasklist.Sort(SortPriorityAsc), "sorting by SortPriorityAsc failed")

		expectTasklist := []string{
			"(A) 2012-01-30 Call Mom @Call @Phone +Family",
			"x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt",
			"x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12",
			"(D) 2013-12-01 Outline chapter 5 @Computer +Novel Level:5 private:false due:2014-02-17",
			"2013-02-22 Pick up milk @GroceryStore",
			"x 2014-01-03 Create golang library @Go +go-todotxt due:2014-01-05",
		}
		checkTaskListOrder(t, actualTasklist, expectTasklist)
	}

	// SortPriorityDesc
	{
		require.NoError(t, actualTasklist.Sort(SortPriorityDesc), "sorting by SortPriorityDesc failed")

		expectTasklist := []string{
			"x 2014-01-03 Create golang library @Go +go-todotxt due:2014-01-05",
			"2013-02-22 Pick up milk @GroceryStore",
			"(D) 2013-12-01 Outline chapter 5 @Computer +Novel Level:5 private:false due:2014-02-17",
			"x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12",
			"x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt",
			"(A) 2012-01-30 Call Mom @Call @Phone +Family",
		}
		checkTaskListOrder(t, actualTasklist, expectTasklist)
	}
}

func TestTaskList_Sort_sort_by_created_date(t *testing.T) {
	t.Parallel()

	// Load the test tasklist
	testTasklist := testLoadFromPath(t, testInputSort)

	const taskID = 6

	// Get 5 tasks from the test tasklist
	actualTasklist := testTasklist[taskID : taskID+5]

	// SortCreatedDateAsc
	{
		require.NoError(t, actualTasklist.Sort(SortCreatedDateAsc), "sorting by SortCreatedDateAsc failed")

		expectTasklist := []string{
			"x 2014-01-03 Create golang library @Go +go-todotxt due:2014-01-05",
			"(A) Call Mom @Call @Phone +Family",
			"2013-02-22 Pick up milk @GroceryStore",
			"x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt",
			"x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12",
		}
		checkTaskListOrder(t, actualTasklist, expectTasklist)
	}

	// SortCreatedDateDesc
	{
		require.NoError(t, actualTasklist.Sort(SortCreatedDateDesc), "sorting by SortCreatedDateDesc failed")

		expectTasklist := []string{
			"x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12",
			"x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt",
			"2013-02-22 Pick up milk @GroceryStore",
			"(A) Call Mom @Call @Phone +Family",
			"x 2014-01-03 Create golang library @Go +go-todotxt due:2014-01-05",
		}
		checkTaskListOrder(t, actualTasklist, expectTasklist)
	}
}

//nolint:paralleltest // do not parallel to avoid race conditions
func TestTaskList_Sort_sort_by_completed_date(t *testing.T) {
	// Load the test tasklist
	testTasklist := testLoadFromPath(t, testInputSort)

	const taskID = 11

	// Get 6 tasks from the test tasklist
	actualTasklist := testTasklist[taskID : taskID+6]

	// SortCompletedDateAsc
	{
		require.NoError(t, actualTasklist.Sort(SortCompletedDateAsc), "sorting by SortCompletedDateAsc failed")

		expectTasklist := []string{
			"x Download Todo.txt mobile app @Phone",
			"x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12",
			"2013-02-22 Pick up milk @GroceryStore",
			"x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt",
			"x 2014-01-03 Create golang library @Go +go-todotxt due:2014-01-05",
			"x 2014-01-04 2014-01-01 Create some more golang library test cases @Go +go-todotxt",
		}
		checkTaskListOrder(t, actualTasklist, expectTasklist)
	}

	// SortCompletedDateDesc
	{
		require.NoError(t, actualTasklist.Sort(SortCompletedDateDesc), "sorting by SortCompletedDateDesc failed")

		expectTasklist := []string{
			"x 2014-01-04 2014-01-01 Create some more golang library test cases @Go +go-todotxt",
			"x 2014-01-03 Create golang library @Go +go-todotxt due:2014-01-05",
			"x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt",
			"2013-02-22 Pick up milk @GroceryStore",
			"x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12",
			"x Download Todo.txt mobile app @Phone",
		}
		checkTaskListOrder(t, actualTasklist, expectTasklist)
	}
}

func TestTaskList_Sort_sort_by_due_date(t *testing.T) {
	t.Parallel()

	// Load the test tasklist
	testTasklist := testLoadFromPath(t, testInputSort)

	const taskID = 17

	// Get 4 tasks from the test tasklist
	actualTasklist := testTasklist[taskID : taskID+4]

	// SortDueDateAsc
	{
		require.NoError(t, actualTasklist.Sort(SortDueDateAsc), "sorting by SortDueDateAsc failed")

		expectTasklist := []string{
			"x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt",
			"x 2014-01-03 Create golang library @Go +go-todotxt due:2014-01-05",
			"x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12",
			"(B) 2013-12-01 Outline chapter 5 @Computer +Novel Level:5 private:false due:2014-02-17",
		}
		checkTaskListOrder(t, actualTasklist, expectTasklist)
	}

	// SortDueDateDesc
	{
		require.NoError(t, actualTasklist.Sort(SortDueDateDesc), "sorting by SortDueDateDesc failed")

		expectTasklist := []string{
			"(B) 2013-12-01 Outline chapter 5 @Computer +Novel Level:5 private:false due:2014-02-17",
			"x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12",
			"x 2014-01-03 Create golang library @Go +go-todotxt due:2014-01-05",
			"x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt",
		}
		checkTaskListOrder(t, actualTasklist, expectTasklist)
	}
}

func TestTaskList_Sort_sort_by_task_id(t *testing.T) {
	t.Parallel()

	// Load the test tasklist
	testTasklist := testLoadFromPath(t, testInputSort)

	const taskID = 21

	// Get 5 tasks from the test tasklist
	actualTasklist := testTasklist[taskID : taskID+5]

	// Pre-sort the tasklist by priority
	require.NoError(t, actualTasklist.Sort(SortPriorityAsc))

	// SortTaskIDAsc
	{
		require.NoError(t, actualTasklist.Sort(SortTaskIDAsc), "sorting by SortTaskIDAsc failed")

		expectTasklist := []string{
			"(B) Task 1",
			"(A) Task 2",
			"Task 3 due:2020-11-11",
			"(C) Task 4 due:2020-12-12",
			"x Task 5",
		}
		checkTaskListOrder(t, actualTasklist, expectTasklist)
	}

	// SortTaskIDDesc
	{
		require.NoError(t, actualTasklist.Sort(SortTaskIDDesc), "sorting by SortTaskIDDesc failed")

		expectTasklist := []string{
			"x Task 5",
			"(C) Task 4 due:2020-12-12",
			"Task 3 due:2020-11-11",
			"(A) Task 2",
			"(B) Task 1",
		}
		checkTaskListOrder(t, actualTasklist, expectTasklist)
	}
}

func TestTaskList_Sort_sort_by_context(t *testing.T) {
	t.Parallel()

	// Load the test tasklist
	testTasklist := testLoadFromPath(t, testInputSort)

	const taskID = 26

	// Get 6 tasks from the test tasklist
	actualTasklist := testTasklist[taskID : taskID+6]

	// Pre-sort the tasklist by creation date
	require.NoError(t, actualTasklist.Sort(SortCreatedDateAsc))

	// SortContextAsc
	{
		require.NoError(t, actualTasklist.Sort(SortContextAsc), "sorting by SortContextAsc failed")

		expectTasklist := []string{
			"2020-12-19 Task 3 @Apple",
			"2020-10-19 Task 2 @Apple @Banana",
			"2020-11-09 Task 1 @Apple @Banana",
			"2020-11-11 Task 6 @Apple @Coconut",
			"2020-11-19 Task 4 @Banana",
			"2020-12-09 Task 5",
		}
		checkTaskListOrder(t, actualTasklist, expectTasklist)
	}

	// SortContextDesc
	{
		require.NoError(t, actualTasklist.Sort(SortContextDesc), "sorting by SortContextDesc failed")

		expectTasklist := []string{
			"2020-12-09 Task 5",
			"2020-11-19 Task 4 @Banana",
			"2020-11-11 Task 6 @Apple @Coconut",
			"2020-10-19 Task 2 @Apple @Banana",
			"2020-11-09 Task 1 @Apple @Banana",
			"2020-12-19 Task 3 @Apple",
		}
		checkTaskListOrder(t, actualTasklist, expectTasklist)
	}
}

func TestTaskList_Sort_sort_by_project(t *testing.T) {
	t.Parallel()

	// Load the test tasklist
	testTasklist := testLoadFromPath(t, testInputSort)

	const taskID = 32

	// Get 6 tasks from the test tasklist
	actualTasklist := testTasklist[taskID : taskID+6]

	// Pre-sort the tasklist by creation date
	require.NoError(t, actualTasklist.Sort(SortCreatedDateAsc), "sorting by SortCreatedDateAsc failed")

	// SortProjectAsc
	{
		require.NoError(t, actualTasklist.Sort(SortProjectAsc), "sorting by SortProjectAsc failed")

		expectTasklist := []string{
			"2020-12-29 Task 3 +Apple",
			"2020-10-09 Task 1 +Apple +Banana",
			"2020-10-19 Task 2 +Apple +Banana",
			"2020-12-19 Task 4 +Banana",
			"2020-11-11 Task 6 +Coconut",
			"2020-12-09 Task 5",
		}
		checkTaskListOrder(t, actualTasklist, expectTasklist)
	}

	// SortProjectDesc
	{
		require.NoError(t, actualTasklist.Sort(SortProjectDesc),
			"sort SortProjectDesc failed")

		expectTasklist := []string{
			"2020-12-09 Task 5",
			"2020-11-11 Task 6 +Coconut",
			"2020-12-19 Task 4 +Banana",
			"2020-10-09 Task 1 +Apple +Banana",
			"2020-10-19 Task 2 +Apple +Banana",
			"2020-12-29 Task 3 +Apple",
		}
		checkTaskListOrder(t, actualTasklist, expectTasklist)
	}
}

func TestTaskList_Sort_sort_by_todo_text(t *testing.T) {
	t.Parallel()

	// Load the test tasklist
	testTasklist := testLoadFromPath(t, testInputSort)

	const taskID = 38

	// Get 5 tasks from the test tasklist
	actualTasklist := testTasklist[taskID : taskID+5]

	// SortTodoTextAsc
	{
		require.NoError(t, actualTasklist.Sort(SortTodoTextAsc), "sorting by SortTodoTextAsc failed")

		expectTasklist := []string{
			"2020-10-09 Task 1 +Apple +Banana",
			"2020-10-19 Task 2 +Apple +Brown",
			"2020-12-29 Task 3 +Apple",
			"2020-12-19 Task 4 +Banana",
			"2020-11-11 Task 5 +Coconut",
		}
		checkTaskListOrder(t, actualTasklist, expectTasklist)
	}

	// SortTodoTextDesc
	{
		require.NoError(t, actualTasklist.Sort(SortTodoTextDesc), "sorting by SortTodoTextDesc failed")

		expectTasklist := []string{
			"2020-11-11 Task 5 +Coconut",
			"2020-12-19 Task 4 +Banana",
			"2020-12-29 Task 3 +Apple",
			"2020-10-19 Task 2 +Apple +Brown",
			"2020-10-09 Task 1 +Apple +Banana",
		}
		checkTaskListOrder(t, actualTasklist, expectTasklist)
	}
}

//nolint:funlen // function is long but leave it as is
func TestTaskList_Sort_sort_by_multiple_flags(t *testing.T) {
	t.Parallel()

	// Load the test tasklist
	testTasklist := testLoadFromPath(t, testInputSort)

	const taskID = 43

	// Get 7 tasks from the test tasklist
	actualTasklist := testTasklist[taskID : taskID+7]

	// SortTodoTextAsc, SortPriorityDesc combo
	{
		require.NoError(t, actualTasklist.Sort(SortTodoTextAsc, SortPriorityDesc),
			"sorting by SortTodoTextAsc and SortPriorityDesc failed")

		expectTasklist := []string{
			"(A) 2020-10-09 Task 1 +Apple +Banana",
			"2020-10-19 Task 2 +Apple +Brown",
			"2020-12-29 Task 3",
			"(C) 2020-12-29 Task 3 +Apple",
			"2020-12-19 Task 4 +Banana",
			"(D) 2020-11-11 Task 5 +Coconut",
			"2020-12-29 Task 6",
		}
		checkTaskListOrder(t, actualTasklist, expectTasklist)
	}

	// SortPriorityAsc, SortTodoTextAsc combo
	{
		require.NoError(t, actualTasklist.Sort(SortPriorityAsc, SortTodoTextAsc),
			"sorting by SortPriorityAsc and SortTodoTextAsc failed")

		expectTasklist := []string{
			"(A) 2020-10-09 Task 1 +Apple +Banana",
			"(C) 2020-12-29 Task 3 +Apple",
			"(D) 2020-11-11 Task 5 +Coconut",
			"2020-10-19 Task 2 +Apple +Brown",
			"2020-12-29 Task 3",
			"2020-12-19 Task 4 +Banana",
			"2020-12-29 Task 6",
		}
		checkTaskListOrder(t, actualTasklist, expectTasklist)
	}

	// SortPriorityAsc, SortCreatedDateAsc, SortTodoTextDesc combo
	{
		require.NoError(t, actualTasklist.Sort(SortPriorityAsc, SortCreatedDateAsc, SortTodoTextDesc),
			"sorting by SortPriorityAsc, SortCreatedDateAsc and SortTodoTextDesc failed")

		expectTasklist := []string{
			"(A) 2020-10-09 Task 1 +Apple +Banana",
			"(C) 2020-12-29 Task 3 +Apple",
			"(D) 2020-11-11 Task 5 +Coconut",
			"2020-10-19 Task 2 +Apple +Brown",
			"2020-12-19 Task 4 +Banana",
			"2020-12-29 Task 6",
			"2020-12-29 Task 3",
		}
		checkTaskListOrder(t, actualTasklist, expectTasklist)
	}

	// SortPriorityAsc, SortProjectAsc, SortTaskIDAsc combo
	{
		require.NoError(t, actualTasklist.Sort(SortPriorityAsc, SortProjectAsc, SortTaskIDAsc),
			"sorting by SortPriorityAsc, SortProjectAsc and SortTaskIDAsc failed")

		expectTasklist := []string{
			"(A) 2020-10-09 Task 1 +Apple +Banana",
			"(C) 2020-12-29 Task 3 +Apple",
			"(D) 2020-11-11 Task 5 +Coconut",
			"2020-10-19 Task 2 +Apple +Brown",
			"2020-12-19 Task 4 +Banana",
			"2020-12-29 Task 3",
			"2020-12-29 Task 6",
		}
		checkTaskListOrder(t, actualTasklist, expectTasklist)
	}
}

func TestTaskList_Sort_error(t *testing.T) {
	t.Parallel()

	testTasklist := testLoadFromPath(t, testInputSort)
	unknownSortType := TaskSortByType(123)

	err := testTasklist.Sort(unknownSortType)

	require.Error(t, err, "unrecognized sort flags should return error")
	require.Contains(t, err.Error(), "unrecognized sort option")
}

func Test_lessStrings(t *testing.T) {
	t.Parallel()

	for index, test := range []struct {
		a    []string
		b    []string
		want bool
	}{
		{[]string{"a", "b", "c"}, []string{"a", "b", "c"}, false},
		{[]string{"a", "b", "c"}, []string{"a", "b"}, false},
		{[]string{"a", "b", "c"}, []string{"a", "c"}, true},
		{[]string{"a", "b", "c"}, []string{"b"}, true},
		{[]string{"a", "b", "c"}, []string{}, true},
		{[]string{"a", "b"}, []string{"a", "b", "c"}, true},
		{[]string{"a", "b"}, []string{"a", "a"}, false},
		{[]string{"a", "b"}, []string{"a", "c"}, true},
		{[]string{"a", "b"}, []string{"b"}, true},
		{[]string{"a", "a"}, []string{"a", "b", "c"}, true},
		{[]string{"a", "a"}, []string{"a", "b"}, true},
		{[]string{"a", "a"}, []string{"a", "a"}, false},
		{[]string{"a", "a"}, []string{"a", "c"}, true},
		{[]string{"a", "a"}, []string{"b"}, true},
		{[]string{"a", "a"}, []string{"b", "a"}, true},
		{[]string{"a", "a"}, []string{"c"}, true},
		{[]string{"a", "c"}, []string{"a", "b", "c"}, false},
		{[]string{"a", "c"}, []string{"a", "c"}, false},
		{[]string{"a", "c"}, []string{"b"}, true},
		{[]string{"a", "c"}, []string{"c"}, true},
		{[]string{"b"}, []string{"a", "b", "c"}, false},
		{[]string{"b"}, []string{"a", "c"}, false},
		{[]string{"b"}, []string{"b"}, false},
		{[]string{"b"}, []string{"b", "a"}, true},
		{[]string{"b"}, []string{"c"}, true},
		{[]string{"b"}, []string{}, true},
		{[]string{"b", "a"}, []string{"a", "b", "c"}, false},
		{[]string{"b", "a"}, []string{"b"}, false},
		{[]string{"b", "a"}, []string{"b", "a"}, false},
		{[]string{"b", "a"}, []string{"c"}, true},
		{[]string{"c"}, []string{"a", "b", "c"}, false},
		{[]string{"c"}, []string{"b"}, false},
		{[]string{}, []string{"a", "b", "c"}, false},
		{[]string{}, []string{"c"}, false},
		{[]string{}, []string{}, false},
	} {
		expect := test.want
		actual := lessStrings(test.a, test.b)

		require.Equal(t, expect, actual,
			"test case #%v failed. lessStrings(%v, %v) = %v, want %v",
			index+1, test.a, test.b, actual, expect,
		)
	}
}

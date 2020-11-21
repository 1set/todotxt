package todotxt

import (
	"fmt"
	"testing"
)

func BenchmarkTaskList_Sort(b *testing.B) {
	testTasklist.LoadFromPath(testInputSort)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = testTasklist.Sort(SortPriorityAsc, SortCreatedDateAsc, SortTodoTextDesc)
	}
}

func TestTaskSortByType(t *testing.T) {
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
	for n, s := range names {
		if ss := n.String(); ss != s {
			t.Errorf("Expected TaskSortByType %v is %q, but got: %q", n, s, ss)
		}
	}
}

func TestTaskSortByPriority(t *testing.T) {
	testTasklist.LoadFromPath(testInputSort)
	taskID := 0

	testTasklist = testTasklist[taskID : taskID+6]

	if err := testTasklist.Sort(SortPriorityAsc); err != nil {
		t.Fatal(err)
	}
	testExpectedList = []string{
		"(A) 2012-01-30 Call Mom @Call @Phone +Family",
		"x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt",
		"x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12",
		"(D) 2013-12-01 Outline chapter 5 @Computer +Novel Level:5 private:false due:2014-02-17",
		"2013-02-22 Pick up milk @GroceryStore",
		"x 2014-01-03 Create golang library @Go +go-todotxt due:2014-01-05",
	}
	checkTaskListOrder(t, testTasklist, testExpectedList)

	if err := testTasklist.Sort(SortPriorityDesc); err != nil {
		t.Fatal(err)
	}
	testExpectedList = []string{
		"x 2014-01-03 Create golang library @Go +go-todotxt due:2014-01-05",
		"2013-02-22 Pick up milk @GroceryStore",
		"(D) 2013-12-01 Outline chapter 5 @Computer +Novel Level:5 private:false due:2014-02-17",
		"x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12",
		"x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt",
		"(A) 2012-01-30 Call Mom @Call @Phone +Family",
	}
	checkTaskListOrder(t, testTasklist, testExpectedList)
}

func TestTaskSortByCreatedDate(t *testing.T) {
	testTasklist.LoadFromPath(testInputSort)
	taskID := 6

	testTasklist = testTasklist[taskID : taskID+5]

	if err := testTasklist.Sort(SortCreatedDateAsc); err != nil {
		t.Fatal(err)
	}
	testExpectedList = []string{
		"x 2014-01-03 Create golang library @Go +go-todotxt due:2014-01-05",
		"(A) Call Mom @Call @Phone +Family",
		"2013-02-22 Pick up milk @GroceryStore",
		"x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt",
		"x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12",
	}
	checkTaskListOrder(t, testTasklist, testExpectedList)

	if err := testTasklist.Sort(SortCreatedDateDesc); err != nil {
		t.Fatal(err)
	}
	testExpectedList = []string{
		"x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12",
		"x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt",
		"2013-02-22 Pick up milk @GroceryStore",
		"(A) Call Mom @Call @Phone +Family",
		"x 2014-01-03 Create golang library @Go +go-todotxt due:2014-01-05",
	}
	checkTaskListOrder(t, testTasklist, testExpectedList)
}

func TestTaskSortByCompletedDate(t *testing.T) {
	testTasklist.LoadFromPath(testInputSort)
	taskID := 11

	testTasklist = testTasklist[taskID : taskID+6]

	if err := testTasklist.Sort(SortCompletedDateAsc); err != nil {
		t.Fatal(err)
	}
	testExpectedList = []string{
		"x Download Todo.txt mobile app @Phone",
		"x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12",
		"2013-02-22 Pick up milk @GroceryStore",
		"x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt",
		"x 2014-01-03 Create golang library @Go +go-todotxt due:2014-01-05",
		"x 2014-01-04 2014-01-01 Create some more golang library test cases @Go +go-todotxt",
	}
	checkTaskListOrder(t, testTasklist, testExpectedList)

	if err := testTasklist.Sort(SortCompletedDateDesc); err != nil {
		t.Fatal(err)
	}
	testExpectedList = []string{
		"x 2014-01-04 2014-01-01 Create some more golang library test cases @Go +go-todotxt",
		"x 2014-01-03 Create golang library @Go +go-todotxt due:2014-01-05",
		"x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt",
		"2013-02-22 Pick up milk @GroceryStore",
		"x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12",
		"x Download Todo.txt mobile app @Phone",
	}
	checkTaskListOrder(t, testTasklist, testExpectedList)
}

func TestTaskSortByDueDate(t *testing.T) {
	testTasklist.LoadFromPath(testInputSort)
	taskID := 17

	testTasklist = testTasklist[taskID : taskID+4]

	if err := testTasklist.Sort(SortDueDateAsc); err != nil {
		t.Fatal(err)
	}
	testExpectedList = []string{
		"x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt",
		"x 2014-01-03 Create golang library @Go +go-todotxt due:2014-01-05",
		"x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12",
		"(B) 2013-12-01 Outline chapter 5 @Computer +Novel Level:5 private:false due:2014-02-17",
	}
	checkTaskListOrder(t, testTasklist, testExpectedList)

	if err := testTasklist.Sort(SortDueDateDesc); err != nil {
		t.Fatal(err)
	}
	testExpectedList = []string{
		"(B) 2013-12-01 Outline chapter 5 @Computer +Novel Level:5 private:false due:2014-02-17",
		"x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12",
		"x 2014-01-03 Create golang library @Go +go-todotxt due:2014-01-05",
		"x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt",
	}
	checkTaskListOrder(t, testTasklist, testExpectedList)
}

func TestTaskSortByTaskID(t *testing.T) {
	testTasklist.LoadFromPath(testInputSort)
	taskID := 21

	testTasklist = testTasklist[taskID : taskID+5]

	if err := testTasklist.Sort(SortPriorityAsc); err != nil {
		t.Fatal(err)
	}

	if err := testTasklist.Sort(SortTaskIDAsc); err != nil {
		t.Fatal(err)
	}
	testExpectedList = []string{
		"(B) Task 1",
		"(A) Task 2",
		"Task 3 due:2020-11-11",
		"(C) Task 4 due:2020-12-12",
		"x Task 5",
	}
	checkTaskListOrder(t, testTasklist, testExpectedList)

	if err := testTasklist.Sort(SortTaskIDDesc); err != nil {
		t.Fatal(err)
	}
	testExpectedList = []string{
		"x Task 5",
		"(C) Task 4 due:2020-12-12",
		"Task 3 due:2020-11-11",
		"(A) Task 2",
		"(B) Task 1",
	}
	checkTaskListOrder(t, testTasklist, testExpectedList)
}

func TestTaskSortByContext(t *testing.T) {
	testTasklist.LoadFromPath(testInputSort)
	taskID := 26

	testTasklist = testTasklist[taskID : taskID+6]

	if err := testTasklist.Sort(SortCreatedDateAsc); err != nil {
		t.Fatal(err)
	}

	if err := testTasklist.Sort(SortContextAsc); err != nil {
		t.Fatal(err)
	}
	testExpectedList = []string{
		"2020-12-19 Task 3 @Apple",
		"2020-10-19 Task 2 @Apple @Banana",
		"2020-11-09 Task 1 @Apple @Banana",
		"2020-11-11 Task 6 @Apple @Coconut",
		"2020-11-19 Task 4 @Banana",
		"2020-12-09 Task 5",
	}
	checkTaskListOrder(t, testTasklist, testExpectedList)

	if err := testTasklist.Sort(SortContextDesc); err != nil {
		t.Fatal(err)
	}
	testExpectedList = []string{
		"2020-12-09 Task 5",
		"2020-11-19 Task 4 @Banana",
		"2020-11-11 Task 6 @Apple @Coconut",
		"2020-10-19 Task 2 @Apple @Banana",
		"2020-11-09 Task 1 @Apple @Banana",
		"2020-12-19 Task 3 @Apple",
	}
	checkTaskListOrder(t, testTasklist, testExpectedList)
}

func TestTaskSortByProject(t *testing.T) {
	testTasklist.LoadFromPath(testInputSort)
	taskID := 32

	testTasklist = testTasklist[taskID : taskID+6]

	if err := testTasklist.Sort(SortCreatedDateAsc); err != nil {
		t.Fatal(err)
	}

	if err := testTasklist.Sort(SortProjectAsc); err != nil {
		t.Fatal(err)
	}
	testExpectedList = []string{
		"2020-12-29 Task 3 +Apple",
		"2020-10-09 Task 1 +Apple +Banana",
		"2020-10-19 Task 2 +Apple +Banana",
		"2020-12-19 Task 4 +Banana",
		"2020-11-11 Task 6 +Coconut",
		"2020-12-09 Task 5",
	}
	checkTaskListOrder(t, testTasklist, testExpectedList)

	if err := testTasklist.Sort(SortProjectDesc); err != nil {
		t.Fatal(err)
	}
	testExpectedList = []string{
		"2020-12-09 Task 5",
		"2020-11-11 Task 6 +Coconut",
		"2020-12-19 Task 4 +Banana",
		"2020-10-09 Task 1 +Apple +Banana",
		"2020-10-19 Task 2 +Apple +Banana",
		"2020-12-29 Task 3 +Apple",
	}
	checkTaskListOrder(t, testTasklist, testExpectedList)
}

func TestTaskSortByTodoText(t *testing.T) {
	testTasklist.LoadFromPath(testInputSort)
	taskID := 38

	testTasklist = testTasklist[taskID : taskID+5]

	if err := testTasklist.Sort(SortTodoTextAsc); err != nil {
		t.Fatal(err)
	}
	testExpectedList = []string{
		"2020-10-09 Task 1 +Apple +Banana",
		"2020-10-19 Task 2 +Apple +Brown",
		"2020-12-29 Task 3 +Apple",
		"2020-12-19 Task 4 +Banana",
		"2020-11-11 Task 5 +Coconut",
	}
	checkTaskListOrder(t, testTasklist, testExpectedList)

	if err := testTasklist.Sort(SortTodoTextDesc); err != nil {
		t.Fatal(err)
	}
	testExpectedList = []string{
		"2020-11-11 Task 5 +Coconut",
		"2020-12-19 Task 4 +Banana",
		"2020-12-29 Task 3 +Apple",
		"2020-10-19 Task 2 +Apple +Brown",
		"2020-10-09 Task 1 +Apple +Banana",
	}
	checkTaskListOrder(t, testTasklist, testExpectedList)
}

func TestTaskSortByMultipleFlags(t *testing.T) {
	testTasklist.LoadFromPath(testInputSort)
	taskID := 43

	testTasklist = testTasklist[taskID : taskID+7]

	if err := testTasklist.Sort(SortTodoTextAsc, SortPriorityDesc); err != nil {
		t.Fatal(err)
	}
	testExpectedList = []string{
		"(A) 2020-10-09 Task 1 +Apple +Banana",
		"2020-10-19 Task 2 +Apple +Brown",
		"2020-12-29 Task 3",
		"(C) 2020-12-29 Task 3 +Apple",
		"2020-12-19 Task 4 +Banana",
		"(D) 2020-11-11 Task 5 +Coconut",
		"2020-12-29 Task 6",
	}
	checkTaskListOrder(t, testTasklist, testExpectedList)

	if err := testTasklist.Sort(SortPriorityAsc, SortTodoTextAsc); err != nil {
		t.Fatal(err)
	}
	testExpectedList = []string{
		"(A) 2020-10-09 Task 1 +Apple +Banana",
		"(C) 2020-12-29 Task 3 +Apple",
		"(D) 2020-11-11 Task 5 +Coconut",
		"2020-10-19 Task 2 +Apple +Brown",
		"2020-12-29 Task 3",
		"2020-12-19 Task 4 +Banana",
		"2020-12-29 Task 6",
	}
	checkTaskListOrder(t, testTasklist, testExpectedList)

	if err := testTasklist.Sort(SortPriorityAsc, SortCreatedDateAsc, SortTodoTextDesc); err != nil {
		t.Fatal(err)
	}
	testExpectedList = []string{
		"(A) 2020-10-09 Task 1 +Apple +Banana",
		"(C) 2020-12-29 Task 3 +Apple",
		"(D) 2020-11-11 Task 5 +Coconut",
		"2020-10-19 Task 2 +Apple +Brown",
		"2020-12-19 Task 4 +Banana",
		"2020-12-29 Task 6",
		"2020-12-29 Task 3",
	}
	checkTaskListOrder(t, testTasklist, testExpectedList)

	if err := testTasklist.Sort(SortPriorityAsc, SortProjectAsc, SortTaskIDAsc); err != nil {
		t.Fatal(err)
	}
	testExpectedList = []string{
		"(A) 2020-10-09 Task 1 +Apple +Banana",
		"(C) 2020-12-29 Task 3 +Apple",
		"(D) 2020-11-11 Task 5 +Coconut",
		"2020-10-19 Task 2 +Apple +Brown",
		"2020-12-19 Task 4 +Banana",
		"2020-12-29 Task 3",
		"2020-12-29 Task 6",
	}
	checkTaskListOrder(t, testTasklist, testExpectedList)
}

func TestTaskSortError(t *testing.T) {
	testTasklist.LoadFromPath(testInputSort)

	if err := testTasklist.Sort(123); err == nil {
		t.Errorf("Expected Sort() to fail because of unrecognized sort option, but it didn't!")
	} else if err.Error() != "unrecognized sort option" {
		t.Error(err)
	}
}

func Test_lessStrings(t *testing.T) {
	tests := []struct {
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
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Case#%d", i+1), func(t *testing.T) {
			if got := lessStrings(tt.a, tt.b); got != tt.want {
				t.Errorf("lessStrings() %v < %v got = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

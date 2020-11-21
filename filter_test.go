package todotxt

import (
	"testing"
	"time"
)

func BenchmarkTaskList_Filter(b *testing.B) {
	testTasklist.LoadFromPath(testInputFilter)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = testTasklist.Filter(FilterNot(FilterCompleted)).Filter(FilterByPriority("A"), FilterByPriority("B"))
	}
}

func TestTaskListFilter(t *testing.T) {
	if err := testTasklist.LoadFromPath(testInputTasklist); err != nil {
		t.Fatal(err)
	}

	// Filter list to get only completed tasks
	completedList := testTasklist.Filter(FilterCompleted)
	testExpected = 33
	testGot = len(completedList)
	if testGot != testExpected {
		t.Errorf("Expected TaskList to contain %d tasks, but got %d", testExpected, testGot)
	}

	// Filter list to get only tasks with a due date
	dueDateList := testTasklist.Filter(FilterHasDueDate)
	testExpected = 26
	testGot = len(dueDateList)
	if testGot != testExpected {
		t.Errorf("Expected TaskList to contain %d tasks, but got %d", testExpected, testGot)
	}

	// Filter list to get only tasks with "B" priority
	prioBList := testTasklist.Filter(FilterByPriority("b"))
	testExpected = 17
	testGot = len(prioBList)
	if testGot != testExpected {
		t.Errorf("Expected TaskList to contain %d tasks, but got %d", testExpected, testGot)
	}
}

func TestTaskListFilterByMultiplePredicates(t *testing.T) {
	if err := testTasklist.LoadFromPath(testInputFilter); err != nil {
		t.Fatal(err)
	}
	now := time.Now()
	testTasklist[0].DueDate = now.AddDate(0, 0, -2)
	testTasklist[1].DueDate = now
	testTasklist[2].DueDate = now.AddDate(0, 0, 1)

	// and
	filteredList := testTasklist.Filter(FilterHasPriority).Filter(FilterOverdue).Filter(FilterCompleted)
	testExpected = 2
	testGot = len(filteredList)
	if testGot != testExpected {
		t.Errorf("Expected TaskList to contain %d tasks, but got %d: [%v]", testExpected, testGot, filteredList.String())
	}

	// or -- filters tasks with priority A or B
	filteredList = testTasklist.Filter(FilterByPriority("a"), FilterByPriority("b"))
	testExpected = 7
	testGot = len(filteredList)
	if testGot != testExpected {
		t.Errorf("Expected TaskList to contain %d tasks, but got %d: [%v]", testExpected, testGot, filteredList.String())
	}

	// or -- filters incompleted tasks that are overdue or have set priority
	filteredList = testTasklist.Filter(FilterNot(FilterCompleted)).Filter(FilterOverdue, FilterHasPriority)
	testExpected = 10
	testGot = len(filteredList)
	if testGot != testExpected {
		t.Errorf("Expected TaskList to contain %d tasks, but got %d: [%v]", testExpected, testGot, filteredList.String())
	}
}

func TestTaskListFilterHelpers(t *testing.T) {
	if err := testTasklist.LoadFromPath(testInputFilter); err != nil {
		t.Fatal(err)
	}

	now := time.Now()
	testTasklist[0].DueDate = now.AddDate(0, 0, -2)
	testTasklist[1].DueDate = now
	testTasklist[2].DueDate = now.AddDate(0, 0, 1)

	testCases := []struct {
		predicate Predicate
		number    int
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
	}

	for i, tt := range testCases {
		filteredList := testTasklist.Filter(tt.predicate)
		testExpected = tt.number
		testGot = len(filteredList)
		if testGot != testExpected {
			t.Errorf("Case #%d, Expected TaskList to contain %d tasks, but got %d: [%v]", i+1, testExpected, testGot, filteredList.String())
		}
	}
}

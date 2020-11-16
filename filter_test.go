package todotxt

import (
	"testing"
	"time"
)

func TestTaskListFilter(t *testing.T) {
	if err := testTasklist.LoadFromPath(testInputTasklist); err != nil {
		t.Fatal(err)
	}

	// Filter list to get only completed tasks
	completedList := testTasklist.Filter(FilterCompleted)
	testExpected = 33
	testGot = len(*completedList)
	if testGot != testExpected {
		t.Errorf("Expected TaskList to contain %d tasks, but got %d", testExpected, testGot)
	}

	// Filter list to get only tasks with a due date
	dueDateList := testTasklist.Filter(FilterHasDueDate)
	testExpected = 26
	testGot = len(*dueDateList)
	if testGot != testExpected {
		t.Errorf("Expected TaskList to contain %d tasks, but got %d", testExpected, testGot)
	}

	// Filter list to get only tasks with "B" priority
	prioBList := testTasklist.Filter(FilterByPriority("b"))
	testExpected = 17
	testGot = len(*prioBList)
	if testGot != testExpected {
		t.Errorf("Expected TaskList to contain %d tasks, but got %d", testExpected, testGot)
	}
}

func TestTaskListFilterHelpers(t *testing.T) {
	if err := testTasklist.LoadFromPath(testInputFilter); err != nil {
		t.Fatal(err)
	}

	now := time.Now()
	testTasklist[0].DueDate = now.AddDate(0, 0, -1)
	testTasklist[1].DueDate = now
	testTasklist[2].DueDate = now.AddDate(0, 0, 1)

	testCases := []struct {
		predicate Predicate
		number    int
	}{
		{FilterCompleted, 9},
		{FilterReverse(FilterCompleted), 17},
		{FilterReverse(FilterReverse(FilterCompleted)), 9},
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
		testGot = len(*filteredList)
		if testGot != testExpected {
			t.Errorf("Case #%d, Expected TaskList to contain %d tasks, but got %d: [%v]", i+1, testExpected, testGot, filteredList.String())
		}
	}
}

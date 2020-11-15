package todotxt

import "testing"

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
	testCases := []struct {
		predicate func(Task) bool
		number    int
	}{
		{FilterCompleted, 9},
		{FilterNotCompleted, 17},
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

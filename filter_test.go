package todotxt

import "testing"

func TestTaskListFilter(t *testing.T) {
	if err := testTasklist.LoadFromPath(testInputTasklist); err != nil {
		t.Fatal(err)
	}

	// Filter list to get only completed tasks
	completedList := testTasklist.Filter(func(t Task) bool { return t.Completed })
	testExpected = 33
	testGot = len(*completedList)
	if testGot != testExpected {
		t.Errorf("Expected TaskList to contain %d tasks, but got %d", testExpected, testGot)
	}

	// Filter list to get only tasks with a due date
	dueDateList := testTasklist.Filter(func(t Task) bool { return t.HasDueDate() })
	testExpected = 26
	testGot = len(*dueDateList)
	if testGot != testExpected {
		t.Errorf("Expected TaskList to contain %d tasks, but got %d", testExpected, testGot)
	}

	// Filter list to get only tasks with "B" priority
	prioBList := testTasklist.Filter(func(t Task) bool {
		return t.HasPriority() && t.Priority == "B"
	})
	testExpected = 17
	testGot = len(*prioBList)
	if testGot != testExpected {
		t.Errorf("Expected TaskList to contain %d tasks, but got %d", testExpected, testGot)
	}
}

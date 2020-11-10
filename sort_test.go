package todotxt

import (
	"testing"
)

var (
	testInputSort = "testdata/sort_todo.txt"
)

func TestTaskSortByPriority(t *testing.T) {
	testTasklist.LoadFromPath(testInputSort)
	taskID := 0

	testTasklist = testTasklist[taskID : taskID+6]

	if err := testTasklist.Sort(SortPriorityAsc); err != nil {
		t.Fatal(err)
	}

	testExpected = "(A) 2012-01-30 Call Mom @Call @Phone +Family"
	testGot = testTasklist[0].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[1] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt"
	testGot = testTasklist[1].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[2] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12"
	testGot = testTasklist[2].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[3] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "(D) 2013-12-01 Outline chapter 5 @Computer +Novel Level:5 private:false due:2014-02-17"
	testGot = testTasklist[3].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[4] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "2013-02-22 Pick up milk @GroceryStore"
	testGot = testTasklist[4].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[5] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x 2014-01-03 Create golang library @Go +go-todotxt due:2014-01-05"
	testGot = testTasklist[5].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[6] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	if err := testTasklist.Sort(SortPriorityDesc); err != nil {
		t.Fatal(err)
	}

	testExpected = "x 2014-01-03 Create golang library @Go +go-todotxt due:2014-01-05"
	testGot = testTasklist[0].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[1] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "2013-02-22 Pick up milk @GroceryStore"
	testGot = testTasklist[1].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[2] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "(D) 2013-12-01 Outline chapter 5 @Computer +Novel Level:5 private:false due:2014-02-17"
	testGot = testTasklist[2].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[3] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12"
	testGot = testTasklist[3].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[4] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt"
	testGot = testTasklist[4].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[5] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "(A) 2012-01-30 Call Mom @Call @Phone +Family"
	testGot = testTasklist[5].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[6] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}
}

func TestTaskSortByCreatedDate(t *testing.T) {
	testTasklist.LoadFromPath(testInputSort)
	taskID := 6

	testTasklist = testTasklist[taskID : taskID+5]

	if err := testTasklist.Sort(SortCreatedDateAsc); err != nil {
		t.Fatal(err)
	}

	testExpected = "x 2014-01-03 Create golang library @Go +go-todotxt due:2014-01-05"
	testGot = testTasklist[0].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[1] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "(A) Call Mom @Call @Phone +Family"
	testGot = testTasklist[1].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[2] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "2013-02-22 Pick up milk @GroceryStore"
	testGot = testTasklist[2].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[3] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt"
	testGot = testTasklist[3].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[4] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12"
	testGot = testTasklist[4].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[5] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	if err := testTasklist.Sort(SortCreatedDateDesc); err != nil {
		t.Fatal(err)
	}

	testExpected = "x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12"
	testGot = testTasklist[0].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[1] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt"
	testGot = testTasklist[1].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[2] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "2013-02-22 Pick up milk @GroceryStore"
	testGot = testTasklist[2].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[3] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "(A) Call Mom @Call @Phone +Family"
	testGot = testTasklist[3].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[4] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x 2014-01-03 Create golang library @Go +go-todotxt due:2014-01-05"
	testGot = testTasklist[4].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[5] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}
}

func TestTaskSortByCompletedDate(t *testing.T) {
	testTasklist.LoadFromPath(testInputSort)
	taskID := 11

	testTasklist = testTasklist[taskID : taskID+6]

	if err := testTasklist.Sort(SortCompletedDateAsc); err != nil {
		t.Fatal(err)
	}

	testExpected = "x Download Todo.txt mobile app @Phone"
	testGot = testTasklist[0].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[1] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12"
	testGot = testTasklist[1].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[2] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "2013-02-22 Pick up milk @GroceryStore"
	testGot = testTasklist[2].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[3] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt"
	testGot = testTasklist[3].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[4] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x 2014-01-03 Create golang library @Go +go-todotxt due:2014-01-05"
	testGot = testTasklist[4].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[5] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x 2014-01-04 2014-01-01 Create some more golang library test cases @Go +go-todotxt"
	testGot = testTasklist[5].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[6] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	if err := testTasklist.Sort(SortCompletedDateDesc); err != nil {
		t.Fatal(err)
	}

	testExpected = "x 2014-01-04 2014-01-01 Create some more golang library test cases @Go +go-todotxt"
	testGot = testTasklist[0].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[1] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x 2014-01-03 Create golang library @Go +go-todotxt due:2014-01-05"
	testGot = testTasklist[1].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[2] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt"
	testGot = testTasklist[2].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[3] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "2013-02-22 Pick up milk @GroceryStore"
	testGot = testTasklist[3].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[4] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12"
	testGot = testTasklist[4].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[5] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x Download Todo.txt mobile app @Phone"
	testGot = testTasklist[5].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[6] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}
}

func TestTaskSortByDueDate(t *testing.T) {
	testTasklist.LoadFromPath(testInputSort)
	taskID := 17

	testTasklist = testTasklist[taskID : taskID+4]

	if err := testTasklist.Sort(SortDueDateAsc); err != nil {
		t.Fatal(err)
	}

	testExpected = "x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt"
	testGot = testTasklist[0].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[1] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x 2014-01-03 Create golang library @Go +go-todotxt due:2014-01-05"
	testGot = testTasklist[1].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[2] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12"
	testGot = testTasklist[2].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[3] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "(B) 2013-12-01 Outline chapter 5 @Computer +Novel Level:5 private:false due:2014-02-17"
	testGot = testTasklist[3].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[4] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	if err := testTasklist.Sort(SortDueDateDesc); err != nil {
		t.Fatal(err)
	}

	testExpected = "(B) 2013-12-01 Outline chapter 5 @Computer +Novel Level:5 private:false due:2014-02-17"
	testGot = testTasklist[0].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[1] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12"
	testGot = testTasklist[1].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[2] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x 2014-01-03 Create golang library @Go +go-todotxt due:2014-01-05"
	testGot = testTasklist[2].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[3] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt"
	testGot = testTasklist[3].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[4] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}
}

func TestTaskSortByTaskID(t *testing.T) {
	testTasklist.LoadFromPath(testInputSort)
	taskID := 21

	testTasklist = testTasklist[taskID : taskID+5]

	if err := testTasklist.Sort(SortTaskIDAsc); err != nil {
		t.Fatal(err)
	}

	testExpected = "(B) Task 1"
	testGot = testTasklist[0].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[1] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "(A) Task 2"
	testGot = testTasklist[1].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[2] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "Task 3 due:2020-11-11"
	testGot = testTasklist[2].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[3] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "Task 4 due:2020-12-12"
	testGot = testTasklist[3].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[4] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "x Task 5"
	testGot = testTasklist[4].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[5] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	if err := testTasklist.Sort(SortTaskIDDesc); err != nil {
		t.Fatal(err)
	}

	testExpected = "x Task 5"
	testGot = testTasklist[0].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[1] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "Task 4 due:2020-12-12"
	testGot = testTasklist[1].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[2] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "Task 3 due:2020-11-11"
	testGot = testTasklist[2].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[3] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "(A) Task 2"
	testGot = testTasklist[3].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[4] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "(B) Task 1"
	testGot = testTasklist[4].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[5] after Sort() to be [%s], but got [%s]", testExpected, testGot)
	}
}

func TestTaskSortError(t *testing.T) {
	testTasklist.LoadFromPath(testInputSort)

	if err := testTasklist.Sort(123); err == nil {
		t.Errorf("Expected Sort() to fail because of unrecognized sort option, but it didn't!")
	} else if err.Error() != "unrecognized sort option" {
		t.Error(err)
	}
}

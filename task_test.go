package todotxt

import (
	"testing"
	"time"
)

var (
	testInputTask = "testdata/task_todo.txt"
)

func BenchmarkParseTask(b *testing.B) {
	s := "x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12   "
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ParseTask(s)
	}
}

func BenchmarkTaskString(b *testing.B) {
	s := "x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12   "
	task, _ := ParseTask(s)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = task.String()
	}
}

func TestNewTask(t *testing.T) {
	task := NewTask()

	testExpected = 0
	testGot = task.ID
	if testGot != testExpected {
		t.Errorf("Expected new Task to have default ID [%d], but got [%d]", testExpected, testGot)
	}

	testExpected = ""
	testGot = task.Original
	if testGot != testExpected {
		t.Errorf("Expected new Task to be empty, but got [%s]", testGot)
	}

	testExpected = ""
	testGot = task.Todo
	if testGot != testExpected {
		t.Errorf("Expected new Task to be empty, but got [%s]", testGot)
	}

	testExpected = false
	testGot = task.HasPriority()
	if testGot != testExpected {
		t.Errorf("Expected new Task to have no priority, but got [%v]", testGot)
	}

	testExpected = 0
	testGot = len(task.Projects)
	if testGot != testExpected {
		t.Errorf("Expected new Task to have %d projects, but got [%d]", testExpected, testGot)
	}

	testExpected = 0
	testGot = len(task.Contexts)
	if testGot != testExpected {
		t.Errorf("Expected new Task to have %d contexts, but got [%d]", testExpected, testGot)
	}

	testExpected = 0
	testGot = len(task.AdditionalTags)
	if testGot != testExpected {
		t.Errorf("Expected new Task to have %d additional tags, but got [%d]", testExpected, testGot)
	}

	testExpected = true
	testGot = task.HasCreatedDate()
	if testGot != testExpected {
		t.Errorf("Expected new Task to have a created date, but got [%v]", testGot)
	}

	testExpected = false
	testGot = task.HasCompletedDate()
	if testGot != testExpected {
		t.Errorf("Expected new Task to not have a completed date, but got [%v]", testGot)
	}

	testExpected = false
	testGot = task.HasDueDate()
	if testGot != testExpected {
		t.Errorf("Expected new Task to not have a due date, but got [%v]", testGot)
	}

	testExpected = false
	testGot = task.Completed
	if testGot != testExpected {
		t.Errorf("Expected new Task to not be completed, but got [%v]", testGot)
	}
}

func TestParseTask(t *testing.T) {
	task, err := ParseTask("x (C) 2014-01-01 @Go due:2014-01-12 Create golang library documentation +go-todotxt   ")
	if err != nil {
		t.Error(err)
	}

	testExpected = "x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12"
	testGot = task.Task()
	if testGot != testExpected {
		t.Errorf("Expected Task to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = 0
	testGot = task.ID
	if testGot != testExpected {
		t.Errorf("Expected Task to have default ID [%d], but got [%d]", testExpected, testGot)
	}

	testExpected = "x (C) 2014-01-01 @Go due:2014-01-12 Create golang library documentation +go-todotxt"
	testGot = task.Original
	if testGot != testExpected {
		t.Errorf("Expected Task to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = "Create golang library documentation"
	testGot = task.Todo
	if testGot != testExpected {
		t.Errorf("Expected Task to be [%s], but got [%s]", testExpected, testGot)
	}

	testExpected = true
	testGot = task.HasPriority()
	if testGot != testExpected {
		t.Errorf("Expected Task to have no priority, but got [%v]", testGot)
	}

	testExpected = "C"
	testGot = task.Priority
	if testGot != testExpected {
		t.Errorf("Expected Task to have priority [%v], but got [%v]", testExpected, testGot)
	}

	testExpected = 1
	testGot = len(task.Projects)
	if testGot != testExpected {
		t.Errorf("Expected Task to have %d projects, but got [%d]", testExpected, testGot)
	}

	testExpected = 1
	testGot = len(task.Contexts)
	if testGot != testExpected {
		t.Errorf("Expected Task to have %d contexts, but got [%d]", testExpected, testGot)
	}

	testExpected = 0
	testGot = len(task.AdditionalTags)
	if testGot != testExpected {
		t.Errorf("Expected Task to have %d additional tags, but got [%d]", testExpected, testGot)
	}

	testExpected = true
	testGot = task.HasCreatedDate()
	if testGot != testExpected {
		t.Errorf("Expected Task to have a created date, but got [%v]", testGot)
	}

	testExpected = false
	testGot = task.HasCompletedDate()
	if testGot != testExpected {
		t.Errorf("Expected Task to not have a completed date, but got [%v]", testGot)
	}

	testExpected = true
	testGot = task.HasDueDate()
	if testGot != testExpected {
		t.Errorf("Expected Task to have a due date, but got [%v]", testGot)
	}

	testExpected = true
	testGot = task.Completed
	if testGot != testExpected {
		t.Errorf("Expected Task to be completed, but got [%v]", testGot)
	}
}

func TestTaskId(t *testing.T) {
	testTasklist.LoadFromPath(testInputTask)

	taskId := 1
	testGot = testTasklist[taskId-1].ID
	if testGot != taskId {
		t.Errorf("Expected Task[%d] to have ID [%d], but got [%d]", taskId, taskId, testGot)
	}

	taskId = 5
	testGot = testTasklist[taskId-1].ID
	if testGot != taskId {
		t.Errorf("Expected Task[%d] to have ID [%d], but got [%d]", taskId, taskId, testGot)
	}

	taskId = 27
	testGot = testTasklist[taskId-1].ID
	if testGot != taskId {
		t.Errorf("Expected Task[%d] to have ID [%d], but got [%d]", taskId, taskId, testGot)
	}
}

func TestTaskString(t *testing.T) {
	testTasklist.LoadFromPath(testInputTask)
	taskId := 1

	testExpected = "2013-02-22 Pick up milk @GroceryStore"
	testGot = testTasklist[taskId-1].String()
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%s], but got [%s]", taskId, testExpected, testGot)
	}
	taskId++

	testExpected = "x Download Todo.txt mobile app @Phone"
	testGot = testTasklist[taskId-1].String()
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%s], but got [%s]", taskId, testExpected, testGot)
	}
	taskId++

	testExpected = "(B) 2013-12-01 Outline chapter 5 @Computer +Novel Level:5 private:false due:2014-02-17"
	testGot = testTasklist[taskId-1].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%s], but got [%s]", taskId, testExpected, testGot)
	}
	taskId++

	testExpected = "x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt"
	testGot = testTasklist[taskId-1].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%s], but got [%s]", taskId, testExpected, testGot)
	}
	taskId++

	testExpected = "x 2014-01-03 2014-01-01 Create some more golang library test cases @Go +go-todotxt"
	testGot = testTasklist[taskId-1].Task()
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%s], but got [%s]", taskId, testExpected, testGot)
	}
}

func TestTaskPriority(t *testing.T) {
	testTasklist.LoadFromPath(testInputTask)
	taskId := 6

	testExpected = "B"
	testGot = testTasklist[taskId-1].Priority
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have priority '%s', but got '%s'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected = "C"
	testGot = testTasklist[taskId-1].Priority
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have priority '%s', but got '%s'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected = "B"
	testGot = testTasklist[taskId-1].Priority
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have priority '%s', but got '%s'", taskId, testExpected, testGot)
	}
	taskId++

	if testTasklist[taskId-1].HasPriority() {
		t.Errorf("Expected Task[%d] to have no priority, but got '%s'", taskId, testTasklist[4].Priority)
	}
}

func TestTaskCreatedDate(t *testing.T) {
	testTasklist.LoadFromPath(testInputTask)
	taskId := 10

	testExpected, err := time.Parse(DateLayout, "2012-01-30")
	if err != nil {
		t.Fatal(err)
	}
	testGot = testTasklist[taskId-1].CreatedDate
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have created date '%s', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected, err = time.Parse(DateLayout, "2013-02-22")
	if err != nil {
		t.Fatal(err)
	}
	testGot = testTasklist[taskId-1].CreatedDate
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have created date '%s', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected, err = time.Parse(DateLayout, "2014-01-01")
	if err != nil {
		t.Fatal(err)
	}
	testGot = testTasklist[taskId-1].CreatedDate
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have created date '%s', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected, err = time.Parse(DateLayout, "2013-12-30")
	if err != nil {
		t.Fatal(err)
	}
	testGot = testTasklist[taskId-1].CreatedDate
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have created date '%s', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected, err = time.Parse(DateLayout, "2014-01-01")
	if err != nil {
		t.Fatal(err)
	}
	testGot = testTasklist[taskId-1].CreatedDate
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have created date '%s', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	if testTasklist[taskId-1].HasCreatedDate() {
		t.Errorf("Expected Task[%d] to have no created date, but got '%v'", taskId, testTasklist[4].CreatedDate)
	}
}

func TestTaskContexts(t *testing.T) {
	testTasklist.LoadFromPath(testInputTask)
	taskId := 16

	testExpected = []string{"Call", "Phone"}
	testGot = testTasklist[taskId-1].Contexts
	if !compareSlices(testGot.([]string), testExpected.([]string)) {
		t.Errorf("Expected Task[%d] to have contexts '%v', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected = []string{"Office"}
	testGot = testTasklist[taskId-1].Contexts
	if !compareSlices(testGot.([]string), testExpected.([]string)) {
		t.Errorf("Expected Task[%d] to have contexts '%v', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected = []string{"Electricity", "Home", "Of_Super-Importance", "Television"}
	testGot = testTasklist[taskId-1].Contexts
	if !compareSlices(testGot.([]string), testExpected.([]string)) {
		t.Errorf("Expected Task[%d] to have contexts '%v', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected = []string{}
	testGot = testTasklist[taskId-1].Contexts
	if !compareSlices(testGot.([]string), testExpected.([]string)) {
		t.Errorf("Expected Task[%d] to have no contexts, but got '%v'", taskId, testGot)
	}
}

func TestTasksProjects(t *testing.T) {
	testTasklist.LoadFromPath(testInputTask)
	taskId := 20

	testExpected = []string{"Gardening", "Improving", "Planning", "Relaxing-Work"}
	testGot = testTasklist[taskId-1].Projects
	if !compareSlices(testGot.([]string), testExpected.([]string)) {
		t.Errorf("Expected Task[%d] to have projects '%v', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected = []string{"Novel"}
	testGot = testTasklist[taskId-1].Projects
	if !compareSlices(testGot.([]string), testExpected.([]string)) {
		t.Errorf("Expected Task[%d] to have projects '%v', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected = []string{}
	testGot = testTasklist[taskId-1].Projects
	if !compareSlices(testGot.([]string), testExpected.([]string)) {
		t.Errorf("Expected Task[%d] to have no projects, but got '%v'", taskId, testGot)
	}
}

func TestTaskDueDate(t *testing.T) {
	testTasklist.LoadFromPath(testInputTask)
	taskId := 23

	testExpected, err := time.Parse(DateLayout, "2014-02-17")
	if err != nil {
		t.Fatal(err)
	}
	testGot = testTasklist[taskId-1].DueDate
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have due date '%s', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	if testTasklist[taskId-1].HasDueDate() {
		t.Errorf("Expected Task[%d] to have no due date, but got '%v'", taskId, testTasklist[taskId-1].DueDate)
	}
}

func TestTaskAddonTags(t *testing.T) {
	testTasklist.LoadFromPath(testInputTask)
	taskId := 25

	testExpected = map[string]string{"Level": "5", "private": "false"}
	testGot = testTasklist[taskId-1].AdditionalTags
	if len(testGot.(map[string]string)) != 2 ||
		!compareMaps(testGot.(map[string]string), testExpected.(map[string]string)) {
		t.Errorf("Expected Task[%d] to have addon tags '%v', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected = map[string]string{"Importance": "Very!"}
	testGot = testTasklist[taskId-1].AdditionalTags
	if len(testGot.(map[string]string)) != 1 ||
		!compareMaps(testGot.(map[string]string), testExpected.(map[string]string)) {
		t.Errorf("Expected Task[%d] to have projects '%v', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected = map[string]string{}
	testGot = testTasklist[taskId-1].AdditionalTags
	if len(testGot.(map[string]string)) != 0 ||
		!compareMaps(testGot.(map[string]string), testExpected.(map[string]string)) {
		t.Errorf("Expected Task[%d] to have no additional tags, but got '%v'", taskId, testGot)
	}
	taskId++

	testExpected = map[string]string{}
	testGot = testTasklist[taskId-1].AdditionalTags
	if len(testGot.(map[string]string)) != 0 ||
		!compareMaps(testGot.(map[string]string), testExpected.(map[string]string)) {
		t.Errorf("Expected Task[%d] to have no additional tags, but got '%v'", taskId, testGot)
	}
}

func TestTaskIsCompleted(t *testing.T) {
	testTasklist.LoadFromPath(testInputTask)
	var (
		taskId   int
		testGot1 bool
		testGot2 bool
	)

	taskId = 31
	testGot1 = testTasklist[taskId-1].Completed
	testGot2 = testTasklist[taskId-1].IsCompleted()
	if testGot1 != testGot2 {
		t.Errorf("Expected Task[%d] to be completed '%v', but got '%v'", taskId, testGot1, testGot2)
	}

	taskId = 32
	testGot1 = testTasklist[taskId-1].Completed
	testGot2 = testTasklist[taskId-1].IsCompleted()
	if testGot1 != testGot2 {
		t.Errorf("Expected Task[%d] to be not completed '%v', but got '%v'", taskId, testGot1, testGot2)
	}
}

func TestTaskCompleted(t *testing.T) {
	testTasklist.LoadFromPath(testInputTask)
	taskId := 29

	testExpected = true
	testGot = testTasklist[taskId-1].Completed
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be completed, but got '%v'", taskId, testGot)
	}
	taskId++

	testExpected = true
	testGot = testTasklist[taskId-1].Completed
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be completed, but got '%v'", taskId, testGot)
	}
	taskId++

	testExpected = true
	testGot = testTasklist[taskId-1].Completed
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be completed, but got '%v'", taskId, testGot)
	}
	taskId++

	testExpected = false
	testGot = testTasklist[taskId-1].Completed
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] not to be completed, but got '%v'", taskId, testGot)
	}
	taskId++

	testExpected = false
	testGot = testTasklist[taskId-1].Completed
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] not to be completed, but got '%v'", taskId, testGot)
	}
}

func TestTaskCompletedDate(t *testing.T) {
	testTasklist.LoadFromPath(testInputTask)
	taskId := 34

	if testTasklist[taskId-1].HasCompletedDate() {
		t.Errorf("Expected Task[%d] to not have a completed date, but got '%v'", taskId, testTasklist[taskId-1].CompletedDate)
	}
	taskId++

	testExpected, err := time.Parse(DateLayout, "2014-01-03")
	if err != nil {
		t.Fatal(err)
	}
	testGot = testTasklist[taskId-1].CompletedDate
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have completed date '%s', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	if testTasklist[taskId-1].HasCompletedDate() {
		t.Errorf("Expected Task[%d] to not have a completed date, but got '%v'", taskId, testTasklist[taskId-1].CompletedDate)
	}
	taskId++

	testExpected, err = time.Parse(DateLayout, "2014-01-02")
	if err != nil {
		t.Fatal(err)
	}
	testGot = testTasklist[taskId-1].CompletedDate
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have completed date '%s', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	testExpected, err = time.Parse(DateLayout, "2014-01-03")
	if err != nil {
		t.Fatal(err)
	}
	testGot = testTasklist[taskId-1].CompletedDate
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have completed date '%s', but got '%v'", taskId, testExpected, testGot)
	}
	taskId++

	if testTasklist[taskId-1].HasCompletedDate() {
		t.Errorf("Expected Task[%d] to not have a completed date, but got '%v'", taskId, testTasklist[taskId-1].CompletedDate)
	}
}

func TestTaskIsOverdue(t *testing.T) {
	testTasklist.LoadFromPath(testInputTask)
	taskId := 40

	testGot = testTasklist[taskId-1].IsOverdue()
	if !testGot.(bool) {
		t.Errorf("Expected Task[%d] to be overdue, but got '%v'", taskId, testGot)
	}
	taskId++

	testGot = testTasklist[taskId-1].IsOverdue()
	if testGot.(bool) {
		t.Errorf("Expected Task[%d] not to be overdue, but got '%v'", taskId, testGot)
	}
	testTasklist[taskId-1].DueDate = time.Now().AddDate(0, 0, 1)
	testGot = testTasklist[taskId-1].Due()
	if testGot.(time.Duration).Hours() < 23 ||
		testGot.(time.Duration).Hours() > 25 {
		t.Errorf("Expected Task[%d] to be due in 24 hours, but got '%v'", taskId, testGot)
	}
	taskId++

	testGot = testTasklist[taskId-1].IsOverdue()
	if !testGot.(bool) {
		t.Errorf("Expected Task[%d] to be overdue, but got '%v'", taskId, testGot)
	}
	testTasklist[taskId-1].DueDate = time.Now().AddDate(0, 0, -3)
	testGot = testTasklist[taskId-1].Due()
	if testGot.(time.Duration).Hours() < 71 ||
		testGot.(time.Duration).Hours() > 73 {
		t.Errorf("Expected Task[%d] to be due since 72 hours, but got '%v'", taskId, testGot)
	}
	taskId++

	testGot = testTasklist[taskId-1].IsOverdue()
	if testGot.(bool) {
		t.Errorf("Expected Task[%d] not to be overdue, but got '%v'", taskId, testGot)
	}
}

func TestTaskComplete(t *testing.T) {
	testTasklist.LoadFromPath(testInputTask)
	taskId := 44

	// first 4 tasks should all match the same tests
	for i := 0; i < 4; i++ {
		testExpected = false
		testGot = testTasklist[taskId-1].Completed
		if testGot != testExpected {
			t.Errorf("Expected Task[%d] not to be completed, but got '%v'", taskId, testGot)
		}
		testGot = testTasklist[taskId-1].HasCompletedDate()
		if testGot != testExpected {
			t.Errorf("Expected Task[%d] not to have a completed date, but got '%v'", taskId, testGot)
		}
		testTasklist[taskId-1].Complete()
		testExpected = true
		testGot = testTasklist[taskId-1].Completed
		if testGot != testExpected {
			t.Errorf("Expected Task[%d] to be completed, but got '%v'", taskId, testGot)
		}
		testGot = testTasklist[taskId-1].HasCompletedDate()
		if testGot != testExpected {
			t.Errorf("Expected Task[%d] to have a completed date, but got '%v'", taskId, testGot)
		}
		testExpected = time.Now().Format(DateLayout)
		testGot = testTasklist[taskId-1].CompletedDate.Format(DateLayout)
		if testGot != testExpected {
			t.Errorf("Expected Task[%d] to have a completed date of '%v', but got '%v'", taskId, testExpected, testGot)
		}
		taskId++
	}

	testExpected = true
	testGot = testTasklist[taskId-1].Completed
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be completed, but got '%v'", taskId, testGot)
	}
	testGot = testTasklist[taskId-1].HasCompletedDate()
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have a completed date, but got '%v'", taskId, testGot)
	}
	testTasklist[taskId-1].Complete()
	testGot = testTasklist[taskId-1].Completed // should be unchanged
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be completed, but got '%v'", taskId, testGot)
	}
	testGot = testTasklist[taskId-1].HasCompletedDate() // should be unchanged
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have a completed date, but got '%v'", taskId, testGot)
	}
	testExpected = "2012-01-01" // should be unchanged
	testGot = testTasklist[taskId-1].CompletedDate.Format(DateLayout)
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have a completed date of '%v', but got '%v'", taskId, testExpected, testGot)
	}
}

func TestTaskReopen(t *testing.T) {
	testTasklist.LoadFromPath(testInputTask)
	taskId := 49

	// the first 2 tasks should match the same tests
	for i := 0; i < 2; i++ {
		testExpected = true
		testGot = testTasklist[taskId-1].Completed
		if testGot != testExpected {
			t.Errorf("Expected Task[%d] to be completed, but got '%v'", taskId, testGot)
		}
		testExpected = false
		testGot = testTasklist[taskId-1].HasCompletedDate()
		if testGot != testExpected {
			t.Errorf("Expected Task[%d] to have a completed date, but got '%v'", taskId, testGot)
		}
		testTasklist[taskId-1].Reopen()
		testExpected = false
		testGot = testTasklist[taskId-1].Completed
		if testGot != testExpected {
			t.Errorf("Expected Task[%d] to not be completed, but got '%v'", taskId, testGot)
		}
		testGot = testTasklist[taskId-1].HasCompletedDate()
		if testGot != testExpected {
			t.Errorf("Expected Task[%d] to not have a completed date, but got '%v'", taskId, testGot)
		}
		taskId++
	}

	// the next 3 tasks should all match the same tests
	for i := 0; i < 3; i++ {
		testExpected = true
		testGot = testTasklist[taskId-1].Completed
		if testGot != testExpected {
			t.Errorf("Expected Task[%d] to be completed, but got '%v'", taskId, testGot)
		}
		testGot = testTasklist[taskId-1].HasCompletedDate()
		if testGot != testExpected {
			t.Errorf("Expected Task[%d] to have a completed date, but got '%v'", taskId, testGot)
		}
		testTasklist[taskId-1].Reopen()
		testExpected = false
		testGot = testTasklist[taskId-1].Completed
		if testGot != testExpected {
			t.Errorf("Expected Task[%d] to not be completed, but got '%v'", taskId, testGot)
		}
		testGot = testTasklist[taskId-1].HasCompletedDate()
		if testGot != testExpected {
			t.Errorf("Expected Task[%d] to not have a completed date, but got '%v'", taskId, testGot)
		}
		taskId++
	}

	testExpected = false
	testGot = testTasklist[taskId-1].Completed
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be completed, but got '%v'", taskId, testGot)
	}
	testGot = testTasklist[taskId-1].HasCompletedDate()
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have a completed date, but got '%v'", taskId, testGot)
	}
	testTasklist[taskId-1].Reopen()
	testGot = testTasklist[taskId-1].Completed // should be unchanged
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be completed, but got '%v'", taskId, testGot)
	}
	testGot = testTasklist[taskId-1].HasCompletedDate() // should be unchanged
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to have a completed date, but got '%v'", taskId, testGot)
	}
}

func compareSlices(list1 []string, list2 []string) bool {
	if len(list1) != len(list2) {
		return false
	}

	for i := range list1 {
		if list1[i] != list2[i] {
			return false
		}
	}

	return true
}

func compareMaps(map1 map[string]string, map2 map[string]string) bool {
	if len(map1) != len(map2) {
		return false
	}

	compare := func(map1 map[string]string, map2 map[string]string) bool {
		for key, value := range map1 {
			if value2, found := map2[key]; !found {
				return false
			} else if value != value2 {
				return false
			}
		}
		return true
	}

	return compare(map1, map2) && compare(map2, map1)
}

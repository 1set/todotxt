package todotxt

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func BenchmarkLoadFromPath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = LoadFromPath(testInputTasklist)
	}
}

func BenchmarkTaskList_String(b *testing.B) {
	taskList, _ := LoadFromPath(testInputTasklist)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = taskList.String()
	}
}

func TestLoadFromFile(t *testing.T) {
	file, err := os.Open(testInputTasklist)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	if testTasklist, err := LoadFromFile(file); err != nil {
		t.Fatal(err)
	} else {
		data, err := ioutil.ReadFile(testExpectedOutput)
		if err != nil {
			t.Fatal(err)
		}
		testExpected = string(data)
		testGot = testTasklist.String()
		if testGot != testExpected {
			t.Errorf("Expected TaskList to be [%s], but got [%s]", testExpected, testGot)
		}
	}

	if testTasklist, err := LoadFromFile(nil); testTasklist != nil || err == nil {
		t.Errorf("Expected LoadFromFile to fail, but got TaskList back: [%s]", testTasklist)
	}
}

func TestLoadFromPath(t *testing.T) {
	if testTasklist, err := LoadFromPath(testInputTasklist); err != nil {
		t.Fatal(err)
	} else {
		data, err := ioutil.ReadFile(testExpectedOutput)
		if err != nil {
			t.Fatal(err)
		}
		testExpected = string(data)
		testGot = testTasklist.String()
		if testGot != testExpected {
			t.Errorf("Expected TaskList to be [%s], but got [%s]", testExpected, testGot)
		}
	}

	if testTasklist, err := LoadFromPath("some_file_that_does_not_exists.txt"); testTasklist != nil || err == nil {
		t.Errorf("Expected LoadFromPath to fail, but got TaskList back: [%s]", testTasklist)
	}
}

func TestWriteFile(t *testing.T) {
	_ = os.Remove(testOutput)
	_, _ = os.Create(testOutput)
	var err error

	fileInput, err := os.Open(testInputTasklist)
	if err != nil {
		t.Fatal(err)
	}
	defer fileInput.Close()
	fileOutput, err := os.OpenFile(testOutput, os.O_RDWR, 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer fileInput.Close()

	if testTasklist, err = LoadFromFile(fileInput); err != nil {
		t.Fatal(err)
	}
	if err = WriteToFile(&testTasklist, fileOutput); err != nil {
		t.Fatal(err)
	}
	fileInput.Close()
	fileOutput, err = os.Open(testOutput)
	if err != nil {
		t.Fatal(err)
	}
	if testTasklist, err = LoadFromFile(fileOutput); err != nil {
		t.Fatal(err)
	}

	data, err := ioutil.ReadFile(testExpectedOutput)
	if err != nil {
		t.Fatal(err)
	}
	testExpected = string(data)
	testGot = testTasklist.String()
	if testGot != testExpected {
		t.Errorf("Expected TaskList to be [%s], but got [%s]", testExpected, testGot)
	}

	if err = WriteToFile(&testTasklist, os.Stdin); err == nil {
		t.Errorf("Expected WriteToFile to fail for Stdin, but it didn't")
	}
}

func TestTaskListWriteFile(t *testing.T) {
	_ = os.Remove(testOutput)
	_, _ = os.Create(testOutput)
	testTasklist := TaskList{}

	fileInput, err := os.Open(testInputTasklist)
	if err != nil {
		t.Fatal(err)
	}
	defer fileInput.Close()
	fileOutput, err := os.OpenFile(testOutput, os.O_RDWR, 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer fileInput.Close()

	if err := testTasklist.LoadFromFile(fileInput); err != nil {
		t.Fatal(err)
	}
	if err := testTasklist.WriteToFile(fileOutput); err != nil {
		t.Fatal(err)
	}
	fileInput.Close()
	fileOutput, err = os.Open(testOutput)
	if err != nil {
		t.Fatal(err)
	}
	if err := testTasklist.LoadFromFile(fileOutput); err != nil {
		t.Fatal(err)
	}

	data, err := ioutil.ReadFile(testExpectedOutput)
	if err != nil {
		t.Fatal(err)
	}
	testExpected = string(data)
	testGot = testTasklist.String()
	if testGot != testExpected {
		t.Errorf("Expected TaskList to be [%s], but got [%s]", testExpected, testGot)
	}
}

func TestWriteFilename(t *testing.T) {
	os.Remove(testOutput)
	var err error

	if testTasklist, err = LoadFromPath(testInputTasklist); err != nil {
		t.Fatal(err)
	}
	if err = WriteToPath(&testTasklist, testOutput); err != nil {
		t.Fatal(err)
	}
	if testTasklist, err = LoadFromPath(testOutput); err != nil {
		t.Fatal(err)
	}

	data, err := ioutil.ReadFile(testExpectedOutput)
	if err != nil {
		t.Fatal(err)
	}
	testExpected = string(data)
	testGot = testTasklist.String()
	if testGot != testExpected {
		t.Errorf("Expected TaskList to be [%s], but got [%s]", testExpected, testGot)
	}
}

func TestTaskListWriteFilename(t *testing.T) {
	os.Remove(testOutput)
	testTasklist := TaskList{}

	if err := testTasklist.LoadFromPath(testInputTasklist); err != nil {
		t.Fatal(err)
	}
	if err := testTasklist.WriteToPath(testOutput); err != nil {
		t.Fatal(err)
	}
	if err := testTasklist.LoadFromPath(testOutput); err != nil {
		t.Fatal(err)
	}

	data, err := ioutil.ReadFile(testExpectedOutput)
	if err != nil {
		t.Fatal(err)
	}
	testExpected = string(data)
	testGot = testTasklist.String()
	if testGot != testExpected {
		t.Errorf("Expected TaskList to be [%s], but got [%s]", testExpected, testGot)
	}
}

func TestNewTaskList(t *testing.T) {
	testTasklist := NewTaskList()

	testExpected = 0
	testGot = len(testTasklist)
	if testGot != testExpected {
		t.Errorf("Expected TaskList to contain %d tasks, but got %d", testExpected, testGot)
	}
}

func TestTaskListCount(t *testing.T) {
	if err := testTasklist.LoadFromPath(testInputTasklist); err != nil {
		t.Fatal(err)
	}

	testExpected = 63
	testGot = len(testTasklist)
	if testGot != testExpected {
		t.Errorf("Expected TaskList to contain %d tasks, but got %d", testExpected, testGot)
	}
}

func TestTaskListAddTask(t *testing.T) {
	if err := testTasklist.LoadFromPath(testInputTasklist); err != nil {
		t.Fatal(err)
	}

	// add new empty task
	task := NewTask()
	testTasklist.AddTask(&task)

	testExpected = 64
	testGot = len(testTasklist)
	if testGot != testExpected {
		t.Errorf("Expected TaskList to contain %d tasks, but got %d", testExpected, testGot)
	}

	taskID := 64
	testExpected = time.Now().Format(DateLayout) + " " // tasks created by NewTask() have their created date set
	testGot = testTasklist[taskID-1].String()
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%s], but got [%s]", taskID, testExpected, testGot)
	}
	testExpected = 64
	testGot = testTasklist[taskID-1].ID
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%d], but got [%d]", taskID, testExpected, testGot)
	}
	taskID++

	// add parsed task
	parsed, err := ParseTask("x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12")
	if err != nil {
		t.Error(err)
	}
	testTasklist.AddTask(parsed)

	testExpected = "x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12"
	testGot = testTasklist[taskID-1].String()
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%s], but got [%s]", taskID, testExpected, testGot)
	}
	testExpected = 65
	testGot = testTasklist[taskID-1].ID
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%d], but got [%d]", taskID, testExpected, testGot)
	}
	taskID++

	// add selfmade task
	createdDate := time.Now()
	testTasklist.AddTask(&Task{
		CreatedDate: createdDate,
		Todo:        "Go shopping..",
		Contexts:    []string{"GroceryStore"},
	})

	testExpected = createdDate.Format(DateLayout) + " Go shopping.. @GroceryStore"
	testGot = testTasklist[taskID-1].String()
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%s], but got [%s]", taskID, testExpected, testGot)
	}
	testExpected = 66
	testGot = testTasklist[taskID-1].ID
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%d], but got [%d]", taskID, testExpected, testGot)
	}
	taskID++

	// add task with explicit ID, AddTask() should ignore this!
	testTasklist.AddTask(&Task{
		ID: 101,
	})

	testExpected = 67
	testGot = testTasklist[taskID-1].ID
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%d], but got [%d]", taskID, testExpected, testGot)
	}
}

func TestTaskListGetTask(t *testing.T) {
	if err := testTasklist.LoadFromPath(testInputTasklist); err != nil {
		t.Fatal(err)
	}

	taskID := 3
	task, err := testTasklist.GetTask(taskID)
	if err != nil {
		t.Error(err)
	}
	testExpected = "(B) 2013-12-01 Outline chapter 5 @Computer +Novel Level:5 private:false due:2014-02-17"
	testGot = task.String()
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%s], but got [%s]", taskID, testExpected, testGot)
	}
	testExpected = 3
	testGot = testTasklist[taskID-1].ID
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%d], but got [%d]", taskID, testExpected, testGot)
	}
}

func TestTaskListUpdateTask(t *testing.T) {
	if err := testTasklist.LoadFromPath(testInputTasklist); err != nil {
		t.Fatal(err)
	}

	taskID := 3
	task, err := testTasklist.GetTask(taskID)
	if err != nil {
		t.Error(err)
	}
	testExpected = "(B) 2013-12-01 Outline chapter 5 @Computer +Novel Level:5 private:false due:2014-02-17"
	testGot = task.String()
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%s], but got [%s]", taskID, testExpected, testGot)
	}
	testExpected = 3
	testGot = testTasklist[taskID-1].ID
	if testGot != testExpected {
		t.Errorf("Expected Task[%d] to be [%d], but got [%d]", taskID, testExpected, testGot)
	}

	task.Priority = "C"
	task.Todo = "Go home!"
	date, err := parseTime("2011-11-11")
	if err != nil {
		t.Error(err)
	}
	task.DueDate = date
	testGot := task

	os.Remove(testOutput)
	if err := testTasklist.WriteToPath(testOutput); err != nil {
		t.Fatal(err)
	}
	if err := testTasklist.LoadFromPath(testOutput); err != nil {
		t.Fatal(err)
	}
	testExpected, err := testTasklist.GetTask(taskID)
	if err != nil {
		t.Error(err)
	}
	if testGot.Task() != testExpected.Task() {
		t.Errorf("Expected Task to be [%v], but got [%v]", testExpected, testGot)
	}
}

func TestTaskListRemoveTaskByID(t *testing.T) {
	if err := testTasklist.LoadFromPath(testInputTasklist); err != nil {
		t.Fatal(err)
	}

	taskID := 10
	if err := testTasklist.RemoveTaskByID(taskID); err != nil {
		t.Error(err)
	}
	testExpected = 62
	testGot = len(testTasklist)
	if testGot != testExpected {
		t.Errorf("Expected TaskList to contain %d tasks, but got %d", testExpected, testGot)
	}
	task, err := testTasklist.GetTask(taskID)
	if err == nil || task != nil {
		t.Errorf("Expected no Task to be found anymore, but got %v", task)
	}

	taskID = 27
	if err := testTasklist.RemoveTaskByID(taskID); err != nil {
		t.Error(err)
	}
	testExpected = 61
	testGot = len(testTasklist)
	if testGot != testExpected {
		t.Errorf("Expected TaskList to contain %d tasks, but got %d", testExpected, testGot)
	}
	task, err = testTasklist.GetTask(taskID)
	if err == nil || task != nil {
		t.Errorf("Expected no Task to be found anymore, but got %v", task)
	}

	taskID = 99
	if err := testTasklist.RemoveTaskByID(taskID); err == nil {
		t.Errorf("Expected no Task to be found for removal")
	}
}

func TestTaskListRemoveTask(t *testing.T) {
	if err := testTasklist.LoadFromPath(testInputTasklist); err != nil {
		t.Fatal(err)
	}

	taskID := 52 // Is "unique" in tasklist
	task, err := testTasklist.GetTask(taskID)
	if err != nil {
		t.Error(err)
	}

	if err := testTasklist.RemoveTask(*task); err != nil {
		t.Error(err)
	}
	testExpected = 62
	testGot = len(testTasklist)
	if testGot != testExpected {
		t.Errorf("Expected TaskList to contain %d tasks, but got %d", testExpected, testGot)
	}
	task, err = testTasklist.GetTask(taskID)
	if err == nil || task != nil {
		t.Errorf("Expected no Task to be found anymore, but got %v", task)
	}

	taskID = 2 // Exists 3 times in tasklist
	task, err = testTasklist.GetTask(taskID)
	if err != nil {
		t.Error(err)
	}

	if err := testTasklist.RemoveTask(*task); err != nil {
		t.Error(err)
	}
	testExpected = 59
	testGot = len(testTasklist)
	if testGot != testExpected {
		t.Errorf("Expected TaskList to contain %d tasks, but got %d", testExpected, testGot)
	}
	task, err = testTasklist.GetTask(taskID)
	if err == nil || task != nil {
		t.Errorf("Expected no Task to be found anymore, but got %v", task)
	}

	if err := testTasklist.RemoveTask(NewTask()); err == nil {
		t.Errorf("Expected no Task to be found for removal")
	}
}

func TestTaskListReadErrors(t *testing.T) {
	if testTasklist, err := LoadFromPath(testInputTasklistCreatedDateError); testTasklist != nil || err == nil {
		t.Errorf("Expected LoadFromPath to fail because of invalid created date, but got TaskList back: [%s]", testTasklist)
	} else if err.Error() != `parsing time "2013-13-01": month out of range` {
		t.Error(err)
	}

	if testTasklist, err := LoadFromPath(testInputTasklistDueDateError); testTasklist != nil || err == nil {
		t.Errorf("Expected LoadFromPath to fail because of invalid due date, but got TaskList back: [%s]", testTasklist)
	} else if err.Error() != `parsing time "2014-02-32": day out of range` {
		t.Error(err)
	}

	if testTasklist, err := LoadFromPath(testInputTasklistCompletedDateError); testTasklist != nil || err == nil {
		t.Errorf("Expected LoadFromPath to fail because of invalid completed date, but got TaskList back: [%s]", testTasklist)
	} else if err.Error() != `parsing time "2014-25-04": month out of range` {
		t.Error(err)
	}

	// really silly test
	if testTasklist, err := LoadFromPath(testInputTasklistScannerError); testTasklist != nil || err == nil {
		t.Errorf("Expected LoadFromPath to fail because of invalid file, but got TaskList back: [%s]", testTasklist)
	} else if err.Error() != `bufio.Scanner: token too long` {
		t.Error(err)
	}
}

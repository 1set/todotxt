package todo

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
//  Tests: Constructors
// ----------------------------------------------------------------------------

func TestLoadFromFile(t *testing.T) {
	t.Parallel()

	// Open test file
	file, err := os.Open(testInputTasklist)
	require.NoError(t, err, "failed to open test file")

	defer file.Close()

	// Test
	actualTasklist, err := LoadFromFile(file)
	require.NoError(t, err, "the LoadFromFile failed to load tasklist from file")

	// Load expected output data
	rawOutput, err := os.ReadFile(testExpectedOutput)
	require.NoError(t, err, "failed to read expected output")

	expectStr := string(rawOutput)
	actualStr := actualTasklist.String()
	require.Equal(t, expectStr, actualStr, "loaded tasklist does not match expected")

	// Try loading from nil
	testTasklist, err := LoadFromFile(nil)

	require.Error(t, err, "expected error when loading from nil")
	require.Nil(t, testTasklist, "returned object should be nil on error")
	assert.Contains(t, err.Error(), "failed to load from file")
	assert.Contains(t, err.Error(), "nil io.Reader")
}

func TestLoadFromPath(t *testing.T) {
	t.Parallel()

	// Load test data
	actualTasklist, err := LoadFromPath(testInputTasklist)
	require.NoError(t, err, "failed to load tasklist from path")

	// Load expected output data
	rawOutput, err := os.ReadFile(testExpectedOutput)
	require.NoError(t, err, "failed to read expected output")

	// Compare
	expectStr := string(rawOutput)
	actualStr := actualTasklist.String()
	require.Equal(t, expectStr, actualStr, "loaded tasklist does not match expected")

	// Try loading from non-existent path
	testTasklist, err := LoadFromPath("some_file_that_does_not_exists.txt")

	require.Error(t, err, "expected error when loading from non-existent path")
	require.Nil(t, testTasklist, "returned object should be nil on error")
}

func TestNewTaskList(t *testing.T) {
	t.Parallel()

	testTasklist := NewTaskList()

	expectLen := 0
	actualLen := len(testTasklist)
	require.Equal(t, expectLen, actualLen, "expected empty tasklist")
}

// ----------------------------------------------------------------------------
//  Tests: Methods
// ----------------------------------------------------------------------------

func TestTaskList_WriteFile(t *testing.T) {
	t.Parallel()

	// Load test data
	expectTasklist := testLoadFromPath(t, testInputTasklist)
	expectStr := expectTasklist.String()

	// Get temporary file path for testing
	pathFileOutput := testGetPathFileTemp(t, testOutput)

	// Open temp file for writing
	const perm = 0o644

	fileOutput, err := os.OpenFile(pathFileOutput, os.O_RDWR|os.O_CREATE, perm)
	require.NoError(t, err, "failed to open temp file for writing")

	// Test writing to file pointer
	require.NoError(t, expectTasklist.WriteToFile(fileOutput),
		"method WriteToFile failed to write to file")

	fileOutput.Close()

	// Read raw tasklist to compare
	rawTaskList, err := os.ReadFile(pathFileOutput)
	require.NoError(t, err, "failed to read saved tasklist")

	actualStr := string(rawTaskList)
	require.Equal(t, expectStr, actualStr, "saved tasklist does not match expected")
}

func TestTaskList_WriteToPath(t *testing.T) {
	t.Parallel()

	// Load test data
	expectTasklist := testLoadFromPath(t, testInputTasklist)
	expectStr := expectTasklist.String()
	// Get temporary file path for testing
	pathFileOutput := testGetPathFileTemp(t, testOutput)

	// Test
	require.NoError(t, expectTasklist.WriteToPath(pathFileOutput),
		"method WriteToPath failed")

	// Reload saved tasklist to compare
	require.NoError(t, expectTasklist.LoadFromPath(pathFileOutput),
		"method LoadFromPath failed to reload saved tasklist")

	// Read saved tasklist to compare
	rawTaskList, err := os.ReadFile(pathFileOutput)
	require.NoError(t, err, "failed to read saved tasklist")

	// Compare
	actualStr := string(rawTaskList)
	require.Equal(t, expectStr, actualStr, "saved tasklist does not match expected")
}

func TestTaskList_Count(t *testing.T) {
	t.Parallel()

	// Load test data
	testTasklist := testLoadFromPath(t, testInputTasklist)

	expectLen := 63
	actualLen := testTasklist.Count()
	require.Equal(t, expectLen, actualLen)
}

func TestTaskList_AddTask(t *testing.T) {
	t.Parallel()

	// Load test data
	testTasklist := testLoadFromPath(t, testInputTasklist)

	// add new empty task
	task := NewTask()
	testTasklist.AddTask(&task)

	taskID := 64

	expectLen := 64
	actualLen := len(testTasklist)
	require.Equal(t, expectLen, actualLen, "number of tasks in tasklist was not as expected")

	expectStr := time.Now().Format(DateLayout) + " " // tasks created by NewTask() have their created date set
	actualStr := testTasklist[taskID-1].String()
	require.Equal(t, expectStr, actualStr, "task was not added to tasklist as expected")

	expectID := 64
	actualID := testTasklist[taskID-1].ID
	require.Equal(t, expectID, actualID, "the ID filed and the index of the task in the tasklist do not match")

	taskID++

	// add parsed task
	parsed, err := ParseTask("x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12")
	require.NoError(t, err, "failed to parse task")

	testTasklist.AddTask(parsed)

	expectStr = "x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12"
	actualStr = testTasklist[taskID-1].String()
	require.Equal(t, expectStr, actualStr, "new task was not added to tasklist as expected")

	expectID = 65
	actualID = testTasklist[taskID-1].ID
	require.Equal(t, expectID, actualID, "the ID filed and the index of the task in the tasklist do not match")

	taskID++

	// add selfmade task
	createdDate := time.Now()

	//nolint:exhaustruct // other fields are missing intentionally
	testTasklist.AddTask(&Task{
		CreatedDate: createdDate,
		Todo:        "Go shopping..",
		Contexts:    []string{"GroceryStore"},
	})

	expectStr = createdDate.Format(DateLayout) + " Go shopping.. @GroceryStore"
	actualStr = testTasklist[taskID-1].String()
	require.Equal(t, expectStr, actualStr, "new task was not added to tasklist as expected")

	expectID = 66
	actualID = testTasklist[taskID-1].ID
	require.Equal(t, expectID, actualID, "the ID filed and the index of the task in the tasklist do not match")
}

//nolint:paralleltest // do not parallel to avoid race conditions
func TestTaskList_AddTask_add_task_with_explicit_id(t *testing.T) {
	// Load test data
	testTasklist := testLoadFromPath(t, testInputTasklist)

	const taskID = 64

	// add task with explicit ID, AddTask() should ignore this!
	//nolint:exhaustruct // missing fields are intentional
	testTasklist.AddTask(&Task{
		ID: 101,
	})

	expectLen := 64
	actualLen := len(testTasklist)
	require.Equal(t, expectLen, actualLen, "adding a task with an explicit ID should not be added to the tasklist")

	expectID := 64
	actualID := testTasklist[taskID-1].ID
	require.Equal(t, expectID, actualID, "the ID filed and the index of the task in the tasklist do not match")
}

//nolint:paralleltest // do not parallel to avoid race conditions
func TestTaskList_GetTask(t *testing.T) {
	// Load test data
	testTasklist := testLoadFromPath(t, testInputTasklist)

	const taskID = 3

	task, err := testTasklist.GetTask(taskID)
	require.NoError(t, err, "failed to get task")

	require.Equal(t, testTasklist[taskID-1], *task, "the GetTask() method returned an unexpected task")

	actualStr := task.String()

	{
		expectID := 3
		actualID := task.ID
		require.Equal(t, expectID, actualID, "the ID filed and the index of the task in the tasklist do not match")
	}
	{
		expectStr := "(B) 2013-12-01 Outline chapter 5 @Computer +Novel Level:5 private:false due:2014-02-17"
		require.Equal(t, expectStr, actualStr, "the GetTask() method returned an unexpected task")
	}
}

//nolint:paralleltest // do not parallel to avoid race conditions
func TestTaskList_update_task(t *testing.T) {
	// Load test data
	testTasklist := testLoadFromPath(t, testInputTasklist)

	const taskID = 3

	task, err := testTasklist.GetTask(taskID)
	require.NoError(t, err, "failed to get task %d", taskID)

	{
		// Assert task contents be ok before update
		expectStr := "(B) 2013-12-01 Outline chapter 5 @Computer +Novel Level:5 private:false due:2014-02-17"
		actualStr := task.String()
		require.Equal(t, expectStr, actualStr, "task.String() did not return expected value. Task ID: %d", taskID)

		expectID := 3
		actualID := testTasklist[taskID-1].ID
		require.Equal(t, expectID, actualID, "task ID did not match expected value. Task ID: %d", taskID)
	}
	{
		// Update/change task properties
		date, err := parseTime("2011-11-11")
		require.NoError(t, err, "failed to parse time during test")

		task.DueDate = date    // 2014-02-17 -> 2011-11-11
		task.Priority = "C"    // from B to C
		task.Todo = "Go home!" // "Outline chapter 5" to "Go home!"

		expectTask := task.Task()

		// Save tasklist to a temporary file path for testing and reload it
		pathTemp := testGetPathFileTemp(t, testOutput)

		require.NoError(t, testTasklist.WriteToPath(pathTemp), "failed to write tasklist to temp file")
		require.NoError(t, testTasklist.LoadFromPath(pathTemp), "failed to load tasklist from temp file")

		loadedTask, err := testTasklist.GetTask(taskID)
		require.NoError(t, err, "failed to get task %d", taskID)

		// Compare task contents
		actualTask := loadedTask.Task()

		require.Equal(t, expectTask, actualTask, "task contents did not match expected value. Task ID: %d", taskID)
	}
}

func TestTaskList_RemoveTaskByID(t *testing.T) {
	t.Parallel()

	// Load test data
	testTasklist := testLoadFromPath(t, testInputTasklist)

	{
		// Remove un-registered task
		taskID := 99
		require.Error(t, testTasklist.RemoveTaskByID(taskID), "removing non-existing task should return error")
	}

	{
		taskID := 10

		// Remove task
		require.NoError(t, testTasklist.RemoveTaskByID(taskID), "failed to remove existing task")

		expectLen := 62
		actualLen := len(testTasklist)
		require.Equal(t, expectLen, actualLen, "unexpected length of tasklist after removing task")

		// Get the task again
		task, err := testTasklist.GetTask(taskID)
		require.Error(t, err, "getting removed task should return error")
		require.Nil(t, task, "getting removed task should return nil")
	}
	{
		taskID := 27

		// Remove task
		require.NoError(t, testTasklist.RemoveTaskByID(taskID), "failed to remove existing task")

		expectLen := 61
		actualLen := len(testTasklist)
		require.Equal(t, expectLen, actualLen, "unexpected length of tasklist after removing task")

		// Get the task again
		task, err := testTasklist.GetTask(taskID)
		require.Error(t, err, "getting removed task should return error")
		require.Nil(t, task, "getting removed task should return nil")
	}
}

func TestTaskList_RemoveTask(t *testing.T) {
	t.Parallel()

	// Load test data
	testTasklist := testLoadFromPath(t, testInputTasklist)

	{
		// Remove un-registered task
		require.Error(t, testTasklist.RemoveTask(NewTask()), "removing un-registered task should be an error")
	}

	{
		// Task ID 52 is "unique" in tasklist
		taskID := 52

		task, err := testTasklist.GetTask(taskID)
		require.NoError(t, err)

		// Remove the task
		require.NoError(t, testTasklist.RemoveTask(*task))

		expectLen := 62
		actualLen := len(testTasklist)
		require.Equal(t, expectLen, actualLen, "number of tasks in tasklist were not as expected")

		// Get the task again
		task, err = testTasklist.GetTask(taskID)
		require.Nil(t, task, "task was not removed from tasklist")
		require.Error(t, err, "getting the removed task from tasklist should be an error")
	}

	{
		// Task ID exists 3 times in tasklist
		taskID := 2

		task, err := testTasklist.GetTask(taskID)
		require.NoError(t, err)

		// Remove the task
		require.NoError(t, testTasklist.RemoveTask(*task))

		expectLen := 59
		actualLen := len(testTasklist)
		require.Equal(t, expectLen, actualLen, "number of tasks in tasklist were not as expected")

		// Get the task again
		task, err = testTasklist.GetTask(taskID)
		require.Error(t, err, "getting the removed task from tasklist should be an error")
		require.Nil(t, task, "task was not removed from tasklist")
	}
}

func TestTaskList_read_errors(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		path      string
		expectMsg string
		errMsg    string
	}{
		{
			testInputTasklistCreatedDateError,
			`parsing time "2013-13-01": month out of range`,
			"Expected LoadFromPath to fail because of invalid created date, but got TaskList back: [%s]",
		},
		{
			testInputTasklistDueDateError,
			`parsing time "2014-02-32": day out of range`,
			"Expected LoadFromPath to fail because of invalid due date, but got TaskList back: [%s]",
		},
		{
			testInputTasklistCompletedDateError,
			`parsing time "2014-25-04": month out of range`,
			"Expected LoadFromPath to fail because of invalid completed date, but got TaskList back: [%s]",
		},
		{
			testInputTasklistScannerError,
			`bufio.Scanner: token too long`,
			"Expected LoadFromPath to fail because of invalid file, but got TaskList back: [%s]",
		},
	} {
		testTasklist, err := LoadFromPath(test.path)

		require.Empty(t, testTasklist, test.errMsg, testTasklist)
		require.Error(t, err)
		require.Contains(t, err.Error(), test.expectMsg)
	}
}

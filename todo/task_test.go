package todo

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
//  Tests: Constructors
// ----------------------------------------------------------------------------

//nolint:paralleltest,funlen // do not parallel to avoid race conditions
func TestNewTask_default_state(t *testing.T) {
	task := NewTask()

	// ID
	expectID := 0
	actualID := task.ID
	require.Equal(t, expectID, actualID, "field ID failed to return expected ID")

	// Original
	expectStrRaw := ""
	actualStrRaw := task.Original
	require.Equal(t, expectStrRaw, actualStrRaw, "field Original failed to return expected string")

	//nolint:godox // the below Todo is not a comment as a TODO task
	// Todo
	expectTodo := ""
	actualTodo := task.Todo
	require.Equal(t, expectTodo, actualTodo, "field Todo failed to return expected string")

	// HasPriority()
	require.False(t, task.HasPriority(),
		"method HasPriority failed to return false. the given task does not have a priority")

	// HasProjects()
	require.False(t, task.HasProjects(),
		"method HasProjects failed to return false. the given task does not have projects")

	// Projects
	expectLenProj := 0
	actualLenProj := len(task.Projects)
	require.Equal(t, expectLenProj, actualLenProj, "field Projects failed to return expected length")

	// HasContexts()
	require.False(t, task.HasContexts(),
		"method HasContexts failed to return false. the given task does not have contexts")

	// Contexts length
	expectLenContexts := 0
	actualLenContexts := len(task.Contexts)
	require.Equal(t, expectLenContexts, actualLenContexts, "field Contexts failed to return expected length")

	// HasAdditionalTags()
	require.False(t, task.HasAdditionalTags(),
		"method HasAdditionalTags failed to return false. the given task does not have additional tags")

	// AdditionalTags length
	expectLenTags := 0
	actualLenTags := len(task.AdditionalTags)
	require.Equal(t, expectLenTags, actualLenTags, "field AdditionalTags failed to return expected length")

	// HasCreatedDate()
	require.True(t, task.HasCreatedDate(),
		"method HasCreatedDate failed to return true. newly created tasks are automatically assigned a creation date")

	// HasCompletedDate()
	require.False(t, task.HasCompletedDate(),
		"method HasCompletedDate failed to return false. the given task does not have a completed date")

	// HasDueDate()
	require.False(t, task.HasDueDate(),
		"method HasDueDate failed to return false. the given task does not have a due date")

	// Completed
	require.False(t, task.Completed,
		"field Completed failed to return false. the given task is not completed")
}

//nolint:funlen // function is long but leave it as is for now
func Test_ParseTask(t *testing.T) {
	t.Parallel()

	// ParseTask()
	task, err := ParseTask(
		"x (C) 2014-01-01 @Go due:2014-01-12 Create golang library documentation +go-todotxt  hello:world not::tag  ",
	)
	require.NoError(t, err, "method ParseTask failed to parse task")

	/* Method tests */

	// Task()
	expectStr := "x (C) 2014-01-01 Create golang library documentation not::tag @Go +go-todotxt hello:world due:2014-01-12"
	actualStr := task.Task()
	require.Equal(t, expectStr, actualStr, "method Task failed to return expected string")

	// ID
	expectID := 0
	actualID := task.ID
	require.Equal(t, expectID, actualID, "field ID failed to return expected ID")

	// Original
	expectRaw := "x (C) 2014-01-01 @Go due:2014-01-12 Create " +
		"golang library documentation +go-todotxt  hello:world not::tag"
	actualRaw := task.Original
	require.Equal(t, expectRaw, actualRaw, "field Original failed to return expected string")

	//nolint:godox // the below is not a comment as a TODO
	// Todo
	expectTask := "Create golang library documentation not::tag"
	actualTask := task.Todo
	require.Equal(t, expectTask, actualTask, "field Todo failed to return expected string")

	// HasPriority()
	require.True(t, task.HasPriority(),
		"method HasPriority failed to return true. the given task has a priority")

	// Priority
	expectPriority := "C"
	actualPriority := task.Priority
	require.Equal(t, expectPriority, actualPriority, "field Priority failed to return expected string")

	// HasProjects()
	require.True(t, task.HasProjects(), "method HasProjects failed to return true. the given task has projects")

	// Projects length
	expectLenProj := 1
	actualLenProj := len(task.Projects)
	require.Equal(t, expectLenProj, actualLenProj, "field Projects failed to return expected length")

	// HasContexts()
	require.True(t, task.HasContexts(), "method HasContexts failed to return true. the given task has contexts")

	// Contexts length
	expectLenContexts := 1
	actualLenContexts := len(task.Contexts)
	require.Equal(t, expectLenContexts, actualLenContexts, "field Contexts failed to return expected length")

	// HasAdditionalTags()
	require.True(t, task.HasAdditionalTags(),
		"method HasAdditionalTags failed to return true. the given task has additional tags")

	// AdditionalTags length
	expectLenTag := 1
	actualLenTag := len(task.AdditionalTags)
	require.Equal(t, expectLenTag, actualLenTag, "field AdditionalTags failed to return expected length")

	// Completed
	require.True(t, task.Completed, "field Completed failed to return true. the given task is completed")

	// HasCreatedDate()
	require.True(t, task.HasCreatedDate(),
		"method HasCreatedDate failed to return true. the given task has a created date")

	// HasDueDate()
	require.True(t, task.HasDueDate(), "method HasDueDate failed to return true. the given task has a due date")

	// HasCompletedDate()
	require.False(t, task.HasCompletedDate(),
		"method HasCompletedDate failed to return false. the given task does not have a completed date")
}

// ----------------------------------------------------------------------------
//  Tests for fields of Task type
// ----------------------------------------------------------------------------

func TestTask_ID(t *testing.T) {
	t.Parallel()

	// Load the test tasklist
	testTasklist := testLoadFromPath(t, testInputTask)

	for _, test := range []struct {
		taskID   int
		expectID int
	}{
		{taskID: 1, expectID: 1},
		{taskID: 5, expectID: 5},
		{taskID: 27, expectID: 27},
	} {
		taskID := test.taskID
		task := testTasklist[taskID-1]

		expect := test.expectID
		actual := task.ID

		require.Equal(t, expect, actual, "field ID of task[%d] failed to return expected ID", taskID)
	}
}

func TestTask_Priority(t *testing.T) {
	t.Parallel()

	// Load the test tasklist
	testTasklist := testLoadFromPath(t, testInputTask)

	// Golden cases
	for _, test := range []struct {
		expect string
		taskID int
	}{
		{taskID: 6, expect: "B"},
		{taskID: 7, expect: "C"},
		{taskID: 8, expect: "B"},
	} {
		taskID := test.taskID
		task := testTasklist[taskID-1]

		expectPriority := test.expect
		actualPriority := task.Priority
		require.Equal(t, expectPriority, actualPriority,
			"field Priority of task[%d] did not return the expected priority: %s", taskID, task.String())
	}

	// Test cases with no priority
	{
		taskID := 9
		task := testTasklist[taskID-1]

		require.Empty(t, task.Priority, "field Priority of task[%d] should be empty: %s", taskID, task.String())
		require.False(t, task.HasPriority(), "method HasPriority of task[%d] should return false: %s", taskID, task.String())
	}
}

func TestTask_CreatedDate(t *testing.T) {
	t.Parallel()

	// Load the test tasklist
	testTasklist := testLoadFromPath(t, testInputTask)

	// Golden cases
	for _, test := range []struct {
		expect string
		taskID int
	}{
		{taskID: 10, expect: "2012-01-30"},
		{taskID: 11, expect: "2013-02-22"},
		{taskID: 12, expect: "2014-01-01"},
		{taskID: 13, expect: "2013-12-30"},
		{taskID: 14, expect: "2014-01-01"},
	} {
		taskID := test.taskID
		task := testTasklist[taskID-1]

		expectTime, err := parseTime(test.expect)
		require.NoError(t, err, "failed to parse time for testing")

		actualTime := task.CreatedDate
		require.Equal(t, expectTime, actualTime,
			"field CreatedDate of task[%d] did not return as expected: %s", taskID, task.String())
	}

	// Missing created date
	{
		taskID := 15
		task := testTasklist[taskID-1]

		require.Empty(t, task.CreatedDate,
			"field CreatedDate of task[%d] should be empty: %s", taskID, task.String())
		require.False(t, task.HasCreatedDate(),
			"method HasCreatedDate of task[%d] should return false: %s", taskID, task.String())
	}
}

func TestTask_Contexts(t *testing.T) {
	t.Parallel()

	// Load the test tasklist
	testTasklist := testLoadFromPath(t, testInputTask)

	{
		taskID := 16
		task := testTasklist[taskID-1]

		expectContexts := []string{"Call", "Phone"}
		actualContexts := task.Contexts
		require.Equal(t, expectContexts, actualContexts,
			"task[%d] did not have the expected contexts: %s", taskID, task.String())
	}
	{
		taskID := 17
		task := testTasklist[taskID-1]

		expectContexts := []string{"Office"}
		actualContexts := task.Contexts
		require.Equal(t, expectContexts, actualContexts,
			"task[%d] did not have the expected contexts: %s", taskID, task.String())
	}
	{
		taskID := 18
		task := testTasklist[taskID-1]

		expectContexts := []string{"Electricity", "Home", "Of_Super-Importance", "Television"}
		actualContexts := testTasklist[taskID-1].Contexts
		require.Equal(t, expectContexts, actualContexts,
			"task[%d] did not have the expected contexts: %s", taskID, task.String())
	}
	{
		taskID := 19
		task := testTasklist[taskID-1]

		require.Empty(t, task.Contexts,
			"task[%d] should not have contexts: %s", taskID, task.String())
	}
}

func TestTask_Projects(t *testing.T) {
	t.Parallel()

	// Load the test tasklist
	testTasklist := testLoadFromPath(t, testInputTask)

	{
		taskID := 20
		task := testTasklist[taskID-1]

		expectProj := []string{"Gardening", "Improving", "Planning", "Relaxing-Work"}
		actualProj := task.Projects
		require.Equal(t, expectProj, actualProj,
			"Task[%d] did not have the expected projects: %s", taskID, task.String())
	}
	{
		taskID := 21
		task := testTasklist[taskID-1]

		expectProj := []string{"Novel"}
		actualProj := task.Projects
		require.Equal(t, expectProj, actualProj,
			"Task[%d] did not have the expected projects: %s", taskID, task.String())
	}
	{
		taskID := 22
		task := testTasklist[taskID-1]

		require.Empty(t, task.Projects,
			"Task[%d] should not have projects: %s", taskID, task.String())
	}
}

//nolint:funlen // leave it as is since it's a test
func TestTask_DueDate(t *testing.T) {
	t.Parallel()

	// Load the test tasklist
	testTasklist := testLoadFromPath(t, testInputTask)

	// Has due date
	{
		taskID := 23
		task := testTasklist[taskID-1]

		expectTime, err := parseTime("2014-02-17")
		require.NoError(t, err, "failed to parse expected time for testing")

		actualTime := task.DueDate

		require.Equal(t, expectTime, actualTime, "task[%d] should have a due date: %s", taskID, task.String())
	}

	// No due date
	{
		taskID := 24
		task := testTasklist[taskID-1]

		// HasDueDate()
		require.False(t, task.HasDueDate(),
			"task[%d] does not have a due date but HasDueDate() returned true: %s", taskID, task.String())
		// IsDueToday()
		require.False(t, task.IsDueToday(),
			"task[%d] should not havebe due today but IsDueToday() returned true: %s", taskID, task.String())
	}

	// Yesterdays task
	{
		task, err := ParseTask(fmt.Sprintf(
			"Hello Yesterday Task due:%s", time.Now().AddDate(0, 0, -1).Format(DateLayout)))
		require.NoError(t, err, "failed to parse task during testing")

		require.Less(t, task.Due(), time.Duration(0),
			"on overdue the duration time should be netagive: %s", task.String())
		require.True(t, task.IsOverdue(),
			"on overdue tasks IsOverdue should return true: %s", task.String())
		require.False(t, task.IsDueToday(),
			"on overdue tasks IsDueToday should return false: %s", task.String())
	}

	// Do it right now
	{
		task, err := ParseTask(fmt.Sprintf("Hello Today Task due:%s", time.Now().Format(DateLayout)))
		require.NoError(t, err, "failed to parse task during testing")

		require.Less(t, task.Due(), 24*time.Hour,
			"on due today tasks duration time should not be greater than one day: %s", task.String())
		require.False(t, task.IsOverdue(),
			"on due today tasks IsOverdue should return false: %s", task.String())
		require.True(t, task.IsDueToday(),
			"on due today tasks IsDueToday should return true: %s", task.String())
	}

	// Hasta mañana. Will do tomorrow
	{
		task, err := ParseTask(fmt.Sprintf("Hello Tomorrow Task due:%s", time.Now().AddDate(0, 0, 1).Format(DateLayout)))
		require.NoError(t, err, "failed to parse task during testing")

		require.Greater(t, task.Due(), 24*time.Hour,
			"on due tomorrow tasks duration time should be greater than one day: %s", task.String())
		require.LessOrEqual(t, task.Due(), 48*time.Hour,
			"on due tomorrow tasks duration time should not be greater than two days: %s", task.String())
		require.False(t, task.IsOverdue(),
			"on due tomorrow tasks IsOverdue should return false: %s", task.String())
		require.False(t, task.IsDueToday(),
			"on due tomorrow tasks IsDueToday should return false: %s", task.String())
	}
}

func TestTask_AddonTags(t *testing.T) {
	t.Parallel()

	// Load the test tasklist
	testTasklist := testLoadFromPath(t, testInputTask)

	// Golden cases
	{
		taskID := 25
		task := testTasklist[taskID-1]

		expectTags := map[string]string{"Level": "5", "private": "false"}
		actualTags := task.AdditionalTags

		require.Equal(t, expectTags, actualTags,
			"task[%d] did not contain the expected additional tags: %s", taskID, task)
	}
	{
		taskID := 26
		task := testTasklist[taskID-1]

		expectTags := map[string]string{"Importance": "Very!"}
		actualTags := testTasklist[taskID-1].AdditionalTags

		require.Equal(t, expectTags, actualTags,
			"task[%d] did not contain the expected additional tags: %s", taskID, task)
	}

	// Empty tag cases
	{
		taskID := 27
		task := testTasklist[taskID-1]

		require.Empty(t, task.AdditionalTags,
			"task[%d] should not contain additional tags: %s", taskID, task)
	}
	{
		taskID := 28
		task := testTasklist[taskID-1]

		require.Empty(t, task.AdditionalTags,
			"task[%d] should not contain additional tags: %s", taskID, task)
	}
}

// ----------------------------------------------------------------------------
//  Tests for Methods of Task type
// ----------------------------------------------------------------------------

func TestTask_IsCompleted(t *testing.T) {
	t.Parallel()

	// Load the test tasklist
	testTasklist := testLoadFromPath(t, testInputTask)

	// Completed case
	{
		taskID := 31
		task := testTasklist[taskID-1]

		testGot1 := task.Completed
		testGot2 := task.IsCompleted()

		require.True(t, testGot2, "task[%d] should be as completed: %s", taskID, task.String())
		require.Equal(t, testGot1, testGot2,
			"task[%d] should be completed and IsCompleted() should return true: %s", taskID, task.String())
	}

	// Not completed case
	{
		taskID := 32
		task := testTasklist[taskID-1]

		testGot1 := task.Completed
		testGot2 := task.IsCompleted()

		require.False(t, testGot2, "task[%d] should not be as completed: %s", taskID, task.String())
		require.Equal(t, testGot1, testGot2,
			"task[%d] should not be completed and IsCompleted() should return false: %s", taskID, task.String())
	}
}

func TestTask_Completed(t *testing.T) {
	t.Parallel()

	// Load the test tasklist
	testTasklist := testLoadFromPath(t, testInputTask)

	for _, test := range []struct {
		taskID int
		expect bool
	}{
		{29, true},
		{30, true},
		{31, true},
		{32, false},
		{33, false},
	} {
		task := testTasklist[test.taskID-1]

		if test.expect {
			require.True(t, task.Completed, "task[%d] should be completed: %s", test.taskID, task.String())
		} else {
			require.False(t, task.Completed, "task[%d] should be not completed: %s", test.taskID, task.String())
		}
	}
}

func TestTask_CompletedDate(t *testing.T) {
	t.Parallel()

	// Load the test tasklist
	testTasklist := testLoadFromPath(t, testInputTask)

	{
		taskID := 34

		require.False(t, testTasklist[taskID-1].HasCompletedDate(),
			"task[%d] should to not have a completed date: %s", taskID, testTasklist[taskID-1].String())
	}
	{
		taskID := 35

		expectCompletedDate, err := parseTime("2014-01-03")
		require.NoError(t, err, "failed to parse time for task[%d]", taskID)

		actualCompletedDate := testTasklist[taskID-1].CompletedDate
		require.Equal(t, expectCompletedDate, actualCompletedDate,
			"task[%d] should have a completed date of %s, but got %s", taskID, expectCompletedDate, actualCompletedDate)
	}
	{
		taskID := 36

		require.False(t, testTasklist[taskID-1].HasCompletedDate(),
			"task[%d] should to not have a completed date: %s", taskID, testTasklist[taskID-1].String())
	}
	{
		taskID := 37

		expectCompletedDate, err := parseTime("2014-01-02")
		require.NoError(t, err, "failed to parse time for task[%d]", taskID)

		actualCompletedDate := testTasklist[taskID-1].CompletedDate
		require.Equal(t, expectCompletedDate, actualCompletedDate,
			"task[%d] should have a completed date of %s, but got %s", taskID, expectCompletedDate, actualCompletedDate)
	}
	{
		taskID := 38

		expectCompletedDate, err := parseTime("2014-01-03")
		require.NoError(t, err, "failed to parse time for task[%d]", taskID)

		actualCompletedDate := testTasklist[taskID-1].CompletedDate
		require.Equal(t, expectCompletedDate, actualCompletedDate,
			"task[%d] should have a completed date of %s, but got %s", taskID, expectCompletedDate, actualCompletedDate)
	}
	{
		taskID := 39

		require.False(t, testTasklist[taskID-1].HasCompletedDate(),
			"task[%d] should to not have a completed date: %s", taskID, testTasklist[taskID-1].String())
	}
}

//nolint:paralleltest // do not parallel to avoid race conditions
func TestTask_IsOverdue(t *testing.T) {
	// Load the test tasklist
	testTasklist := testLoadFromPath(t, testInputTask)

	{
		taskID := 40

		require.True(t, testTasklist[taskID-1].IsOverdue(),
			"task[%d] should be as overdue: %s", taskID, testTasklist[taskID-1].String())
	}
	{
		taskID := 41

		require.False(t, testTasklist[taskID-1].IsOverdue(),
			"task[%d] should not be as overdue: %s", taskID, testTasklist[taskID-1].String())

		// Update the due date to be now
		testTasklist[taskID-1].DueDate = time.Now()

		dueHours := testTasklist[taskID-1].Due().Hours()
		require.True(t, dueHours > 23.0 && dueHours < 25.0,
			"task[%d] should be due in 24 hours: %s", taskID, testTasklist[taskID-1].String())
	}
	{
		taskID := 42

		require.True(t, testTasklist[taskID-1].IsOverdue(),
			"task[%d] should be as overdue: %s", taskID, testTasklist[taskID-1].String())

		// Update the due date to be 4 days ago
		testTasklist[taskID-1].DueDate = time.Now().AddDate(0, 0, -4)

		dueHours := testTasklist[taskID-1].Due().Hours()

		require.True(t, dueHours < 71 || dueHours > 73,
			"task[%d] should to be due since 72 hours: due hours: %v, task: %s",
			taskID, dueHours, testTasklist[taskID-1].String(),
		)

		testTasklist[taskID-1].DueDate = time.Now().AddDate(0, 0, 2)

		dueHours = testTasklist[taskID-1].Due().Hours()
		require.True(t, dueHours > 71 || dueHours < 73,
			"task[%d] should be due in 72 hours: due hours: %v, task: %s",
			taskID, dueHours, testTasklist[taskID-1].String(),
		)
	}
	{
		taskID := 43

		require.False(t, testTasklist[taskID-1].IsOverdue(),
			"task[%d] should not be as overdue: %s", taskID, testTasklist[taskID-1].String())
	}
}

func TestTask_Complete(t *testing.T) {
	t.Parallel()

	// Load the test tasklist
	testTasklist := testLoadFromPath(t, testInputTask)

	taskID := 44

	// first 4 tasks should all match the same tests (which are not completed tasks)
	for i := 0; i < 4; i++ {
		tmpTask := testTasklist[taskID-1]

		require.False(t, tmpTask.Completed,
			"task[%d] should not be as completed: %s", taskID, tmpTask.String())
		require.False(t, tmpTask.HasCompletedDate(),
			"task[%d] should not have a completed date: %s", taskID, tmpTask.String())

		tmpTask.Complete() // close the task right now

		require.True(t, tmpTask.Completed,
			"tasks invoked with the Complete() method should be marked as completed: %s", tmpTask.String())
		require.True(t, tmpTask.HasCompletedDate(),
			"tasks invoked with the Complete() method should have a completed date: %s", tmpTask.String())

		expectDate := time.Now().Format(DateLayout)
		actualTime := tmpTask.CompletedDate.Format(DateLayout)
		require.Equal(t, expectDate, actualTime,
			"tasks invoked with the Complete() method should have a completed date of Now(): %s", tmpTask.String())

		taskID++
	}

	// Closing a closed already task (should not change anything)
	{
		tmpTask := testTasklist[taskID-1]

		require.True(t, tmpTask.Completed,
			"task[%d] should be as completed: %s", taskID, tmpTask.String())
		require.True(t, tmpTask.HasCompletedDate(),
			"task[%d] should have a completed date: %s", taskID, tmpTask.String())

		tmpTask.Complete() // close the task right now

		require.True(t, tmpTask.Completed,
			"tasks invoked with the Complete() method should be marked as completed: %s", tmpTask.String())
		require.True(t, tmpTask.HasCompletedDate(),
			"tasks invoked with the Complete() method should have a completed date: %s", tmpTask.String())

		// completed date should not be changed
		expectDate := "2012-01-01"
		actualTime := tmpTask.CompletedDate.Format(DateLayout)
		require.Equal(t, expectDate, actualTime,
			"the Complete() method should not overwrite a completed date: %s", tmpTask.String())
	}
}

//nolint:funlen // leave it as is since it's a test
func TestTask_Reopen(t *testing.T) {
	t.Parallel()

	// Load the test tasklist
	testTasklist := testLoadFromPath(t, testInputTask)

	taskID := 49

	// the first 2 tasks should match the same tests (completed flag but with no completed date)
	for i := 0; i < 2; i++ {
		tmpTask := testTasklist[taskID-1]

		require.True(t, tmpTask.Completed,
			"task[%d] should be as completed: %s", taskID, tmpTask.String())
		require.False(t, tmpTask.HasCompletedDate(),
			"task[%d] should not have a completed date before reopening: %s", taskID, tmpTask.String())

		tmpTask.Reopen() // reopen the task

		require.False(t, tmpTask.Completed,
			"reopened task should not be as completed: %s", taskID, tmpTask.String())
		require.False(t, tmpTask.HasCompletedDate(),
			"reopened task should not have a completed date: %s", taskID, tmpTask.String())

		taskID++
	}

	// the next 3 tasks should all match the same tests (completed flag with completed date)
	for i := 0; i < 3; i++ {
		tmpTask := testTasklist[taskID-1]

		require.True(t, tmpTask.Completed,
			"task[%d] should be as completed: %s", taskID, tmpTask.String())
		require.True(t, tmpTask.HasCompletedDate(),
			"task[%d] should have a completed date before reopening: %s", taskID, tmpTask.String())

		tmpTask.Reopen() // reopen the task

		require.False(t, tmpTask.Completed,
			"reopened task[%d] should not be as completed: %s", taskID, tmpTask.String())
		require.False(t, tmpTask.HasCompletedDate(),
			"reopened task[%d] should not have a completed date: %s", taskID, tmpTask.String())

		taskID++
	}

	// Reopening an uncompleted task (should not change anything)
	{
		tmpTask := testTasklist[taskID-1]

		require.False(t, tmpTask.Completed,
			"the task[%d] should not be as completed: %s", taskID, tmpTask.String())
		require.False(t, tmpTask.HasCompletedDate(),
			"the task[%d] should not have a completed date: %s", taskID, tmpTask.String())

		tmpTask.Reopen() // reopen the task

		// these two should not be changed
		require.False(t, tmpTask.Completed,
			"reopened task[%d] should not be as completed: %s", taskID, tmpTask.String())
		require.False(t, tmpTask.HasCompletedDate(),
			"reopened task[%d] should not have a completed date: %s", taskID, tmpTask.String())
	}
}

func TestTask_String(t *testing.T) {
	t.Parallel()

	// Load the test tasklist
	testTasklist := testLoadFromPath(t, testInputTask)

	for _, test := range []struct {
		expectStr string
		taskID    int
	}{
		{taskID: 1, expectStr: "2013-02-22 Pick up milk @GroceryStore"},
		{taskID: 2, expectStr: "x Download Todo.txt mobile app @Phone"},
		{taskID: 3, expectStr: "(B) 2013-12-01 Outline chapter 5 @Computer +Novel Level:5 private:false due:2014-02-17"},
		{taskID: 4, expectStr: "x 2014-01-02 (B) 2013-12-30 Create golang library test cases @Go +go-todotxt"},
		{taskID: 5, expectStr: "x 2014-01-03 2014-01-01 Create some more golang library test cases @Go +go-todotxt"},
	} {
		taskID := test.taskID
		task := testTasklist[taskID-1]

		expect := test.expectStr
		actual := task.String()
		require.Equal(t, expect, actual, "method String of task[%d] failed to return expected string", taskID)

		// Method Task() should return the same string
		require.Equal(t, expect, task.Task(), "method Task of task[%d] failed to return expected string", taskID)
	}
}

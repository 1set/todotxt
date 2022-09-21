package todo

import (
	"os"
	"path/filepath"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ============================================================================
//  Helper functions/variables for Testing
// ============================================================================

const (
	testInputTask                       = "testdata/task_todo.txt"
	testInputSort                       = "testdata/sort_todo.txt"
	testInputFilter                     = "testdata/filter_todo.txt"
	testInputTasklist                   = "testdata/tasklist_todo.txt"
	testInputTasklistCreatedDateError   = "testdata/tasklist_createdDate_error.txt"
	testInputTasklistDueDateError       = "testdata/tasklist_dueDate_error.txt"
	testInputTasklistCompletedDateError = "testdata/tasklist_completedDate_error.txt"
	testInputTasklistScannerError       = "testdata/tasklist_scanner_error.txt"
	testOutput                          = "testdata/output_todo.txt"
	testExpectedOutput                  = "testdata/expected_todo.txt"
)

var _ = func() interface{} {
	RemoveCompletedPriority = false

	return nil
}()

// It holds the tasklist for each test file. The lists will be loaded via
// `testLoadFromPath`.
//
//nolint:gochecknoglobals // it is intentionally global
var taskLists = map[string]TaskList{}

// clientMutex is a mutex used to avoid concurrent map writes.
//
//nolint:gochecknoglobals // it is intentionally global
var clientMutex struct {
	sync.Mutex
}

// ----------------------------------------------------------------------------
//  Helper functions
// ----------------------------------------------------------------------------

// It errors if the tasks of given tasklist is not equal to the expected tasklist.
func checkTaskListOrder(t *testing.T, gotList TaskList, expStrList []string) {
	t.Helper()

	assert.Equal(t, len(expStrList), len(gotList), "tasklist length mismatch")

	for i, expectTask := range expStrList {
		actualTask := gotList[i].Task()

		assert.Equal(t, expectTask, actualTask, "task mismatch")
	}
}

// It returns the loaded tasklist for the given path. It is equivalent to
// `LoadFromPath` but it caches the loaded tasklist.
//
// Using this function will avoid accessing the same file multiple times in
// parallel tests.
func testLoadFromPath(t *testing.T, path string) TaskList {
	t.Helper()

	// Return cached TaskList if it exists
	if taskList, ok := taskLists[path]; ok {
		return taskList
	}

	clientMutex.Lock() // Lock

	// Load TaskList from file and cache it
	taskList, err := LoadFromPath(path)
	require.NoError(t, err, "failed to load tasklist")

	taskLists[path] = taskList

	clientMutex.Unlock() // Unlock

	return taskList
}

// It returns the absolute path of the given file path as a temporary file.
// Each subsequent call to testGetPathFileTemp returns a unique directory.
//
// The pathSub is used to create the temporary file in a temporary directory.
// The directory is automatically removed by Cleanup when the test and all its
// subtests complete. If the directory creation fails, TempDir terminates the
// test by calling Fatal.
//
//nolint:unparam // currently `pathSub` always receives `testOutput` but leave it as a parameter for future use
func testGetPathFileTemp(t *testing.T, pathSub string) string {
	t.Helper()

	pathDirTemp := t.TempDir()
	pathFileFull := filepath.Join(pathDirTemp, t.Name(), pathSub)
	pathDirFull := filepath.Dir(pathFileFull)

	require.NoError(t, os.MkdirAll(pathDirFull, PermReadWriteExec))

	return pathFileFull
}

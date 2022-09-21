package todo

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
//  RemoveCompletedPriority
// ----------------------------------------------------------------------------

//nolint:paralleltest // do not parallel to avoid race conditions
func TestRemoveCompletedPriority(t *testing.T) {
	oldRemoveCompletedPriority := RemoveCompletedPriority

	// Reset the flag to the default state
	defer func() {
		RemoveCompletedPriority = oldRemoveCompletedPriority
	}()

	task, err := ParseTask("(A) Hello World @Work")
	require.NoError(t, err, "failed to parse task during test setup")

	// Pre-condition check
	{
		expectStr := "(A) Hello World @Work"
		actualStr := task.String()
		require.Equal(t, expectStr, actualStr, "expected task to be unchanged")

		expectLenSeg := 3
		actualLenSeg := len(task.Segments())
		require.Equal(t, expectLenSeg, actualLenSeg, "expected task to have 3 segments")
	}

	// Sets Task.Completed to 'true' if the task was not already completed.
	// Also sets Task.CompletedDate to time.Now().
	task.Complete()

	// RemoveCompletedPriority as false
	{
		RemoveCompletedPriority = false

		// Set completed date to an old date
		task.CompletedDate = time.Date(2020, 11, 30, 0, 0, 0, 0, time.Local)

		expectStr := "x 2020-11-30 (A) Hello World @Work"
		actualStr := task.String()
		require.Equal(t, expectStr, actualStr, "expected task to have completed flag and date with priority")

		expectLenSeg := 5
		actualLenSeg := len(task.Segments())
		require.Equal(t, expectLenSeg, actualLenSeg,
			"expected task to have 5 segments with priority in case RemoveCompletedPriority is false")
	}

	// RemoveCompletedPriority as true
	{
		RemoveCompletedPriority = true

		expectStr := "x 2020-11-30 Hello World @Work"
		actualStr := task.String()
		require.Equal(t, expectStr, actualStr, "expected task to have completed flag and date but no priority")

		expectLenSeg := 4
		actualLenSeg := len(task.Segments())
		require.Equal(t, expectLenSeg, actualLenSeg,
			"expected task to remove segment for priority in case RemoveCompletedPriority is true")
	}
}

// ----------------------------------------------------------------------------
//  WriteToFile()
// ----------------------------------------------------------------------------

// no parallel due to wrigint to file.
func TestWriteFile_goledn(t *testing.T) {
	t.Parallel()

	// Load the TaskList from file
	expectTasklist := testLoadFromPath(t, testInputTasklist)

	// Get temporary file path for testing
	pathFileOutpt := testGetPathFileTemp(t, testOutput)

	// Open the temporary file for WriteToFile
	const perm = 0o644

	fileOutput, err := os.OpenFile(pathFileOutpt, os.O_RDWR|os.O_CREATE, perm)
	require.NoError(t, err, "failed to open temporary file")

	defer fileOutput.Close()

	// Save the TaskList to a temporary file
	require.NoError(t, WriteToFile(&expectTasklist, fileOutput),
		"the WriteToFile failed to write to temporary file")

	// Reload the TaskList from the temporary file
	actualTasklist, err := LoadFromPath(pathFileOutpt)
	require.NoError(t, err, "failed to load saved tasklist")

	expect := expectTasklist.String()
	actual := actualTasklist.String()
	require.Equal(t, expect, actual, "the saved tasklist is not equal to the original tasklist")
}

func TestWriteFile_fail_to_write(t *testing.T) {
	t.Parallel()

	// Load the TaskList from file
	tempTasklist := testLoadFromPath(t, testInputTasklist)

	require.Error(t, WriteToFile(&tempTasklist, os.Stdin),
		"the WriteToFile should fail to write to stdin")
}

// ----------------------------------------------------------------------------
//  WriteToPath()
// ----------------------------------------------------------------------------

func TestWriteToPath(t *testing.T) {
	t.Parallel()

	// Load the TaskList from file
	expectTasklist := testLoadFromPath(t, testInputTasklist)

	// Get temporary file path for testing
	pathFileOutpt := testGetPathFileTemp(t, testOutput)

	// Test
	require.NoError(t, WriteToPath(&expectTasklist, pathFileOutpt))

	actualTasklist, err := LoadFromPath(pathFileOutpt)
	require.NoError(t, err, "failed to load saved tasklist")

	expect := expectTasklist.String()
	actual := actualTasklist.String()
	require.Equal(t, expect, actual, "the saved tasklist is not equal to the original tasklist")
}

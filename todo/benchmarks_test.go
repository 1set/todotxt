package todo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// ============================================================================
//  Benchmarks for Functions
// ============================================================================

func BenchmarkLoadFromPath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = LoadFromPath(testInputTasklist)
	}
}

func BenchmarkParseTask(b *testing.B) {
	const s = "x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12   "

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = ParseTask(s)
	}
}

// ============================================================================
//  Benchmarks for Methods
// ============================================================================

// ----------------------------------------------------------------------------
//  Task
// ----------------------------------------------------------------------------

func BenchmarkTask_String(b *testing.B) {
	s := "x (C) 2014-01-01 Create golang library documentation @Go +go-todotxt due:2014-01-12   "

	task, err := ParseTask(s)
	require.NoError(b, err, "failed to parse task during benchmark")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = task.String()
	}
}

// ----------------------------------------------------------------------------
//  TaskList
// ----------------------------------------------------------------------------

func BenchmarkTaskList_Filter(b *testing.B) {
	testTasklist, err := LoadFromPath(testInputFilter)
	require.NoError(b, err)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = testTasklist.Filter(FilterNot(FilterCompleted)).Filter(FilterByPriority("A"), FilterByPriority("B"))
	}
}

func BenchmarkTaskList_Sort(b *testing.B) {
	// Load the test tasklist
	testTasklist, err := LoadFromPath(testInputSort)
	require.NoError(b, err, "loading test tasklist failed during benchmark")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = testTasklist.Sort(SortPriorityAsc, SortCreatedDateAsc, SortTodoTextDesc)
	}
}

func BenchmarkTaskList_String(b *testing.B) {
	taskList, err := LoadFromPath(testInputTasklist)
	require.NoError(b, err, "failed to load tasklist from path")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = taskList.String()
	}
}

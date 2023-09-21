package todo_test

import (
	"fmt"
	"log"

	"github.com/KEINOS/go-todotxt/todo"
)

// ============================================================================
//  TaskList
// ============================================================================

// ============================================================================
//  TaskList
// ============================================================================

// ----------------------------------------------------------------------------
//  TaskList.CustomSort
// ----------------------------------------------------------------------------

func ExampleTaskList_CustomSort() {
	tasks, err := todo.LoadFromString(`
		Task 3
		Task 1
		Task 4
		Task 2
	`)
	if err != nil {
		log.Fatal(err)
	}

	// customFunc returns true if taskA is less than taskB.
	customFunc := func(a, b todo.Task) bool {
		// Compare strings of the text part of the task.
		return a.Todo < b.Todo
	}

	tasks.CustomSort(customFunc)

	fmt.Println(tasks[0].Todo)
	fmt.Println(tasks[1].Todo)
	fmt.Println(tasks[2].Todo)
	fmt.Println(tasks[3].Todo)
	// Output:
	// Task 1
	// Task 2
	// Task 3
	// Task 4
}

// ----------------------------------------------------------------------------
//  TaskList.Filter
// ----------------------------------------------------------------------------

func ExampleTaskList_Filter_from_file() {
	var tasklist todo.TaskList
	if err := tasklist.LoadFromPath("testdata/filter_todo.txt"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Before:", tasklist[0].String())

	// Get tasks that are not overdue and are priority A or B.
	tasklist = tasklist.Filter(todo.FilterNot(todo.FilterOverdue)).Filter(
		todo.FilterByPriority("A"), todo.FilterByPriority("B"),
	)

	fmt.Println("After :", tasklist[0].String())
	// Output:
	// Before: This is a task should be due yesterday due:2020-11-15
	// After : (A) Call Mom @Call @Phone +Family
}

func ExampleTaskList_Filter_from_string() {
	tasks, err := todo.LoadFromString(`
		(A) Call Mom @Phone +Family
		x (A) Schedule annual checkup +Health
		(B) Outline chapter 5 +Novel @Computer
		(C) Add cover sheets @Office +TPSReports
		Plan backyard herb garden @Home
		Pick up milk @GroceryStore
		Research self-publishing services +Novel @Computer
		x Download Todo.txt mobile app @Phone
	`)
	if err != nil {
		log.Fatal(err)
	}

	// AND filter.  Get tasks that are not completed AND has priority.
	nearTopTasks := tasks.Filter(todo.FilterNotCompleted).Filter(todo.FilterHasPriority)

	// OR filter. Filter the above tasks by priority "A" OR "C".
	nearTopTasks = nearTopTasks.Filter(
		todo.FilterByPriority("A"),
		todo.FilterByPriority("C"),
	)

	fmt.Println(nearTopTasks.String())

	// Output:
	// (A) Call Mom @Phone +Family
	// (C) Add cover sheets @Office +TPSReports
}

// ----------------------------------------------------------------------------
//  TaskList.Sort
// ----------------------------------------------------------------------------

func ExampleTaskList_Sort() {
	var tasklist todo.TaskList
	if err := tasklist.LoadFromPath("testdata/sort_todo.txt"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Before #1:", tasklist[0].Projects)
	fmt.Println("Before #2:", tasklist[1].Projects)
	fmt.Println("Before #3:", tasklist[2].Projects)

	// sort tasks by project and then priority in ascending order
	if err := tasklist.Sort(todo.SortProjectAsc, todo.SortPriorityAsc); err != nil {
		log.Fatal(err)
	}

	fmt.Println("After  #1:", tasklist[0].Projects)
	fmt.Println("After  #2:", tasklist[1].Projects)
	fmt.Println("After  #3:", tasklist[2].Projects)
	// Output:
	// Before #1: []
	// Before #2: [go-todotxt]
	// Before #3: [go-todotxt]
	// After  #1: [Apple]
	// After  #2: [Apple]
	// After  #3: [Apple]
}

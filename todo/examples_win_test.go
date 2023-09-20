//go:build windows
// +build windows

package todo_test

import (
	"fmt"
	"log"
	"strings"

	"github.com/KEINOS/go-todotxt/todo"
)

func ExampleLoadFromPath_win() {
	taskListRaw, err := todo.LoadFromPath("testdata/todo.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Usually, `fmt.Println(taskListRaw)` or `fmt.Println(taskListRaw.String())`
	// is enough to print the task list.
	//
	// However, `todo.TaskList.String()` returns the task list with the end of
	// line character that differs from OS to OS. "\n" for Unix-like and "\r\n"
	// for Windows.
	//
	// Since the test is created on Unix-like OS (`\n`), simply printing the
	// task list will fail on Windows (See issue #14).
	//
	// So in this example we split the task list by the end of line character
	// (`todo.NewLine`) to print each task to keep the compatibility of the test
	// between OSes.
	taskList := strings.Split(taskListRaw.String(), todo.NewLine)

	for _, task := range taskList {
		fmt.Println(task)
	}
	// Output:
	// (A) Call Mom @Phone +Family
	// (A) Schedule annual checkup +Health
	// (B) Outline chapter 5 @Computer +Novel
	// (C) Add cover sheets @Office +TPSReports
	// Plan backyard herb garden @Home
	// Pick up milk @GroceryStore
	// Research self-publishing services @Computer +Novel
	// x Download Todo.txt mobile app @Phone
}

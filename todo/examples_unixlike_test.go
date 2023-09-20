//go:build !windows
// +build !windows

package todo_test

import (
	"fmt"
	"log"

	"github.com/KEINOS/go-todotxt/todo"
)

func ExampleLoadFromPath_win() {
	taskListRaw, err := todo.LoadFromPath("testdata/todo.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(taskListRaw.String())
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

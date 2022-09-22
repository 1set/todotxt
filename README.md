<!-- markdownlint-disable MD033 MD050 -->
# go-todotxt

[![go1.16+](https://img.shields.io/badge/Go-1.16+-blue?logo=go)](https://github.com/KEINOS/go-todotxt/blob/main/.github/workflows/unit-tests.yml#L81 "Supported versions")
[![Go Reference](https://pkg.go.dev/badge/github.com/KEINOS/go-todotxt.svg)](https://pkg.go.dev/github.com/KEINOS/go-todotxt/todo "View document")
[![License](https://img.shields.io/github/license/KEINOS/go-todotxt)](https://github.com/KEINOS/go-todotxt/blob/master/LICENSE)

`github.com/KEINOS/go-todotxt` is a **Go package that parses and manipulates tasks and to-do lists in [todo.txt format](https://github.com/todotxt/todo.txt)** by [Gina Trapani](https://github.com/ginatrapani).

It implements the custom user sort functionality as well.

> __Note__ This package is based on [**todotxt**](https://github.com/1set/todotxt) from [Kevin Tang](https://github.com/vt128) and [**go-todotxt**](https://github.com/JamesClonk/go-todotxt) from [Fabio Berchtold](https://github.com/JamesClonk).

## Usage

```go
go get github.com/KEINOS/go-todotxt
```

```go
import "github.com/KEINOS/go-todotxt/todo"
```

```go
// Load a todo.txt formatted file
tasksAll, err := todo.LoadFromPath("my_todo.txt");
if err != nil {
    log.Fatal(err)
}

// Retrieve uncompleted tasks, with due dates or priorities from the task list.
// - AND filter:
//     TaskList.Filter(todo.FilterDueToday).Filter(todo.FilterHasPriority)
// - OR filter:
//     TaskList.Filter(todo.FilterDueToday, todo.FilterHasPriority)
tasksToday := tasksAll.Filter(todo.FilterNotCompleted).Filter(
    todo.FilterDueToday,
    todo.FilterHasPriority,
)

// Sort the tasks by priority then project name.
if err := tasksToday.Sort(todo.SortPriorityAsc, todo.SortProjectAsc); err != nil {
    log.Fatal(err)
}

// Print each task info and set as completed.
for i, task := range tasksToday {
    fmt.Println(task.Todo)     // Print the task to be done
    fmt.Println(task.Priority) // Print its priority (if any)
    fmt.Println(task.Projects) // Print its projects name (if any)
    fmt.Println(task.Contexts) // Print its contexts (if any)

    tasks[i].Complete() // oh really?
}

// Save the tasks to a different file.
if err = tasks.WriteToPath("today-todo.txt"); err != nil {
    log.Fatal(err)
}
```

```go
func ExampleTaskList_CustomSort() {
    tasks := TaskList{
        Task{Todo: "Task 3"},
        Task{Todo: "Task 1"},
        Task{Todo: "Task 4"},
        Task{Todo: "Task 2"},
    }

    customFunc := func(a, b Task) bool {
        return strings.Compare(a.Todo, b.Todo) < 0
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
```

For more examples and details, please check the [Go Doc](https://pkg.go.dev/github.com/KEINOS/go-todotxt/todo#pkg-examples).

## Todo.txt format

| <img src="https://raw.githubusercontent.com/todotxt/todo.txt/master/description.svg" width="100%" > |
| :---: |
| (Image from: [https://github.com/todotxt/todo.txt](https://github.com/todotxt/todo.txt) ) |

## Contributing

[![go1.16+](https://img.shields.io/badge/Go-1.16+-blue?logo=go)](https://github.com/KEINOS/go-todotxt/blob/main/.github/workflows/unit-tests.yml#L81 "Supported versions")
[![Go Reference](https://pkg.go.dev/badge/github.com/KEINOS/go-todotxt.svg)](https://pkg.go.dev/github.com/KEINOS/go-todotxt/todo "View document")

Any contribution for the better is welcome. Please feel free to open an issue or a pull request.

- Branch to PR:
  - `main` ([Draft PR](https://github.blog/2019-02-14-introducing-draft-pull-requests/) is recommended)
- [Open an issue](https://github.com/KEINOS/go-todotxt/issues)
  - Please attach a simple and reproducible test code if possible. This helps us alot and to fix the issue faster.
- [CI](https://en.wikipedia.org/wiki/Continuous_integration)/[CD](https://en.wikipedia.org/wiki/Continuous_delivery):
  - The below tests will run on Push/Pull Request via GitHub Actions. You need to pass all the tests before requesting a review.
    - Unit testing on various Go versions (1.15 ... latest)
    - Unit testing on various platforms (Linux, macOS, Windows)
    - Static analysis/lint check by [golangci-lint](https://golangci-lint.run/).
      - Configuration: [.golangci.yml](./.golangci.yml)

> __Note__ : The branch `original` is a copy from the `master` branch of the [upstream repo](https://github.com/1set/todotxt). This is for the purpose of keeping the original code as is and contribute to the upstream. DO NOT PR to the `original` branch.

## Statuses

[![UnitTests](https://github.com/KEINOS/go-todotxt/actions/workflows/unit-tests.yml/badge.svg)](https://github.com/KEINOS/go-todotxt/actions/workflows/unit-tests.yml)
[![golangci-lint](https://github.com/KEINOS/go-todotxt/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/KEINOS/go-todotxt/actions/workflows/golangci-lint.yml)
[![CodeQL-Analysis](https://github.com/KEINOS/go-todotxt/actions/workflows/codeQL-analysis.yml/badge.svg)](https://github.com/KEINOS/go-todotxt/actions/workflows/codeQL-analysis.yml)
[![PlatformTests](https://github.com/KEINOS/go-todotxt/actions/workflows/platform-tests.yml/badge.svg)](https://github.com/KEINOS/go-todotxt/actions/workflows/platform-tests.yml "Tests on Win, macOS and Linux")

[![codecov](https://codecov.io/gh/KEINOS/go-todotxt/branch/main/graph/badge.svg?token=JVY7WUeUFz)](https://codecov.io/gh/KEINOS/go-todotxt)
[![Go Report Card](https://goreportcard.com/badge/github.com/KEINOS/go-todotxt)](https://goreportcard.com/report/github.com/KEINOS/go-todotxt)

## License and Credits

- MIT License. Copyright (c) 2022 [KEINOS and the go-todotxt contributors](https://github.com/KEINOS/go-todotxt/graphs/contributors) with all the respect to Kevin Tang, Fabio Berchtold and Gina Trapani.
- This package is based on the below packages and ideas:
  - Mother/Upstream: [**todotxt**](https://github.com/1set/todotxt) authored by [Kevin Tang](https://github.com/vt128) @ [MIT](https://github.com/1set/todotxt/blob/master/LICENSE)
  - Grand Mother/Most upstream: [**go-todotxt**](https://github.com/JamesClonk/go-todotxt) authored by [Fabio Berchtold](https://github.com/JamesClonk) @ [MPL-2.0](https://github.com/JamesClonk/go-todotxt/blob/master/LICENSE)
  - Origin: [**todo.txt**](https://github.com/todotxt/todo.txt) is an awesome task format. Initially designed by [Gina Trapani](https://github.com/ginatrapani). @ [GPL-3.0](https://github.com/todotxt/todo.txt/blob/master/LICENSE)

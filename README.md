# todotxt

Yet another a Go library for Gina Trapani's todo.txt files. ✅

[![PkgGoDev](https://pkg.go.dev/badge/github.com/1set/todotxt)](https://pkg.go.dev/github.com/1set/todotxt)
[![License](https://img.shields.io/github/license/1set/todotxt)](https://github.com/1set/todotxt/blob/master/LICENSE)
[![GitHub Action Workflow](https://github.com/1set/todotxt/workflows/build/badge.svg)](https://github.com/1set/todotxt/actions?workflow=build)
[![Go Report Card](https://goreportcard.com/badge/github.com/1set/todotxt)](https://goreportcard.com/report/github.com/1set/todotxt)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/448bb14650984e21a7c41bcac221a1ba)](https://www.codacy.com/gh/1set/todotxt/dashboard)
[![Codecov](https://codecov.io/gh/1set/todotxt/branch/master/graph/badge.svg)](https://codecov.io/gh/1set/todotxt)

## Features

Based on [**go-todotxt**](https://github.com/JamesClonk/go-todotxt) from [Fabio Berchtold](https://github.com/JamesClonk) with:

- [x] Go mod support
- [x] Segments for task string
- [x] Task due today is not overdue
- [x] Negative `Due()` for overdue tasks
- [x] Support multiple options for sorting and filtering
- [x] More sorting options: by ID, text, context, project
- [x] Preset filters

## Usage

A quick start example:

```go
import (
	todo "github.com/1set/todotxt"
)

// ...

if tasklist, err := todo.LoadFromPath("todo.txt"); err != nil {
    log.Fatal(err)
} else {
    tasks := tasklist.Filter(todo.FilterNotCompleted).Filter(todo.FilterDueToday, todo.FilterHasPriority)
    _ = tasks.Sort(todo.SortPriorityAsc, todo.SortProjectAsc)
    for i, t := range tasks {
        fmt.Println(t.Todo)
        // oh really?
        tasks[i].Complete()
    }
    if err = tasks.WriteToPath("today-todo.txt"); err != nil {
        log.Fatal(err)
    }
}
```

For more examples and details, please check the [Go Doc](https://pkg.go.dev/github.com/1set/todotxt).

## Credits

- Original [**go-todotxt**](https://github.com/JamesClonk/go-todotxt) authored by [Fabio Berchtold](https://github.com/JamesClonk)
- Cool idea about [**todo.txt**](https://github.com/todotxt/todo.txt) from [Gina Trapani](http://todotxt.org/)

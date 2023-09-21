# Usages

These are examples of how to use this package.
For more examples and details, please check the [Go Doc](https://pkg.go.dev/github.com/KEINOS/go-todotxt/todo#pkg-examples).

## Load Tasks

```go
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
```

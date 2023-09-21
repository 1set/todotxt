<!-- markdownlint-disable MD033 MD050 -->
# go-todotxt

[![go1.18+](https://img.shields.io/badge/Go-1.18+-blue?logo=go)](https://github.com/KEINOS/go-todotxt/blob/main/.github/workflows/unit-tests.yml#L81 "Supported versions")
[![Go Reference](https://pkg.go.dev/badge/github.com/KEINOS/go-todotxt.svg)](https://pkg.go.dev/github.com/KEINOS/go-todotxt/todo "View document")
[![License](https://img.shields.io/github/license/KEINOS/go-todotxt)](https://github.com/KEINOS/go-todotxt/blob/master/LICENSE)

`github.com/KEINOS/go-todotxt` is a Go package for parsing and editing todo.txt files, a [text format for task annotations](https://github.com/todotxt/todo.txt) designed by [Gina Trapani](https://github.com/ginatrapani).

> __Note__: This package is based on the following packages with **custom user sort functionality**.
>
> - [**todotxt**](https://github.com/1set/todotxt) from [Kevin Tang](https://github.com/vt128)
> - [**go-todotxt**](https://github.com/JamesClonk/go-todotxt) from [Fabio Berchtold](https://github.com/JamesClonk)

## Usage

```go
// Download the package.
go get "github.com/KEINOS/go-todotxt"
```

```go
// Import the package.
import "github.com/KEINOS/go-todotxt/todo"
```

```go
func Example() {
    // Load tasks from a string. You can also load from a file by using LoadFromFile().
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

    // AND filter.  Get tasks that have priority A and are not completed.
    nearTopTasks := tasks.Filter(todo.FilterHasPriority).Filter(todo.FilterNotCompleted)

    fmt.Println(nearTopTasks.String())

    // Output:
    // (A) Call Mom @Phone +Family
    // (B) Outline chapter 5 @Computer +Novel
    // (C) Add cover sheets @Office +TPSReports
}
```

## Todo.txt format

![](https://raw.githubusercontent.com/todotxt/todo.txt/master/description.svg)

- Image from: [https://github.com/todotxt/todo.txt](https://github.com/todotxt/todo.txt)

## Contributing

[![go1.18+](https://img.shields.io/badge/Go-1.18+-blue?logo=go)](https://github.com/KEINOS/go-todotxt/blob/main/.github/workflows/unit-tests.yml#L81 "Supported versions")
[![Go Reference](https://pkg.go.dev/badge/github.com/KEINOS/go-todotxt.svg)](https://pkg.go.dev/github.com/KEINOS/go-todotxt/todo "View document")

Any contribution for the better is welcome. We provide full code coverage of unit tests, so feel free to refactor or play around with the code.

- Branch to PR:
  - `main` ([Draft PR](https://github.blog/2019-02-14-introducing-draft-pull-requests/) is recommended)
- [Open an issue](https://github.com/KEINOS/go-todotxt/issues)
  - Please attach a simple and reproducible test code if possible. This helps us alot and to fix the issue faster.
- [CI](https://en.wikipedia.org/wiki/Continuous_integration)/[CD](https://en.wikipedia.org/wiki/Continuous_delivery):
  - The below tests will run on Push/Pull Request via GitHub Actions. You need to pass all the tests before requesting a review.
    - Unit testing on various Go versions (1.18 ... latest)
    - Unit testing on various platforms (Linux, macOS, Windows)
    - Static analysis/lint check by [golangci-lint](https://golangci-lint.run/).
      - Configuration: [.golangci.yml](./.golangci.yml)
  - To **run tests locally**, we provide a convenient [Makefile](./Makefile). Please run the below command to run all the tests (`docker` and `compose` are required).

    ```bash
    # Runs unit tests on Go 1.18 to latest and `golangci-lint` check.
    make test
    ```

> __Note__ : Please **DO NOT PR to the `original` branch** but to `main` branch. The branch `original` is simply a copy from the `master` branch of the [upstream repo](https://github.com/1set/todotxt). This is for the purpose of keeping the original code as is and contribute to the upstream.

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

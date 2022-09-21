/*
Package todo parses and manipulates tasks and to-do lists in todo.txt format
by Gina Trapani.
*/
package todo

// Go generate directives.
//
// These will generate the stringer implementations for TaskSortByType and
// TaskSegmentType types.
// Note that to call `go generate ./...` you need `stringer` command installed.
// You can use `docker compose run go_generate` for convenience.
//
//go:generate stringer -type TaskSortByType -trimprefix Sort -output tasksortbytype_string.go
//go:generate stringer -type TaskSegmentType -trimprefix Segment -output tasksegmenttype_string.go

package todo

import "strings"

// ----------------------------------------------------------------------------
//  Type: Predicate
// ----------------------------------------------------------------------------

// Predicate is a function type that takes a task as an input and returns a bool.
//
// Which is used to filter a TaskList as a "Functional Options Pattern" like style.
type Predicate func(Task) bool

// ----------------------------------------------------------------------------
//  Constructors
// ----------------------------------------------------------------------------

// FilterByContext returns a filter for tasks that have the given context.
// String comparison in the filters is case-insensitive.
func FilterByContext(context string) Predicate {
	return func(t Task) bool {
		for _, c := range t.Contexts {
			if strings.EqualFold(c, context) {
				return true
			}
		}

		return false
	}
}

// FilterByPriority returns a filter for tasks that have the given priority.
// String comparison in the filters is case-insensitive.
func FilterByPriority(priority string) Predicate {
	priority = strings.ToUpper(priority)

	return func(t Task) bool {
		return t.Priority == priority
	}
}

// FilterByProject returns a filter for tasks that have the given project.
// String comparison in the filters is case-insensitive.
func FilterByProject(project string) Predicate {
	return func(t Task) bool {
		for _, p := range t.Projects {
			if strings.EqualFold(p, project) {
				return true
			}
		}

		return false
	}
}

// FilterNot returns a reversed filter for existing predicate.
func FilterNot(predicate Predicate) Predicate {
	return func(t Task) bool {
		return !predicate(t)
	}
}

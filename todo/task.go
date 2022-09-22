package todo

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// ----------------------------------------------------------------------------
//  Type: Task
// ----------------------------------------------------------------------------

// Task represents a todo.txt task entry.
//
// The 'Contexts' and 'Projects' are both used to categorize tasks.
// The difference is that 'Contexts' are used to categorize tasks by location or
// situation where you'll work on the job, while 'Projects' are used to categorize
// tasks by project.
//
// For the "todo.txt" format specification see:
// https://github.com/todotxt/todo.txt#todotxt-format-rules
//
//nolint:godox // The todo in the comment below is not a TODO.
type Task struct {
	DueDate        time.Time         // DueDate is the due date calculated from the 'due:' tag.
	CompletedDate  time.Time         // CompletedDate is the date the task was completed.
	CreatedDate    time.Time         // CreatedDate is the date the task was created.
	AdditionalTags map[string]string // AdditionalTags of the task in a key:value format (e.g. "due:2012-12-12")
	Original       string            // Original raw task text.
	Priority       string            // Priority of the task in (A)-(Z) range.
	Todo           string            // Todo part of task text.
	Contexts       []string          // Contexts of the task (e.g. @MyContext).
	Projects       []string          // Projects of the task (e.g. +MyProject).
	ID             int               // ID of the task internaly.
	Completed      bool              // Completed flag. If true, the task has been completed.
}

// ----------------------------------------------------------------------------
//  Constructors
// ----------------------------------------------------------------------------

// NewTask creates a new empty Task with default values. (CreatedDate is set to Now()).
func NewTask() Task {
	task := new(Task)
	task.CreatedDate = time.Now()

	return *task
}

// ParseTask parses the input text string into a Task struct.
func ParseTask(text string) (*Task, error) {
	var err error

	oriText := strings.Trim(text, whitespaces)
	task := new(Task)
	task.Original = oriText
	task.Todo = oriText

	// Check for completed
	if completedRx.MatchString(oriText) {
		if err := parseCompleted(oriText, task); err != nil {
			return nil, errors.Wrap(err, "failed to parse task")
		}
	}

	// Check for priority
	if priorityRx.MatchString(oriText) {
		parsePriority(oriText, task)
	}

	// Check for created date
	if createdDateRx.MatchString(oriText) {
		if err := parseCreatedDate(oriText, task); err != nil {
			return nil, errors.Wrap(err, "failed to parse task")
		}
	}

	// Check for contexts
	if contextRx.MatchString(oriText) {
		task.Contexts = getSlice(oriText, contextRx)
		task.Todo = contextRx.ReplaceAllString(task.Todo, emptyStr) // Remove from Todo text
	}

	// Check for projects
	if projectRx.MatchString(oriText) {
		task.Projects = getSlice(oriText, projectRx)
		task.Todo = projectRx.ReplaceAllString(task.Todo, emptyStr) // Remove from Todo text
	}

	// Check for additional tags
	if addonTagRx.MatchString(oriText) {
		if err := parseAdditionalTags(oriText, task); err != nil {
			return nil, errors.Wrap(err, "failed to parse task")
		}
	}

	// Trim any remaining whitespaces from Todo text
	task.Todo = strings.Trim(task.Todo, "\t\n\r\f ")

	return task, err
}

// ----------------------------------------------------------------------------
//  Methods
// ----------------------------------------------------------------------------

// Complete sets Task.Completed to 'true' if the task was not already completed.
// Also sets Task.CompletedDate to time.Now().
func (task *Task) Complete() {
	if !task.Completed {
		task.Completed = true
		task.CompletedDate = time.Now()
	}
}

// Due returns the duration left until due date from now. The duration is negative
// if the task is overdue.
//
// Just as with IsOverdue(), this function does also not take the Completed flag
// into consideration. You should check Task.Completed first if needed.
func (task *Task) Due() time.Duration {
	return time.Until(task.DueDate.AddDate(0, 0, 1))
}

// HasAdditionalTags returns true if the task has any additional tags.
func (task *Task) HasAdditionalTags() bool {
	return len(task.AdditionalTags) > 0
}

// HasCompletedDate returns true if the task has a completed date.
func (task *Task) HasCompletedDate() bool {
	return !task.CompletedDate.IsZero() && task.Completed
}

// HasContexts returns true if the task has any contexts.
func (task *Task) HasContexts() bool {
	return len(task.Contexts) > 0
}

// HasCreatedDate returns true if the task has a created date.
func (task *Task) HasCreatedDate() bool {
	return !task.CreatedDate.IsZero()
}

// HasDueDate returns true if the task has a due date.
func (task *Task) HasDueDate() bool {
	return !task.DueDate.IsZero()
}

// HasPriority returns true if the task has a priority.
func (task *Task) HasPriority() bool {
	return isNotEmpty(task.Priority)
}

// HasProjects returns true if the task has any projects.
func (task *Task) HasProjects() bool {
	return len(task.Projects) > 0
}

// IsCompleted returns true if the task has already been completed.
func (task *Task) IsCompleted() bool {
	return task.Completed
}

// IsDueToday returns true if the task is due todasy.
func (task *Task) IsDueToday() bool {
	if task.HasDueDate() {
		due := task.Due()

		return 0 < due && due <= oneDay
	}

	return false
}

// IsOverdue returns true if due date is in the past.
//
// This function does not take the Completed flag into consideration.
// You should check Task.Completed first if needed.
func (task *Task) IsOverdue() bool {
	if task.HasDueDate() {
		return task.Due() < 0
	}

	return false
}

// Reopen sets Task.Completed to 'false' if the task was completed.
// Also resets Task.CompletedDate.
func (task *Task) Reopen() {
	if task.Completed {
		task.Completed = false
		task.CompletedDate = time.Time{} // time.IsZero() value
	}
}

// String returns a complete task string in todo.txt format.
//
// Contexts, Projects and additional tags are alphabetically sorted,
// and appended at the end in the following order:
// Contexts, Projects, Tags
//
// For example:
//
//	"(A) 2013-07-23 Call Dad @Home @Phone +Family due:2013-07-31 customTag1:Important!"
//
//nolint:cyclop // complexity is high (=15), but leave it as is for now
func (task Task) String() string {
	var strBld strings.Builder

	if task.Completed {
		strBld.WriteString("x ")

		if task.HasCompletedDate() {
			strBld.WriteString(fmt.Sprintf("%s ", task.CompletedDate.Format(DateLayout)))
		}
	}

	if task.HasPriority() && (!task.Completed || !RemoveCompletedPriority) {
		strBld.WriteString(fmt.Sprintf("(%s) ", task.Priority))
	}

	if task.HasCreatedDate() {
		strBld.WriteString(fmt.Sprintf("%s ", task.CreatedDate.Format(DateLayout)))
	}

	strBld.WriteString(task.Todo)

	if task.HasContexts() {
		sort.Strings(task.Contexts)

		for _, context := range task.Contexts {
			strBld.WriteString(fmt.Sprintf(" @%s", context))
		}
	}

	if task.HasProjects() {
		sort.Strings(task.Projects)

		for _, project := range task.Projects {
			strBld.WriteString(fmt.Sprintf(" +%s", project))
		}
	}

	if task.HasAdditionalTags() {
		// Sort map alphabetically by keys
		keys := make([]string, 0, len(task.AdditionalTags))

		for key := range task.AdditionalTags {
			keys = append(keys, key)
		}

		sort.Strings(keys)

		for _, key := range keys {
			strBld.WriteString(fmt.Sprintf(" %s:%s", key, task.AdditionalTags[key]))
		}
	}

	if task.HasDueDate() {
		strBld.WriteString(fmt.Sprintf(" due:%s", task.DueDate.Format(DateLayout)))
	}

	return strBld.String()
}

// Task returns a complete task string in todo.txt format.
//
// It is an alias of String(). See *Task.String() for further information.
func (task *Task) Task() string {
	return task.String()
}

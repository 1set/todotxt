package todotxt

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"
)

var (
	// DateLayout is used for formatting time.Time into todo.txt date format and vice-versa.
	DateLayout = "2006-01-02"

	priorityRx = regexp.MustCompile(`^(x|x \d{4}-\d{2}-\d{2}|)\s*\(([A-Z])\)\s+`) // Match priority: '(A) ...' or 'x (A) ...' or 'x 2012-12-12 (A) ...'
	// Match created date: '(A) 2012-12-12 ...' or 'x 2012-12-12 (A) 2012-12-12 ...' or 'x (A) 2012-12-12 ...'or 'x 2012-12-12 2012-12-12 ...' or '2012-12-12 ...'
	createdDateRx   = regexp.MustCompile(`^(\([A-Z]\)|x \d{4}-\d{2}-\d{2} \([A-Z]\)|x \([A-Z]\)|x \d{4}-\d{2}-\d{2}|)\s*(\d{4}-\d{2}-\d{2})\s+`)
	completedRx     = regexp.MustCompile(`^x\s+`)                       // Match completed: 'x ...'
	completedDateRx = regexp.MustCompile(`^x\s*(\d{4}-\d{2}-\d{2})\s+`) // Match completed date: 'x 2012-12-12 ...'
	addonTagRx      = regexp.MustCompile(`(^|\s+)([^:\s]+):([^:\s]+)`)  // Match additional tags date: '... due:2012-12-12 ...'
	contextRx       = regexp.MustCompile(`(^|\s+)@(\S+)`)               // Match contexts: '@Context ...' or '... @Context ...'
	projectRx       = regexp.MustCompile(`(^|\s+)\+(\S+)`)              // Match projects: '+Project...' or '... +Project ...')
)

// Task represents a todo.txt task entry.
type Task struct {
	ID             int    // Internal task ID.
	Original       string // Original raw task text.
	Todo           string // Todo part of task text.
	Priority       string
	Projects       []string
	Contexts       []string
	AdditionalTags map[string]string // Addon tags will be available here.
	CreatedDate    time.Time
	DueDate        time.Time
	CompletedDate  time.Time
	Completed      bool
}

// NewTask creates a new empty Task with default values. (CreatedDate is set to Now())
func NewTask() Task {
	task := Task{}
	task.CreatedDate = time.Now()
	return task
}

// Task returns a complete task string in todo.txt format.
// See *Task.String() for further information.
func (task *Task) Task() string {
	return task.String()
}

// String returns a complete task string in todo.txt format.
//
// Contexts, Projects and additional tags are alphabetically sorted,
// and appended at the end in the following order:
// Contexts, Projects, Tags
//
// For example:
//  "(A) 2013-07-23 Call Dad @Home @Phone +Family due:2013-07-31 customTag1:Important!"
func (task Task) String() string {
	var sb strings.Builder

	if task.Completed {
		sb.WriteString("x ")
		if task.HasCompletedDate() {
			sb.WriteString(fmt.Sprintf("%s ", task.CompletedDate.Format(DateLayout)))
		}
	}

	if task.HasPriority() && (!task.Completed || !RemoveCompletedPriority) {
		sb.WriteString(fmt.Sprintf("(%s) ", task.Priority))
	}

	if task.HasCreatedDate() {
		sb.WriteString(fmt.Sprintf("%s ", task.CreatedDate.Format(DateLayout)))
	}

	sb.WriteString(task.Todo)

	if task.HasContexts() {
		sort.Strings(task.Contexts)
		for _, context := range task.Contexts {
			sb.WriteString(fmt.Sprintf(" @%s", context))
		}
	}

	if task.HasProjects() {
		sort.Strings(task.Projects)
		for _, project := range task.Projects {
			sb.WriteString(fmt.Sprintf(" +%s", project))
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
			sb.WriteString(fmt.Sprintf(" %s:%s", key, task.AdditionalTags[key]))
		}
	}

	if task.HasDueDate() {
		sb.WriteString(fmt.Sprintf(" due:%s", task.DueDate.Format(DateLayout)))
	}

	return sb.String()
}

// ParseTask parses the input text string into a Task struct.
func ParseTask(text string) (*Task, error) {
	var err error

	oriText := strings.Trim(text, whitespaces)
	task := Task{}
	task.Original = oriText
	task.Todo = oriText

	// Check for completed
	if completedRx.MatchString(oriText) {
		task.Completed = true
		// Check for completed date
		if completedDateRx.MatchString(oriText) {
			if date, err := parseTime(completedDateRx.FindStringSubmatch(oriText)[1]); err == nil {
				task.CompletedDate = date
			} else {
				return nil, err
			}
		}

		// Remove from Todo text
		task.Todo = completedDateRx.ReplaceAllString(task.Todo, emptyStr) // Strip CompletedDate first, otherwise it wouldn't match anymore (^x date...)
		task.Todo = completedRx.ReplaceAllString(task.Todo, emptyStr)     // Strip 'x '
	}

	// Check for priority
	if priorityRx.MatchString(oriText) {
		task.Priority = priorityRx.FindStringSubmatch(oriText)[2]
		task.Todo = priorityRx.ReplaceAllString(task.Todo, emptyStr) // Remove from Todo text
	}

	// Check for created date
	if createdDateRx.MatchString(oriText) {
		if date, err := parseTime(createdDateRx.FindStringSubmatch(oriText)[2]); err == nil {
			task.CreatedDate = date
			task.Todo = createdDateRx.ReplaceAllString(task.Todo, emptyStr) // Remove from Todo text
		} else {
			return nil, err
		}
	}

	// function for collecting projects/contexts as slices from text
	getSlice := func(rx *regexp.Regexp) []string {
		matches := rx.FindAllStringSubmatch(oriText, -1)
		slice := make([]string, 0, len(matches))
		seen := make(map[string]bool, len(matches))
		for _, match := range matches {
			word := strings.Trim(match[2], whitespaces)
			if _, found := seen[word]; !found {
				slice = append(slice, word)
				seen[word] = true
			}
		}
		sort.Strings(slice)
		return slice
	}

	// Check for contexts
	if contextRx.MatchString(oriText) {
		task.Contexts = getSlice(contextRx)
		task.Todo = contextRx.ReplaceAllString(task.Todo, emptyStr) // Remove from Todo text
	}

	// Check for projects
	if projectRx.MatchString(oriText) {
		task.Projects = getSlice(projectRx)
		task.Todo = projectRx.ReplaceAllString(task.Todo, emptyStr) // Remove from Todo text
	}

	// Check for additional tags
	if addonTagRx.MatchString(oriText) {
		matches := addonTagRx.FindAllStringSubmatch(oriText, -1)
		tags := make(map[string]string, len(matches))
		for _, match := range matches {
			key, value := match[2], match[3]
			if key == "due" { // due date is a known addon tag, it has its own struct field
				if date, err := parseTime(value); err == nil {
					task.DueDate = date
				} else {
					return nil, err
				}
			} else if isNotEmpty(key) && isNotEmpty(value) {
				tags[key] = value
			}
		}
		task.AdditionalTags = tags
		task.Todo = addonTagRx.ReplaceAllString(task.Todo, emptyStr) // Remove from Todo text
	}

	// Trim any remaining whitespaces from Todo text
	task.Todo = strings.Trim(task.Todo, "\t\n\r\f ")

	return &task, err
}

// HasProjects returns true if the task has any projects.
func (task *Task) HasProjects() bool {
	return len(task.Projects) > 0
}

// HasContexts returns true if the task has any contexts.
func (task *Task) HasContexts() bool {
	return len(task.Contexts) > 0
}

// HasAdditionalTags returns true if the task has any additional tags.
func (task *Task) HasAdditionalTags() bool {
	return len(task.AdditionalTags) > 0
}

// HasPriority returns true if the task has a priority.
func (task *Task) HasPriority() bool {
	return isNotEmpty(task.Priority)
}

// HasCreatedDate returns true if the task has a created date.
func (task *Task) HasCreatedDate() bool {
	return !task.CreatedDate.IsZero()
}

// HasCompletedDate returns true if the task has a completed date.
func (task *Task) HasCompletedDate() bool {
	return !task.CompletedDate.IsZero() && task.Completed
}

// IsCompleted returns true if the task has already been completed.
func (task *Task) IsCompleted() bool {
	return task.Completed
}

// Complete sets Task.Completed to 'true' if the task was not already completed.
// Also sets Task.CompletedDate to time.Now()
func (task *Task) Complete() {
	if !task.Completed {
		task.Completed = true
		task.CompletedDate = time.Now()
	}
}

// Reopen sets Task.Completed to 'false' if the task was completed.
// Also resets Task.CompletedDate.
func (task *Task) Reopen() {
	if task.Completed {
		task.Completed = false
		task.CompletedDate = time.Time{} // time.IsZero() value
	}
}

// HasDueDate returns true if the task has a due date.
func (task *Task) HasDueDate() bool {
	return !task.DueDate.IsZero()
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

// IsDueToday returns true if the task is due todasy.
func (task *Task) IsDueToday() bool {
	if task.HasDueDate() {
		due := task.Due()
		return 0 < due && due <= oneDay
	}
	return false
}

// Due returns the duration left until due date from now. The duration is negative if the task is overdue.
//
// Just as with IsOverdue(), this function does also not take the Completed flag into consideration.
// You should check Task.Completed first if needed.
func (task *Task) Due() time.Duration {
	return time.Until(task.DueDate.AddDate(0, 0, 1))
}

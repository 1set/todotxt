package todo

import (
	"os"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// ============================================================================
//  Declaration of global private/public function, variable, constants
// ============================================================================

// ----------------------------------------------------------------------------
//  Constants
// ----------------------------------------------------------------------------

const (
	oneDay      = 24 * time.Hour
	whitespaces = "\t\n\r "
	emptyStr    = ""

	// PermReadWrite represents the permission bits for "rw- r-- ---".
	PermReadWrite = 0o640
	// PermReadWriteExec represents the permission bits for "rwx r-x r-x".
	PermReadWriteExec = 0o755
	// DateLayout is used for formatting time.Time into todo.txt date format and vice-versa.
	DateLayout = "2006-01-02"
)

// ----------------------------------------------------------------------------
//  Global variables
// ----------------------------------------------------------------------------

// IgnoreComments can be set to 'false', in order to revert to a more standard
// behaviour of todo.txt.
// The todo.txt format does not define comments.
//
//nolint:gochecknoglobals // global variable is intentional
var (
	// IgnoreComments is used to switch ignoring of comments (lines starting
	// with "#"). If this is set to 'false', then lines starting with "#" will
	// be parsed as tasks.
	IgnoreComments = true

	// RemoveCompletedPriority is used to switch discarding priority on task
	// completion like many todo.txt clients do. If this is set to 'false', then
	// the priority of completed task will be kept as it is.
	RemoveCompletedPriority = true
)

var (
	// Match priority: '(A) ...' or 'x (A) ...' or 'x 2012-12-12 (A) ...'.
	priorityRx = regexp.MustCompile(`^(x|x \d{4}-\d{2}-\d{2}|)\s*\(([A-Z])\)\s+`)
	// Match created date:
	//   '(A) 2012-12-12 ...' or 'x 2012-12-12 (A) 2012-12-12 ...'
	// or
	//   'x (A) 2012-12-12 ...' or 'x 2012-12-12 2012-12-12 ...' or '2012-12-12 ...'.
	createdDateRx = regexp.MustCompile(
		`^(\([A-Z]\)|x \d{4}-\d{2}-\d{2} \([A-Z]\)|x \([A-Z]\)|x \d{4}-\d{2}-\d{2}|)\s*(\d{4}-\d{2}-\d{2})\s+`,
	)
	// Match completed: 'x ...'.
	completedRx = regexp.MustCompile(`^x\s+`)
	// Match completed date: 'x 2012-12-12 ...'.
	completedDateRx = regexp.MustCompile(`^x\s*(\d{4}-\d{2}-\d{2})\s+`)
	// Match additional tags date: '... due:2012-12-12 ...'.
	addonTagRx = regexp.MustCompile(`(^|\s+)([^:\s]+):([^:\s]+)`)
	// Match contexts: '@Context ...' or '... @Context ...'.
	contextRx = regexp.MustCompile(`(^|\s+)@(\S+)`)
	// Match projects: '+Project...' or '... +Project ...'.
	projectRx = regexp.MustCompile(`(^|\s+)\+(\S+)`)
)

// ----------------------------------------------------------------------------
//  Public functions
// ----------------------------------------------------------------------------

// WriteToFile writes a TaskList to *os.File.
//
// Using *os.File instead of a filename allows to also use os.Stdout.
//
// Note: Comments from original file will be omitted and not written to target *os.File,
// if IgnoreComments is set to 'true'.
func WriteToFile(tasklist *TaskList, file *os.File) error {
	return tasklist.WriteToFile(file)
}

// WriteToPath writes a TaskList to the specified file (most likely called "todo.txt").
func WriteToPath(tasklist *TaskList, filename string) error {
	return tasklist.WriteToPath(filename)
}

// ----------------------------------------------------------------------------
//  Private functions
// ----------------------------------------------------------------------------

// It will collect projects/contexts from txtOrig and returns them as a slice.
func getSlice(txtOrig string, rx *regexp.Regexp) []string {
	matches := rx.FindAllStringSubmatch(txtOrig, -1)
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

// isEmpty checks if the string is empty.
func isEmpty(s string) bool {
	return len(s) == 0
}

// isNotEmpty checks if the string is not empty.
func isNotEmpty(s string) bool {
	return len(s) > 0
}

// lessStrings checks if the string slices a is exactly less than b in lexicographical
// order.
//
//nolint:cyclop,varnamelen // complexity=11 and param name a,b is short but it's ok
func lessStrings(a, b []string) bool {
	// set lenMin to the length of a by default
	lenA, lenB, lenMin := len(a), len(b), len(a)

	switch {
	case lenA == 0 && lenB == 0:
		return false
	case lenA == 0 && lenB > 0:
		return false
	case lenA > 0 && lenB == 0:
		return true
	case lenA > lenB:
		// swap lenMin from lenA to lenB
		lenMin = lenB
	}

	for i := 0; i < lenMin; i++ {
		if a[i] < b[i] {
			return true
		}

		if a[i] > b[i] {
			return false
		}
	}

	// Note that lenA == lenB will always be false here
	return lenA < lenB
}

func parseAdditionalTags(txtOrig string, task *Task) error {
	matches := addonTagRx.FindAllStringSubmatch(txtOrig, -1)
	tags := make(map[string]string, len(matches))

	for _, match := range matches {
		key, value := match[2], match[3]

		// due date is a known addon tag, it has its own struct field
		if key == "due" {
			date, err := parseTime(value)
			if err != nil {
				return errors.Wrap(err, "failed to parse time of due date")
			}

			task.DueDate = date
		} else if isNotEmpty(key) && isNotEmpty(value) {
			// add other tags rather than due date to the map
			tags[key] = value
		}
	}

	task.AdditionalTags = tags
	task.Todo = addonTagRx.ReplaceAllString(task.Todo, emptyStr) // Remove from Todo text

	return nil
}

func parseCompleted(txtOrig string, task *Task) error {
	task.Completed = true

	// Check for completed date
	if completedDateRx.MatchString(txtOrig) {
		date, err := parseTime(completedDateRx.FindStringSubmatch(txtOrig)[1])
		if err != nil {
			return errors.Wrap(err, "failed to parse completed date")
		}

		task.CompletedDate = date
	}

	/* Remove from Todo text */
	// Strip CompletedDate first, otherwise it wouldn't match anymore (^x date...)
	task.Todo = completedDateRx.ReplaceAllString(task.Todo, emptyStr)
	// Strip 'x '
	task.Todo = completedRx.ReplaceAllString(task.Todo, emptyStr)

	return nil
}

func parseCreatedDate(txtOrig string, task *Task) error {
	date, err := parseTime(createdDateRx.FindStringSubmatch(txtOrig)[2])
	if err != nil {
		return errors.Wrap(err, "failed to parse time of created date")
	}

	task.CreatedDate = date
	task.Todo = createdDateRx.ReplaceAllString(task.Todo, emptyStr) // Remove from Todo text

	return nil
}

func parsePriority(txtOrig string, task *Task) {
	task.Priority = priorityRx.FindStringSubmatch(txtOrig)[2]
	task.Todo = priorityRx.ReplaceAllString(task.Todo, emptyStr) // Remove from Todo text
}

// parseTime parses a string as a local time into a time.Time struct.
func parseTime(s string) (time.Time, error) {
	//nolint:gosmopolitan //
	parsed, err := time.ParseInLocation(DateLayout, s, time.Local)
	if err != nil {
		return time.Time{}, errors.Wrap(err, "failed to parse time")
	}

	return parsed, nil
}

func sortByDate(asc bool, hasDate1, hasDate2 bool, date1, date2 time.Time) bool {
	// ASC
	if asc {
		if hasDate1 && hasDate2 {
			return date1.Before(date2)
		}

		return hasDate2
	}

	// DESC
	if hasDate1 && hasDate2 {
		return date1.After(date2)
	}

	return !hasDate2
}

package todotxt

import (
    "strings"
    "time"
)

var (
	emptyStr    string
	whitespaces = "\t\n\r "
	oneDay      = 24 * time.Hour
)

// ignoreLine uses the IgnoreComments variable and the presence of
//      a '#' to detect if a line should be ignored. Empty lines are
//      ignored automatically.
func ignoreLine(line string) bool {
    if isEmpty(line) {
        return true
    } else if IgnoreComments && strings.HasPrefix(line, "#") {
        return true
    }
    return false
}

// isEmpty checks if the string is empty.
func isEmpty(s string) bool {
	return len(s) == 0
}

// isNotEmpty checks if the string is not empty.
func isNotEmpty(s string) bool {
	return len(s) > 0
}

func parseTime(s string) (time.Time, error) {
	return time.ParseInLocation(DateLayout, s, time.Local)
}

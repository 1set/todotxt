package todotxt

import "time"

var (
	emptyStr    string
	whitespaces = "\t\n\r "
	oneDay      = 24 * time.Hour
)

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

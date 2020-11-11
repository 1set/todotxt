package todotxt

var (
	emptyStr    string
	whitespaces = "\t\n\r "
)

// isEmpty checks if the string is empty.
func isEmpty(s string) bool {
	return len(s) == 0
}

// isNotEmpty checks if the string is not empty.
func isNotEmpty(s string) bool {
	return len(s) > 0
}

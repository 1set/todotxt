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

// lessStrings checks if the string slices a is exactly less than b.
func lessStrings(a, b []string) bool {
	la, lb, min := len(a), len(b), 0
	if la == 0 && lb == 0 {
		return false
	} else if la == 0 && lb > 0 {
		return true
	} else if la > 0 && lb == 0 {
		return false
	}

	if min = la; la > lb {
		min = lb
	}
	for i := 0; i < min; i++ {
		if a[i] < b[i] {
			return true
		} else if a[i] > b[i] {
			return false
		}
	}

	if la == lb {
		return false
	}
	return la < lb
}

package todo

// ----------------------------------------------------------------------------
//  TaskList.Filter()
// ----------------------------------------------------------------------------

// Filter filters the current TaskList for the given predicate, and returns a
// new TaskList. The original TaskList is not modified.
//
//	For the Predicate type filters see the todo/filters.go file.
func (tasklist TaskList) Filter(filter Predicate, filters ...Predicate) TaskList {
	combined := []Predicate{filter}
	combined = append(combined, filters...)

	var newList TaskList

	for _, task := range tasklist {
		for _, filt := range combined {
			// Append tasks to the new list if the filter returns true.
			if filt(task) {
				newList = append(newList, task)

				break
			}
		}
	}

	return newList
}

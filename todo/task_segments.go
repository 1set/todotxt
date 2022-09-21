package todo

import (
	"fmt"
	"sort"
)

// Segments returns a segmented task string in todo.txt format. The order of
// segments is the same as String().
//
//nolint:funlen, cyclop // length is 77 and complexity is 15 but leave it as is for now
func (task *Task) Segments() []*TaskSegment {
	var segs []*TaskSegment

	newBasicTaskSeg := func(t TaskSegmentType, s string) *TaskSegment {
		return &TaskSegment{
			Type:      t,
			Originals: []string{s},
			Display:   s,
		}
	}
	newTaskSeg := func(t TaskSegmentType, so, sd string) *TaskSegment {
		return &TaskSegment{
			Type:      t,
			Originals: []string{so},
			Display:   sd,
		}
	}

	if task.Completed {
		segs = append(segs, newBasicTaskSeg(SegmentIsCompleted, "x"))

		if task.HasCompletedDate() {
			segs = append(segs, newBasicTaskSeg(SegmentCompletedDate, task.CompletedDate.Format(DateLayout)))
		}
	}

	if task.HasPriority() && (!task.Completed || !RemoveCompletedPriority) {
		segs = append(segs, newTaskSeg(SegmentPriority, task.Priority, fmt.Sprintf("(%s)", task.Priority)))
	}

	if task.HasCreatedDate() {
		segs = append(segs, newBasicTaskSeg(SegmentCreatedDate, task.CreatedDate.Format(DateLayout)))
	}

	segs = append(segs, newBasicTaskSeg(SegmentTodoText, task.Todo))

	if task.HasContexts() {
		sort.Strings(task.Contexts)

		for _, context := range task.Contexts {
			segs = append(segs, newTaskSeg(SegmentContext, context, fmt.Sprintf("@%s", context)))
		}
	}

	if task.HasProjects() {
		sort.Strings(task.Projects)

		for _, project := range task.Projects {
			segs = append(segs, newTaskSeg(SegmentProject, project, fmt.Sprintf("+%s", project)))
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
			val := task.AdditionalTags[key]

			segs = append(segs, &TaskSegment{
				Type:      SegmentTag,
				Originals: []string{key, val},
				Display:   fmt.Sprintf("%s:%s", key, val),
			})
		}
	}

	if task.HasDueDate() {
		segs = append(segs, newBasicTaskSeg(SegmentDueDate, fmt.Sprintf("due:%s", task.DueDate.Format(DateLayout))))
	}

	return segs
}

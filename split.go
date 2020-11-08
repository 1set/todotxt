package todotxt

import (
	"fmt"
	"sort"
)

type TaskSegmentType int

const (
	TaskSegment_IsCompleted TaskSegmentType = iota + 1
	TaskSegment_CompletedDate
	TaskSegment_Priority
	TaskSegment_CreatedDate
	TaskSegment_Todo
	TaskSegment_Context
	TaskSegment_Project
	TaskSegment_Tag
	TaskSegment_DueDate
)

type TaskSegment struct {
	Type     TaskSegmentType
	Original string
	Display  string
}

func (task *Task) Split() []*TaskSegment {
	var segs []*TaskSegment
	newBasicTaskSeg := func(t TaskSegmentType, s string) *TaskSegment {
		return &TaskSegment{
			Type:     t,
			Original: s,
			Display:  s,
		}
	}
	newTaskSeg := func(t TaskSegmentType, so, sd string) *TaskSegment {
		return &TaskSegment{
			Type:     t,
			Original: so,
			Display:  sd,
		}
	}

	if task.Completed {
		segs = append(segs, newBasicTaskSeg(TaskSegment_IsCompleted, "x"))
		if task.HasCompletedDate() {
			segs = append(segs, newBasicTaskSeg(TaskSegment_CompletedDate, task.CompletedDate.Format(DateLayout)))
		}
	}

	if task.HasPriority() {
		segs = append(segs, newTaskSeg(TaskSegment_Priority, task.Priority, fmt.Sprintf("(%s)", task.Priority)))
	}

	if task.HasCreatedDate() {
		segs = append(segs, newBasicTaskSeg(TaskSegment_CreatedDate, task.CreatedDate.Format(DateLayout)))
	}

	segs = append(segs, newBasicTaskSeg(TaskSegment_Todo, task.Todo))

	if len(task.Contexts) > 0 {
		sort.Strings(task.Contexts)
		for _, context := range task.Contexts {
			segs = append(segs, newTaskSeg(TaskSegment_Context, context, fmt.Sprintf("@%s", context)))
		}
	}

	if len(task.Projects) > 0 {
		sort.Strings(task.Projects)
		for _, project := range task.Projects {
			segs = append(segs, newTaskSeg(TaskSegment_Project, project, fmt.Sprintf("+%s", project)))
		}
	}

	if len(task.AdditionalTags) > 0 {
		// Sort map alphabetically by keys
		keys := make([]string, 0, len(task.AdditionalTags))
		for key := range task.AdditionalTags {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			segs = append(segs, newBasicTaskSeg(TaskSegment_Tag, fmt.Sprintf("%s:%s", key, task.AdditionalTags[key])))
		}
	}

	if task.HasDueDate() {
		segs = append(segs, newBasicTaskSeg(TaskSegment_DueDate, fmt.Sprintf("due:%s", task.DueDate.Format(DateLayout))))
	}
	return segs
}

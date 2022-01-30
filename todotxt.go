package todotxt

import (
	"bufio"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"strings"

	ys "github.com/1set/gut/ystring"
)

// TaskList represents a list of todo.txt task entries.
// It is usually loaded from a whole todo.txt file.
type TaskList []Task

// IgnoreComments can be set to 'false', in order to revert to a more standard todo.txt behaviour.
// The todo.txt format does not define comments.
var (
	// IgnoreComments is used to switch ignoring of comments (lines starting with "#").
	// If this is set to 'false', then lines starting with "#" will be parsed as tasks.
	IgnoreComments = true

	// RemoveCompletedPriority is used to switch discarding priority on task completion like many todo.txt clients do.
	// If this is set to 'false', then the priority of completed task will be kept as it is.
	RemoveCompletedPriority = true
)

// NewTaskList creates a new empty TaskList.
func NewTaskList() TaskList {
	return TaskList{}
}

// String returns a complete list of tasks in todo.txt format.
func (tasklist TaskList) String() string {
	var sb strings.Builder
	for _, task := range tasklist {
		sb.WriteString(task.String())
		sb.WriteString(ys.NewLine)
	}
	return sb.String()
}

// AddTask appends a Task to the current TaskList and takes care to set the Task.ID correctly, modifying the Task by the given pointer!
func (tasklist *TaskList) AddTask(task *Task) {
	task.ID = 0
	for _, t := range *tasklist {
		if t.ID > task.ID {
			task.ID = t.ID
		}
	}
	task.ID++

	*tasklist = append(*tasklist, *task)
}

// GetTask returns a Task by given task 'id' from the TaskList. The returned Task pointer can be used to update the Task inside the TaskList.
// Returns an error if Task could not be found.
func (tasklist *TaskList) GetTask(id int) (*Task, error) {
	for i := range *tasklist {
		if ([]Task(*tasklist))[i].ID == id {
			return &([]Task(*tasklist))[i], nil
		}
	}
	return nil, errors.New("task not found")
}

// RemoveTaskByID removes any Task with given Task 'id' from the TaskList.
// Returns an error if no Task was removed.
func (tasklist *TaskList) RemoveTaskByID(id int) error {
	var newList TaskList

	found := false
	for _, t := range *tasklist {
		if t.ID != id {
			newList = append(newList, t)
		} else {
			found = true
		}
	}
	if !found {
		return errors.New("task not found")
	}

	*tasklist = newList
	return nil
}

// RemoveTask removes any Task from the TaskList with the same String representation as the given Task.
// Returns an error if no Task was removed.
func (tasklist *TaskList) RemoveTask(task Task) error {
	var newList TaskList

	found := false
	for _, t := range *tasklist {
		if t.String() != task.String() {
			newList = append(newList, t)
		} else {
			found = true
		}
	}
	if !found {
		return errors.New("task not found")
	}

	*tasklist = newList
	return nil
}

// LoadFrom loads a TaskList from an io.Reader
//
// os.File and os.Stdin are interfaces that include io.Reader, which is also bufio.Scanner's input type
//
// Note: This will clear the current TaskList and replace its contents with the reader's output.
func (tasklist *TaskList) LoadFrom(tasks io.Reader) error {
	*tasklist = []Task{} // Empty task list

	taskID := 1
	scanner := bufio.NewScanner(tasks)

    // scan in each line
	for scanner.Scan() {
		text := strings.Trim(scanner.Text(), whitespaces) // Read line
		task, err := ParseTask(text)

		if err != nil {
            return err
		} else if task == nil {
            // if task pointer is empty do not create task
            continue 
        }

		task.ID = taskID

		*tasklist = append(*tasklist, *task)
		taskID++
	}

	return scanner.Err()
}

// LoadFromFile loads a TaskList from *os.File using LoadFrom
func (tasklist *TaskList) LoadFromFile(file *os.File) error {
	// passthrough method to generic form
	return tasklist.LoadFrom(file)
}

// LoadFromPath loads a TaskList from a file (most likely called "todo.txt").
func (tasklist *TaskList) LoadFromPath(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return tasklist.LoadFrom(file)
}

// WriteTo
//
// Using *os.File instead of a filename allows to also use os.Stdout.
//
// Note: Comments from original file will be omitted and not written to target *os.File, if IgnoreComments is set to 'true'.

// WriteToFile writes a TaskList to *os.File.
//
func (tasklist *TaskList) WriteToFile(file *os.File) error {
	writer := bufio.NewWriter(file)
	if _, err := writer.WriteString(tasklist.String()); err != nil {
		return err
	}
	return writer.Flush()
}

// WriteToPath writes a TaskList to the specified file (most likely called "todo.txt").
func (tasklist *TaskList) WriteToPath(filename string) error {
	return ioutil.WriteFile(filename, []byte(tasklist.String()), 0640)
}

// LoadFrom creates and returns a new TaskList
func LoadFrom(todos io.Reader) (TaskList, error) {
	newTaskList := TaskList{}
	if err := newTaskList.LoadFrom(todos); err != nil {
		return nil, err
	}
	return tasklist, nil
}

// LoadFromFile creates a new TaskList and loads it from an *os.File using TaskList.LoadFrom
func LoadFromFile(file *os.File) (TaskList, error) {
	// passthrough to generic function
	return LoadFrom(file)
}

// WriteToFile writes a TaskList to *os.File.
//
// Using *os.File instead of a filename allows to also use os.Stdout.
//
// Note: Comments from original file will be omitted and not written to target *os.File, if IgnoreComments is set to 'true'.
func WriteToFile(tasklist *TaskList, file *os.File) error {
	return tasklist.WriteToFile(file)
}

// LoadFromPath loads and returns a TaskList from a file (most likely called "todo.txt").
func LoadFromPath(filename string) (TaskList, error) {
	tasklist := TaskList{}
	if err := tasklist.LoadFromPath(filename); err != nil {
		return nil, err
	}
	return tasklist, nil
}

// WriteToPath writes a TaskList to the specified file (most likely called "todo.txt").
func WriteToPath(tasklist *TaskList, filename string) error {
	return tasklist.WriteToPath(filename)
}

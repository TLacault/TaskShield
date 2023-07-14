package task

import (
	comment "github.com/TLacault/TaskShield/server/project/comment"
	label "github.com/TLacault/TaskShield/server/project/label"
	person "github.com/TLacault/TaskShield/server/project/person"
)

/* ********************************************************************************************* */

type Status string

const (
	TODO       Status = "ToDo"
	INPROGRESS Status = "InProgress"
	DONE       Status = "Done"
)

type Task struct {
	title       string
	description string
	status      Status
	assignees   []person.Person
	comments    []comment.Comment
	labels      []label.Label
}

/* ********************************************************************************************* */

func Default() Task {
	task := Task{
		title:       "Default Title",
		description: "Default Description",
		status:      TODO,
		assignees:   make([]person.Person, 0),
		comments:    make([]comment.Comment, 0),
		labels:      make([]label.Label, 0),
	}
	return task
}

func New(title string, description string, status Status) Task {
	task := Default()
	task.title = title
	task.description = description
	task.status = status
	return task
}

/* ********************************************************************************************* */

func (t *Task) GetTitle() string               { return t.title }
func (t *Task) GetDescription() string         { return t.description }
func (t *Task) GetStatus() Status              { return t.status }
func (t *Task) GetAssignees() []person.Person  { return t.assignees }
func (t *Task) GetComments() []comment.Comment { return t.comments }
func (t *Task) GetLabels() []label.Label       { return t.labels }

/* ********************************************************************************************* */

func (t *Task) AddAssignee(p person.Person) {
	t.assignees = append(t.assignees, p)
}

// func (t *Task) AddComment(c comment.Comment) {
// 	t.comments = append(t.comments, c)
// }

func (t *Task) AddComment(c string, p person.Person) {
	comment := comment.New(c, p)
	t.comments = append(t.comments, comment)
}

// func (t *Task) AddLabel(l label.Label) {
// 	t.labels = append(t.labels, l)
// }

func (t *Task) AddLabel(l string, p label.Priority) {
	label := label.New(l, p)
	t.labels = append(t.labels, label)
}

/* ********************************************************************************************* */

func (t *Task) ToString() string {
	var str string
	str += t.title + ";" + t.description + ";" + string(t.status) + ";"
	str += "{"
	for _, a := range t.assignees {
		str += "[" + a.ToString() + "]"
	}
	str += "}{"
	for _, c := range t.comments {
		str += "[" + c.ToString() + "]"
	}
	str += "}{"
	for _, l := range t.labels {
		str += "[" + l.ToString() + "]"
	}
	str += "}"
	return str
}

func (t *Task) ToStringFormat() string {
	var str string
	str += "[Title]: " + t.title + "\n"
	str += "[Description]: " + t.description + "\n"
	str += "[Status]: " + string(t.status) + "\n"
	if len(t.assignees) > 0 {
		str += "[Assignees]:\n"
		for _, a := range t.assignees {
			str += "\t" + a.ToStringFormat() + "\n"
		}
	}
	if len(t.comments) > 0 {
		str += "[Comments]:\n"
		for _, c := range t.comments {
			str += "\t" + c.ToStringFormat() + "\n"
		}
	}
	if len(t.labels) > 0 {
		str += "[Labels]:\n"
		for _, l := range t.labels {
			str += "\t" + l.ToStringFormat() + "\n"
		}
	}
	return str
}

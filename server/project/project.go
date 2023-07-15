package project

import (
	"errors"
	"strings"

	// comment "github.com/TLacault/TaskShield/server/project/comment"
	person "github.com/TLacault/TaskShield/server/project/person"
	task "github.com/TLacault/TaskShield/server/project/task"
)

type Project struct {
	name        string
	description string
	tasks       []task.Task
	members     []person.Person
}

/* ********************************************************************************************* */

func New(name string, description string) Project {
	project := Project{
		name:        name,
		description: description,
		tasks:       make([]task.Task, 0),
		members:     make([]person.Person, 0),
	}
	return project
}

/* ********************************************************************************************* */

func (p *Project) GetName() string        { return p.name }
func (p *Project) GetDescription() string { return p.description }
func (p *Project) GetTasks() []task.Task  { return p.tasks }

func (p *Project) GetMembers() []person.Person { return p.members }

func (p *Project) GetMemberByName(name string) (person.Person, error) {
	// check if name is in format "firstName lastName"
	if len(strings.Split(name, " ")) != 2 {
		return person.Person{}, errors.New("Invalid name format")
	}
	firstName := strings.Split(name, " ")[0]
	lastName := strings.Split(name, " ")[1]

	for _, m := range p.members {
		if m.GetFirstName() == firstName && m.GetLastName() == lastName {
			return m, nil
		}
	}
	return person.Person{}, errors.New("No match for " + name + " in project " + p.name)
}

/* ********************************************************************************************* */

/* ********************************************************************************************* */

func (p *Project) ToString() string {
	var str string
	str += p.name + ";" + p.description + ";"
	str += "\n{"

	for _, m := range p.members {
		str += "[" + m.ToString() + "]"
	}

	str += "}\n{"

	for _, t := range p.tasks {
		str += "[" + t.ToString() + "]"
	}

	str += "}"
	return str
}

func (p *Project) ToStringFormat() string {
	var str string
	str += "====[" + p.name + "]====\n"
	str += p.description + "\n"

	str += "\n==[MEMBERS]==\n"
	for _, m := range p.members {
		str += m.ToStringFormat() + "\n"
	}

	str += "\n==[TASKS]==\n"
	for _, t := range p.tasks {
		str += t.ToStringFormat() + "\n"
	}
	return str
}

/* ********************************************************************************************* */

func (p *Project) AddTask(t task.Task) {
	p.tasks = append(p.tasks, t)
}

func (p *Project) AddMember(m person.Person) {
	p.members = append(p.members, m)
}

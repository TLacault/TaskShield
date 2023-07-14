package project

import (
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

func (p *Project) GetName() string             { return p.name }
func (p *Project) GetDescription() string      { return p.description }
func (p *Project) GetTasks() []task.Task       { return p.tasks }
func (p *Project) GetMembers() []person.Person { return p.members }

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

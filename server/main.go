package main

import (
	"fmt"

	project "github.com/TLacault/TaskShield/server/project"
	"github.com/TLacault/TaskShield/server/project/label"
	person "github.com/TLacault/TaskShield/server/project/person"
	task "github.com/TLacault/TaskShield/server/project/task"
)

func main() {
	project := project.New("TaskShield", "Task management & Secure messaging application")

	// add tasks
	project.AddTask(task.New(
		"Define requirements",
		"Define primary requirements to meet before deadline",
		task.INPROGRESS,
	))
	project.AddTask(task.New(
		"Design architecture",
		"Design architecture of the application",
		task.TODO,
	))

	// add members
	project.AddMember(person.New("Tim", "Lacault", person.HOST))
	project.AddMember(person.New("John", "Doe", person.COWORKER_READ))
	project.AddMember(person.New("Some", "One", person.COWORKER_READ_WRITE))

	// modify task
	project.GetTasks()[0].AddAssignee(project.GetMembers()[0])
	project.GetTasks()[0].AddAssignee(project.GetMembers()[1])
	project.GetTasks()[0].AddAssignee(project.GetMembers()[2])

	project.GetTasks()[0].AddComment("This is a comment", project.GetMembers()[0])
	project.GetTasks()[0].AddComment("This is another comment", project.GetMembers()[1])

	project.GetTasks()[0].AddLabel("Documentation", label.ASAP)

	// print project
	fmt.Print(project.ToStringFormat())
	fmt.Print("----------------------------------------------------\n\n")
	fmt.Println(project.ToString())
}

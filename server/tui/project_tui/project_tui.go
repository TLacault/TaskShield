package project_tui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	project "github.com/TLacault/TaskShield/server/project"
	member_tui "github.com/TLacault/TaskShield/server/tui/member_tui"
	utils "github.com/TLacault/TaskShield/server/tui/utils"
)

func ProjectHeader(p *project.Project, index ...int) {
	if len(index) == 0 {
		if len(p.GetDescription()) > 0 {
			fmt.Print("╭", strings.Repeat("─", len(p.GetName())+2), "┬", strings.Repeat("─", len(p.GetDescription())+2), "╮\n")
			fmt.Print("│ ", p.GetName(), " │ ", p.GetDescription(), " │\n")
			fmt.Print("╰", strings.Repeat("─", len(p.GetName())+2), "┴", strings.Repeat("─", len(p.GetDescription())+2), "╯\n")
		} else {
			fmt.Print("╭", strings.Repeat("─", len(p.GetName())+2), "╮\n")
			fmt.Print("│ ", p.GetName(), " │\n")
			fmt.Print("╰", strings.Repeat("─", len(p.GetName())+2), "╯\n")
		}
	} else {
		if len(p.GetDescription()) > 0 {
			fmt.Print("╭", strings.Repeat("─", len(strconv.Itoa(index[0]))+2), "╮╭", strings.Repeat("─", len(p.GetName())+2), "┬", strings.Repeat("─", len(p.GetDescription())+2), "╮\n")
			fmt.Print("│ ", strconv.Itoa(index[0]), " ││ ", p.GetName(), " │ ", p.GetDescription(), " │\n")
			fmt.Print("╰", strings.Repeat("─", len(strconv.Itoa(index[0]))+2), "╯╰", strings.Repeat("─", len(p.GetName())+2), "┴", strings.Repeat("─", len(p.GetDescription())+2), "╯\n")
		} else {
			fmt.Print("╭", strings.Repeat("─", len(strconv.Itoa(index[0]))+2), "╮╭", strings.Repeat("─", len(p.GetName())+2), "╮\n")
			fmt.Print("│ ", strconv.Itoa(index[0]), " ││ ", p.GetName(), " │\n")
			fmt.Print("╰", strings.Repeat("─", len(strconv.Itoa(index[0]))+2), "╯╰", strings.Repeat("─", len(p.GetName())+2), "╯\n")
		}
	}
}

func ListProjects(p *[]project.Project) {
	println()
	for i, project := range *p {
		ProjectHeader(&project, i+1)
	}
}

func NewProject(p *[]project.Project) {
	println()
	println("╭─────────────╮ ╭───┬────────╮")
	println("│ New Project │ │ 0 │ Cancel │")
	println("╰─────────────╯ ╰───┴────────╯\n")

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Chose a name for your project:\n> ")
	scanner.Scan()
	name := scanner.Text()
	if name == "0" {
		fmt.Print("\n[project creation was cancelled]\n")
		return
	}

	fmt.Print("\nWrite a description for your project (optional):\n> ")
	scanner.Scan()
	description := scanner.Text()
	if description == "0" {
		fmt.Print("\n[project creation was cancelled]\n")
		return
	}

	*p = append(*p, project.New(name, description))
	fmt.Print("\nYour project has been created:\n")
	ProjectHeader(&(*p)[len(*p)-1])
}

func SelectProject(p *[]project.Project) {
	for {
		if len(*p) == 0 {
			fmt.Println("\n[no projects have been created]")
			utils.CLEAR_SCREEN = false
			return
		}
		utils.ClearScreen()

		println("╭──────────────────╮ ╭───┬──────╮")
		println("│ Select a Project │ │ 0 │ Back │")
		println("╰──────────────────╯ ╰───┴──────╯")

		ListProjects(p)

		fmt.Print("\n> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			// fmt.Println("[error]: ", err)
			SelectProject(p)
			return
		}

		number, err := strconv.Atoi(line)
		if err != nil {
			// fmt.Println("[error]: ", err)
			SelectProject(p)
			return
		}

		if number == 0 {
			return
		}

		if number < 0 || number > len(*p) {
			// fmt.Println("[error]: Invalid project number")
			SelectProject(p)
			return
		}
		ProjectView(&(*p)[number-1])
	}
}

func ProjectOverview(p *project.Project) {
	if len(p.GetMembers()) != 0 {
		println("\nMEMBERS:")
		member_tui.ListMembers(p.GetMembers())
	}
	if len(p.GetTasks()) != 0 {
		println("\nTASKS:")
		// ListTasks(p.GetTasks())
	}
}

func ProjectView(p *project.Project) {
	for {
		utils.ClearScreen()
		if utils.CLEAR_SCREEN {
			ProjectHeader(p)
			println("╭───┬──────────╮ ╭───┬─────────╮ ╭───┬───────╮ ╭───┬──────╮")
			println("│ 1 │ Overview │ │ 2 │ Members │ │ 3 │ Tasks │ │ 4 │ Back │")
			println("╰───┴──────────╯ ╰───┴─────────╯ ╰───┴───────╯ ╰───┴──────╯")
		}

		fmt.Print("\n> ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "1":
			ProjectOverview(p)
			utils.CLEAR_SCREEN = false
			break
		case "2":
			utils.CLEAR_SCREEN = true
			member_tui.SelectMember(p)
			break
		// case "3":
		// 	EditTasks(p)
		case "4":
			utils.CLEAR_SCREEN = true
			return
		default:
			continue
		}
	}
}

package member_tui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	project "github.com/TLacault/TaskShield/server/project"
	person "github.com/TLacault/TaskShield/server/project/person"
	utils "github.com/TLacault/TaskShield/server/tui/utils"
)

func MemberHeader(p *person.Person, index ...int) {
	if len(index) == 0 {
		fmt.Print("╭", strings.Repeat("─", len(p.GetRole())+2), "┬", strings.Repeat("─", len(p.GetFirstName())+2+len(p.GetLastName())+1), "╮\n")
		fmt.Print("│ ", p.GetRole(), " │ ", p.GetFirstName(), " ", p.GetLastName(), " │\n")
		fmt.Print("╰", strings.Repeat("─", len(p.GetRole())+2), "┴", strings.Repeat("─", len(p.GetFirstName())+2+len(p.GetLastName())+1), "╯\n")
	} else {
		fmt.Print("╭", strings.Repeat("─", len(strconv.Itoa(index[0]))+2), "╮╭", strings.Repeat("─", len(p.GetRole())+2), "┬", strings.Repeat("─", len(p.GetFirstName())+2+len(p.GetLastName())+1), "╮\n")
		fmt.Print("│ ", strconv.Itoa(index[0]), " ││ ", p.GetRole(), " │ ", p.GetFirstName(), " ", p.GetLastName(), " │\n")
		fmt.Print("╰", strings.Repeat("─", len(strconv.Itoa(index[0]))+2), "╯╰", strings.Repeat("─", len(p.GetRole())+2), "┴", strings.Repeat("─", len(p.GetFirstName())+2+len(p.GetLastName())+1), "╯\n")
	}
}

func ListMembers(p []person.Person) {
	println()
	for i, person := range p {
		MemberHeader(&person, i+1)
	}
}

func NewMember(p *project.Project) {
	utils.ClearScreen()
	println("╭────────────╮ ╭───┬────────╮")
	println("│ New Member │ │ 0 │ Cancel │")
	println("╰────────────╯ ╰───┴────────╯\n")

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("First name:\n> ")
	scanner.Scan()
	firstName := scanner.Text()
	if firstName == "0" {
		fmt.Print("\n[member creation was cancelled]\n")
		return
	}

	fmt.Print("\nLast name:\n> ")
	scanner.Scan()
	lastName := scanner.Text()
	if lastName == "0" {
		fmt.Print("\n[member creation was cancelled]\n")
		return
	}

	fmt.Print("\nChose a role:\n")
	println("╭───┬──────╮ ╭───┬──────────────────────╮ ╭────┬────────────────────────────╮")
	println("│ H │ Host │ │ R │ Coworker Read Access │ │ RW │ Coworker Read+Write Access │")
	println("╰───┴──────╯ ╰───┴──────────────────────╯ ╰────┴────────────────────────────╯")
	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		switch input {
		case "H", "h":
			p.AddMember(person.New(firstName, lastName, person.HOST))
			return
		case "R", "r":
			p.AddMember(person.New(firstName, lastName, person.COWORKER_READ))
			return
		case "RW", "rw":
			p.AddMember(person.New(firstName, lastName, person.COWORKER_READ_WRITE))
			return
		case "0":
			return
		default:
			println()
			continue
		}
	}
}

func SelectMember(p *project.Project) {
	for {
		utils.ClearScreen()
		println("╭─────────────────╮ ╭───┬──────╮")
		println("│ Select a Member │ │ 0 │ Back │")
		println("╰─────────────────╯ ╰───┴──────╯")

		if len(p.GetMembers()) != 0 {
			ListMembers(p.GetMembers())
		}

		println()
		println("╭───┬────────────────╮")
		println("│ + │ Add New Member │")
		println("╰───┴────────────────╯")

		fmt.Print("\n> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		if input == "0" {
			return
		}

		if input == "+" {
			NewMember(p)
		}
		member, err := strconv.Atoi(input)
		if err != nil || member < 1 || member > len(p.GetMembers()) {
			continue
		}
	}
}

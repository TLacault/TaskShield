package tui

import (
	"bufio"
	"fmt"
	"os"

	project "github.com/TLacault/TaskShield/server/project"
	project_tui "github.com/TLacault/TaskShield/server/tui/project_tui"
	utils "github.com/TLacault/TaskShield/server/tui/utils"
)

func Start() {
	// DATA
	projects := make([]project.Project, 0)

	// CLI
	scanner := bufio.NewScanner(os.Stdin)

	for {
		utils.ClearScreen()
		if utils.CLEAR_SCREEN {
			println("╭───────────────────────┬───────────────────────────────────────────╮")
			println("│ Welcome to TaskShield │ Secure Project Management & Messaging App │")
			println("╰───────────────────────┴───────────────────────────────────────────╯")
			println("╭───┬─────────────╮ ╭───┬────────────────╮")
			println("│ 1 │ New Project │ │ 2 │ Select Project │")
			println("╰───┴─────────────╯ ╰───┴────────────────╯")

		}
		fmt.Print("\n> ")

		scanner.Scan()
		input := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Println("[error]: ", err)
			break
		}

		switch input {
		case "1":
			project_tui.NewProject(&projects)
			utils.CLEAR_SCREEN = false
		case "2":
			utils.CLEAR_SCREEN = true
			project_tui.SelectProject(&projects)
		case "quit":
			break
		default:
			continue
		}

	}
}

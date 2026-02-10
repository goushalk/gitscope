package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/goushalk/gitscope/internal/api"
	"github.com/goushalk/gitscope/internal/logic"
	"github.com/goushalk/gitscope/internal/ui"
)

func main() {
	username := flag.String("user", "", "GitHub username")
	cli := flag.Bool("cli", false, "print plain CLI output (no TUI)")
	// git := 
	flag.Parse()

	if *username == "" {
		fmt.Println("provide username")
		os.Exit(1)
	}
	
	
	logic.Banner(*username)
	events, err := api.UserBasedActivity(*username)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if *cli{
		fmt.Println()
		logic.Cli(events)
		return
	} 

	var rows []table.Row
	for _, e := range events {
		rows = append(rows, table.Row{
			logic.EventAction(e.Type),
			e.Repo.Name,
			e.CreatedAt,
		})
	}

	t := ui.NewTable(rows)
	m := ui.Model{Table: t}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

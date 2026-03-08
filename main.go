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
	// Input flags.
	username := flag.String("user", "", "GitHub username")
	cli := flag.Bool("cli", false, "print plain CLI output (no TUI)")
	flag.Usage = func() {
		flag.PrintDefaults()
	}
	flag.Parse()

	if *username == "" {
		flag.Usage()
		os.Exit(1)
	}

	// Fetch recent public events for the target user.
	events, err := api.UserBasedActivity(*username)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(events) == 0 {
		fmt.Printf("No username: %s exist", *username)
		os.Exit(1)
	}

	// Non-interactive output mode.
	if *cli {
		fmt.Println(logic.Banner(*username))
		fmt.Println()
		logic.Cli(events)
		return
	}

	// Interactive TUI mode: map events into table rows.
	var rows []table.Row
	for _, e := range events {
		rows = append(rows, table.Row{
			logic.EventAction(e.Type),
			e.Repo.Name,
			e.CreatedAt,
		})
	}

	t := ui.NewTable(rows)
	m := ui.NewModel(t, *username)

	if _, err := tea.NewProgram(m, tea.WithAltScreen(), tea.WithMouseAllMotion()).Run(); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

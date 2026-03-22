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

	//DOC:: Input flags.
	//TODO: : Add flag for authenticated users (--auth) .

	username := flag.String("user", "", "GitHub username")
	cli := flag.Bool("cli", false, "print plain CLI output (no TUI)")
	IsAuth := flag.Bool("auth", false, "get info as authenticated user")
	Json := flag.Bool("json", false, "stdout as json format")

	//NOTE: : unauthenticated is the default fetch

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "A tool to fetch GitHub user info.\n")

		fmt.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults()

		fmt.Fprintln(os.Stderr, "\nExamples:")
		fmt.Fprintln(os.Stderr, "  gitscope -user=torvalds")
		fmt.Fprintln(os.Stderr, "  gitscope -user=torvalds -cli for cli version")
	}

	flag.Parse()

	if *IsAuth {

	}

	if *username == "" {
		flag.Usage()
		os.Exit(1)
	}

	//DOC: : Fetch recent public events for the target user.

	events, err := api.UserBasedActivity(*username)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(events) == 0 {
		fmt.Printf("No username: %s exist", *username)
		os.Exit(1)
	}

	//DONE: : json output

	if *Json {
		jdata, err := logic.JsonOutput(events)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(jdata)
		return

	}

	//DOC: : Non-interactive output mode.

	if *cli {
		fmt.Println(logic.Banner(*username))
		fmt.Println()
		logic.Cli(events)
		return
	}

	//DOC: : Interactive TUI mode: map events into table rows.

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

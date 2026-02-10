package logic

import(
	"github.com/goushalk/gitscope/internal/api"
	"fmt"
)

func Cli (events []api.GitHubEvent) {
	for _, e := range events{
		action := EventAction(e.Type)
		fmt.Printf(
		"%-15s %-20s %s\n",
			action,
			e.Repo.Name,
			e.CreatedAt,		)
	}
}



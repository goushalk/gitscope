package logic

import (
	"fmt"
	"strings"
	"time"

	"github.com/goushalk/gitscope/internal/api"
)

// Cli prints a readable, line-oriented summary of GitHub events.
func Cli(events []api.GitHubEvent) {
	for _, e := range events {
		switch e.Type {

		case "PushEvent":
			fmt.Printf(
				"[Push]     %s → %s (%s) at %s\n",
				e.Repo.Name,
				extractBranch(e.Payload.Ref),
				shortSHA(e.Payload.Head),
				formatTime(e.CreatedAt),
			)
			fmt.Println()
		case "CreateEvent":
			handleCreate(e)

			fmt.Println()
		case "DeleteEvent":
			fmt.Printf(
				"[Delete]   %s %s in %s at %s\n",
				e.Payload.RefType,
				e.Payload.Ref,
				e.Repo.Name,
				formatTime(e.CreatedAt),
			)

			fmt.Println()
		case "WatchEvent":
			fmt.Printf(
				"[Star]     %s at %s\n",
				e.Repo.Name,
				formatTime(e.CreatedAt),
			)

			fmt.Println()
		case "ForkEvent":
			fmt.Printf(
				"[Fork]     %s at %s\n",
				e.Repo.Name,
				formatTime(e.CreatedAt),
			)

			fmt.Println()
		case "IssuesEvent":
			fmt.Printf(
				"[Issue]    %s at %s\n",
				e.Repo.Name,
				formatTime(e.CreatedAt),
			)

			fmt.Println()
		case "PullRequestEvent":
			fmt.Printf(
				"[PR]       %s at %s\n",
				e.Repo.Name,
				formatTime(e.CreatedAt),
			)

			fmt.Println()
		default:
			fmt.Printf(
				"[Other]    %s in %s at %s\n",
				e.Type,
				e.Repo.Name,
				formatTime(e.CreatedAt),
			)
			fmt.Println()
		}
	}
}

// handleCreate renders CreateEvent variants (branch, tag, repository).
func handleCreate(e api.GitHubEvent) {
	switch e.Payload.RefType {

	case "branch":
		fmt.Printf(
			"[Create]   Branch %s in %s at %s\n",
			e.Payload.Ref,
			e.Repo.Name,
			formatTime(e.CreatedAt),
		)

	case "repository":
		fmt.Printf(
			"[Create]   Repository %s at %s\n",
			e.Repo.Name,
			formatTime(e.CreatedAt),
		)

	case "tag":
		fmt.Printf(
			"[Create]   Tag %s in %s at %s\n",
			e.Payload.Ref,
			e.Repo.Name,
			formatTime(e.CreatedAt),
		)

	default:
		fmt.Printf(
			"[Create]   %s in %s at %s\n",
			e.Payload.RefType,
			e.Repo.Name,
			formatTime(e.CreatedAt),
		)
	}
}

// shortSHA returns a 7-char abbreviated SHA when possible.
func shortSHA(sha string) string {
	if len(sha) >= 7 {
		return sha[:7]
	}
	return sha
}

// extractBranch returns the trailing branch segment from a Git ref.
func extractBranch(ref string) string {
	// refs/heads/master → master
	parts := strings.Split(ref, "/")
	return parts[len(parts)-1]
}

// formatTime converts RFC3339 input to a compact local-readable format.
func formatTime(t string) string {
	parsed, err := time.Parse(time.RFC3339, t)
	if err != nil {
		return t
	}
	return parsed.Format("02 Jan 2006 15:04")
}

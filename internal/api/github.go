package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type GitHubEvent struct {
	ID        string  `json:"id"`
	Type      string  `json:"type"`
	Repo      Repo    `json:"repo"`
	Payload   Payload `json:"payload"`
	CreatedAt string  `json:"created_at"`
}

type Repo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Payload struct {
	Action   string   `json:"action,omitempty"`
	Ref      string   `json:"ref,omitempty"`       // IMPORTANT
	RefType  string   `json:"ref_type,omitempty"`  // branch / tag / repository
	Head     string   `json:"head,omitempty"`      // commit SHA for PushEvent
	Commits  []Commit `json:"commits,omitempty"`
}

type Commit struct {
	SHA     string `json:"sha,omitempty"`
	Message string `json:"message,omitempty"`
}
func UserBasedActivity(username string) ([]GitHubEvent, error) {
	url := fmt.Sprintf(
		"https://api.github.com/users/%s/events/public",
		username,
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "github-activity-cli")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	var events []GitHubEvent
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return nil, err
	}

	return events, nil
}


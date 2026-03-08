# gitscope

`gitscope` is a terminal tool for viewing a GitHub user’s recent public activity.

It supports two output modes:
- Interactive TUI table (default)
- Plain CLI text output (`-cli`)

## Features

- Fetches recent public events from `GET /users/{username}/events/public`
- Maps GitHub event types to readable action labels
- Renders event list in an interactive terminal table (Bubble Tea)
- Provides a non-interactive CLI mode for quick scripting/inspection

## Requirements

- Go `1.25.6+` (as declared in [`go.mod`](/home/jellybean/github/personal/gitscope/go.mod))
- Internet access to `api.github.com`

## Installation

### Build locally

```bash
go build -o gitscope .
```

### Install with Go

```bash
go install github.com/goushalk/gitscope@latest
```

If needed, ensure `$(go env GOPATH)/bin` is on your `PATH`.

## Usage

```bash
gitscope -user <github-username>
```

### Flags

- `-user string` GitHub username (required)
- `-cli` Print plain CLI output instead of TUI

### Examples

```bash
# Interactive table
gitscope -user torvalds

# Plain terminal output
gitscope -user torvalds -cli
```

## Output Modes

### TUI mode (default)

- Shows ASCII banner + table with columns: `Event`, `Repo`, `Time`
- Table is keyboard-navigable via Bubble Tea
- Quit keys: `q` or `ctrl+c`

### CLI mode (`-cli`)

- Prints a banner followed by one line per event
- Formats common event types (`PushEvent`, `CreateEvent`, `DeleteEvent`, etc.)
- Shows formatted timestamps when possible (`02 Jan 2006 15:04`)

## Behavior and Limitations

- Only public activity is shown (from GitHub public events API).
- If `-user` is missing, usage is printed and process exits with status `1`.
- If request fails or GitHub returns non-`200`, an error is printed and process exits with status `1`.
- Current API requests are unauthenticated, so stricter rate limits may apply.
- GitHub event API is recent-activity oriented (not full history).

## Project Structure

- [`main.go`](/home/jellybean/github/personal/gitscope/main.go): CLI flags, mode switch, app entrypoint
- [`internal/api/github.go`](/home/jellybean/github/personal/gitscope/internal/api/github.go): GitHub API fetch + event models
- [`internal/logic/events.go`](/home/jellybean/github/personal/gitscope/internal/logic/events.go): event type to action label mapping
- [`internal/logic/cliMode.go`](/home/jellybean/github/personal/gitscope/internal/logic/cliMode.go): CLI rendering logic
- [`internal/logic/banner.go`](/home/jellybean/github/personal/gitscope/internal/logic/banner.go): banner text
- [`internal/ui/table.go`](/home/jellybean/github/personal/gitscope/internal/ui/table.go): table styling/configuration
- [`internal/ui/model.go`](/home/jellybean/github/personal/gitscope/internal/ui/model.go): Bubble Tea model/update/view

## Development

Run directly:

```bash
go run . -user <github-username>
```

Format and vet:

```bash
gofmt -w .
go vet ./...
```

Run tests (if/when test files are added):

```bash
go test ./...
```

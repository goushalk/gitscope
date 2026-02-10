# gitscope

Terminal UI for browsing a GitHub user's recent activity. Supports a TUI table view or plain CLI output.

## Requirements

- Go 1.25+

## Install

### Build from source

```bash
go build -o gitscope .
```

### Install with Go

```bash
go install github.com/goushalk/gitscope@latest
```

After installing, ensure your Go bin directory is on your `PATH`:

```bash
# macOS/Linux
export PATH="$PATH:$(go env GOPATH)/bin"
```

## Usage

```bash
gitscope -user <github-username>
```

### Options

- `-user` GitHub username (required)
- `-cli` Print plain CLI output (no TUI)

### Examples

```bash
gitscope -user torvalds
gitscope -user torvalds -cli
```

## Notes

- If the username has no recent activity, the app exits with a message.
- TUI mode uses arrow keys to navigate the table and `q`/`ctrl+c` to quit.

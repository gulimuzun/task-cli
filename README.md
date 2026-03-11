# task-cli

A small Go CLI for tracking tasks in a local JSON file.

## Features

- Add tasks with an auto-incrementing ID
- Update task description
- Mark task status: `todo`, `in-progress`, `done`
- Delete tasks
- List tasks (filtered by status)
- Zero dependencies (standard library only)

## How it works

Tasks are stored in a file named `task.json` **in your current working directory**.

- If `task.json` is missing, it will be created automatically by `add`, `update`, `mark`, and `delete`.
- `list` requires `task.json` to already exist.

Each task record looks like this:

```json
[
  {
    "id": 1,
    "description": "Buy milk",
    "status": "todo",
    "createdAt": "2026-03-11 11:22:14",
    "updatedAt": "2026-03-11 11:22:14"
  }
]
```

## Requirements

- Go (the project `go.mod` targets Go `1.25.4`)

## Install / Build

Clone the repo, then from the project root.

### Option A: Build a local binary

```bash
go build -o task-cli ./cmd
```

This creates an executable named `task-cli` in the current directory. Run it as:

```bash
./task-cli help
```

### Option B: Install to your PATH (recommended)

Install the binary into your Go bin directory (`GOBIN`, or `$(go env GOPATH)/bin` if `GOBIN` is not set):

```bash
go install ./cmd
```

Make sure the Go bin directory is in your `PATH`, for example in `~/.zshrc`:

```bash
export PATH="$(go env GOPATH)/bin:$PATH"
```

After that, you can run the command anywhere (no `./`):

```bash
cmd help
```

> Note: because the entrypoint is currently in the `cmd/` directory, `go install ./cmd` produces a binary named `cmd`.
> If you prefer the command name to be `task-cli`, either use Option A, or move the entrypoint to `cmd/task-cli/` later.

### Run without building

```bash
go run ./cmd -- help
```

## Usage

```text
task-cli <command> [arguments]
```

### Commands

#### `help`

Show help.

```bash
task-cli help
```

#### `add <description> [<todo|in-progress|done>]`

Add a new task. If status is omitted, it defaults to `todo`.

```bash
task-cli add "Buy milk"
task-cli add "Write report" in-progress
```

#### `update <id> <text>`

Update a task's description.

```bash
task-cli update 1 "Buy oat milk"
```

#### `mark <id> <todo|in-progress|done>`

Update a task's status.

```bash
task-cli mark 1 done
```

#### `delete <id>`

Delete a task by ID.

```bash
task-cli delete 1
```

#### `list [todo|in-progress|done|all]`

List tasks filtered by status. Default is `all`.

```bash
task-cli list
task-cli list todo
task-cli list in-progress
task-cli list done
task-cli list all
```

## Output format

For each task, the CLI prints:

- `<description> <status>`
- `ID: <id>`
- `Updated at:<timestamp>`
- `Created at:<timestamp>`

## Project structure

```text
cmd/
  main.go                # CLI entrypoint and argument parsing
internal/
  help/                  # help text output
  list/                  # list command implementation
  status/                # task storage (add/update/delete) + Task model
  errorPrint/            # error printing + exit(1)
```

## License

See `LICENSE`.

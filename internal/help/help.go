package help

import "fmt"

// Help prints task-cli usage information to stdout.
//
// This project currently stores tasks in a local JSON file named task.json (in the current
// working directory). Task statuses are: todo, in-progress, and done.
func Help() {
	fmt.Print(`task-cli - a simple task tracker

USAGE
  task-cli <command> [arguments]

COMMANDS
  add <description> [<todo|in-progress|done>]
      Add a new task. ID is auto-incremented and timestamps are set automatically.

  mark <id> <todo|in-progress|done>
      Update a task's status.

  update <id> <text>
      Update a task's description.

  delete <id>
      Delete a task by ID.

  list [todo|in-progress|done|all]
      List tasks filtered by status (default: all).

  help
      Show this help message.

NOTES
  - Data file: task.json (created automatically if missing)
  - Time format: 2006-01-02 15:04:05
`)
}
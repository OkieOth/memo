# TL;DR
Memo is a tool to manage on-the-fly memos from the terminal. It's still heavy WIP. 
Right now only the add command is implemented

This repository contains a spare-time project.

# Usage
```bash
go run main.go help
The tool supports the creation, search and list of on-the-fly
textual memos from the terminal.

Usage:
  memo [flags]
  memo [command]

Available Commands:
  add         Adds a new memo entry
  completion  Generate the autocompletion script for the specified shell
  del         Delete a memo from one our all storage targets
  done        A brief description of your command
  find        A brief description of your command
  help        Help about any command
  list        A brief description of your command

Flags:
  -h, --help     help for memo
  -t, --toggle   Help message for toggle

Use "memo [command] --help" for more information about a command.


# run tests
go test -cover ./... && echo ":)"
```
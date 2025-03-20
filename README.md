# Task Tracker

A simple command-line task management application written in Go.

## Project URL
https://roadmap.sh/projects/task-tracker

## Overview

Task Tracker is a lightweight CLI tool that helps you manage your tasks. Features include:

- Add new tasks
- Update existing tasks
- Delete tasks
- List all tasks with optional filtering
- Mark tasks as "in progress" or "done"
- Persistent storage using JSON

## Installation

```bash
git clone https://github.com/yourusername/task-tracker.git
cd task-tracker
go build
```

## Project Structure

```
task-tracker/
├── main.go           # Entry point and command handler
├── models/
│   └── task-list.go  # Task data structure and business logic
└── tasks.json        # Data storage file (created automatically)
```

## Usage

### Add a task

```bash
./task-tracker add "Complete the project documentation"
```

### Update a task

```bash
./task-tracker update 1 "Update the project documentation with examples"
```

### Delete a task

```bash
./task-tracker delete 1
```

### List all tasks

```bash
./task-tracker list
```

### List tasks with filter

```bash
./task-tracker list todo
./task-tracker list in-progress
./task-tracker list done
```

### Mark task as in progress

```bash
./task-tracker mark-in-progress 1
```

### Mark task as done

```bash
./task-tracker mark-done 1
```

## Data Structure

Tasks are stored with the following properties:

- **ID**: Unique identifier for the task
- **Description**: Task details
- **Status**: Current state (todo, in-progress, done)
- **CreatedAt**: Timestamp when task was created
- **UpdatedAt**: Timestamp when task was last modified

## Storage

All tasks are stored in a JSON file (`tasks.json`) in the application directory. This file is created automatically when the first task is added.

## Error Handling

The application provides error messages for common issues:
- Invalid command syntax
- Task ID not found
- File read/write errors


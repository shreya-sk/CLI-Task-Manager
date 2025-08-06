# Task Manager CLI

[![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat-square&logo=go)](https://golang.org/doc/go1.23)
[![License](https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square)](LICENSE)

A command-line task management application built in Go for learning and productivity.

## Features

- Create, read, update, and delete tasks
- Search and filter tasks by title, category, or priority
- Color-coded interface for improved readability
- Detailed statistics on task completion rates and progress
- Categorize tasks for better organization
- Set priority levels to focus on what matters most
- Persistent storage with automatic JSON backup

## Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/CLI-Task-Manager.git

# Navigate to the directory
cd CLI-Task-Manager

# Build the project
go build -o task-manager

# Run the application
./task-manager
```

## Usage

The application presents a simple menu interface:

```
-------------------------
Task Manager
-------------------------
1. Add Task
2. Show Tasks
3. Complete Task
4. Delete Task
5. Search Task
6. Completion Stats
7. Exit
-------------------------
```

### Managing Tasks
- **Add Task**: Enter title, category (personal, work, fitness), and priority (P1/P2/P3)
- **Show Tasks**: View all tasks in a formatted table with status
- **Complete/Delete**: Mark tasks as done or remove them by ID
- **Search**: Find tasks by title, category, or priority
- **Statistics**: View completion rates and productivity metrics

## Docker Versions

This project has multiple branches demonstrating different approaches:

- **main**: Basic Go application
- **docker-optimization**: Containerized version with Docker optimization techniques

Check the `docker-optimization` branch for advanced Docker implementation including multi-stage builds and volume mounting.

## Technical Details

- **Language**: Go 1.23
- **Dependencies**: fatih/color for terminal colors
- **Storage**: JSON file-based persistence
- **Data Structure**: Tasks with ID, title, status, category, priority, and timestamps

## Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b new-feature`
3. Commit your changes: `git commit -am 'Add new feature'`
4. Push the branch: `git push origin new-feature`
5. Submit a pull request

## License

This project is licensed under the MIT License.

## Acknowledgements

- The Go team for an excellent programming language
- [fatih/color](https://github.com/fatih/color) for terminal color capabilities

# CLI Task Manager

A command-line task management application built in Go that demonstrates containerization with Docker and data persistence through volume mounting.

## Features

- Create, read, update, and delete tasks
- Color-coded terminal interface
- Task categorization and priority levels
- Search and filter functionality
- Completion statistics and progress tracking
- Persistent JSON storage

## Quick Start

### Prerequisites
- Docker installed on your system

### Running with Docker

Build the image:
```bash
docker build -t task-manager .
```

Run with persistent storage:
```bash
mkdir -p data
docker run -it -v $(pwd)/data:/app/data task-manager
```

Your tasks will be saved to `./data/tasks.json` and persist between container runs.

### Local Development

If you prefer to run without Docker:
```bash
go run main.go task.go
```

## Docker Optimization Journey

This project demonstrates progressive Docker optimization:

| Version | Size | Description |
|---------|------|-------------|
| Basic | 1.3GB | Single-stage build with full Go toolchain |
| Alpine | 18.6MB | Multi-stage build with Alpine Linux base |
| Scratch | 5.26MB | Ultra-minimal build with scratch base |

**99.6% size reduction** achieved through multi-stage builds.

### Build Variants

Basic Docker build:
```bash
docker build -t task-manager:basic -f Dockerfile.basic .
```

Optimized Alpine build:
```bash
docker build -t task-manager:alpine .
```

Ultra-minimal scratch build:
```bash
docker build -t task-manager:scratch -f Dockerfile.scratch .
```

## Volume Mounting Implementation

### The Challenge
Initially, volume mounting presented a conflict where the mount would overwrite the container's binary files.

### The Solution
Discovered that volume mounts replace entire directories, not merge content. Fixed by:
- Mounting host directory to `/app/data` instead of `/app`
- Modifying application to save files to `data/tasks.json`
- Keeping the binary safe at `/app/main`

### Key Learning
Volume mounting connects host and container file systems, enabling data persistence while maintaining application functionality.

## Usage

The application provides an interactive menu system:

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

### Adding Tasks
When creating tasks, you can specify:
- Task title
- Category (personal, work, fitness, etc.)
- Priority level (P1, P2, P3)

### Task Management
- View tasks in a formatted table with completion status
- Mark tasks as complete or delete them
- Search tasks by title, category, or priority
- View completion statistics and trends

## Technical Implementation

### Architecture
- **Language**: Go 1.23
- **Dependencies**: fatih/color for terminal colors
- **Storage**: JSON file-based persistence
- **Containerization**: Multi-stage Docker builds

### Data Persistence
Tasks are stored in JSON format with the following structure:
```json
{
  "id": 1,
  "title": "Example Task",
  "completed": false,
  "createdAt": "2025-01-01T10:00:00Z",
  "category": "work",
  "priority": "P1"
}
```

### Docker Multi-Stage Build
The optimized Dockerfile uses two stages:
1. **Builder stage**: Full Go environment for compilation
2. **Runtime stage**: Minimal base image with only the binary

This approach reduces image size by 99.6% while maintaining full functionality.

## Development Notes

### Cross-Compilation
The Docker build includes cross-compilation flags for Linux compatibility:
```dockerfile
RUN CGO_ENABLED=0 GOOS=linux go build -a -o main .
```

### Volume Mount Strategy
```bash
# Host directory maps to container data directory
-v $(pwd)/data:/app/data
```

This ensures data persistence without interfering with the application binary.

## Build Information

- **Go Version**: 1.23
- **Base Images**: golang:1.23 (builder), alpine:latest or scratch (runtime)
- **Final Size**: 5.26MB (scratch), 18.6MB (alpine)
- **Architecture**: linux/amd64

---

This project demonstrates fundamental Docker concepts including containerization, optimization techniques, and data persistence patterns commonly used in production environments.
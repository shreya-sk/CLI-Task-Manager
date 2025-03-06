# 📋 Task Manager CLI

[![Go Version](https://img.shields.io/badge/Go-1.19+-00ADD8?style=flat-square&logo=go)](https://golang.org/doc/go1.19)
[![License](https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square)](LICENSE)

A powerful, feature-rich command-line task manager built with Go. Efficiently manage your tasks, set priorities, categorize tasks, track completion metrics, and more—all from your terminal.

## ✨ Features

- 📝 **Create, read, update, and delete tasks**
- 🔍 **Search and filter tasks** by title, category, or priority
- 🌈 **Color-coded interface** for improved readability
- 📊 **Detailed statistics** on task completion rates and progress
- 🏷️ **Categorize tasks** for better organization
- 🔢 **Set priority levels** to focus on what matters most
- 💾 **Persistent storage** with automatic JSON backup

## 🖥️ Screenshots

![Screenshot 2025-03-06 at 4 43 11 pm](https://github.com/user-attachments/assets/e1cdd27f-2b47-4d27-a53e-106b59e55651)

![Screenshot 2025-03-06 at 4 43 39 pm](https://github.com/user-attachments/assets/0737b1d0-6363-40e7-ba1b-223d1721a001)



## 🚀 Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/go-task-manager.git

# Navigate to the directory
cd go-task-manager

# Build the project
go build -o task-manager

# Run the application
./task-manager
```

## 📖 Usage

Once running, the application presents a simple menu:

```
-------------------------
Task Manager
-------------------------
1. Add Task
2. List Tasks
3. Complete Task
4. Delete Task
5. Show Statistics
6. Search Tasks
7. Exit
-------------------------
```

### Adding a Task

Select option `1` and follow the prompts to enter task details:
- Task title
- Category (e.g., personal, work, fitness)
- Priority (e.g., high, medium, low)

### Managing Tasks

- **List Tasks (Option 2)**: View all tasks in a neatly formatted table
- **Complete Task (Option 3)**: Mark a task as completed by entering its ID
- **Delete Task (Option 4)**: Remove a task by entering its ID

### Statistics & Insights

Option `5` provides detailed statistics about your tasks:
- Overall completion rate
- Weekly and monthly completion metrics
- Task distribution by category
- Productivity trends

### Search & Filter

Option `6` allows you to search and filter tasks by:
- Title keywords
- Category
- Priority level

## 🔧 Technical Details

The application uses:
- Go's built-in JSON marshaling/unmarshaling for data persistence
- Structured error handling for robustness
- The `fatih/color` package for the colorful UI
- Efficient data structures for task management

## 🛠️ Customization

You can easily customize the application by modifying:
- Categories in the `addTask` function
- Priority levels in the `addTask` function
- Colors in the UI by changing the color functions

## 🤝 Contributing

Contributions are welcome! Feel free to:
1. Fork the repository
2. Create a feature branch: `git checkout -b new-feature`
3. Commit your changes: `git commit -am 'Add a new feature'`
4. Push the branch: `git push origin new-feature`
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgements

- The Go team for an amazing programming language
- [fatih/color](https://github.com/fatih/color) for terminal color capabilities

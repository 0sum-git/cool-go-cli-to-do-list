# cool CLI To-Do List in Go

A command-line (CLI) task list application developed in Go, featuring an interactive interface and persistent storage.

![](https://github.com/rodrigofernandesribeiro/cool-go-cli-to-do-list/blob/main/demotodo.gif)

## 📋 Features

- Add tasks
- Remove tasks
- Toggle task completion status
- Interactive interface with dynamic borders
- JSON persistent storage
- UTF-8 character support
- Keyboard navigation
- Real-time task list updates

## 🛠️ Technologies Used

- **Go (Golang)** - Main programming language
- **github.com/eiannone/keyboard** - Keyboard input capture library
- **encoding/json** - JSON data handling
- **bufio** - User input reading
- **os** - File manipulation
- **strings** - String manipulation
- **unicode/utf8** - UTF-8 character support
- **fmt** - Output formatting

## 🔧 Data Structures

- **Doubly Linked List** - Custom implementation for task management
  - `Node` - List node structure
  
  - `DoublyLinkedList` - Main list structure
  
  - `task` - Task data structure
  
    

## 📥 Installation

1. Ensure Go is installed on your system.
2. Clone the repository:

   ```bash
   git clone https://github.com/0sum-git/go-cli-todo
   ```

3. Install dependencies:

   ```go
   go mod tidy
   ```



## 🚀 How to Run



### Development Mode
Navigate to the project directory and run:

go run cmd/todolist/main.go



### Build and Install
To create an executable and install it system-wide:

### For Linux:
go build -o todo cmd/todolist/main.go
sudo mv todo /usr/local/bin/

Now you can run the todo list from anywhere by typing `todo` on the terminal.



### For Windows:
go build -o todo.exe cmd/todolist/main.go



## 📝 How to Use

- **a** - Add new task
- **r** - Enter removal mode
- **1-9** - Toggle task completion (using corresponding number)
- **0** - Exit program or current mode

### Add Mode
- Type task description and press Enter
- Tasks are displayed in real-time as you add them
- Type "0" to exit add mode

### Remove Mode
- Press the number corresponding to the task you want to remove
- Press "0" to exit remove mode

## 💾 Storage

Tasks are automatically saved to a `tasks.json` file in JSON format. Example:

[
    {
        "desc": "Buy groceries",
        "completed": false
    },
    {
        "desc": "Write code",
        "completed": true
    }
]


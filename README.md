```markdown
# Golang To-Do App Documentation

## Table of Contents

1. [Overview]
2. [Prerequisites]
3. [Project Setup and Folder Structure]
4. [Backend (Golang) Code]
   - [Main Entry Point (`main.go`)]
   - [Model Definition (`internal/models/todo.go`)]
   - [Database Initialization (`internal/db/db.go`)]
   - [Router and API Handlers (`internal/router/router.go`)]
5. [Frontend Code]
   - [HTML Template (`templates/index.html`)]
   - [CSS Styling (`static/css/style.css`)]
   - [JavaScript Logic (`static/js/app.js`)]
6. [Troubleshooting and Common Issues]
7. [Future Enhancements]
8. [Conclusion]



## Overview

This project is a simple **To-Do List** application built with:

- **Backend:** Golang using the Gin framework and GORM for ORM with SQLite as the database.
- **Frontend:** HTML, CSS, and vanilla JavaScript.
- **Styling:** CSS with Font Awesome icons for interactive buttons.

The app demonstrates full-stack development with basic CRUD operations:
- **Create:** Add new tasks.
- **Read:** List all tasks.
- **Update:** Mark tasks as complete.
- **Delete:** Remove tasks.



## Prerequisites

- **Go Language:** Install from [https://go.dev/dl/](https://go.dev/dl/).
- **Git:** For version control.
- **SQLite:** (Used via GORM with the SQLite driver).
- **C Compiler:** Required if using a CGO-based SQLite driver (or use a pure Go driver).
- **Text Editor/IDE:** VS Code or similar.
- **Internet Connection:** To load external resources (e.g., Font Awesome CDN).



## Project Setup and Folder Structure

Create a project folder with the following structure:


Golang-ToDo-App/
├── static/
│   ├── css/
│   │   └── style.css          # CSS styling for the app
│   └── js/
│       └── app.js             # JavaScript logic for dynamic behavior
├── templates/
│   └── index.html             # Main HTML page template
├── internal/
│   ├── db/
│   │   └── db.go              # Database initialization and migration
│   ├── models/
│   │   └── todo.go            # Definition of the Todo model
│   └── router/
│       └── router.go          # API routes and handlers using Gin
├── main.go                    # Main entry point for the Golang server
├── go.mod                     # Go module file
├── go.sum                     # Go module dependencies


**Notes:**
- CSS files reside in `static/css/` and JavaScript files in `static/js/`.
- HTML templates are placed in the `templates/` folder.
- Backend code is organized under `internal/` for modularity.

---

## Backend (Golang) Code

### Main Entry Point (`main.go`)

This file initializes the Gin server, connects to the database, sets up routes, and starts the HTTP server.

```go
package main

import (
    "fmt"

    "github.com/your-username/Go-tossssssdo-list/internal/db"
    "github.com/your-username/Go-tossssssdo-list/internal/router"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Connect to the database and run migrations
    db.InitDB()

    // Setup API routes and HTML rendering routes
    router.SetupRoutes(r)

    fmt.Println("Server running on http://localhost:8080")
    r.Run(":8080")
}
```

### Model Definition (`internal/models/todo.go`)

Defines the `Todo` model using GORM, inheriting default fields from `gorm.Model`.

```go
package models

import "gorm.io/gorm"

type Todo struct {
    gorm.Model
    Title  string `gorm:"not null"`
    Status string `gorm:"not null"`
}
```

### Database Initialization (`internal/db/db.go`)

Initializes the SQLite database using GORM and auto-migrates the `Todo` model.

```go
package db

import (
    "fmt"
    "log"

    "gorm.io/driver/sqlite"
    "gorm.io/gorm"

    "github.com/your-username/Go-tossssssdo-list/internal/models"
)

var Database *gorm.DB

func InitDB() {
    var err error
    // Open (or create) the SQLite database file "todos.db"
    Database, err = gorm.Open(sqlite.Open("todos.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    fmt.Println("Connected to SQLite DB successfully")

    // Auto-migrate the schema for the Todo model
    err = Database.AutoMigrate(&models.Todo{})
    if err != nil {
        log.Fatal("Migration failed:", err)
    }
    fmt.Println("Database migration completed")
}
```

### Router and API Handlers (`internal/router/router.go`)

Defines the API endpoints for CRUD operations and renders the HTML page.

```go
package router

import (
    "net/http"

    "github.com/your-username/Go-tossssssdo-list/internal/db"
    "github.com/your-username/Go-tossssssdo-list/internal/models"
    "github.com/gin-gonic/gin"
)

// SetupRoutes initializes all routes
func SetupRoutes(r *gin.Engine) {
    r.LoadHTMLGlob("templates/*")   // Load HTML templates
    r.Static("/static", "./static") // Serve static assets

    // Define routes
    r.GET("/", ShowIndex)             // Render the index page
    r.GET("/todos", GetTodos)         // API: Get all todos
    r.POST("/todos", AddTodo)         // API: Add a new todo
    r.PUT("/todos/:id", UpdateTodo)   // API: Mark a todo as complete
    r.DELETE("/todos/:id", DeleteTodo) // API: Delete a todo
}

// ShowIndex renders the main HTML page
func ShowIndex(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", nil)
}

// GetTodos returns all todos in JSON format
func GetTodos(c *gin.Context) {
    var todos []models.Todo
    db.Database.Find(&todos)
    c.JSON(http.StatusOK, todos)
}

// AddTodo creates a new todo
func AddTodo(c *gin.Context) {
    var input struct {
        Title string `json:"title"`
    }

    // Bind JSON input to the struct
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    // Validate that title is not empty
    if input.Title == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
        return
    }

    // Create a new Todo with a default status of "pending"
    todo := models.Todo{
        Title:  input.Title,
        Status: "pending",
    }
    db.Database.Create(&todo)
    c.JSON(http.StatusOK, todo)
}

// UpdateTodo marks a todo as completed
func UpdateTodo(c *gin.Context) {
    id := c.Param("id")
    var todo models.Todo
    if err := db.Database.First(&todo, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
        return
    }

    todo.Status = "completed"
    db.Database.Save(&todo)
    c.JSON(http.StatusOK, todo)
}

// DeleteTodo removes a todo by its ID
func DeleteTodo(c *gin.Context) {
    id := c.Param("id")
    if err := db.Database.Delete(&models.Todo{}, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
```

---

## Frontend Code

### HTML Template (`templates/index.html`)

This file defines the user interface, including the input form, task list, and links to CSS and JavaScript.

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>To-Do List</title>
    <!-- Link to CSS; ensure the correct path -->
    <link rel="stylesheet" href="/static/css/style.css">
    <!-- Font Awesome CSS for icons -->
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.4.2/css/all.min.css">
</head>
<body>
    <h1>My To-Do List</h1>

    <!-- Input form is placed above the list -->
    <form id="todo-form">
        <input type="text" id="todo-input" placeholder="New task" required>
        <button type="submit">Add</button>
    </form>

    <ul id="todo-list"></ul>

    <!-- Load JavaScript at the end of the body -->
    <script src="/static/js/app.js"></script>
</body>
</html>
```

### CSS Styling (`static/css/style.css`)

Enhances the visual appearance of the app with layout, colors, and hover effects.

```css
/* General Body Styling */
body {
    font-family: Arial, sans-serif;
    background-color: #f8f9fa;
    display: flex;
    flex-direction: column;
    align-items: center;
    padding: 20px;
}

/* Heading Styling */
h1 {
    margin-top: 20px;
    margin-bottom: 20px;
}

/* Form styling */
form#todo-form {
    display: flex;
    gap: 10px;
    margin-bottom: 20px;
}

#todo-input {
    padding: 10px;
    font-size: 16px;
    border: 1px solid #ccc;
    border-radius: 5px;
}

/* List Styling */
ul#todo-list {
    list-style: none;
    padding: 0;
    width: 100%;
    max-width: 500px;
}

ul#todo-list li {
    display: flex;
    justify-content: space-between;
    align-items: center;
    background-color: #fff;
    padding: 10px;
    margin-bottom: 10px;
    border-radius: 5px;
    box-shadow: 0 2px 5px rgba(0,0,0,0.1);
}

/* Button Group for Complete and Delete */
.btn-group {
    display: flex;
    gap: 8px;
}

/* Button Styling */
.btn {
    background: none;
    border: none;
    cursor: pointer;
    font-size: 20px;  /* Increase icon size */
    padding: 5px;
}

.complete-btn {
    color: #28a745;
}

.delete-btn {
    color: #dc3545;
}

.btn:hover {
    opacity: 0.8;
    transform: scale(1.1);
    transition: transform 0.2s;
}
```

### JavaScript Logic (`static/js/app.js`)

Handles dynamic behavior such as adding tasks, fetching tasks from the backend, and handling Complete/Delete actions.

```javascript
// Wait until the DOM is fully loaded
document.addEventListener("DOMContentLoaded", function () {
    loadTodos();

    // Add event listener for new task submission
    document.getElementById("todo-form").addEventListener("submit", async function (event) {
        event.preventDefault();

        const inputField = document.getElementById("todo-input");
        const title = inputField.value.trim();
        
        if (!title) {
            alert("Task title cannot be empty");
            return;
        }

        try {
            const response = await fetch("/todos", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ title })
            });

            const data = await response.json();
            if (response.ok) {
                addTodoToUI(data);
                inputField.value = "";
            } else {
                alert("Error: " + data.error);
            }
        } catch (error) {
            console.error("Failed to send task:", error);
        }
    });
});

// Function to fetch and load all todos from the backend
async function loadTodos() {
    try {
        const response = await fetch("/todos");
        const todos = await response.json();

        const todoList = document.getElementById("todo-list");
        todoList.innerHTML = "";  // Clear existing list

        todos.forEach(addTodoToUI);
    } catch (error) {
        console.error("Error loading todos:", error);
    }
}

// Function to add a single todo item to the UI
function addTodoToUI(todo) {
    console.log("Adding todo:", todo);
    const todoList = document.getElementById("todo-list");
    const li = document.createElement("li");
    li.classList.add("todo-item");

    // Ensure the todo has an ID
    if (!todo.ID) {
        console.error("Error: Todo ID is missing", todo);
        return;
    }

    li.innerHTML = `
        <span>${todo.Title} - ${todo.Status}</span>
        <div class="btn-group">
            <button class="btn complete-btn" onclick="markComplete(${todo.ID})">
                <i class="fas fa-check"></i>
            </button>
            <button class="btn delete-btn" onclick="deleteTodo(${todo.ID})">
                <i class="fas fa-trash"></i>
            </button>
        </div>
    `;

    li.setAttribute("id", `todo-${todo.ID}`);
    todoList.appendChild(li);
}

// Function to mark a task as complete
async function markComplete(id) {
    try {
        const response = await fetch(`/todos/${id}`, { method: "PUT" });
        const updatedTodo = await response.json();

        if (response.ok) {
            document.querySelector(`#todo-${id} span`).textContent = `${updatedTodo.Title} - ${updatedTodo.Status}`;
        } else {
            alert("Error: " + updatedTodo.message);
        }
    } catch (error) {
        console.error("Error updating todo:", error);
    }
}

// Function to delete a task
async function deleteTodo(id) {
    try {
        const response = await fetch(`/todos/${id}`, { method: "DELETE" });
        const result = await response.json();

        if (response.ok) {
            document.getElementById(`todo-${id}`).remove();
        } else {
            alert("Error: " + result.message);
        }
    } catch (error) {
        console.error("Error deleting todo:", error);
    }
}
```

---

## Troubleshooting and Common Issues

1. **DOM Not Loaded (Null Errors):**  
   - **Issue:** JavaScript runs before DOM elements are available.  
   - **Solution:** Ensure that the `<script>` tags are placed at the end of the `<body>` section or use the `defer` attribute in the `<head>`.

2. **Missing Icons:**  
   - **Issue:** Font Awesome icons may not display if the CDN link is incorrect or blocked.  
   - **Solution:** Use the Font Awesome CSS CDN link in the `<head>` and ensure it is loaded properly.

3. **API Data Issues:**  
   - **Issue:** New tasks might not appear or `/todos` endpoint may not return the expected data.  
   - **Solution:** Add logging in both the backend (using `fmt.Println`) and frontend (using `console.log`) to verify that data is being transmitted correctly.

4. **Incorrect File Paths:**  
   - **Issue:** CSS or JavaScript files are not loaded due to wrong paths.  
   - **Solution:** Verify that the paths in `index.html` match your folder structure (e.g., `/static/css/style.css` and `/static/js/app.js`).

---

## Future Enhancements

- **Persistent Storage Improvements:**  
  Use a production-ready database and add robust error handling.

- **Enhanced UI/UX:**  
  Consider integrating a front-end framework (e.g., React, Vue) or a CSS framework (e.g., Bootstrap) for improved styling and responsiveness.

- **Authentication:**  
  Add user authentication to manage individual to-do lists.

- **Real-time Updates:**  
  Implement WebSockets for real-time task updates across devices.

- **Testing:**  
  Add unit and integration tests for both backend and frontend functionality.

---

## Conclusion

This documentation provides a comprehensive guide for setting up, running, and maintaining the Golang To-Do App with a JavaScript frontend. It covers:

- **Backend Development:** Using Golang, Gin, and GORM for REST API creation and database interactions.
- **Frontend Development:** Using HTML, CSS, and vanilla JavaScript to create a dynamic user interface.
- **File Organization and Troubleshooting:** Detailed instructions and common issues with solutions.
- **Future Enhancements:** Suggestions for scaling and improving the application.

By following this guide, you will have a modular, maintainable full-stack application. Feel free to update and extend this documentation as your project evolves.

Happy coding!
```

Feel free to modify any sections as needed for your project.
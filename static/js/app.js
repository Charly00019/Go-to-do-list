document.addEventListener("DOMContentLoaded", function () {
    // ✅ Ensure the form exists before adding the event listener
    const todoForm = document.getElementById("todo-form");
    if (!todoForm) {
        console.error("Error: todo-form not found!");
        return;
    }

    todoForm.addEventListener("submit", async function (event) {
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

    // ✅ Load todos only after the page is fully loaded
    loadTodos();
});


async function loadTodos() {
    try {
        const response = await fetch("/todos");
        const todos = await response.json();

        const todoList = document.getElementById("todo-list");
        todoList.innerHTML = "";

        todos.forEach(addTodoToUI);
    } catch (error) {
        console.error("Error loading todos:", error);
    }
}

function addTodoToUI(todo) {
    const todoList = document.getElementById("todo-list");
    const li = document.createElement("li");
    li.classList.add("todo-item");

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

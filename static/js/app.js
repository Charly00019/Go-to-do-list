document.addEventListener("DOMContentLoaded", fetchTodos);

function fetchTodos() {
    fetch("/todos")
        .then(response => response.json())
        .then(todos => {
            const todoList = document.getElementById("todoList");
            todoList.innerHTML = "";
            todos.forEach(todo => {
                const li = document.createElement("li");
                li.innerHTML = `
                    ${todo.title} - ${todo.status}
                    <button onclick="markComplete(${todo.ID})">✔</button>
                    <button onclick="deleteTodo(${todo.ID})">❌</button>
                `;
                todoList.appendChild(li);
            });
        });
}

function addTodo() {
    const todoInput = document.getElementById("todoInput");
    const title = todoInput.value.trim();
    if (!title) return alert("Enter a todo!");

    fetch("/todos", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ title: title })
    })
    .then(() => {
        todoInput.value = "";
        fetchTodos();
    });
}

function markComplete(id) {
    fetch(`/todos/${id}`, { method: "PUT" })
    .then(() => fetchTodos());
}

function deleteTodo(id) {
    fetch(`/todos/${id}`, { method: "DELETE" })
    .then(() => fetchTodos());
}

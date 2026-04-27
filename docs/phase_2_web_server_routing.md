# 🚦 Phase 2: Web Server & Routing (The "URLconf")

In Django, handling a request involves `urls.py` (Routing) and `views.py` (Logic). In Go, we combine these using **Handlers**.

## 1. The Standard Library: `net/http`

Go has a powerful web server built right into its standard library. You don't *need* a framework like Django or FastAPI to get started, though we will use one later to make life easier (the "DRF" of Go).

### Django vs Go Request Flow

| Step | Django | Go |
| :--- | :--- | :--- |
| **Server** | Gunicorn / Runserver | `http.ListenAndServe` |
| **Router** | `urls.py` (`path()`) | `http.HandleFunc()` |
| **View** | `def my_view(request):` | `func(w http.ResponseWriter, r *http.Request)` |
| **Response** | `return HttpResponse("...")` | `fmt.Fprintf(w, "...")` |

---

## 2. Key Concepts: The "Handler"

In Go, a web view is just a function that takes two arguments:
1.  **`w http.ResponseWriter`**: Think of this as your "Response Object." You "write" data to it (headers, status codes, and the body).
2.  **`r *http.Request`**: This is your "Request Object." It contains the URL, headers, and any body data sent by the user.

### Example Handler:
```go
func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the Go-Social Home Page!")
}
```

---

## 3. The "URLconf": Mapping Paths

To tell Go which function to run for a specific URL, we use `http.HandleFunc`.

```go
http.HandleFunc("/", homeHandler)
```
This is exactly like `path('', home_view)` in Django.

---

## 4. Let's Build it!

Update your `main.go` to include a web server. We will use the `net/http` package.

### Your task:
Modify your `main.go` to look like this:

```go
package main

import (
	"fmt"
	"net/http" // New import for web functionality
)

// This is your View
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Go-Social!")
}

func main() {
	// 1. Mapping the path (Routing)
	http.HandleFunc("/", homeHandler)

	// 2. Starting the server
	fmt.Println("Server starting on :8080...")
	err := http.ListenAndServe(":8080", nil)
	
	// Error handling: If the server fails to start, we want to know why.
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
```

### 🧐 Detailed Breakdown of `main.go`:
-   **`http.ListenAndServe(":8080", nil)`**: This is the big one. It starts a server on port 8080. The `nil` tells it to use the default router (which we just configured with `HandleFunc`).
-   **Error Handling**: Go doesn't use "Try/Except" like Python. Instead, functions return an `error` value, and you check if it's `nil`. This is the "Go Way."

---

## 🔍 Deep Dive: Understanding the Syntax

Since you're new to Go, let's break down the "weird" parts of that code block.

### 1. `w http.ResponseWriter` and `r *http.Request`
In Django, your view looks like `def home(request):`. In Go, we have two parts:

*   **`w`**: The "Writer". This represents the **Response**. You use this to send data back to the user.
*   **`r`**: The **Request**. This contains everything the user sent (Headers, URL, Cookies, Form data).
*   **Why `w` and `r`?**: These are just variable names. You could name them `response` and `request`, but the Go community loves short, 1-letter names for local variables.

### 2. The `*` (The Pointer)
You noticed `r` has a `*http.Request` while `w` does not. 

*   **What is `*`?**: It's a **Pointer**. 
*   **In Django/Python**: When you pass an object to a function, Python automatically passes a "reference" to it.
*   **In Go**: You have to be explicit. `*http.Request` means "don't give me a copy of the whole request, give me the **memory address** of the original request."
*   **Why?**: A request can be huge (it might have a file upload). Copying it would be slow. Using a pointer (`*`) is lightning fast because we're just passing a small memory address.
*   **Why doesn't `w` have a `*`?**: `http.ResponseWriter` is an **Interface**. In Go, interfaces are "magic" containers that handle the pointers for you behind the scenes. Just remember: Requests get a `*`, ResponseWriters don't.

### 3. `Println` vs `Fprintf`
Both come from the `fmt` (format) package, but they have different targets:

*   **`fmt.Println("...")`**: Prints to your **Terminal/Console**. It stands for "Print Line".
*   **`fmt.Fprintf(w, "...")`**: The **F** stands for "File" (or any "Writer"). It means "Print **to** this specific thing." By passing `w` as the first argument, we are telling Go: "Don't print this to my console; print it into the response that goes back to the browser."

### 4. `err := http.ListenAndServe(":8080", nil)`

*   **`:=`**: This is the "Short Variable Declaration." It's like saying `err = ...` in Python, but it also handles the "Type" for you automatically. Use `:=` for the first time you create a variable.
*   **`ListenAndServe`**: This function does two things: 1. It binds to the port `:8080`. 2. It starts a loop that waits for requests forever.
*   **The `nil`**: In Django, the `URLconf` is your router. In Go, `nil` here tells the server to use the "Default ServeMux" (the built-in router where we registered our paths using `http.HandleFunc`). 



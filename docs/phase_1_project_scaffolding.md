### 🏗️ Phase 1: Project Scaffolding (The "Startproject")

In Django, you'd run `django-admin startproject`. In Go, we initialize a **Module**.

#### 1. What is a Go Module?
Think of a Go Module as your `venv` (virtual environment) and `requirements.txt` merged into one. It defines the root of your project and tracks every dependency you install.

#### 2. Initialize your project
Open your terminal in the `go` directory and run this command:

```bash
go mod init go-social
```

*   **Breakdown:** `go mod init` is the command, and `go-social` is the name of your module (like your Django project name).
*   **Result:** This will create a `go.mod` file. Take a look at it once it's created—it's very minimal right now!

#### 3. Create the Entry Point
Now, create a file named `main.go`. This is your `manage.py` and `wsgi.py` equivalent—it's where everything starts.

Copy this into your `main.go` (I won't write the file for you, as requested, but here is the template to explain it):

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go-Social!")
}
```

*   **`package main`**: This tells Go that this file should be compiled into an executable program. In Django, logic is spread out, but in Go, there's always one `main` package that starts it all.
*   **`import "fmt"`**: This is like `import sys` or `import os`. `fmt` (format) is the standard library for printing text to the console.
*   **`func main()`**: This is the "Main Entry Point." When you run your program, Go looks for this exact function.

---

### Your Turn
1. Run `go mod init go-social` in your terminal.
2. Create `main.go` and add the code above.
3. Once done, run your program using:
   ```bash
   go run main.go
   ```

Tell me when you see `Hello, Go-Social!` in your terminal, and we'll move on to building an actual web server!


---

## Possible Error

```bash
C:\Users\Destiny Franks\Desktop\go>go mod init go-social
go: creating new go.mod: module go-social

C:\Users\Destiny Franks\Desktop\go>go run main.go
# command-line-arguments
.\main.go:5:11: syntax error: unexpected {, expected (
.\main.go:6:13: syntax error: unexpected ( in parameter list; possibly missing comma or )
.\main.go:7:1: syntax error: non-declaration statement outside function body
```

The error `unexpected {, expected (` on line 5 is telling you exactly what's wrong. In Go, functions **must** have parentheses `()` after the name, even if they don't take any arguments.

### The Fix

In your `main.go`, change:
```go
func main {  // ❌ Missing parentheses
```
to:
```go
func main() { // ✅ Added parentheses
```

### Why does this happen?
Unlike Python, where you might see `def my_func:` and it would be a syntax error immediately, Go's parser explicitly looks for the parameter list `()` after the function name before it allows the opening brace `{`.

**Go Punctuation Rule:**
In Go, the opening brace `{` **must** be on the same line as the function declaration. You can't do this:
```go
func main()
{ // ❌ This will also cause a syntax error!
}
```

Go automatically inserts semicolons behind the scenes, so the brace position matters!

---

### Your Turn
Fix that line to `func main() {` and try running `go run main.go` again. You'll be greeted with "hello, go social" shortly!
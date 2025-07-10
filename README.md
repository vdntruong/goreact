# Go-React Starter Project

This project provides a foundational example of integrating a **React.js frontend** with a **Go (Golang) backend**. It demonstrates how to effectively serve a bundled React user interface using Go's native HTTP server capabilities, with build automation managed by a `Makefile`.

## Project Structure

* `ui/`: Contains the React.js frontend application.
    * `src/`: React source code.
    * `package.json`: Frontend dependencies and build scripts.
* `be/`: Contains the Go (Golang) backend application.
    * `main.go`: Go server entry point.
    * `go.mod`: Go module dependencies.
    * `static/`: **Output directory for bundled React UI files.**
* `Makefile`: Automates build and run tasks for both frontend and backend.
* `Dockerfile`: For containerizing the application (optional, but good to mention if present).
* `.gitignore`: Specifies files and directories that Git should ignore.

## Prerequisites

Before you begin, ensure you have the following installed:

* **Go (Golang)**: [Installation Guide](https://go.dev/doc/install)
* **Node.js & npm**: [Installation Guide](https://nodejs.org/en/download/) (npm is installed with Node.js)
* **direnv**: [Installation Guide](https://direnv.net/docs/installation.html) (Recommended for local environment variable management)
* **make**: Most Unix-like systems (Linux, macOS) have `make` pre-installed. For Windows, you might need to install [MinGW](http://www.mingw.org/) or use [WSL](https://learn.microsoft.com/en-us/windows/wsl/install).

## Getting Started

1.  **Clone the repository:**
    ```bash
    git clone [your-repo-url]
    cd project/
    ```

2.  **Allow `direnv`:**
    If you're using `direnv` (highly recommended for local development), allow it to load environment variables for the project:
    ```bash
    direnv allow
    ```
    (Ensure your `.envrc` in the project root points to your `env.example` or `.env` file for your Go backend configuration.)

## Running the Project with `make`

The `Makefile` simplifies the process of setting up, building, and running both your frontend and backend.

### `make build-and-run`

This is the primary command to get your application up and running. It performs the following steps:

1.  Installs Node.js dependencies for the UI.
2.  Builds the React frontend and places the static files into `be/static/`.
3.  Builds the Go backend executable.
4.  Starts the Go backend server.

To run the project:

```bash
make build-and-run
```

Once the Go server starts (you should see Go server starting on :8080), open your web browser and navigate to:

http://localhost:8080

You should see your React UI loaded, and it will interact with the Go backend's API endpoints (e.g., /api/hello).

### Other make commands:

- `make setup`: Installs frontend (npm) dependencies.
- `make build-ui`: Builds the React frontend only.
- `make build-be`: Builds the Go backend executable only.
- `make run-be`: Runs the Go backend server (assumes it's already built).
- `make clean`: Cleans up build artifacts (e.g., node_modules, Parcel cache, be/static, Go executable).

# FastAPI-Todo Documentation

## Installation

1. Clone the repository: `git clone <repo-url>`
2. Navigate to the project root: `cd practice`
3. Install dependencies: `uv sync` (assumes uv is installed)
4. For task-specific: Navigate to `task1-testing` if needed.

## Running the Application

- Activate virtual environment: `source .venv/bin/activate`
- Run the server: `make start` (starts uvicorn on http://0.0.0.0:8000)
- Access Swagger UI: Open http://localhost:8000/docs in browser.

## API Documentation

The API provides CRUD operations for tasks:

- **GET /tasks**: Retrieve all tasks.
- **POST /tasks**: Create a new task (JSON: {"id": int, "title": str, "completed": bool}).
- **PUT /tasks/{task_id}**: Update a task by ID.
- **DELETE /tasks/{task_id}**: Delete a task by ID.
  Full OpenAPI spec available at /openapi.json.

## Usage Examples

- Create task: `curl -X POST "http://localhost:8000/tasks" -H "Content-Type: application/json" -d '{"id":1,"title":"Buy groceries","completed":false}'`
- Get tasks: `curl http://localhost:8000/tasks`
- Update: `curl -X PUT "http://localhost:8000/tasks/1" -H "Content-Type: application/json" -d '{"id":1,"title":"Buy groceries updated","completed":true}'`
- Delete: `curl -X DELETE "http://localhost:8000/tasks/1"`

## Troubleshooting

- **ModuleNotFoundError**: Ensure venv is activated and dependencies installed via `uv sync`.
- **Port in use**: Change port in Makefile or kill process on 8000.
- **Tests fail**: Check pytest.ini and run `make test` after sync.
- If API returns 404: Verify task ID exists; for other issues, check server logs.

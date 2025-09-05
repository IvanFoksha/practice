# FastAPI-Todo Documentation (ENG)

## Installation

1. Clone the repository: `git clone https://github.com/IvanFoksha/practice`
2. Navigate to the project root: `cd practice && make start-project`
3. Install dependencies: `uv sync` (assumes uv is installed)
4. For task-specific: Navigate to `task1-testing` if needed.

## Running the Application

- Activate virtual environment: `source .venv/bin/activate` (if not activated)
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

# Документация FastAPI-Todo (RU)

## Установка

1. Клонируйте репозиторий: `git clone https://github.com/IvanFoksha/practice`
2. Перейдите в корень проекта: `cd practice && make start-project`
3. Установите зависимости: `uv sync` (предполагается, что uv установлен)
4. Для конкретной задачи: Перейдите в `task1-testing`, если нужно.

## Запуск приложения

- Активируйте виртуальное окружение: `source .venv/bin/activate` (если не активировано)
- Запустите сервер: `make start` (запускает uvicorn на http://0.0.0.0:8000)
- Доступ к Swagger UI: Откройте http://localhost:8000/docs в браузере.

## Документация API

API предоставляет CRUD-операции для задач:

- **GET /tasks**: Получить все задачи.
- **POST /tasks**: Создать новую задачу (JSON: {"id": int, "title": str, "completed": bool}).
- **PUT /tasks/{task_id}**: Обновить задачу по ID.
- **DELETE /tasks/{task_id}**: Удалить задачу по ID.
  Полная спецификация OpenAPI доступна по /openapi.json.

## Примеры использования

- Создать задачу: `curl -X POST "http://localhost:8000/tasks" -H "Content-Type: application/json" -d '{"id":1,"title":"Купить продукты","completed":false}'`
- Получить задачи: `curl http://localhost:8000/tasks`
- Обновить: `curl -X PUT "http://localhost:8000/tasks/1" -H "Content-Type: application/json" -d '{"id":1,"title":"Купить продукты обновлено","completed":true}'`
- Удалить: `curl -X DELETE "http://localhost:8000/tasks/1"`

## Устранение неисправностей

- **ModuleNotFoundError**: Убедитесь, что venv активировано и зависимости установлены через `uv sync`.
- **Порт занят**: Измените порт в Makefile или завершите процесс на 8000.
- **Тесты не проходят**: Проверьте pytest.ini и запустите `make test` после sync.
- Если API возвращает 404: Проверьте существование ID задачи; для других проблем смотрите логи сервера.

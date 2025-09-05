import pytest
from fastapi.testclient import TestClient
from task1_testing.main import app, tasks

client = TestClient(app)


@pytest.fixture(autouse=True)
def clear_tasks():
    tasks.clear()  # Очищаем tasks перед каждым тестом


def test_get_tasks_empty():
    response = client.get("/tasks")
    assert response.status_code == 200
    assert response.json() == []


def test_create_task():
    task_data = {"id": 1, "title": "Test Task", "completed": False}
    response = client.post("/tasks", json=task_data)
    assert response.status_code == 200
    assert response.json() == task_data
    assert len(tasks) == 1


def test_update_task():
    # Сначала создаём задачу
    task_data = {"id": 1, "title": "Original", "completed": False}
    client.post("/tasks", json=task_data)
    updated_data = {"id": 1, "title": "Updated", "completed": True}
    response = client.put("/tasks/1", json=updated_data)
    assert response.status_code == 200
    assert response.json() == updated_data
    assert tasks[0].title == "Updated"
    assert tasks[0].completed is True


def test_update_task_not_found():
    update_json = {"id": 999, "title": "Nonexistent", "completed": False}
    response = client.put("/tasks/999", json=update_json)
    assert response.status_code == 200  # FastAPI возвращает 200
    # с ошибкой, можно скорректировать на 404 в app
    assert response.json() == {"error": "Task not found"}


def test_delete_task():
    # Создаём задачу
    task_data = {"id": 1, "title": "To Delete", "completed": False}
    client.post("/tasks", json=task_data)
    response = client.delete("/tasks/1")
    assert response.status_code == 200
    assert response.json() == task_data
    assert len(tasks) == 0


def test_delete_task_not_found():
    response = client.delete("/tasks/999")
    assert response.status_code == 200
    assert response.json() == {"error": "Task not found"}


def test_create_task_invalid_data():
    invalid_data = {"id": "not_int", "title": "Invalid"}  # Некорректный тип id
    response = client.post("/tasks", json=invalid_data)
    assert response.status_code == 422  # Unprocessable Entity от Pydantic

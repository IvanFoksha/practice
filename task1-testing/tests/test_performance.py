import pytest
from fastapi.testclient import TestClient
from main import app, tasks

client = TestClient(app)

@pytest.fixture(autouse=True)
def clear_tasks():
    tasks.clear()


def test_performance_create_tasks(benchmark):
    def create_tasks():
        for i in range(1000):
            client.post("/tasks", json={
                "id": i,
                "title": f"Task {i}",
                "completed": False
            })

    benchmark(create_tasks)  # Измеряем время на 1000 create


def test_performance_get_tasks(benchmark):
    # Подготавливаем 1000 задач
    for i in range(1000):
        client.post("/tasks", json={
            "id": i,
            "title": f"Task {i}",
            "completed": False
        })

    def get_tasks():
        client.get("/tasks")

    benchmark(get_tasks)  # Измеряем время на get всех задач

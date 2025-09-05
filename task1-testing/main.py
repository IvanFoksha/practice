from fastapi import FastAPI
from pydantic import BaseModel

app = FastAPI()


class Task(BaseModel):
    id: int
    title: str
    completed: bool = False


tasks: list[Task] = []


@app.get("/tasks")
async def get_tasks():
    return tasks


@app.post("/tasks")
async def create_task(task: Task):
    tasks.append(task)
    return task


@app.put("/tasks/{task_id}")
async def update_task(task_id: int, task: Task):
    for i, t in enumerate(tasks):
        if t.id == task_id:
            tasks[i] = task
            return task
        return {"error": "Task not found"}


@app.delete("/tasks/{task_id}")
async def delete_task(task_id: int):
    for i, t in enumerate(tasks):
        if t.id == task_id:
            return tasks.pop(i)
        return {"error": "Task not found"}

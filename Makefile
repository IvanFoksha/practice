.PHONY: setup test clean


start-project:
	@echo "Starting project..."
	@bash uv venv
	@bash source .venv/bin/activate
	@bash uv sync

build:
	@echo "Building task1-testing..."
	@bash build.sh

test:
	@echo "Running tests in task1_testing..."
	@cd task1_testing && pytest

clean:
	@echo "Cleaning task1-testing..."
	@deactivate && rm -rf .venv

start:
	@echo "Starting task1_testing..."
	@cd task1_testing && uvicorn main:app --reload --host 0.0.0.0 --port 8000

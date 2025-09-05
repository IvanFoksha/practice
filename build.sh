#!/bin/bash
uv venv
source .venv/bin/activate
uv sync
cd task1-testing
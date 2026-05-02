# 🧠 Cortex — AI Codebase Intelligence Engine

Cortex is a system that analyzes entire code repositories and generates structured, AI-powered insights about the codebase, its components, and overall architecture.

Instead of manually exploring hundreds of files, Cortex helps you **understand any project in minutes**.

---

# 🚀 What It Does

* Accepts a public GitHub repository URL
* Clones and processes the codebase
* Parses and analyzes source files
* Uses AI to generate:

  * File-level summaries
  * Project-level explanation
* Returns a structured understanding of the system

---

# 🧱 Architecture Overview

```text
Client
  ↓
Go API (Core Engine)
  ↓
Worker Pool (Concurrency Layer)
  ↓
Python AI Service
  ↓
Response (Summaries + Insights)
```

---

# ⚙️ Tech Stack

### Backend (Core Engine)

* Go (Golang)
* Concurrency (goroutines, channels, worker pools)

### AI Layer

* Python
* LLM API (OpenAI or equivalent)

### Infrastructure

* Docker
* AWS EC2 (deployment)

---

# 🔥 Key Features

* ⚡ Concurrent file processing using worker pools
* 🧠 AI-powered code understanding
* 🧱 Clean system architecture (Go + Python separation)
* 📦 Scalable pipeline design (ready for queues, RAG, etc.)

---

# 📌 API Endpoints

## Analyze Repository

```http
POST /analyze
```

### Request

```json
{
  "repo_url": "https://github.com/user/repo"
}
```

### Response (example)

```json
{
  "project_summary": "This project is a REST API service handling user authentication...",
  "files": [
    {
      "file": "main.go",
      "summary": "Entry point of the application..."
    }
  ]
}
```

---

# 🛠️ How It Works

1. User submits a GitHub repository URL
2. The system clones the repository locally
3. Files are scanned and filtered
4. Files are split into manageable chunks
5. Worker pool processes files concurrently
6. Each chunk is sent to the AI service
7. AI returns summaries
8. System aggregates results into a final explanation

---

# 🧪 Development Phases

## Phase 1 — Core Pipeline

* Repo cloning
* File scanning
* Basic processing

## Phase 2 — Concurrency Layer

* Worker pool
* Parallel file processing

## Phase 3 — AI Integration

* Python AI service
* File + project summaries

---

# 💣 Design Principles

* Separation of concerns (Go = orchestration, Python = AI)
* Concurrency-first backend design
* Incremental system building (no overengineering)
* Real-world system architecture

---

# 🚀 Getting Started (Local)

### Run Go Service

```bash
go run ./cmd
```

### Test Endpoint

```bash
curl -X POST http://localhost:8080/analyze \
-H "Content-Type: application/json" \
-d '{"repo_url":"https://github.com/user/repo"}'
```

---

# 📦 Deployment

* Dockerized application
* Hosted on AWS EC2
* Exposed via public API

---

# 🔮 Future Improvements

* Upload ZIP repositories
* Codebase visualization (flowcharts / graphs)
* Chat with codebase (RAG-based)
* Persistent storage of analysis
* Multi-repo comparison

---

# 🎯 Why This Project

Cortex is designed to demonstrate:

* Advanced Go concurrency patterns
* Distributed system thinking
* AI integration in backend systems
* Real-world problem solving

---

# ⚡ Final Note

Cortex is not just a project — it's a system designed to explore how modern backend engineering and AI can work together to solve real developer problems.

---

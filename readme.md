<div align="center">

# 🧠 Cortex
### Codebase Intelligence System

**Paste a GitHub URL. Understand the entire codebase instantly.**

[![FastAPI](https://img.shields.io/badge/FastAPI-0.136-009688?style=flat-square&logo=fastapi)](https://fastapi.tiangolo.com)
[![Python](https://img.shields.io/badge/Python-3.11-3776AB?style=flat-square&logo=python)](https://python.org)
[![LangChain](https://img.shields.io/badge/LangChain-latest-1C3C3C?style=flat-square)](https://langchain.com)
[![Qdrant](https://img.shields.io/badge/Qdrant-Vector_DB-DC244C?style=flat-square)](https://qdrant.tech)
[![Redis](https://img.shields.io/badge/Redis-Queue-FF4438?style=flat-square&logo=redis)](https://redis.io)

[Demo](#) · [API Docs](#api-reference) · [Report Bug](https://github.com/0xvicky/Cortex/issues)

</div>

---

## What is Cortex?

Cortex is an AI-powered codebase intelligence system. You give it a public GitHub repository URL — it processes the entire codebase, generates a structured summary, and lets you have a conversation with the code.

No more spending hours reading through unfamiliar repos. Ask Cortex what you want to know.

---

## The Problem I Solved

The naive RAG approach for code understanding is simple: loop through every file, send each one to an LLM for summarization, then index everything. 

**The problem:** a repo with 200 files = 200 LLM calls. Slow, expensive, and completely unpredictable in cost.

**My solution:** a custom heuristic file-scoring algorithm that identifies the most important files *before* touching the LLM — reducing API calls from **O(n) to O(1)**, regardless of repo size.

---

## How the Heuristic Works

Every file in the repository gets scored across three dimensions:

| Signal | Logic |
|---|---|
| **Directory depth** | Files closer to the root score higher — entry points, configs, and core logic live near the top |
| **Naming conventions** | Standard important filenames (`main.py`, `app.js`, `index.ts`) score higher. README files get the highest score — they're gold |
| **File size** | Files within a specific size range are preferred. Too small = boilerplate. Too large = generated or bundled |

Top 10 files are selected. **1 LLM call. That's it.**

---

## Two Pipelines, Two Speeds

```
┌─────────────────────────────────────────────────────────────┐
│                        GitHub Repo URL                       │
└──────────────────────────┬──────────────────────────────────┘
                           │
              ┌────────────┴────────────┐
              │                         │
              ▼                         ▼
   ⚡ Fast Pipeline              🔍 RAG Pipeline
   (Quick Summary)              (Code Q&A)
              │                         │
   Heuristic scoring           Full repo chunked
   Top 10 files selected       & embedded into
   → Single LLM call           Qdrant vector DB
              │                         │
              ▼                         ▼
   Structured repo summary     Context-aware answers
   in seconds                  to any code question
```

### ⚡ Fast Pipeline
Uses the heuristic scoring to grab the top 10 most important files and sends them in a single LLM call. Clean, structured project summary — delivered in seconds.

### 🔍 RAG Pipeline
The full codebase is chunked, embedded using HuggingFace sentence transformers, and stored in Qdrant. When you ask a question, semantically relevant chunks are retrieved and passed to the LLM for precise, context-aware answers.

---

## Architecture

```
┌─────────────┐     HTTP      ┌──────────────────────────────────────┐
│   React     │ ──────────── │           FastAPI Backend              │
│  Frontend   │              │                                        │
└─────────────┘              │  ┌─────────┐       ┌──────────────┐  │
                             │  │  Auth   │       │  Ingest API  │  │
                             │  │ (JWT)   │       │              │  │
                             │  └─────────┘       └──────┬───────┘  │
                             │                           │           │
                             │                    ┌──────▼───────┐  │
                             │                    │  RQ Worker   │  │
                             │                    │  (async job) │  │
                             │                    └──────┬───────┘  │
                             │                           │           │
                             └───────────────────────────┼───────────┘
                                                         │
                          ┌──────────────────────────────┼──────────────────┐
                          │                              │                  │
                    ┌─────▼──────┐              ┌───────▼──────┐   ┌───────▼──────┐
                    │   Redis    │              │    Qdrant    │   │    Groq      │
                    │  (Queue +  │              │  (Vectors)   │   │    LLM       │
                    │  Storage)  │              └──────────────┘   └──────────────┘
                    └────────────┘
```

---

## Tech Stack

**Backend**
- [FastAPI](https://fastapi.tiangolo.com) — async API framework
- [LangChain](https://langchain.com) — LLM orchestration and RAG pipeline
- [Qdrant](https://qdrant.tech) — vector database for semantic search
- [HuggingFace Sentence Transformers](https://huggingface.co) — code embeddings
- [Groq](https://groq.com) — LLM inference
- [Redis](https://redis.io) + [RQ](https://python-rq.org) — async background job queue
- [SQLite](https://sqlite.org) + [SQLAlchemy](https://sqlalchemy.org) — user auth storage
- [GitPython](https://gitpython.readthedocs.io) — repo cloning

**Frontend**
- [React](https://react.dev) — UI framework
- [Tailwind CSS](https://tailwindcss.com) — styling

**Infrastructure**
- [Docker](https://docker.com) — containerization
- [Render](https://render.com) — deployment
- [Qdrant Cloud](https://cloud.qdrant.io) — managed vector DB
- [Redis Cloud](https://redis.io/cloud) — managed Redis

---

## Getting Started

### Prerequisites

- Python 3.11+
- Docker & Docker Compose
- A [Groq API key](https://console.groq.com)
- A [Qdrant Cloud](https://cloud.qdrant.io) account (free tier)
- A [Redis Cloud](https://redis.io/cloud) account (free tier)

### Local Setup

**1. Clone the repo**
```bash
git clone https://github.com/0xvicky/Cortex.git
cd Cortex/backend
```

**2. Set up environment variables**
```bash
cp .env.example .env
```

Fill in your `.env`:
```env
GROQ_API_KEY=your_groq_api_key
QDRANT_URL=https://xxxx.aws.cloud.qdrant.io
QDRANT_API_KEY=your_qdrant_api_key
REDIS_URL=rediss://default:password@xyz.redis.cloud:6379
JWT_SECRET=your_random_secret_key
GOOGLE_CLIENT_ID=your_google_client_id
GOOGLE_CLIENT_SECRET=your_google_client_secret
FRONTEND_URL=http://localhost:5173
```

**3. Run with Docker**
```bash
docker compose up --build
```

This starts:
- `api` — FastAPI server on `http://localhost:8000`
- `worker` — RQ background worker for async repo processing

**4. Start the frontend**
```bash
cd ../frontend
npm install
npm run dev
```

Visit `http://localhost:5173`

---

## API Reference

| Method | Endpoint | Description |
|---|---|---|
| `POST` | `/auth/login` | Login with Google OAuth |
| `POST` | `/ingest` | Submit a GitHub repo URL for processing |
| `GET` | `/job/{job_id}` | Poll job status |
| `POST` | `/query` | Ask a question about a processed repo |
| `DELETE` | `/job/{job_id}` | Delete a job and clean up vectors |

Full interactive docs available at `http://localhost:8000/docs`

---

## Key Design Decisions

**Why heuristic scoring instead of processing all files?**
Cost and speed predictability. A 10-file repo and a 500-file repo both cost exactly 1 LLM call for summarization. The heuristic is accurate enough for the vast majority of repos — most of what matters in any codebase lives near the root.

**Why async background jobs with RQ?**
Cloning and embedding a full repository can take 30-120 seconds. Doing this synchronously in an HTTP request would timeout. RQ lets the API return immediately with a `job_id` while the worker processes in the background. The client polls for status.

**Why Qdrant for vector storage?**
Per-job collection isolation — each repo gets its own Qdrant collection, making cleanup simple and preventing cross-repo retrieval pollution.

---

## What I Learned

Agentic systems aren't just about prompting an LLM. A lot of the real engineering happens *before* the LLM call — how you filter, rank, chunk, and retrieve data matters more than the prompt itself. The heuristic scoring is the core insight of this project: **smart data selection beats brute-force processing**.

---

## Author

**Vivek Tyagi** — [GitHub](https://github.com/0xvicky) · [LinkedIn](https://linkedin.com/in/your-linkedin)

---

<div align="center">
  <sub>Built with curiosity and too many LLM API calls 🧠</sub>
</div>
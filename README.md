# Go RAG (Retrieval-Augmented Generation)

This project is a Retrieval-Augmented Generation (RAG) application built with Go, PostgreSQL (pgvector), and local LLMs via Ollama. 

It allows you to chat with your own text or markdown documents by performing semantic search and providing context-aware answers.

## Prerequisites

To run this project, make sure you have the following installed on your machine:
- [Go](https://go.dev/) (1.20 or newer)
- [Docker & Docker Compose](https://www.docker.com/) (For the PostgreSQL vector database)
- [Ollama](https://ollama.ai/) (For running local LLMs)

## Setup & Installation

**1. Create Document Folders**
Before running the application, you need to create the folders where your documents will be stored. Create a `documents` folder in the root of the project, and a `processed` folder inside it:
```bash
mkdir -p documents/processed
```

**2. Start the Database**
The project uses PostgreSQL with the `pgvector` extension to store vector embeddings. Start the database using Docker Compose:
```bash
docker-compose up -d
```
*Note: You might need to run the SQL commands inside `create-table.sql` on your PostgreSQL instance to create the necessary tables.*

**3. Pull Required Models (Ollama)**
By default, the project uses `gemma3:latest` for generation and `nomic-embed-text` for vector embeddings. Pull these models via Ollama:
```bash
ollama run gemma3:latest
ollama pull nomic-embed-text
```

**4. Environment Variables (.env)**
Create a `.env` file in the root directory and configure it as follows:
```env
OPENAI_BASE_URL=http://localhost:11434/v1
OPENAI_API_KEY=dummy-key
OPENAI_MODEL=gemma3:latest
SYSTEM_PROMPT_FILE=./prompts/system-custom.md
DATABASE_URL=postgres://rag:rag@localhost:5432/rag?sslmode=disable

EMBEDDING_BASE_URL=http://localhost:11434/v1
EMBEDDING_DIM=768
EMBEDDING_MODEL=nomic-embed-text
INGEST_DIR=./documents
PROCESSED_DIR=./documents/processed

HTTP_ADDR=:8080
IMAGES_DIR=./documents/images
VISION_MODEL=mistral-small3.1
```

## Running the Application

Once everything is set up, you can start the application by running:

```bash
go run ./cmd/rag
```

If it starts successfully, you will see output similar to this:
```
[rag] 2026/07/30 13:59:15 watching ./documents for new documents
[rag] 2026/07/30 13:59:15 vector store ready
Chat session started. Type Q to quit
> 
```

## How to Use

**1. Document Ingestion**
To allow the model to answer questions based on your data, simply drag and drop your text or markdown files (`.txt` or `.md`) into the `documents/` folder. The application actively monitors this folder while running:
- When a new file is detected, it chunks the text, creates embeddings, and saves them to the database.
- You will see an `ingested filename.md` log in your terminal when it completes.
- The processed file is automatically moved to the `documents/processed/` folder to prevent reprocessing and clutter.

**2. Chatting**
Type your questions next to the `> ` prompt in the terminal and press Enter. The model will find the most relevant context from your ingested documents and generate an answer based on them.

To exit the application, simply type `Q` and press Enter.

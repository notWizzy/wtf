# ⚡ wtf

> Runs your command, catches the error, and explains it in plain English using a local LLM — right in your terminal.

---

## The Problem

A command fails. You get a wall of red text. You copy it, switch to your browser, open ChatGPT, paste it, wait for an answer, switch back to your terminal.

Every. Single. Time.

`wtf` cuts out all of that. It runs your command, captures the failure, and explains it instantly — no browser, no copy-pasting, no context switching.

---

## Demo

```bash
wtf python3 migrate.py

⚡ wtf  $ python3 migrate.py

WHAT HAPPENED: Python could not connect to the database because the connection was refused.
WHY: The PostgreSQL server is not running on the expected port.
FIX: Start the PostgreSQL service with `brew services start postgresql`.
```

---

## How It Works

1. You run `wtf <your command>`
2. `wtf` executes it and captures the stderr output
3. If the command fails, the error is sent to a local LLM running via Ollama
4. A plain English explanation is printed directly in your terminal

Everything runs locally on your machine. No API keys. No internet required. No data leaves your computer.

---

## Tech Stack

- **Go** — compiled to a single binary, fast startup
- **Cobra** — CLI argument parsing
- **Ollama** — local LLM inference (llama3.2)
- **Lip Gloss** — terminal output styling

---

## Prerequisites

- [Go 1.21+](https://golang.org/dl/)
- [Ollama](https://ollama.com) with `llama3.2` pulled

```bash
ollama pull llama3.2
```

---

## Installation

```bash
git clone https://github.com/notWizzy/wtf.git
cd wtf
make install
```

---

## Usage

```bash
# any command that might fail
wtf python3 script.py
wtf go build ./...
wtf npm run build
wtf ./deploy.sh
```

---

## Project Structure
wtf/

├── cmd/root.go              # CLI command, argument parsing, orchestration

├── internal/

│   ├── llm/llm.go           # Ollama HTTP client and prompt engineering

│   └── output/output.go     # Terminal output styling with Lip Gloss

├── main.go                  # Entry point

└── Makefile                 # Build and install commands

---

## Local Development

```bash
# Run without installing
go run main.go python3 script.py

# Build binary
make build

# Run tests
make test
```
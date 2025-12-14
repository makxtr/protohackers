# Protohackers Solutions

Solutions for https://protohackers.com challenges

## Project Structure

```
protohackers/
├── 0:SmokeTest/
│   ├── go/            # Go implementation
│   ├── elixir/        # Elixir implementation
│   └── rust/          # Rust implementation
├── 1:TaskName/
│   ├── go/
│   └── elixir/
├── fly.toml           # Common Fly.io config
└── Makefile           # Deployment helper
```

## Deployment

Deploy any task implementation to Fly.io:

```bash
# List all available tasks and implementations
make list

# Deploy task 0 (Go version, default)
make deploy TASK=0

# Deploy task 0 (Elixir version)
make deploy TASK=0 LANG=elixir

# Deploy task 1 (Rust version)
make deploy TASK=1 LANG=rust
```

**Default language is `go` if LANG is not specified.**

## Adding New Tasks

### New task (Go):
```bash
mkdir -p "N:TaskName/go/cmd/server"
cp Dockerfile.template "N:TaskName/go/Dockerfile"
# Add your Go code
make deploy TASK=N
```

### Additional language implementation:
```bash
mkdir -p "N:TaskName/elixir"
# Add Dockerfile and code for Elixir
make deploy TASK=N LANG=elixir
```
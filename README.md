# Learning-goLang

### Students REST API in Go: my first production-style Go service

![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white) 

A small CRUD API that follows the standard Go project layout: `cmd/` for entry points and `internal/` for config, handlers, storage, and types. It uses `net/http` method-based routing, struct validation, config loaded from flags and env, and graceful shutdown on SIGINT/SIGTERM.

| Method | Path | Action |
|---|---|---|
| `POST` | `/api/students` | Create a student (validated with `go-playground/validator`) |
| `GET` | `/api/students/{id}` | Get a student by id |

```
cmd/students-api/main.go      Server bootstrap and graceful shutdown
internal/config               YAML/env config (cleanenv)
internal/http/handlers        Student handlers
internal/utils/response       JSON response helpers
```

```bash
go run cmd/students-api/main.go -config config/local.yaml
```

> The SQLite storage package that `main.go` imports (`internal/storage/sqlite`) isn't committed yet, so add it before building.

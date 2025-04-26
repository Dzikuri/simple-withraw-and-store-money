# Simple Withdraw and Store Money API

## Folder Structure

```markdown
.
├── api/
│ ├── handler/
│ │ └── nasabah_handler.go
│ ├── middleware/
│ │ └── log_middleware.go
│ ├── router/
│ │ └── router.go
│ └── api.go
├── build/
│ ├── tmp/
│ │ ├── bin/
│ │ └── logs/
│ │ └── build-errors.log
│ └── .gitignore
├── cmd/
│ └── server/
│ └── main.go
├── config/
│ └── config.go
├── docker/
│ ├── .dockerignore
│ ├── docker-compose.yaml
│ └── Dockerfile
├── model
├── repository/
│ └── nasabah_repository.go
├── service/
│ └── register_service.go
├── storages/
│ ├── /public
│ └── .gitignore
├── tests/
├── util/
│ ├── database.go
│ ├── parsing.go
│ ├── password.go
│ ├── string.go
│ └── validation.go
├── .air.linux.toml
├── .air.windows.toml
├── .editorconfig
├── .env
├── .env.example
├── .gitignore
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

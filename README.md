# Simple Withdraw and Store Money API

This is a simple API for withdraw and store money with nasabah (customer). The API uses PostgreSQL as database, also uses Zerolog for logging and uses Docker for deployment.

## Pre-request

-   > = Go v1.23
-   Able to run `MAKEFILE` for connivent
-   [Golang AIR](https://github.com/air-verse/air) for live reloading
-   [Golang Migrate](https://github.com/golang-migrate/migrate/) - for database migration
-   PostgreSQL >= version 14

## Installation

-   Clone the project
-   Execute command `make tidy`

```bash
make tidy
```

-   Create database in PostgreSQL with any name you want, for example `simple_withdraw_and_store_money` or anything else. The name of the database will be used in the configuration file `.env`
-   Copy `.env.example` to `.env`

```bash
cp .env.example .env
```

-   Fill the empty value of `.env` with your configuration

```conf
API_PORT=8089
APP_ENV=development
APP_DEBUG=true

DB_HOST=127.0.0.1
DB_USERNAME=
DB_PASSWORD=
DB_PORT=
DB_NAME=
```

-   Run migration using command `make migration/up`

```bash
make migration/up
```

-   If you want to run once the application you can use command `make run` or if you want using `AIR` just use `make run/{os}/live` change the `{os}` with your OS, the supported for now only Windows and Linux

-   Another option deploy using Docker. You can use the Docker Compose. Make sure to fill the `.env` configuration

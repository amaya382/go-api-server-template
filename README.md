# go-api-server-template

## Components
* Language: golang 1.12
* Architecture: Clean Architecture (≒Onion Architecture)
* Dependency Management: go module
* API Doc: [swaggo/swag](https://github.com/swaggo/swag)
* Dev Environment: VSCode settings
* Framework/Library, good stuff
    * Config: [spf13/viper](https://github.com/spf13/viper)
    * WAF: [gin-gonic/gin](https://github.com/gin-gonic/gin)
        * Validation: [go-playground/validator](https://github.com/go-playground/validator)
    * ORM: []()
    * Response Marshaling: [liip/sheriff](https://github.com/liip/sheriff)


## Placeholders
* Module Name: `github.com/amaya382/go-api-server-template`
    * (Should be `github.com/amaya382/example`)
* Application Name: `example`
* Database: PostgreSQL


## Architecture
```
/
┣ .vscode/          # VSCode-related settings
┣ cmd/              # Program entry points (main.go)
┃  ┣ example/       # Program entry point for main app
┃  ┗ migration/     # Program entry point for migration tool
┣ config/           # Config loader
┣ docs/             # Code-First API specs
┣ domain/           # Domain layer, in terms of DDD
┃  ┣ model/         # Domain models
┃  ┣ repository/    # Repositoy interfaces
┃  ┗ service/       # Domain services
┣ infrastructure/   # Infrastructure layer, in terms of DDD
┃  ┣ gormpostgres/  # Repository impl by gorm and postgres (Example)
┃  ┣ mock/          # Repository mock impl (Example)
┃  ┗ someapi/       # Repository impl by some APIs (Example)
┣ interface/        # Interface layer, in terms of DDD
┃  ┗ gin            # Interface by gin (Example)
┃     ┣ handler/    # HTTP handlers
┃     ┣ middleware/ # Middlewares
┃     ┗ router/     # Routing definitions
┃        ┗ v1/      # Routing definitions on ver 1 APIs (Example)
┣ usecase/          # Use Case layer, in terms of DDD
┃  ┃                # (Some people say "Application layer")
┃  ┗ http/          # Use case for a http server (Example)
┣ util/             # Essential utilities
┣ example.toml      # Settings for app
┗ Makefile          # For build
```

## Rules
### Repository
* Create: Add an entry w/o id into DB (Like HTTP POST method). Return a pointer of a single value.
* Replace: Add an entry w/ id into DB (Like HTTP PUT method). Idempotent. Return pointers of an old value and a new value. If on a system which manages id by itself (e.g. DB sequence), creating a new entry w/ a non-existing id by *Replace* will fail.
* Get: Search a single entry by id. Return a pointer of a single value. If not found, return a NotFoundError.
* GetByX: Search a single entry by X. Return a pointer of a single value. If not found, return a NotFoundError.
* List: List all entries. Return a list of a pointer of a value. If not found, return an empty list.
* ListByX: List all entries filtered by condition X. Return a list of a pointer of a value. If not found, return an empty list.
* Delete: Idempotent. Return nil.
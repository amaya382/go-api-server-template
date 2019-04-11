# go-api-server-template

## Components
* Language: golang 1.12
* Architecture: Clean Architecture (≒Onion Architecture)
* Dependency management: go module
* API Doc: [swaggo/swag](https://github.com/swaggo/swag)
* Dev Environment: VSCode settings
* Framework/Library
    * Config: [spf13/viper](https://github.com/spf13/viper)
    * WAF: [gin-gonic/gin](https://github.com/gin-gonic/gin)
        * Validation: [go-playground/validator](https://github.com/go-playground/validator)
    * ORM: []()
    * Response Marshaling: [liip/sheriff](https://github.com/liip/sheriff)


## Placeholders
* Application Name: `example`
* Database: PostgreSQL

## Rules
### Repository
* Create: Add an entry w/o id into DB (Like HTTP POST method). Return a pointer of a single value.
* Replace: Add an entry w/ id into DB (Like HTTP PUT method). Idempotent. Return pointers of an old value and a new value.
* Get: Search a single entry by id. Return a pointer of a single value. If not found, return a NotFoundError.
* GetByX: Search a single entry by X. Return a pointer of a single value. If not found, return a NotFoundError.
* List: List all entries. Return a list of a pointer of a value. If not found, return an empty list.
* ListByX: List all entries filtered by condition X. Return a list of a pointer of a value. If not found, return an empty list.
* Delete: Idempotent. Return nil.
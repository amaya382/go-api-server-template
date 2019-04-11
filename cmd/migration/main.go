package main

import "github.com/amaya382/go-api-server-template/infrastructure/gormpostgres"

func main() {
	gormpostgres.Migrate()
}

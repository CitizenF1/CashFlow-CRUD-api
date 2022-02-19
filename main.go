package main

import (
	"cashflow/db"
	"cashflow/router"
)

func main() {
	r := router.SetupRouter()

	db.Connect()
	r.Run()
}

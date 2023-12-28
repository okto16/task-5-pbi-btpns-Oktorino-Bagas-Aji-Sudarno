package main

import (
    "golang-api/database"
    "golang-api/router"
)

func main() {
    database.Connect()
    database.Migrate()

    r := router.SetupRouter()
    r.Run(":8080")
}

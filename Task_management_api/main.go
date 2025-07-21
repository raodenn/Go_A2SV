package main

import (
	"Task_management_api/data"
	"Task_management_api/router"
)

func main() {
	data.ConnectToMongoDB()

	r := router.SetupRouter()
	r.Run(":8080")
}

package main

import "Task_management_api/router"

func main() {
	router := router.SetupRouter()

	router.Run(":8080")
}

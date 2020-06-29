package main

import (
	"limakcv/src/app"
	"limakcv/src/config"
	"os"
)

func main() {
	config := config.GetConfig()
	port := os.Getenv("PORT")
	app := &app.App{}
	app.Initialize(config)
	app.Run(":" + port) // for cloud
	// app.Run(":8116") //for local

}

package main

import (
	"backend/controllers"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	controllers.Start()
}

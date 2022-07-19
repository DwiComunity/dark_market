package main

import (
	"fmt"
	"os"

	"github.com/crownss/dark_market/config"
	"github.com/crownss/dark_market/routers"
	"github.com/joho/godotenv"
)

func main(){
	env := godotenv.Load(".env")
	if env != nil {
		fmt.Println("Cannot load environment")
	}
	config.InitDB()
	e := routers.Router()
	e.Run(":" + os.Getenv("PORT"))
}
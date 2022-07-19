package main

import (
	"fmt"

	"github.com/crownss/dark_market/config"
	"github.com/joho/godotenv"
)

func main(){
	env := godotenv.Load(".env")
	if env != nil {
		fmt.Println("Cannot load environment")
	}
	config.InitDB()
	// e := routes.New()
	// e.Start(":" + os.Getenv("PORT"))
}
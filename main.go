package main

import (
	"fmt"
	"log"
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
	log.Println("\n\n\t\t\tRemember !\n\tYOU ARE NOT REQUIRED TO FILL RUN_HOST OR RUN_PORT IN .env\n\tby the default it will be use http://localhost:8000\n\n\t")
	if os.Getenv("RUN_HOST") != "" && os.Getenv("RUN_PORT") != ""{
		e.Run(os.Getenv("RUN_HOST")+ ":" + os.Getenv("RUN_PORT"))
	}
	e.Run(":8000")

}
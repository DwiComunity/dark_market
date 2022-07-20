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
	log.Println("\n\n\t\t\tRemember !\n\tyou're not requiered to fill RUN_HOST and RUN_PORT in .env file\n\t the default running on http://localhost:8000\n\t")
	e := routers.Router()
	if os.Getenv("RUN_HOST")!=""&&os.Getenv("RUN_PORT")!=""{
		e.Run(os.Getenv("RUN_HOST")+":"+os.Getenv("RUN_PORT"))
	}
	e.Run("localhost:8000")
}
package main

import (
	"log"
	"os"

	"github.com/crownss/dark_market/config"
	"github.com/crownss/dark_market/helpers"
	"github.com/crownss/dark_market/routers"
)

func main() {
	//gin.SetMode(gin.ReleaseMode) //uncomment for release mode
	helpers.Env(".env")
	helpers.WG.Add(2)
	go helpers.StartContainer(os.Getenv("CONTAINER_ID"))
	go func() { config.InitDB(); helpers.WG.Done() }()
	helpers.WG.Wait()
	e := routers.Router()
	log.Println("\n\n\t\t\tRemember !\n\tYOU ARE NOT REQUIRED TO FILL RUN_HOST OR RUN_PORT IN .env\n\tby the default it will be use http://localhost:8000\n\n\t")

	if os.Getenv("RUN_HOST") != "" && os.Getenv("RUN_PORT") != "" {
		e.Run(os.Getenv("RUN_HOST") + ":" + os.Getenv("RUN_PORT"))
	}
	e.Run("localhost:8000")
}
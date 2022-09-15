package main

import (
	"context"
	"log"
	"os"

	"github.com/crownss/dark_market/config"
	"github.com/crownss/dark_market/helpers"
	"github.com/crownss/dark_market/routers"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/joho/godotenv"
)

func main() {
	//gin.SetMode(gin.ReleaseMode) //uncomment for release mode
	Env(".env")
	helpers.WG.Add(2)
	go StartContainer(os.Getenv("CONTAINER_ID"))
	go func() { config.InitDB(); helpers.WG.Done() }()
	helpers.WG.Wait()
	e := routers.Router()
	log.Println("\n\n\t\t\tRemember !\n\tYOU ARE NOT REQUIRED TO FILL RUN_HOST OR RUN_PORT IN .env\n\tby the default it will be use http://localhost:8000\n\n\t")

	if os.Getenv("RUN_HOST") != "" && os.Getenv("RUN_PORT") != "" {
		e.Run(os.Getenv("RUN_HOST") + ":" + os.Getenv("RUN_PORT"))
	}
	e.Run("localhost:8000")
}

func Env(file string) error {
	env := godotenv.Load(file)
	if env != nil {
		log.Fatal("cannot load env file with error:\n", env.Error())
	}
	return env
}

func StartContainer(s string) string {
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Fatal(err.Error())
	}
	cli.ContainerStart(context.Background(), s, types.ContainerStartOptions{})
	log.Println("container starting with id:", s)
	helpers.WG.Done()
	return s
}

package helpers

import (
	"context"
	"log"
	"sync"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/joho/godotenv"
)

var WG = sync.WaitGroup{}

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
	WG.Done()
	return s
}
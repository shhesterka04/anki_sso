package main

import (
	"anki_sso/internal/config"
	"fmt"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg) //TODO: удалить

	//TODO: логгер
	//TODO: приложение
	//TODO: grpc
}

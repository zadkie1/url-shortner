package main

import (
	"fmt"

	"github.com/zadkie1/url-shortner/internal/config"
)

func main() {

	//config
	cfg := config.MustLoad()

	fmt.Println(cfg)

	//init logger

	//init storage

	//init router

	//run server

}

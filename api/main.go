package main

import (
	"github.com/eovinicius/short-link/api/config"
	"github.com/eovinicius/short-link/api/router"
)

func main() {
	config.IntializeSetup()
	router.Initialize()
}

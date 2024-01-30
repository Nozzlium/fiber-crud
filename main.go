package main

import (
	"fiber-crud/config"
	"fiber-crud/lib"
)

func main() {
	cfg, err := config.GetViper()
	if err != nil {
		panic(err)
	}
	db, err := lib.GetDatabase(cfg)
	if err != nil {
		panic(err)
	}
	app := lib.NewApp(db)
	app.Listen("localhost:1723")
	if err != nil {
		panic(err)
	}
}

package main

import (
	"fiberun/internal/router"
	"fiberun/module/base"
	"fiberun/module/conf"
	"log"
)

func init() {
	//init config
	conf.LoadConfig()
	//init database
	dsn := conf.Get("db.dsn")
	base.ConnectDB(dsn)
}

func main() {
	engine := base.FiberEngine()

	router.SetupRoutes(engine)

	log.Fatal(engine.Listen(conf.Get("fiber.addr")))
}

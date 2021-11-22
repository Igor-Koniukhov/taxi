package main

import (
	"taxi/configs"
)

var (
	app configs.AppConfig
	err error
)


func main()  {

	Run(&app)


}
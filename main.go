package main

import (
	"demo/router"
	"demo/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()

	r := router.Router()
	r.Run(":8081")
}

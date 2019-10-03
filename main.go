package main

import (
	_ "user/conf"
	_ "user/database"
	"user/router"
	orm "user/database"
	. "user/conf"
)

func main() {

	defer orm.Eloquent.Close()
	defer orm.Redis.Close()

	router := router.InitRouter()
	router.Run( ":" + ConfigApi["port"] )
	return
}
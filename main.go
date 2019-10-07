package main

import (
	. "user/conf"
	_ "user/conf"
	_ "user/database"
	orm "user/database"
	"user/router"
)


func main() {
	defer orm.Eloquent.Close()
	defer orm.Redis.Close()

	router := router.InitRouter()
	router.Run( ":" + ConfigApi["port"] )
	return
}
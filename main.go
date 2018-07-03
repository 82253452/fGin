package main

import (
	. "fGin/config"
	. "fGin/router"
	_"fGin/subRouter"
)

func main() {
	defer Db.Close()
	Router.Run(":8009")
}

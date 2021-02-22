package main

import (
	"github.com/EikoNakashima/test-go2/db"
	"github.com/EikoNakashima/test-go2/server"
)

func main() {

	db.Init()
	server.Init()
}

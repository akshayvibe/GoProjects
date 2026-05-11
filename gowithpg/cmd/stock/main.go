package main

import (
	"gowithpg/config"
	"gowithpg/internal/db/postgres"
	"net/http"
)
func main(){

	//loading configuration
	cfg:=config.MustLoad()

	//database setup 
	storage,err:=postgres.New(cfg)

	server:=&http.Server{
		
	}
}


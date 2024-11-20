package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/parthkapoor-dev/go-server/cmd/api"
	"github.com/parthkapoor-dev/go-server/db"
)

func main(){

	db , err := db.NewMySQLStorage(*&mysql.Config{
		User : "root",
		Passwd: "asd",
		Addr: "127.0.0.1:3306",
		DBName: "go-server",
		Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	}) 

	server := api.APIServerSetup(":8080" , nil)
	if err := server.Run(); err != nil {
		log.Fatal()
	}
}
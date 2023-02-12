package main

import "github.com/HiDoYa/rent-cat/app"

// @title           Rent Cat
// @version         1.0
// @description     Manages expense sharing
// @host            localhost:8080
// @BasePath        /api/v1
// @securityDefinitions.basic  BasicAuth
func main(){
	app.Launch()
}
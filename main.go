package main

import (
	"log"
	db "profiles/Db"
	handler "profiles/Handler"
	repository "profiles/Repository"
	service "profiles/Service"
	"profiles/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(" .env file not found — using system environment variables")
	} else {
		log.Println(" .env file loaded successfully")
	}

	db.DbInit()

	userRepo := repository.NewUserRepo(db.DB)
	userService := service.NewUserService(userRepo)
	handler := handler.NewUserHandler(userService)

	profileRepo := repository.NewProfileRepo(db.DB)
	profileService := service.NewProfileService(profileRepo)
	Profilehandler := handler.NewProfileHandler(profileService)

	r := gin.Default()
	routes.Routes(r, handler, Profilehandler)

	r.Run(":8080")

}

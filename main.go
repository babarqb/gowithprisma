package main

import (
	"fmt"
	"gowithprisma/config"
	"gowithprisma/controller"
	"gowithprisma/helper"
	"gowithprisma/repository"
	"gowithprisma/router"
	"gowithprisma/service"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("👉 Could not load environment variables", err)
	}

	fmt.Printf("Start Server! " + os.Getenv("PORT") + "\n")

	//handle db connection
	db, err := config.ConnectDB()
	helper.ErrorPanic(err)
	defer db.Prisma.Disconnect()

	// repository
	postRepository := repository.NewPostRepository(db)

	// service
	postService := service.NewPostServiceImpl(postRepository)

	// controller
	postController := controller.NewPostController(postService)

	//router
	routes := router.NewRouter(postController)

	server := &http.Server{
		Addr:           ":" + os.Getenv("PORT"),
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server_err := server.ListenAndServe()
	if server_err != nil {
		panic(server_err)
	}

}

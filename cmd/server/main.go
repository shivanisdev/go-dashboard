package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/letstalkndev/go-dashboard/internal/blog"
	"github.com/letstalkndev/go-dashboard/internal/config"
	"github.com/letstalkndev/go-dashboard/internal/user"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Version indicates the current version of the application.
var Version = "1.0.0"

func main() {

	cfg := config.LoadConfiguration("../../config/local.json")

	address := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)

	fmt.Println(address)

	mongoClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		fmt.Println("Error creating client connecting", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = mongoClient.Connect(ctx)
	if err != nil {
		fmt.Println("Error connecting", err)
	}
	if err := mongoClient.Ping(context.TODO(), readpref.Primary()); err != nil {
		// Can't connect to Mongo server
		fmt.Println("Not connected")
		log.Fatal(err)
	}
	db := mongoClient.Database(cfg.DbName)

	srv := &http.Server{
		Handler: buildHandler(db),
		Addr:    address,
	}
	log.Fatal(srv.ListenAndServe())

}

func buildHandler(db *mongo.Database) http.Handler {
	router := mux.NewRouter()
	blog.RegisterHandlers(router)

	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	user.RegisterHandlers(router, userService)

	http.Handle("/", router)
	return router
}

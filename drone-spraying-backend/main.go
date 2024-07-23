package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/yourusername/drone-spraying-backend/handlers"
)

var client *mongo.Client

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("MongoDB bağlantısı başarısız: %v", err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("MongoDB'ye ping atılamadı: %v", err)
	} else {
		log.Println("MongoDB bağlantısı başarılı.")
	}
	defer client.Disconnect(ctx)

	r := mux.NewRouter()
	handlers.SetClient(client)
	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.HandleFunc("/logout", handlers.Logout).Methods("POST")
	r.HandleFunc("/register", handlers.Register).Methods("POST")

	handler := cors.Default().Handler(r)
	log.Fatal(http.ListenAndServe(":8080", handler))
}

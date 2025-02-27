package main

import (
	"cat_cafe/controllers"
	"cat_cafe/repository"
	"cat_cafe/routes"

	"context"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// ชื่อมต่อMongoDB
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("MongoDB Connect Error:", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("🔴 MongoDB Ping Failed:", err)
	} else {
		log.Println("✅ Connected to MongoDB successfully!")
	}

	db := client.Database("cat_cafe")

	//ส่ง Database ไปให้ Repository
	repo := repository.NewCatRepository(db)

	// ส่ง Repository ไปให้ Controller
	catController := controllers.NewCatController(repo)

	//กำหนด API Routes และเริ่มต้น API
	r := gin.Default()
	routes.SetupRoutes(r, catController)
	r.Run(":9090")
}

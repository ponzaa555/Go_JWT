package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	routes "github.com/ponzaa555/Go_JWT/routes"

	"github.com/gin-gonic/gin"
	"github.com/ponzaa555/Go_JWT/database"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func main() {

	port := os.Getenv("PORT")

	coll := database.DBinstance()

	restaurant_id := "40356018"

	var result bson.M
	fmt.Println("Pass Database Connection")
	err := coll.FindOne(context.TODO(), bson.D{{"restaurant_id", restaurant_id}}).
		Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the restaurant_id %s\n", restaurant_id)
	}
	if err != nil {
		panic(err)
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoute(router)
	routes.UserRoute(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-1"}) // setHeader
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-2"}) // setHeader
	})

	router.Run(":" + port)
}

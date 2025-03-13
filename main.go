package main

import (
	"os"

	"github.com/ponzaa555/Go_JWT/database"
	routes "github.com/ponzaa555/Go_JWT/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")

	// restaurant_id := "40356018"

	// var result bson.M
	// fmt.Println("Pass Database Connection")
	// err := coll.FindOne(context.TODO(), bson.D{{"restaurant_id", restaurant_id}}).
	// 	Decode(&result)
	// if err == mongo.ErrNoDocuments {
	// 	fmt.Printf("No document was found with the restaurant_id %s\n", restaurant_id)
	// }
	// if err != nil {
	// 	panic(err)
	// }
	// jsonData, err := json.MarshalIndent(result, "", "    ")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%s\n", jsonData)

	//check connect Database
	_ = database.DBinstance()

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
	// defer database.CloseDB()
}

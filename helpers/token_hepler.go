package helpers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/ponzaa555/Go_JWT/database"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type SignedDetail struct {
	Email      string
	First_Name string
	Last_Name  string
	Uid        string
	User_type  string
	jwt.StandardClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

var Secret_key string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(email string, firstName string, lastName string, userType string, uid string) (token string, refreshToken string, err error) {
	err = nil
	claim := &SignedDetail{
		Email:      email,
		First_Name: firstName,
		Last_Name:  lastName,
		User_type:  userType,
		Uid:        uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refrestClaim := &SignedDetail{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	// create jwt token
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(Secret_key))
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refrestClaim).SignedString([]byte(Secret_key))
	if err != nil {
		log.Panic(err)
		return
	}
	return token, refreshToken, err
}

func UpdateAllToken(token string, refreshToken string, userId string) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	var updateObject bson.D

	updateObject = append(updateObject, bson.E{Key: "token", Value: token})
	updateObject = append(updateObject, bson.E{Key: "refresh_token", Value: refreshToken})

	UpdateAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObject = append(updateObject, bson.E{Key: "update_at", Value: UpdateAt})
	// update MogoDB
	filter := bson.M{"user_id": userId}
	opt := options.UpdateOne().SetUpsert(true)

	// upsert data
	_, err := userCollection.UpdateOne(
		ctx,
		filter,
		bson.D{
			{"$set", updateObject},
		},
		opt,
	)
	defer cancel()
	if err != nil {
		log.Panic(err)
		return
	}
}

func ValidateToken(clientToken string) (claim *SignedDetail, msg string) {
	token, err := jwt.ParseWithClaims(
		clientToken,
		&SignedDetail{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(Secret_key), nil
		})
	if err != nil {
		msg = err.Error()
		return
	}
	claim, ok := token.Claims.(*SignedDetail)
	if !ok {
		msg = fmt.Sprintf("token is invalid")
		msg = err.Error()
		return
	}
	if claim.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("token is expired")
		msg = err.Error()
		return
	}
	return claim, msg
}

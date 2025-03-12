package helpers

import (
	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/ponzaa555/Go_JWT/database"
	"go.mongodb.org/mongo-driver/v2/mongo"
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

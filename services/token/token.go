package token

import (
	"context"
	// "fmt"
	"log"
	"time"

	models "ecommerce/models/data"
	db "ecommerce/models/db"
	types "ecommerce/types"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MyCustomClaims struct {
	Email            *string `json:"email"`
	First_Name       *string `json:"first_name"`
	Last_Name        *string `json:"last_name"`
	Uid              string  `json:"uid"`
	jwt.RegisteredClaims
}

func TokenGenerator(user models.User) (token string, refreshToken string) {
	//TODO: change time to check token
	claims := &MyCustomClaims {
		Email:      user.Email,
		First_Name: user.First_Name,
		Last_Name:  user.Last_Name,
		Uid:        user.User_ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(24 * time.Hour)),
		},
	}

	refreshclaims := &MyCustomClaims {
		RegisteredClaims: jwt.RegisteredClaims {
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(168 * time.Hour)),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(types.SECRET_KEY))
	if err != nil {
		log.Println(err.Error())
		return 
	}

	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshclaims).SignedString([]byte(types.SECRET_KEY))
	if err != nil {
		log.Println(err.Error())
		return
	}

	return token, refreshToken
}


func ValidateToken(tokenString string) (claims *MyCustomClaims, msg string) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(types.SECRET_KEY), nil
	}, jwt.WithLeeway(5*time.Second))
	
	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*MyCustomClaims)
	if !ok {
		msg = "The Token is invalid"
		return
	}

	currentTime := time.Now().Local().Unix()
	expirationTime := claims.ExpiresAt.Time.Unix()
	if  expirationTime < currentTime {
		msg = "token is expired"
		return
	}
	return claims, msg
}


func UpdateAllTokens(ctx context.Context, token string, freshToken string, userID string) bool {
	userCollection := db.GetMongoClient().CreateCollection(types.UserCollectionName)
	var updateobj primitive.D
	updateTime := time.Now().Local()

	// E represents a BSON element for a D. It is usually used inside a D
	updateobj = append(updateobj, bson.E{Key: "token", Value: token})
	updateobj = append(updateobj, bson.E{Key: "refresh_token", Value: freshToken})
	updateobj = append(updateobj, bson.E{Key: "updated_at", Value: updateTime})

	filter := bson.M{"user_id": userID}
	// D is a slice and M is a map
	_, err := userCollection.UpdateOne(ctx, filter, bson.D {
		{Key: "$set", Value: updateobj},
	})

	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}
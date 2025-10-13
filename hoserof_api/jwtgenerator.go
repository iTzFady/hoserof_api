package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func jwtGenerator(id string)(string,error){
	fmt.Println("generating jwt")
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading.env file")
		log.Fatal("Error loading .env file")
	}

	// Access environment variables
	var (
  key []byte
  t   *jwt.Token
  s   string
	)

	key = []byte(os.Getenv("JWT_KEY"))
	t = jwt.NewWithClaims(jwt.SigningMethodHS256, 
  	jwt.MapClaims{ 
    	"user_ID":id,
  })
s ,err= t.SignedString(key)
if err!=nil{
	fmt.Println("failed to generate token")
	fmt.Println(err)
	return "", errors.New("failed to generate token")
} 
fmt.Println(s)
return s,nil
}
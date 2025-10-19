package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func withTokenEndPoint(c *gin.Context){
	studentToken:=c.GetHeader("Authorization")
	if studentToken==""{
		c.AbortWithStatusJSON(401, gin.H{"error": "no token"})
	}
	token, err := jwt.Parse(studentToken, func(token *jwt.Token) (interface{}, error) {
            return key, nil
        })

        if err != nil || !token.Valid {
            c.AbortWithStatusJSON(401, gin.H{"error": "invalid token"})
        }
}

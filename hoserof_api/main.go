package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct{
	StudentId string `json:"user_ID"`
	StudentPassword string `json:"user_password"`
}
type UserInFirestore struct{
	FirestoreStudentId string `firestore:"student_ID"`
	FirestoreStudentPassword string `firestore:"student_password"`
}
func assignStudentToStruct(c *gin.Context){
	fmt.Println("request made")
	var user User
	if err:=c.ShouldBindJSON(&user);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"couldn't format json"})
	}
	studentLogin,err:=studentLoginToFirestore(user)
	if err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"error":"user not found"})
	}
	c.JSON(http.StatusOK,gin.H{"student_token":studentLogin})
}

func main(){
	router := gin.Default()
	fmt.Println("working...")
	router.POST("/login",assignStudentToStruct)
	router.Run("localhost:3000")
}
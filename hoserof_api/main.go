package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

type User struct{
	StudentId string `json:"user_ID"`
	StudentPassword string `json:"user_password"`
}
type UserData struct{
	StudentId string `json:"user_ID"`
	StudentName string `json:"student_name"`
	StudentClass string `json:"student_class"`
}
type UserInFirestore struct{
	FirestoreStudentId string `firestore:"student_ID"`
	FirestoreStudentPassword string `firestore:"student_password"`
	FirestoreStudentName string `firestore:"student_name"`
	FirestoreStudentClass string `firestore:"student_class"`
}
func assignStudentToStruct(c *gin.Context){
	fmt.Println("request made")
	var user User
	if err:=c.ShouldBindJSON(&user);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"couldn't format json"})
	}
	studentLogin,userDataptr,err:=studentLoginToFirestore(user)
	if err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"error":"user not found"})
	}
	c.JSON(http.StatusOK,gin.H{"student_token":studentLogin,"student_data":*userDataptr})
}
	var ctx= context.Background()
	var sa = option.WithCredentialsFile("C:\\hoserof_api\\hoserof_api\\hoserof_fb_json.json")
	var app, errfb = firebase.NewApp(ctx, nil, sa)
	var client, errfs = app.Firestore(ctx)
	func checkFirebaseRunning(efb error,efs error){
	if efb != nil {
  		log.Fatalln(efb)
	}
	
	if efs != nil {
  		log.Fatalln(efs)
	}
	}	
func main(){
	checkFirebaseRunning(errfb,errfs)
	defer client.Close()
	router := gin.Default()
	fmt.Println("working...")
	router.POST("/login",assignStudentToStruct)
	router.POST("/signup",signup)
	router.GET("/loginwithtoken",withTokenEndPoint)
	router.Run("localhost:3000")
}
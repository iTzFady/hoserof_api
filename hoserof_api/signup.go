package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)
type NewUser struct{
	NewStudentID string `json:"student_id"`
	NewStudentName string `json:"student_name"`
	NewStudentPhoneNumber uint64 `json:"student_phonenumber"`
	NewStudentPassword []byte `json:"student_password"`
	NewStudentAge uint8 `json:"student_age"`
	NewStudentGrade uint8 `json:"student_grade"`
	NewStudentClass string `json:"student_class"`
}
func signup(c *gin.Context){
	fmt.Println("signup request made")
	var newUser NewUser
	if err:=c.ShouldBindJSON(&newUser);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"couldn't format json"})
	}
	hashedPassword, err := bcrypt.GenerateFromPassword(newUser.NewStudentPassword, bcrypt.DefaultCost)
	if err != nil {
    fmt.Println(err)
	}
	newUser.NewStudentPassword=hashedPassword
	signedUp:=signUpToFirestore(newUser)
	if !signedUp{
		c.JSON(http.StatusBadRequest,gin.H{"error":"failed to sign up"})
	}else {
		c.JSON(http.StatusOK,gin.H{"success":"signed up"})
	}
}
func signUpToFirestore(n NewUser)bool{
	_,err:=client.Collection("students").Doc(n.NewStudentID).Set(ctx,map[string]interface{}{
		"student_name":n.NewStudentName,
		"student_id":n.NewStudentID,
		"student_phonenumber":n.NewStudentPhoneNumber,
		"student_password":n.NewStudentPassword,
		"student_age":n.NewStudentAge,
		"student_grade":n.NewStudentGrade,
		"student_class":n.NewStudentClass,
	})
	if err!=nil{
		fmt.Println("failed to signup to firestore")
	}
	
	return true
}
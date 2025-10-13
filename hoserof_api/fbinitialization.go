package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/option"
)
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func FirebaseInitialization()(*firestore.Client,context.Context){
	fmt.Println("initializing firebase")
	ctx := context.Background()
	sa := option.WithCredentialsFile("C:\\hoserof_api\\hoserof_api\\hoserof_fb_json.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
  		log.Fatalln(err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
  		log.Fatalln(err)
	}
	return client,ctx
}
func studentLoginToFirestore(user User)(string,error){
	fmt.Println("function called")
	client,ctx:=FirebaseInitialization()
	defer client.Close()
	studentsInFirestore:=client.Collection("students")
	studentDocument:=studentsInFirestore.Doc(user.StudentId)
	
	if studentDocument==nil{
		fmt.Println("user not found")
		return "",errors.New("user not found")
	}
	studentDocumentAsMap,err:=studentDocument.Get(ctx)
	if err!=nil{
		fmt.Println("failed to map user")
		return "",errors.New("failed to map user")
	}
	var userInFirestore UserInFirestore
	if err:=studentDocumentAsMap.DataTo(&userInFirestore);err!=nil{
		fmt.Println("failed to assign user to struct")
		return "",errors.New("failed to assign user to struct")
	}
	var (studentToken string; 
		err1 error;)
		fmt.Println(user.StudentPassword,userInFirestore.FirestoreStudentPassword)
	if user.StudentId==userInFirestore.FirestoreStudentId&&CheckPasswordHash(user.StudentPassword,userInFirestore.FirestoreStudentPassword){
		fmt.Println("user verified")
		studentToken,err1=jwtGenerator(user.StudentId)
		if err1!=nil{
			fmt.Println("failed to return token")
			return "",errors.New("failed to return token")
		}
	} else{
		fmt.Println("user access denied. invalid id/password")
	}
	fmt.Println("successfully performed operations")
	return studentToken,nil
}
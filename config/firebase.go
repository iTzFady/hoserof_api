package config

import (
	"context"
	"encoding/base64"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var ctx context.Context
var App *firebase.App
var DB *firestore.Client
var opt option.ClientOption

func InitFirebase() {
	ctx = context.Background()

	encodedCredentials := os.Getenv("FIREBASE_CREDENTIALS_BASE64")
	if encodedCredentials == "" {
		log.Fatal("FIREBASE_CREDENTIALS_BASE64 environment variable not set")
	}

	jsonBytes, err := base64.StdEncoding.DecodeString(encodedCredentials)
	if err != nil {
		log.Fatal("Error decoding FIREBASE_CREDENTIALS_BASE64:", err)
	}

	opt := option.WithCredentialsJSON(jsonBytes)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("Firebase initialization failed: %v", err)
	}
	App = app

	client, err := App.Firestore(ctx)
	if err != nil {
		log.Fatalf("Firestore connection failed: %v", err)
	}
	DB = client

	log.Println("Firebase initialized successfully!")
}

package models

import "github.com/dgrijalva/jwt-go"

type SignUp struct {
	FirstName   string `json:"first_name" bson:"first_name"`
	LastName    string `json:"last_name" bson:"last_name"`
	Username    string `json:"username" bson:"username"`
	Age         int32  `json:"age" bson:"age"`
	PhoneNumber string `json:"phone_number" bson:"phone_number"`
	Password    string `json:"password" bson:"password"`
	Gmail       string `json:"gmail" bson:"gmail"`
}

type Login struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Token    string `json:"token" bson:"token"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type Token struct {
	Token   string `json:"token"`
	Success bool   `json:"success"`
}

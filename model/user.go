package model

import "github.com/dgrijalva/jwt-go"

type User struct {
	ID    string `bson:"_id,omitempty" json:"id"`
	Name  string `bson:"name" json:"name"`
	Email string `bson:"email" json:"email"`
	Role  string `bson:"role" json:"role"`
}

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

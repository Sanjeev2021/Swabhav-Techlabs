package main

import (
	jwt "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secretkey")
var alluser []User
var allContact []Contact
var allContactdetails []ContactDetails

//var flagfordelete int

type User struct {
	UID      string `json:"uid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Fname    string `json:"fname"`
	IsActive string `json:"isActive"`
}
type Claims struct {
	UID      string
	Username string
	Role     string
	Fname    string
	jwt.StandardClaims
}

var adminUsers = map[string]string{
	"user1":    "pass1",
	"user2":    "pass2",
	"user3":    "pass3",
	"yashshah": "hello",
}

type Contact struct {
	CID   string `json:"cid"`
	UID   string `json:"uid"`
	Cname string `json:"Cname"`
}
type ContactDetails struct {
	CDID   string `json:"CDID"`
	CID    string `json:"cid"`
	Type   string `json:"Type"`
	Number string `json:"Number"`
}

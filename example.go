package main

import (
	"fmt"
	"go-jwt/jwt"
	"time"
)

func main() {
	secret := "randpmly generated code"
	type Meta struct {
		Name  string
		Email string
	}

	EncodePayload := jwt.Payload{
		Sub: "123",
		Exp: time.Now().Unix() + 100000,
		Public: Meta{
			Name:  "Murphy",
			Email: "Murphy@jwt.com",
		},
	}

	token := jwt.Encode(EncodePayload, secret)

	// prints out your token
	fmt.Println(token)
	// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjMiLCJleHAiOjE1NDEyNzAyNzAsInB1YmxpYyI6eyJOYW1lIjoiTXVycGh5IiwiRW1haWwiOiJNdXJwaHlAand0LmNvbSJ9fQ==.dkzber79rM7gubpPCaAkjz0gFjxndbMCk6zQWrswkzE=

	fmt.Println(jwt.Decode(token, secret))
	// 123 1541270653   map[Name:Murphy Email:Murphy@jwt.com]} <nil>
}

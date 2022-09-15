package main

import (
	"errors"
	"fmt"
	"github.com/eininst/go-jwt"
	"time"
)

func main() {
	j := jwt.New("m9561kcc59534db1203e6572f7d5dq1x")

	data := map[string]interface{}{
		"id":   123,
		"name": "wzq",
	}

	token := j.CreateToken(data, time.Second*10)
	fmt.Println("Create a token:", token)

	var ps map[string]interface{}

	er := j.ParseToken(token, &ps)
	if errors.Is(er, jwt.Expired) {
		fmt.Println("my is expired")
		return
	}

	fmt.Println("Parse a token:", ps)
}

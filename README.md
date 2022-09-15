# Jwt

[![Build Status](https://travis-ci.org/ivpusic/grpool.svg?branch=master)](https://github.com/infinitasx/easi-go-aws)

`A simple jwt`

## ⚙ Installation

```text
go get -u github.com/eininst/go-jwt
```

## Create a token

```go
package main

import (
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
}

```

## Parse a token

```go
func main() {
var ps map[string]interface{}

er := j.ParseToken(token, &ps)

if errors.Is(er, jwt.Expired) {
fmt.Println("my is expired")
return
}

fmt.Println("Parse a token:", ps)
}
```

> Resullt:

```text
Create a token: FXLRoc+EhSJLLm0ZrhxRwJilwynP3pohKTS4NsKyxIfabXw8qN7auU+WZzDdcplPUGrMyggOmlN9f5EwDj0ZsqiAtTRlMQpeQ78azCDxCq0=
Parse a token: map[id:123 name:wzq]
```

> See [examples](/examples)

## License

*MIT*
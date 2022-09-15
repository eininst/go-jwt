package jwt

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"
)

type Jwt struct {
	SecretKey string
}

type Token struct {
	Data string `json:"data"`
	Exp  int64  `json:"exp"`
}

const Expired = JwtError("jwt: token is expired") // nolint:errname

type JwtError string

func (e JwtError) Error() string { return string(e) }

func (JwtError) RedisError() {}

func New(secretKey string) *Jwt {
	l := len(secretKey)
	if l < 32 {
		f := fmt.Sprintf("Invalid secretKey size %v, Cannot be less than 32", l)
		panic(JwtError(f))
	}
	return &Jwt{SecretKey: secretKey}
}

func (j *Jwt) CreateToken(data interface{}, expire time.Duration) string {
	d, _ := json.Marshal(data)
	b, _ := json.Marshal(&Token{
		Data: string(d),
		Exp:  time.Now().UnixNano() + int64(expire),
	})
	result, err := AesEncrypt(b, []byte(j.SecretKey))
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(result)
}

func (j *Jwt) ParseToken(token string, v interface{}) error {
	b, _ := base64.StdEncoding.DecodeString(token)
	origData, err := AesDecrypt(b, []byte(j.SecretKey))
	if err != nil {
		return err
	}
	var tk Token
	err = json.Unmarshal(origData, &tk)
	if err != nil {
		return err
	}

	if time.Now().UnixNano() > tk.Exp {
		return Expired
	}

	return json.Unmarshal([]byte(tk.Data), &v)
}

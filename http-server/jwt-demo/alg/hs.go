package alg

import (
	"github.com/lestrrat-go/jwx/jwa"
	jwt2 "github.com/lestrrat-go/jwx/jwt"
	"time"
)

func Hs256Encode(sign []byte, claims map[string]interface{}) ([]byte, error) {
	token := jwt2.New()
	token.Set(jwt2.IssuedAtKey, time.Now())
	token.Set(jwt2.ExpirationKey, time.Now().Add(time.Hour*1))
	token.Set(jwt2.NotBeforeKey, time.Now().Add(time.Second*1))
	if claims != nil {
		for k, v := range claims {
			if err := token.Set(k, v); err != nil {
				return nil, err
			}
		}
	}

	if bytes, err := jwt2.Sign(token, jwa.HS256, sign); err != nil {
		return nil, err
	} else {
		return bytes, nil
	}
}

package alg

import (
	jwt2 "github.com/lestrrat-go/jwx/jwt"
	"testing"
	"time"
)

const TIME_FORMAT = "2006-01-02 15:04:05"

func TestHs256Encode(t *testing.T) {

	encode, err := Hs256Encode([]byte("nelson"), nil)
	if err != nil {
		t.Fail()
		return
	}
	t.Logf("encode: %s", encode)
	token, err := jwt2.Parse(encode)
	if err != nil {
		t.Logf("err: %v", err)
	}

	t.Logf("now : %s", time.Now().Format(TIME_FORMAT))
	t.Logf("after 1 hour : %s", time.Now().Add(time.Hour).Format(TIME_FORMAT))
	t.Logf("%s: %s", jwt2.IssuerKey, token.IssuedAt().Local().Format(TIME_FORMAT))
	t.Logf("%s: %s", jwt2.ExpirationKey, token.Expiration().Local().Format(TIME_FORMAT))
	t.Logf("%s: %s", jwt2.NotBeforeKey, token.NotBefore().Local().Format(TIME_FORMAT))
}

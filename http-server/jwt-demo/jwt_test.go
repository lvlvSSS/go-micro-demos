package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jws"
	jwt2 "github.com/lestrrat-go/jwx/jwt"
	"regexp"
	"strings"
	"testing"
)

func TestJwt(t *testing.T) {
	token := jwt.New(jwt.SigningMethodHS384)
	token.Claims = jwt.RegisteredClaims{
		ID:      "1",
		Subject: "第一个主题",
	}
	signedString, _ := token.SignedString([]byte("2233"))
	t.Logf("token : %s", signedString)
	decodeString, err := base64.StdEncoding.DecodeString("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9")
	if err != nil {
		t.Logf("decode err: %v", err)
		return
	}
	t.Logf("decode: %s", decodeString)

	decodeString, err = base64.StdEncoding.DecodeString("eyJzdWIiOiLnrKzkuIDkuKrkuLvpopgiLCJqdGkiOiIxIn0")

	t.Logf("decode: %s", decodeString)
}

func TestJwt1(t *testing.T) {
	token := jwt2.New()
	token.Set(jwt2.IssuerKey, `github.com/lestrrat-go/jwx`)
	token.Expiration()
	token.Set(jwt2.JwtIDKey, "1")
	token.Set(jwt2.AudienceKey, "nelson")
	serialized, _ := jwt2.Sign(token, jwa.HS256, []byte("nelson"))
	t.Logf("token: %s", string(serialized))
	token, _ = jwt2.ParseString("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsibmVsc29uIl0sImlzcyI6ImdpdGh1Yi5jb20vbGVzdHJyYXQtZ28vand4IiwianRpIjoiMSJ9.7miBPHN8rYQ0NzsHAjCl1jbMPamHuehyWhJaal548tA")
	get, _ := token.Get(jwt2.AudienceKey)
	t.Logf("%s: %s", jwt2.AudienceKey, get)
	sign, _ := jws.Sign([]byte("{\"iss\":\"github.com/lestrrat-go/jwx\"}"), jwa.HS256, []byte("nelson"))
	t.Logf("jws sign : %s", sign)

	decodeString, _ := base64.StdEncoding.DecodeString("eyJpc3MiOiJodHRwOi8vc2hhb2Jhb2Jhb2VyLmNuIiwiYXVkIjoiaHR0cDovL3NoYW9iYW9iYW9lci5jbi93ZWJ0ZXN0L2p3dF9hdXRoLyIsImp0aSI6IjRmMWcyM2ExMmFhIiwiaWF0IjoxNTM0MDcwNTQ3LCJuYmYiOjE1MzQwNzA2MDcsImV4cCI6MTUzNDA3NDE0NywidWlkIjoxLCJkYXRhIjp7InVuYW1lIjoic2hhb2JhbyIsInVFbWFpbCI6InNoYW9iYW9iYW9lckAxMjYuY29tIiwidUlEIjoiMHhBMCIsInVHcm91cCI6Imd1ZXN0In19")

	t.Logf("decode: %s", decodeString)

	str := `
{
  "iss": "http://shaobaobaoer.cn",
  "aud": "http://shaobaobaoer.cn/webtest/jwt_auth/",
  "jti": "4f1g23a12aa",
  "iat": 1534070547,
  "nbf": 1534070607,
  "exp": 1534074147,
  "uid": 1,
  "data": {
    "uname": "shaobao",
    "uEmail": "shaobaobaoer@126.com",
    "uID": "0xA0",
    "uGroup": "guest"
  }
}`
	str = strings.Replace(str, "\n\r", "", -1)
	/*	str = strings.Replace(str, "\n", "", -1)
		str = strings.Replace(str, "\r", "", -1)
		str = strings.Replace(str, "\t", "", -1)
		str = strings.Replace(str, " ", "", -1)*/

	compile, _ := regexp.Compile(`\s+`)
	str = compile.ReplaceAllString(str, "")
	t.Logf("str origin: %s", str)
	t.Logf("str: %s", base64.StdEncoding.EncodeToString([]byte(str)))
	hmacSha256 := GenHmacSha256("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ", "123456")
	t.Logf("GenHmacSha256 : %s", hmacSha256)
	t.Logf("GenHmacSha256 encoded : %s", base64.StdEncoding.EncodeToString([]byte(hmacSha256)))

}

func GenHmacSha256(message string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	sha := hex.EncodeToString(h.Sum(nil))
	fmt.Printf("sha:%s\n", sha)
	return Base64UrlSafeEncode(h.Sum(nil))
}

func Base64UrlSafeEncode(source []byte) string {
	byteArr := base64.StdEncoding.EncodeToString(source)
	safeUrl := strings.Replace(string(byteArr), "/", "_", -1)
	safeUrl = strings.Replace(safeUrl, "+", "-", -1)
	safeUrl = strings.Replace(safeUrl, "=", "", -1)
	return safeUrl
}

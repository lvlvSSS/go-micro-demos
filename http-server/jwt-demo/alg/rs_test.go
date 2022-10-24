package alg

import (
	"crypto/rsa"
	"strings"
	"testing"
)

func TestRsEncode(t *testing.T) {
	encode, err := RsEncode()
	if err != nil {
		t.Logf("err : %v", err)
		t.Fail()
		return
	}
	t.Logf("encode: %s", encode)
}

func TestRSA(t *testing.T) {
	reader := strings.NewReader(PRIVATE_KEY)
	key, err := rsa.GenerateKey(reader, 256)
	if err != nil {
		t.Logf("err : %s", err)
		return
	}
	t.Logf("key : %s", key.Public())
}

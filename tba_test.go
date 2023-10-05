package netsuiteodbc

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
	"testing"
	"time"
)

func TestNewTBAConnectionStringer(t *testing.T) {
	cfg := Config{
		ConnectionString: "DSN=NetSuite;",
	}
	got := NewConnectionStringer(cfg)

	if got.err != nil {
		t.Errorf("expected no error on ConnectionStringer, got: %s", got.err)
	}

	if got.config != cfg {
		t.Errorf("config not set on ConnectionStringer")
	}

	if dsn, ok := got.params["DSN"]; !ok || dsn != "NetSuite" {
		t.Errorf("params not set correctly")
	}

	if got.randBytes != rand.Reader {
		t.Errorf("randBytes not set correctly")
	}

}

func newTestTBA() *ConnectionStringer {
	return &ConnectionStringer{
		params: map[string]string{
			"DSN": "NetSuite",
		},
		config: Config{
			ConsumerKey:    "71cc02b731f05895561ef0862d71553a3ac99498a947c3b7beaf4a1e4a29f7c4",
			ConsumerSecret: "7278da58caf07f5c336301a601203d10a58e948efa280f0618e25fcee1ef2abd",
			TokenID:        "89e08d9767c5ac85b374415725567d05b54ecf0960ad2470894a52f741020d82",
			TokenSecret:    "060cd9ab3ffbbe1e3d3918e90165ffd37ab12acc76b4691046e2d29c7d7674c2",
			AccountId:      "1234567",
		},
		now:       func() time.Time { return time.Unix(1439829974, 0) },
		randBytes: hex.NewDecoder(strings.NewReader("ea86cc2aad2d998f3295539d124035")),
	}
}

func TestTBAConnectionStringer_Nonce(t *testing.T) {
	auth := newTestTBA()
	got, err := auth.Nonce()
	if err != nil {
		t.Fatalf("expected no error, got: %s", err.Error())
	}

	if got != "6obMKq0tmY8ylVOdEkA1" {
		t.Errorf("unexpected Nonce, got: %s", got)
	}

}

func TestTBAConnectionStringer_ConnectionString(t *testing.T) {
	auth := newTestTBA()
	c, err := auth.ConnectionString()
	if err != nil {
		t.Fatalf("expected no error from ConnectionString, got: %s", err.Error())
	}
	t.Logf("conn string: %s", c)
	m, err := connStringToParameterMap(c)

	if err != nil {
		t.Fatalf("expected no error from connParams, got: %s", err.Error())
	}
	if uid, ok := m["UID"]; !ok || uid != "TBA" {
		t.Errorf("unexpected TBA, got: %s", uid)
	}

	if dsn, ok := m["DSN"]; !ok || dsn != "NetSuite" {
		t.Errorf("unexpected DSN, got: %s", dsn)
	}

	if pwd, ok := m["PWD"]; !ok || pwd != "1234567&71cc02b731f05895561ef0862d71553a3ac99498a947c3b7beaf4a1e4a29f7c4&89e08d9767c5ac85b374415725567d05b54ecf0960ad2470894a52f741020d82&6obMKq0tmY8ylVOdEkA1&1439829974&FCghIZqXNetuZY8ILWOFH0ucdfzQOmAuL+q+kF21zPs=&HMAC-SHA256" {
		t.Errorf("unexpected PWD, got: %s", pwd)
	}
}

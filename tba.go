package netsuiteodbc

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"time"
)

const nonceLength = 15

// NewTBAConnectionStringer implements unixodbc.ConnectionStringFactory, generating connection strings for connecting
// to a NetSuite ODBC database
func NewTBAConnectionStringer(connStr string, config *TokenConfig) *TBAConnectionStringer {
	t := &TBAConnectionStringer{
		config:    config,
		now:       time.Now,
		randBytes: rand.Reader,
	}
	t.params, t.err = connStringToParameterMap(connStr)
	return t
}

type TBAConnectionStringer struct {
	err       error
	config    *TokenConfig
	params    parameterMap
	now       func() time.Time
	randBytes io.Reader
}

// ConnectionString generates a connection string with a dynamic token used for credentials
func (s *TBAConnectionStringer) ConnectionString() (string, error) {
	if s.err != nil {
		return "", s.err
	}
	token, err := s.Token()
	if err != nil {
		return "", err
	}
	s.params["UID"] = "TBA"
	s.params["PWD"] = token
	return s.params.String(), nil
}

func (s *TBAConnectionStringer) Nonce() (string, error) {
	b := make([]byte, nonceLength)
	if _, err := io.ReadFull(s.randBytes, b); err != nil {
		return "", err
	} else {
		return base64.StdEncoding.EncodeToString(b), nil
	}
}

func (s *TBAConnectionStringer) Token() (string, error) {
	nonce, err := s.Nonce()
	if err != nil {
		return "", err
	}

	base := fmt.Sprintf("%s&%s&%s&%s&%d",
		s.config.AccountId,
		s.config.ConsumerKey,
		s.config.TokenID,
		nonce,
		s.now().Unix())
	secret := []byte(s.config.ConsumerSecret + "&" + s.config.TokenSecret)
	h := hmac.New(sha256.New, secret)
	h.Write([]byte(base))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return base + "&" + signature + "&HMAC-SHA256", nil
}

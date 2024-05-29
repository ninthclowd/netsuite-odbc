package netsuiteodbc

import (
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"time"
)

const nonceLength = 15

// NewConnectionStringer creates a new ConnectionStringer used for connecting to a NetSuite ODBC database with TBA
func NewConnectionStringer(config Config) *ConnectionStringer {
	t := &ConnectionStringer{
		config:    config,
		now:       time.Now,
		randBytes: rand.Reader,
	}
	t.params, t.err = connStringToParameterMap(config.ConnectionString)
	return t
}

// ConnectionStringer implements unixodbc.ConnectionStringFactory, generating connection strings for connecting
// to a NetSuite ODBC database
type ConnectionStringer struct {
	err       error
	config    Config
	params    parameterMap
	now       func() time.Time
	randBytes io.Reader
}

// ConnectionString generates a connection string with a dynamic token used for credentials
func (s *ConnectionStringer) ConnectionString(ctx context.Context) (string, error) {
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

func (s *ConnectionStringer) Nonce() (string, error) {
	b := make([]byte, nonceLength)
	if _, err := io.ReadFull(s.randBytes, b); err != nil {
		return "", err
	} else {
		return base64.StdEncoding.EncodeToString(b), nil
	}
}

func (s *ConnectionStringer) Token() (string, error) {
	nonce, err := s.Nonce()
	if err != nil {
		return "", err
	}

	base := fmt.Sprintf("%s&%s&%s&%s&%d",
		s.config.AccountId,
		s.config.ConsumerKey,
		s.config.TokenId,
		nonce,
		s.now().Unix())
	secret := []byte(s.config.ConsumerSecret + "&" + s.config.TokenSecret)
	h := hmac.New(sha256.New, secret)
	h.Write([]byte(base))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return base + "&" + signature + "&HMAC-SHA256", nil
}

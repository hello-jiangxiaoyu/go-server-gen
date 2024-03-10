package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/go-jose/go-jose/v3"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

const (
	DefaultDir  = "data"
	DefaultKid  = "default-kid"
	DefaultPath = DefaultDir + "/" + DefaultKid + ".key"
)

// SigneTokenString 签名token
func SigneTokenString(claims jwt.Claims) (string, error) {
	pemString, err := getJWK()
	if err != nil {
		return "", err
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(pemString)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token.Header["kid"] = DefaultKid
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// GetJWKPublic 获取jwk公钥
func GetJWKPublic() (*jose.JSONWebKeySet, error) {
	pemString, err := getJWK()
	if err != nil {
		if pemString, err = GenerateKey(); err != nil {
			return nil, err
		}
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(pemString)
	if err != nil {
		return nil, err
	}

	jwk := jose.JSONWebKey{
		Key:                       key.Public(),
		KeyID:                     DefaultKid,
		Algorithm:                 "RS256",
		Use:                       "sig",
		Certificates:              []*x509.Certificate{},
		CertificateThumbprintSHA1: []uint8{},
	}
	var jwkSet jose.JSONWebKeySet
	jwkSet.Keys = append(jwkSet.Keys, jwk) // 只生成一个jwk
	return &jwkSet, nil
}

// GetJWKPrivate 获取私钥
func GetJWKPrivate() (*rsa.PrivateKey, error) {
	pemString, err := getJWK()
	if err != nil {
		return nil, err
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(pemString)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// GenerateKey 生成jwk key并写入文件
func GenerateKey() ([]byte, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}
	payload := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	if err = setJWK(payload); err != nil {
		return nil, err
	}

	return payload, err
}

// getJWK 获取jwk
func getJWK() ([]byte, error) {
	pemString, err := os.ReadFile(DefaultPath)
	if err != nil {
		return nil, err
	}

	return pemString, nil
}

// setJWK 保存jwk私钥
func setJWK(payload []byte) error {
	if _, err := os.ReadDir(DefaultDir); err != nil {
		if err = os.MkdirAll(DefaultDir, 0700); err != nil {
			return err
		}
	}

	if err := os.WriteFile(DefaultPath, payload, 0400); err != nil {
		return err
	}
	return nil
}

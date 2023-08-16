package pkg

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
)

type KeyLoader interface {
	LoadPrivateKey() (*ecdsa.PrivateKey, error)
	LoadPublicKey() (*ecdsa.PublicKey, error)
}

type TokenGenerator interface {
	GenerateToken(privateKey *ecdsa.PrivateKey, claims jwt.Claims) (string, error)
	VerifyToken(tokenString string, publicKey *ecdsa.PublicKey) (*jwt.Token, error)
}

type KeyLoaderImpl struct {
	PrivateKeyPath string
	PublicKeyPath  string
}

func (k *KeyLoaderImpl) LoadPrivateKey() (*ecdsa.PrivateKey, error) {
	keyData, err := os.ReadFile(k.PrivateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key file: %v", err)
	}

	block, _ := pem.Decode(keyData)
	if block == nil || block.Type != "EC PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing EC private key")
	}

	key, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func (k *KeyLoaderImpl) LoadPublicKey() (*ecdsa.PublicKey, error) {
	keyData, err := os.ReadFile(k.PublicKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read public key file: %v", err)
	}

	block, _ := pem.Decode(keyData)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing public key")
	}

	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	ecdsaKey, ok := key.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("parsed key is not an ECDSA public key")
	}

	return ecdsaKey, nil
}

type TokenGeneratorImpl struct{}

func (tg *TokenGeneratorImpl) GenerateToken(privateKey *ecdsa.PrivateKey, claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(privateKey)
}

func (tg *TokenGeneratorImpl) VerifyToken(tokenString string, publicKey *ecdsa.PublicKey) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
}

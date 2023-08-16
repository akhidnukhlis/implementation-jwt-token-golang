package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"playground/implementation-jwt-token-golang/config"
	"playground/implementation-jwt-token-golang/pkg"
	"time"
)

func main() {
	// Read file
	file, err := config.ReadConfigFromFile()
	if err != nil {
		fmt.Println("Failed to read file:", err)
		return
	}

	tokenGenerator := &pkg.TokenGeneratorImpl{}

	// Create instances of the refactored components
	keyLoader := &pkg.KeyLoaderImpl{
		PrivateKeyPath: file.PrivateKey,
		PublicKeyPath:  file.PublicKey,
	}

	// Load private and public keys
	privateKey, err := keyLoader.LoadPrivateKey()
	if err != nil {
		fmt.Println("Error loading private key:", err)
		return
	}

	publicKey, err := keyLoader.LoadPublicKey()
	if err != nil {
		fmt.Println("Error loading public key:", err)
		return
	}

	// Set token claims
	claims := jwt.StandardClaims{
		Id:        "ABCDE",
		Subject:   "TEST-TOKEN",
		Issuer:    "https://github.com/akhidnukhlis",
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(), // Token will expire in 1 hour
	}

	// Generate token
	tokenString, err := tokenGenerator.GenerateToken(privateKey, claims)
	if err != nil {
		fmt.Println("Error generate token:", err)
		return
	}

	// Verify token
	parsedToken, err := tokenGenerator.VerifyToken(tokenString, publicKey)
	if err != nil {
		fmt.Println("Error verifying token:", err)
		return
	}

	fmt.Println("######################")
	fmt.Println("Token:", tokenString)
	fmt.Println("Token verified:", parsedToken.Valid)
	fmt.Println("######################")
}

package authen

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const (
	accessTokenDuration  = 15 * time.Minute
	refreshTokenDuration = 7 * 24 * time.Hour
)

type (
	Verifier interface {
		Verify(string) (*Claims, error)
	}

	Signer interface {
		Sign(jwt.Claims) (string, error)
	}

	Claims struct {
		UserID   string `json:"user_id"`
		Username string `json:"username"`
		jwt.RegisteredClaims
	}
)

type TokenGenerator struct {
	signingMethod jwt.SigningMethod

	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func NewTokenGenerator(privateKeyPath, publicKeyPath string) (*TokenGenerator, error) {
	privateKeyBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("could not read private key: %w", err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("could not parse private key: %w", err)
	}

	publicKeyBytes, err := os.ReadFile(publicKeyPath)
	if err != nil {
		return nil, fmt.Errorf("could not read public key: %w", err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("could not parse public key: %w", err)
	}

	return &TokenGenerator{
		privateKey: privateKey,
		publicKey:  publicKey,

		signingMethod: jwt.SigningMethodRS256,
	}, nil
}

func (g *TokenGenerator) GenerateTokenPair(userID string, username string) (string, string, error) {
	accessClaims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	accessTokenString, err := g.Sign(accessClaims)
	if err != nil {
		return "", "", err
	}

	refreshClaims := jwt.RegisteredClaims{
		ID:        uuid.NewString(),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshTokenDuration)),
	}
	refreshTokenString, err := g.Sign(refreshClaims)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

func (g *TokenGenerator) ValidateToken(token string) error {
	if _, err := g.Verify(token); err != nil {
		return err
	}
	return nil
}

func (g *TokenGenerator) Sign(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(g.signingMethod, claims)
	signedToken, err := token.SignedString(g.privateKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (g *TokenGenerator) Verify(s string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(s, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if token.Method != g.signingMethod {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return g.publicKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

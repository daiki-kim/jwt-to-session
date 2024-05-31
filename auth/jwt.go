package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

const (
	Subject                = "AccessToken"
	Issuer                 = "https://idp.alamoa.io"
	Audience               = "https://api.alamoa.io"
	TokenExpiration        = time.Minute * time.Duration(1)
	RefreshTokenExpiration = time.Hour * time.Duration(1)
)

var (
	tokenSignKey          = []byte("secretkey")
	TokenVerifyKey        = []byte("secretkey")
	refreshTokenSignKey   = []byte("refreshTokenSecretkey")
	RefreshTokenVerifyKey = []byte("refreshTokenSecretkey")
)

type Claim struct {
	Email string
	jwt.StandardClaims
}

func NewClaim(email string) *Claim {
	return &Claim{
		Email: email,
	}
}

func (c *Claim) GenerateToken() (token string, err error) {
	claims := jwt.StandardClaims{
		Id:        uuid.New().String(),
		Issuer:    Issuer,
		Subject:   Subject,
		Audience:  Audience,
		IssuedAt:  time.Now().Local().Unix(),
		ExpiresAt: time.Now().Local().Add(TokenExpiration).Unix(),
	}
	if err = copier.CopyWithOption(c, claims, copier.Option{IgnoreEmpty: true, DeepCopy: true}); err != nil {
		return "", err
	}

	generatedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	token, err = generatedToken.SignedString(tokenSignKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (c *Claim) GenerateRefreshToken() (token string, err error) {
	claims := jwt.StandardClaims{
		Id:        uuid.New().String(),
		Issuer:    Issuer,
		Subject:   Subject,
		Audience:  Audience,
		IssuedAt:  time.Now().Local().Unix(),
		ExpiresAt: time.Now().Local().Add(RefreshTokenExpiration).Unix(),
	}
	if err = copier.CopyWithOption(c, claims, copier.Option{IgnoreEmpty: true, DeepCopy: true}); err != nil {
		return "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	token, err = refreshToken.SignedString(refreshTokenSignKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (c *Claim) UpdateRefreshToken() (token string, err error) {
	claims := jwt.StandardClaims{
		Issuer:   Issuer,
		Subject:  Subject,
		Audience: Audience,
	}
	if err = copier.CopyWithOption(c, claims, copier.Option{IgnoreEmpty: true, DeepCopy: true}); err != nil {
		return "", err
	}
	c.Id = uuid.New().String()
	c.IssuedAt = time.Now().Local().Unix()
	updatedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	token, err = updatedToken.SignedString(refreshTokenSignKey)
	if err != nil {
		return token, err
	}

	return token, nil
}

func ValidateToken(token string, verifyKey []byte) (*Claim, error) {
	parsedToken, err := jwt.ParseWithClaims(
		token,
		&Claim{},
		func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(*Claim)
	if !ok {
		err = errors.New("could not parse claims")
		return nil, err
	}

	return claims, nil
}

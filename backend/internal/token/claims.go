package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserClaims struct {
	ID primitive.ObjectID `json:"id"`
	Email string `json:"email"`
	Photo *string `json:"photo"`
	jwt.RegisteredClaims
}

func NewUserClaims(id primitive.ObjectID, email string, duration time.Duration) (*UserClaims, error) {
	tokenID, err := uuid.NewRandom()

	if err != nil {
		return nil, fmt.Errorf("error generating token ID: %w", err)
	}

	return &UserClaims{
		Email: email,
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ID: tokenID.String(),
			Subject: email,
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}, nil
}
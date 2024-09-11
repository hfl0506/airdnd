package handler

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GenericRes[T any] struct {
	Success   bool        `json:"success"`
	Data      interface{} `json:"data"`
	Total     *int64      `json:"total"`
	NextPage  *int16      `json:"nextPage"`
	TotalPage *int16      `json:"totalPage"`
}

type ListingAndReviewsRecord struct {
	ID         string               `json:"id"`
	ListingUrl string               `json:"listingUrl"`
	Name       string               `json:"name"`
	Summary    string               `json:"summary"`
	Images     string               `json:"images"`
	Address    string               `json:"address"`
	Bedrooms   int64                `json:"bedrooms"`
	Beds       int64                `json:"beds"`
	RoomType   string               `json:"roomType"`
	BedType    string               `json:"bedType"`
	Price      primitive.Decimal128 `json:"price"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterReq struct {
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Photo    *string `json:"photo"`
}

type ListCategoryRes struct {
	RoomType     []string `json:"roomType"`
	PropertyType []string `json:"propertyType"`
	BedType      []string `json:"bedType"`
}

type AuthUserRes struct {
	SessionID             string    `json:"sessionId"`
	AccessToken           string    `json:"accessToken"`
	RefreshToken          string    `json:"refreshToken"`
	AccessTokenExpiresAt  time.Time `json:"accessTokenExpiresAt"`
	RefreshTokenExpiresAt time.Time `json:"refreshTokenExpiresAt"`
	User                  UserRes   `json:"user"`
}

type UserRes struct {
	Name  string  `json:"name"`
	Email string  `json:"email"`
	Photo *string `json:"photo"`
}

type CheckInDetailRes struct {
	Adult  int64 `json:"adult"`
	Child  int64 `json:"child"`
	Infant int64 `json:"infant"`
}

type BookingRes struct {
	ID        primitive.ObjectID `json:"id"`
	RoomID    string             `json:"roomId"`
	CheckIn   CheckInDetailRes   `json:"checkIn"`
	StartDate time.Time          `json:"startDate"`
	EndDate   time.Time          `json:"endDate"`
	CreatedAt time.Time          `json:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt"`
}

type BookingReq struct {
	RoomID    string           `json:"roomId"`
	CheckIn   CheckInDetailRes `json:"checkIn"`
	StartDate string           `json:"startDate"`
	EndDate   string           `json:"endDate"`
	Price     float64          `json:"price"`
}

type RefreshReq struct {
	RefreshToken string `json:"refreshToken"`
}

type RefreshRes struct {
	AccessToken          string    `json:"accessToken"`
	AccessTokenExpiresAt time.Time `json:"accessTokenExpiresAt"`
}

type WishListReq struct {
	RoomID string `json:"roomId"`
	Liking bool   `json:"liking"`
}
type RoomRes struct {
	ID           string               `json:"id"`
	Name         string               `json:"name"`
	Summary      string               `json:"summary"`
	Space        string               `json:"space"`
	Description  string               `json:"description"`
	PropertyType string               `json:"propertyType"`
	RoomType     string               `json:"roomType"`
	BedType      string               `json:"bedType"`
	Accommodates int64                `json:"accommodates"`
	Bedrooms     int64                `json:"bedrooms"`
	Beds         int64                `json:"beds"`
	Bathrooms    primitive.Decimal128 `json:"bathrooms"`
	Amenities    []string             `json:"amenities"`
	Price        primitive.Decimal128 `json:"price"`
	Image        string               `json:"image"`
	Address      string               `json:"address"`
}

type ReviewDetail struct {
	ID        primitive.ObjectID `json:"id"`
	User      UserRes            `json:"user"`
	Comment   string             `json:"comment"`
	Rating    int32              `json:"rating"`
	CreatedAt time.Time          `json:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt"`
}

type RoomWithReviewsRes struct {
	ID           string               `json:"id"`
	Name         string               `json:"name"`
	Summary      string               `json:"summary"`
	Space        string               `json:"space"`
	Description  string               `json:"description"`
	PropertyType string               `json:"propertyType"`
	RoomType     string               `json:"roomType"`
	BedType      string               `json:"bedType"`
	Accommodates int64                `json:"accommodates"`
	Bedrooms     int64                `json:"bedrooms"`
	Beds         int64                `json:"beds"`
	Bathrooms    primitive.Decimal128 `json:"bathrooms"`
	Amenities    []string             `json:"amenities"`
	Price        primitive.Decimal128 `json:"price"`
	Image        string               `json:"image"`
	Address      string               `json:"address"`
	Reviews      []ReviewDetail       `json:"reviews"`
}

type WishListItemRes struct {
	ID        primitive.ObjectID `json:"id"`
	Room      RoomRes            `json:"room"`
	CreatedAt time.Time          `json:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt"`
}

type CreateReviewReq struct {
	RoomID  string `json:"roomId"`
	Comment string `json:"comment"`
	Rating  int32  `json:"rating"`
}

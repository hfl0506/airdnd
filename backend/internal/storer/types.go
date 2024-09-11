package storer

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// listing and reviews model
type ListingAndReviewsImage struct {
	ThumbnailUrl string `bson:"thumbnail_url"`
	MediumUrl    string `bson:"medium_url"`
	PictureUrl   string `bson:"picture_url"`
	XlPictureUrl string `bson:"xl_picture_url"`
}

type ListingAndReviewsHost struct {
	HostId                 string   `bson:"host_id"`
	HostUrl                string   `bson:"host_url"`
	HostName               string   `bson:"host_name"`
	HostLocation           string   `bson:"host_location"`
	HostAbout              string   `bson:"host_about"`
	HostThumbnailUrl       string   `bson:"host_thumbnail_url"`
	HostPictureUrl         string   `bson:"host_picture_url"`
	HostNeighbourhood      string   `bson:"host_neighbourhood"`
	HostIsSuperhost        bool     `bson:"host_is_superhost"`
	HostHasProfilePic      bool     `bson:"host_has_profile_pic"`
	HostIdentityVerified   bool     `bson:"host_identity_verified"`
	HostListingsCount      int64    `bson:"host_listings_count"`
	HostTotalListingsCount int64    `bson:"host_total_listings_count"`
	HostVerifications      []string `bson:"host_verifications"`
}

type AddressLocation struct {
	Type            string    `bson:"type"`
	Coordinates     []float64 `bson:"coordinates"`
	IsLocationExact bool      `bson:"is_location_exact"`
}

type ListingAndReviewsAddress struct {
	Street         string          `bson:"street"`
	Suburb         string          `bson:"suburb"`
	GovernmentArea string          `bson:"government_area"`
	Market         string          `bson:"market"`
	Country        string          `bson:"country"`
	CountryCode    string          `bson:"country_code"`
	Location       AddressLocation `bson:"location"`
}

type ListingAndReviewsAvailability struct {
	Availability30  int64 `bson:"availability_30"`
	Availability60  int64 `bson:"availability_60"`
	Availability90  int64 `bson:"availability_90"`
	Availability365 int64 `bson:"availability_365"`
}

type ListingAndReviewsScores struct {
	ReviewScoresAccuracy      int64 `bson:"review_scores_accuracy"`
	ReviewScoresCleanliness   int64 `bson:"review_scores_cleanliness"`
	ReviewScoresCheckin       int64 `bson:"review_scores_checkin"`
	ReviewScoresCommunication int64 `bson:"review_scores_communication"`
	ReviewScoresLocation      int64 `bson:"review_scores_location"`
	ReviewScoresValue         int64 `bson:"review_scores_value"`
	ReviewScoresRating        int64 `bson:"review_scores_rating"`
}

type ReviewsRecord struct {
	ID           string    `bson:"_id"`
	Date         time.Time `bson:"date"`
	ListingID    string    `bson:"listing_id"`
	ReviewerID   string    `bson:"reviewer_id"`
	ReviewerName string    `bson:"reviewer_name"`
	Comments     string    `bson:"comments"`
}

type ListingAndReviews struct {
	ID                   string                        `bson:"_id,omitempty"`
	ListingUrl           string                        `bson:"listing_url"`
	Name                 string                        `bson:"name"`
	Summary              string                        `bson:"summary"`
	Space                string                        `bson:"space"`
	Description          string                        `bson:"description"`
	NeighborhoodOverview string                        `bson:"neighborhood_overview"`
	Notes                string                        `bson:"notes"`
	Transit              string                        `bson:"transit"`
	Access               string                        `bson:"access"`
	Interaction          string                        `bson:"interaction"`
	HouseRules           string                        `bson:"house_rules"`
	PropertyType         string                        `bson:"property_type"`
	RoomType             string                        `bson:"room_type"`
	BedType              string                        `bson:"bed_type"`
	MinimumNights        string                        `bson:"minimum_nights"`
	MaximumNights        string                        `bson:"maximum_nights"`
	CancellationPolicy   string                        `bson:"cancellation_policy"`
	LastScraped          time.Time                     `bson:"last_scraped"`
	CalendarLastScraped  time.Time                     `bson:"calendar_last_scraped"`
	Accommodates         int64                         `bson:"accommodates"`
	Bedrooms             int64                         `bson:"bedrooms"`
	Beds                 int64                         `bson:"beds"`
	NumberOfReviews      int64                         `bson:"number_of_reviews"`
	Bathrooms            primitive.Decimal128          `bson:"bathrooms"`
	Amenities            []string                      `bson:"amenities"`
	Price                primitive.Decimal128          `bson:"price"`
	ExtraPeople          primitive.Decimal128          `bson:"extra_people"`
	GuestsIncluded       primitive.Decimal128          `bson:"guests_included"`
	Images               ListingAndReviewsImage        `bson:"images"`
	Host                 ListingAndReviewsHost         `bson:"host"`
	Address              ListingAndReviewsAddress      `bson:"address"`
	Availability         ListingAndReviewsAvailability `bson:"availability"`
	ReviewScores         ListingAndReviewsScores       `bson:"review_scores"`
	Reviews              []ReviewsRecord               `bson:"reviews"`
}

type ListingPaginateRes struct {
	Data  []ListingAndReviews
	Total int64
}

type ListCategoryRes struct {
	RoomType     []string `bson:"roomType"`
	PropertyType []string `bson:"propertyType"`
	BedType      []string `bson:"bedType"`
}

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	Photo     *string            `bson:"photo"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type Session struct {
	ID           string    `bson:"id"`
	UserEmail    string    `bson:"user_email"`
	RefreshToken string    `bson:"refresh_token"`
	IsRevoked    bool      `bson:"is_revoked"`
	CreatedAt    time.Time `bson:"created_at"`
	ExpiresAt    time.Time `bson:"expires_at"`
}

type CheckInDetail struct {
	Adult  int64 `bson:"adult"`
	Child  int64 `bson:"child"`
	Infant int64 `bson:"infant"`
}
type Booking struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id"`
	RoomID    string             `bson:"room_id"`
	CheckIn   CheckInDetail      `bson:"check_in"`
	StartDate time.Time          `bson:"start_date"`
	Price     float64            `bson:"price"`
	EndDate   time.Time          `bson:"end_date"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type CheckInDetailRes struct {
	Adult  int64 `json:"adult"`
	Child  int64 `json:"child"`
	Infant int64 `json:"infant"`
}

type RoomRes struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type ListBookingRes struct {
	ID        primitive.ObjectID `json:"id"`
	Room      RoomRes            `json:"room"`
	CheckIn   CheckInDetailRes   `json:"checkIn"`
	Price     float64            `json:"price"`
	StartDate time.Time          `json:"startDate"`
	EndDate   time.Time          `json:"endDate"`
	CreatedAt time.Time          `json:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt"`
}

type DeleteBookingReq struct {
	ID     primitive.ObjectID `bson:"_id"`
	UserID primitive.ObjectID `bson:"user_id"`
}

type WishList struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id"`
	RoomID    string             `bson:"room_id"`
	Like      bool               `bson:"like"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type WishListItem struct {
	ID        primitive.ObjectID `bson:"_id"`
	Room      Room               `bson:"room"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type WishListPayload struct {
	RoomID string             `bson:"roomId"`
	Liking bool               `bson:"liking"`
	UserID primitive.ObjectID `bson:"user_id"`
}

type RoomID struct {
	RoomID string `bson:"roomId"`
}

type Room struct {
	ID           string               `bson:"_id"`
	Name         string               `bson:"name"`
	Summary      string               `bson:"summary"`
	Space        string               `bson:"space"`
	Description  string               `bson:"description"`
	PropertyType string               `bson:"property_type"`
	RoomType     string               `bson:"room_type"`
	BedType      string               `bson:"bed_type"`
	Accommodates int64                `bson:"accommodates"`
	Bedrooms     int64                `bson:"bedrooms"`
	Beds         int64                `bson:"beds"`
	Bathrooms    primitive.Decimal128 `bson:"bathrooms"`
	Amenities    []string             `bson:"amenities"`
	Price        primitive.Decimal128 `bson:"price"`
	Image        string               `bson:"image"`
	Address      string               `bson:"address"`
}

type ReviewUser struct {
	Name  string  `bson:"name"`
	Email string  `bson:"email"`
	Photo *string `bson:"photo"`
}

type ReviewDetail struct {
	ID        primitive.ObjectID `bson:"_id"`
	User      ReviewUser         `bson:"user"`
	Comment   string             `bson:"comment"`
	Rating    int32              `bson:"rating"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type RoomWithReviews struct {
	ID           string               `bson:"_id"`
	Name         string               `bson:"name"`
	Summary      string               `bson:"summary"`
	Space        string               `bson:"space"`
	Description  string               `bson:"description"`
	PropertyType string               `bson:"property_type"`
	RoomType     string               `bson:"room_type"`
	BedType      string               `bson:"bed_type"`
	Accommodates int64                `bson:"accommodates"`
	Bedrooms     int64                `bson:"bedrooms"`
	Beds         int64                `bson:"beds"`
	Bathrooms    primitive.Decimal128 `bson:"bathrooms"`
	Amenities    []string             `bson:"amenities"`
	Price        primitive.Decimal128 `bson:"price"`
	Image        string               `bson:"image"`
	Address      string               `bson:"address"`
	Reviews      []ReviewDetail       `bson:"reviews"`
}

type Review struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	RoomID    string             `bson:"room_id"`
	UserID    primitive.ObjectID `bson:"user_id"`
	Comment   string             `bson:"comment"`
	Rating    int32              `bson:"rating"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type DeleteReviewReq struct {
	ID     primitive.ObjectID `bson:"_id"`
	UserID primitive.ObjectID `bson:"user_id"`
}

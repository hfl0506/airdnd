package server

import (
	"backend/internal/storer"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	storer *storer.MongoStorer
}

func NewServer(storer *storer.MongoStorer) *Server {
	return &Server{
		storer: storer,
	}
}

func (s *Server) ListAllListingAndReviews(ctx context.Context, p storer.PaginateParams) (storer.ListingPaginateRes, error) {
	return s.storer.ListingAndReviewsRecords(ctx, p)
}

func (s *Server) ListCategory(ctx context.Context) (storer.ListCategoryRes, error) {
	return s.storer.ListCategory(ctx)
}

func (s *Server) CreateUser(ctx context.Context, user *storer.User) (*mongo.InsertOneResult, error) {
	return s.storer.CreateUser(ctx, user)
}

func (s *Server) GetUserByEmail(ctx context.Context, email string) (*storer.User, error) {
	return s.storer.GetUserByEmail(ctx, email)
}

func (s *Server) CreateSession(ctx context.Context, session *storer.Session) (*storer.Session, error) {
	return s.storer.CreateSession(ctx, session)
}

func (s *Server) GetSession(ctx context.Context, id string) (*storer.Session, error) {
	return s.storer.GetSession(ctx, id)
}

func (s *Server) DeleteSession(ctx context.Context, id string) error {
	return s.storer.DeleteSession(ctx, id)
}

func (s *Server) CreateBooking(ctx context.Context, booking *storer.Booking) (*storer.Booking, error) {
	return s.storer.CreateBooking(ctx, booking)
}

func (s *Server) ListBookings(ctx context.Context, userId primitive.ObjectID) ([]storer.ListBookingRes, error) {
	return s.storer.ListBookings(ctx, userId)
}

func (s *Server) DeleteBooking(ctx context.Context, params storer.DeleteBookingReq) error {
	return s.storer.DeleteBooking(ctx, params)
}

func (s *Server) UpsertWishList(ctx context.Context, params storer.WishListPayload) error {
	return s.storer.UpsertWishList(ctx, params)
}

func (s *Server) ListWishListIds(ctx context.Context, userId primitive.ObjectID) ([]string, error) {
	return s.storer.ListWishListIds(ctx, userId)
}

func (s *Server) ListWishListItems(ctx context.Context, userId primitive.ObjectID) ([]storer.WishListItem, error) {
	return s.storer.GetWishListItems(ctx, userId)
}

func (s *Server) GetRoomById(ctx context.Context, id string) (*storer.RoomWithReviews, error) {
	return s.storer.GetRoomById(ctx, id)
}

func (s *Server) CreateReview(ctx context.Context, payload storer.Review) error {
	return s.storer.CreateReview(ctx, payload)
}

func (s *Server) DeleteReview(ctx context.Context, payload storer.DeleteReviewReq) error {
	return s.storer.DeleteReview(ctx, payload)
}

package handler

import (
	"backend/internal/server"
	"backend/internal/storer"
	"backend/internal/token"
	"backend/internal/util"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type handler struct {
	ctx        context.Context
	server     *server.Server
	tokenMaker *token.JWTMaker
}

func NewHandler(server *server.Server, secretKey string) *handler {
	return &handler{
		ctx:        context.TODO(),
		server:     server,
		tokenMaker: token.NewJWTMaker(secretKey),
	}
}

func (h *handler) listAllListingAndReviews(w http.ResponseWriter, r *http.Request) {
	p := storer.PaginateParams{}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))

	p.Page = page
	if err != nil {
		p.Page = 1
	}

	nextPage := int16(p.Page + 1)

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))

	p.Limit = limit
	if err != nil {
		p.Limit = 10
	}

	tabID := r.URL.Query().Get("tab_id")

	if tabID != "" {
		p.TabID = tabID
	}

	listings, err := h.server.ListAllListingAndReviews(h.ctx, p)

	if err != nil {
		http.Error(w, "error list listings and reviews", http.StatusInternalServerError)
		return
	}

	totalPage := int16(math.Ceil(float64(listings.Total) / float64(p.Limit)))

	var res []ListingAndReviewsRecord

	for _, l := range listings.Data {
		res = append(res, toListingAndReviewsRes(l))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(GenericRes[[]ListingAndReviewsRecord]{
		Data:      res,
		Total:     &listings.Total,
		Success:   true,
		NextPage:  &nextPage,
		TotalPage: &totalPage,
	})
}

func (h *handler) listCategory(w http.ResponseWriter, r *http.Request) {
	categories, err := h.server.ListCategory(h.ctx)

	if err != nil {
		log.Println("err", err)
		http.Error(w, "error list category", http.StatusInternalServerError)
		return
	}

	res := toListCategoryRes(categories)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(GenericRes[ListCategoryRes]{
		Data:    res,
		Success: true,
	})
}

func toListingAndReviewsRes(l storer.ListingAndReviews) ListingAndReviewsRecord {
	return ListingAndReviewsRecord{
		ID:         l.ID,
		ListingUrl: l.ListingUrl,
		Name:       l.Name,
		Summary:    l.Summary,
		Images:     l.Images.PictureUrl,
		Address:    l.Address.Street,
		Bedrooms:   l.Bedrooms,
		Beds:       l.Beds,
		RoomType:   l.RoomType,
		BedType:    l.BedType,
		Price:      l.Price,
	}
}

func toListCategoryRes(c storer.ListCategoryRes) ListCategoryRes {
	return ListCategoryRes{
		RoomType:     c.RoomType,
		PropertyType: c.PropertyType,
		BedType:      c.BedType,
	}
}

func (h *handler) login(w http.ResponseWriter, r *http.Request) {
	var l LoginReq

	if err := json.NewDecoder(r.Body).Decode(&l); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	found, err := h.server.GetUserByEmail(h.ctx, l.Email)

	if found == nil {
		http.Error(w, "user email not found", http.StatusInternalServerError)
		return
	}

	if err != nil {
		http.Error(w, "get user by email failed", http.StatusInternalServerError)
		return
	}

	err = util.CheckPassword(l.Password, found.Password)

	if err != nil {
		http.Error(w, "email or password invalid", http.StatusUnauthorized)
		return
	}

	accessToken, accessClaims, err := h.tokenMaker.CreateToken(found.ID, found.Email, ACCESS_TOKEN_DURATION)

	if err != nil {
		http.Error(w, "error creating token", http.StatusInternalServerError)
		return
	}

	refreshToken, refreshClaims, err := h.tokenMaker.CreateToken(found.ID, found.Email, REFRESH_TOKEN_DURATION)

	if err != nil {
		http.Error(w, "error creating token", http.StatusInternalServerError)
		return
	}

	session, err := h.server.CreateSession(h.ctx, &storer.Session{
		ID:           refreshClaims.RegisteredClaims.ID,
		UserEmail:    found.Email,
		RefreshToken: refreshToken,
		IsRevoked:    false,
		ExpiresAt:    refreshClaims.RegisteredClaims.ExpiresAt.Time,
		CreatedAt:    time.Now(),
	})

	if err != nil {
		http.Error(w, "error creating session", http.StatusInternalServerError)
		return
	}

	res := AuthUserRes{
		SessionID:             session.ID,
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  accessClaims.RegisteredClaims.ExpiresAt.Time,
		RefreshTokenExpiresAt: refreshClaims.RegisteredClaims.ExpiresAt.Time,
		User:                  toUserRes(found),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(GenericRes[AuthUserRes]{
		Data:    res,
		Success: true,
	})
}

func (h *handler) register(w http.ResponseWriter, r *http.Request) {
	var reg RegisterReq

	if err := json.NewDecoder(r.Body).Decode(&reg); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	found, err := h.server.GetUserByEmail(h.ctx, reg.Email)

	if found != nil {
		http.Error(w, "user email already in used", http.StatusInternalServerError)
		return
	}

	if err != nil {
		http.Error(w, "get user by email failed", http.StatusInternalServerError)
		return
	}

	hashed, err := util.HashPassword(reg.Password)

	if err != nil {
		http.Error(w, "hash password failed", http.StatusInternalServerError)
		return
	}

	reg.Password = hashed

	created := toStorerUser(reg)

	_, err = h.server.CreateUser(h.ctx, created)

	if err != nil {
		http.Error(w, "create user failed", http.StatusInternalServerError)
		return
	}

	accessToken, accessClaims, err := h.tokenMaker.CreateToken(created.ID, created.Email, ACCESS_TOKEN_DURATION)

	if err != nil {
		http.Error(w, "error creating token", http.StatusInternalServerError)
		return
	}

	refreshToken, refreshClaims, err := h.tokenMaker.CreateToken(created.ID, created.Email, REFRESH_TOKEN_DURATION)

	if err != nil {
		http.Error(w, "error creating token", http.StatusInternalServerError)
		return
	}

	session, err := h.server.CreateSession(h.ctx, &storer.Session{
		ID:           refreshClaims.RegisteredClaims.ID,
		UserEmail:    created.Email,
		RefreshToken: refreshToken,
		IsRevoked:    false,
		ExpiresAt:    refreshClaims.RegisteredClaims.ExpiresAt.Time,
		CreatedAt:    time.Now(),
	})

	if err != nil {
		http.Error(w, "error creating session", http.StatusInternalServerError)
		return
	}

	res := AuthUserRes{
		SessionID:             session.ID,
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  accessClaims.RegisteredClaims.ExpiresAt.Time,
		RefreshTokenExpiresAt: refreshClaims.RegisteredClaims.ExpiresAt.Time,
		User:                  toUserRes(created),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(GenericRes[AuthUserRes]{
		Data:    res,
		Success: true,
	})
}

func (h *handler) refresh(w http.ResponseWriter, r *http.Request) {
	var req RefreshReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	refreshClaims, err := h.tokenMaker.VerifyToken(req.RefreshToken)

	if err != nil {
		http.Error(w, "error verifying token", http.StatusUnauthorized)
		return
	}

	session, err := h.server.GetSession(h.ctx, refreshClaims.RegisteredClaims.ID)

	if err != nil {
		http.Error(w, "error getting session", http.StatusInternalServerError)
		return
	}

	if session.IsRevoked {
		http.Error(w, "session revoked", http.StatusUnauthorized)
		return
	}

	if session.UserEmail != refreshClaims.Email {
		http.Error(w, "invalid session", http.StatusUnauthorized)
		return
	}

	accessToken, accessClaims, err := h.tokenMaker.CreateToken(refreshClaims.ID, refreshClaims.Email, ACCESS_TOKEN_DURATION)

	if err != nil {
		http.Error(w, "error creating token", http.StatusInternalServerError)
		return
	}

	res := RefreshRes{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessClaims.RegisteredClaims.ExpiresAt.Time,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(GenericRes[RefreshRes]{
		Data:    res,
		Success: true,
	})
}

func (h *handler) logout(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(authKey{}).(*token.UserClaims)

	err := h.server.DeleteSession(h.ctx, claims.RegisteredClaims.ID)

	if err != nil {
		http.Error(w, "error deleting session", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) getMe(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(authKey{}).(*token.UserClaims)

	res, err := h.server.GetUserByEmail(h.ctx, claims.Email)

	if err != nil {
		http.Error(w, "error to fetch user data", http.StatusInternalServerError)
		return
	}

	user := toUserRes(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(GenericRes[UserRes]{
		Data:    user,
		Success: true,
	})
}

func toStorerUser(u RegisterReq) *storer.User {
	return &storer.User{
		ID:        primitive.NewObjectID(),
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		Photo:     u.Photo,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func toUserRes(u *storer.User) UserRes {
	return UserRes{
		Name:  u.Name,
		Email: u.Email,
		Photo: u.Photo,
	}
}

func (h *handler) createBooking(w http.ResponseWriter, r *http.Request) {
	var req BookingReq
	claims := r.Context().Value(authKey{}).(*token.UserClaims)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	startDate, err := util.ParseTime(req.StartDate)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "parse time failed", http.StatusBadRequest)
		return
	}

	endDate, err := util.ParseTime(req.EndDate)

	if err != nil {
		http.Error(w, "parse time failed", http.StatusBadRequest)
		return
	}

	created, err := h.server.CreateBooking(h.ctx, &storer.Booking{
		ID:        primitive.NewObjectID(),
		UserID:    claims.ID,
		RoomID:    req.RoomID,
		CheckIn:   toStorerCheckInDetail(req.CheckIn),
		Price:     req.Price,
		StartDate: startDate,
		EndDate:   endDate,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		http.Error(w, "create booking failed", http.StatusInternalServerError)
		return
	}

	res := toBookingRes(created)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(GenericRes[BookingRes]{
		Data:    res,
		Success: true,
	})

}

func (h *handler) listBookings(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(authKey{}).(*token.UserClaims)

	results, err := h.server.ListBookings(h.ctx, claims.ID)

	if err != nil {
		log.Println("error", err)
		http.Error(w, "list bookings by user id failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(GenericRes[[]storer.ListBookingRes]{
		Data:    results,
		Success: true,
	})
}

func (h *handler) deleteBooking(w http.ResponseWriter, r *http.Request) {
	bookingId := chi.URLParam(r, "bookingId")
	claims := r.Context().Value(authKey{}).(*token.UserClaims)

	mongoParsedId, err := primitive.ObjectIDFromHex(bookingId)

	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	req := storer.DeleteBookingReq{
		ID:     mongoParsedId,
		UserID: claims.ID,
	}

	err = h.server.DeleteBooking(h.ctx, req)

	if err != nil {
		http.Error(w, "delete booking failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(GenericRes[string]{
		Data:    "delete booking success",
		Success: true,
	})

}

func toStorerCheckInDetail(checkIn CheckInDetailRes) storer.CheckInDetail {
	return storer.CheckInDetail{
		Adult:  checkIn.Adult,
		Child:  checkIn.Child,
		Infant: checkIn.Infant,
	}
}

func toCheckInDetail(ci storer.CheckInDetail) CheckInDetailRes {
	return CheckInDetailRes{
		Adult:  ci.Adult,
		Child:  ci.Child,
		Infant: ci.Infant,
	}
}

func toBookingRes(b *storer.Booking) BookingRes {
	return BookingRes{
		ID:        b.ID,
		RoomID:    b.RoomID,
		CheckIn:   toCheckInDetail(b.CheckIn),
		StartDate: b.StartDate,
		EndDate:   b.EndDate,
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
	}
}

func (h *handler) upsertWishList(w http.ResponseWriter, r *http.Request) {
	var req WishListReq
	claims := r.Context().Value(authKey{}).(*token.UserClaims)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	err := h.server.UpsertWishList(h.ctx, storer.WishListPayload{
		UserID: claims.ID,
		RoomID: req.RoomID,
		Liking: req.Liking,
	})

	if err != nil {
		http.Error(w, "liking room failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(GenericRes[string]{
		Data:    "upsert room to wish list",
		Success: true,
	})
}

func (h *handler) listWishListIds(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(authKey{}).(*token.UserClaims)

	results, err := h.server.ListWishListIds(h.ctx, claims.ID)

	if err != nil {
		http.Error(w, "list wish list by user id failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(GenericRes[[]string]{
		Data:    results,
		Success: true,
	})
}

func (h *handler) getRoomById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "roomId")

	room, err := h.server.GetRoomById(h.ctx, id)

	if err != nil {
		fmt.Println("err: ", err)
		http.Error(w, "get room by id failed", http.StatusInternalServerError)
		return
	}

	res := toRoomWithReviewsRes(room)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(GenericRes[RoomWithReviewsRes]{
		Data:    res,
		Success: true,
	})
}

func (h *handler) getWishListItems(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(authKey{}).(*token.UserClaims)

	wishlists, err := h.server.ListWishListItems(h.ctx, claims.ID)

	if err != nil {
		http.Error(w, "list wish list item failed", http.StatusInternalServerError)
		return
	}

	res := make([]WishListItemRes, len(wishlists))

	for i, item := range wishlists {
		res[i] = toWishListItemRes(item)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(GenericRes[[]WishListItemRes]{
		Data:    res,
		Success: true,
	})
}

func toRoomRes(s *storer.Room) RoomRes {
	return RoomRes{
		ID:           s.ID,
		Name:         s.Name,
		Summary:      s.Summary,
		Space:        s.Space,
		Description:  s.Description,
		PropertyType: s.PropertyType,
		RoomType:     s.RoomType,
		BedType:      s.BedType,
		Accommodates: s.Accommodates,
		Bedrooms:     s.Bedrooms,
		Beds:         s.Beds,
		Bathrooms:    s.Bathrooms,
		Amenities:    s.Amenities,
		Price:        s.Price,
		Image:        s.Image,
		Address:      s.Address,
	}
}

func toReviewDetail(r *storer.ReviewDetail) ReviewDetail {
	return ReviewDetail{
		ID:      r.ID,
		Comment: r.Comment,
		Rating:  r.Rating,
		User: UserRes{
			Name:  r.User.Name,
			Email: r.User.Email,
			Photo: r.User.Photo,
		},
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}
}

func toRoomWithReviewsRes(s *storer.RoomWithReviews) RoomWithReviewsRes {
	fmt.Println("len of reviews", len(s.Reviews))
	res := make([]ReviewDetail, len(s.Reviews))

	for i, review := range s.Reviews {
		res[i] = toReviewDetail(&review)
	}

	return RoomWithReviewsRes{
		ID:           s.ID,
		Name:         s.Name,
		Summary:      s.Summary,
		Space:        s.Space,
		Description:  s.Description,
		PropertyType: s.PropertyType,
		RoomType:     s.RoomType,
		BedType:      s.BedType,
		Accommodates: s.Accommodates,
		Bedrooms:     s.Bedrooms,
		Beds:         s.Beds,
		Bathrooms:    s.Bathrooms,
		Amenities:    s.Amenities,
		Price:        s.Price,
		Image:        s.Image,
		Address:      s.Address,
		Reviews:      res,
	}
}

func toWishListItemRes(s storer.WishListItem) WishListItemRes {
	return WishListItemRes{
		ID:        s.ID,
		Room:      toRoomRes(&s.Room),
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}

// review
func (h *handler) createReview(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(authKey{}).(*token.UserClaims)

	var req CreateReviewReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	err := h.server.CreateReview(h.ctx, storer.Review{
		ID:        primitive.NewObjectID(),
		RoomID:    req.RoomID,
		UserID:    claims.ID,
		Comment:   req.Comment,
		Rating:    req.Rating,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		http.Error(w, "create review failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(GenericRes[string]{
		Data:    "create review success",
		Success: true,
	})
}

func (h *handler) deleteReview(w http.ResponseWriter, r *http.Request) {
	reviewId := chi.URLParam(r, "reviewId")
	claims := r.Context().Value(authKey{}).(*token.UserClaims)

	mongoParsedId, err := primitive.ObjectIDFromHex(reviewId)

	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	err = h.server.DeleteReview(h.ctx, storer.DeleteReviewReq{
		ID:     mongoParsedId,
		UserID: claims.ID,
	})

	if err != nil {
		http.Error(w, "delete review failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(GenericRes[string]{
		Data:    "delete review success",
		Success: true,
	})
}

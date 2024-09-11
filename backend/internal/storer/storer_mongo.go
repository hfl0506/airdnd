package storer

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStorer struct {
	db *mongo.Database
}

func NewMongoStorer(db *mongo.Database) *MongoStorer {
	return &MongoStorer{
		db: db,
	}
}

type PaginateParams struct {
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
	TabID string `json:"tab_id"`
}

func (ms *MongoStorer) ListingAndReviewsRecords(ctx context.Context, params PaginateParams) (ListingPaginateRes, error) {
	var results []ListingAndReviews
	offset := (params.Page - 1) * params.Limit

	query := bson.D{}

	if params.TabID != "" {
		query = bson.D{
			{Key: "property_type", Value: params.TabID},
		}
	}

	projection := bson.D{
		{Key: "id", Value: 1},
		{Key: "listing_url", Value: 1},
		{Key: "name", Value: 1},
		{Key: "summary", Value: 1},
		{Key: "images", Value: 1},
		{Key: "address", Value: 1},
		{Key: "bedrooms", Value: 1},
		{Key: "beds", Value: 1},
		{Key: "room_type", Value: 1},
		{Key: "bed_type", Value: 1},
		{Key: "price", Value: 1},
	}

	totalCount, err := ms.db.Collection("listingsAndReviews").CountDocuments(ctx, query)

	if err != nil {
		return ListingPaginateRes{}, err
	}

	cursor, err := ms.db.Collection("listingsAndReviews").Find(ctx, query, options.Find().SetSkip(int64(offset)).SetLimit(int64(params.Limit)).SetProjection(projection))

	if err != nil {
		return ListingPaginateRes{}, err
	}

	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &results); err != nil {
		return ListingPaginateRes{}, err
	}

	return ListingPaginateRes{
		Data:  results,
		Total: totalCount,
	}, nil
}

func (ms *MongoStorer) ListCategory(ctx context.Context) (ListCategoryRes, error) {
	pipeline := mongo.Pipeline{
		{
			{Key: "$group", Value: bson.D{
				{Key: "_id", Value: nil},
				{Key: "roomType", Value: bson.D{
					{Key: "$addToSet", Value: "$room_type"},
				}},
				{Key: "propertyType", Value: bson.D{
					{Key: "$addToSet", Value: "$property_type"},
				}},
				{Key: "bedType", Value: bson.D{
					{Key: "$addToSet", Value: "$bed_type"},
				}},
			}},
		},
		{
			{Key: "$set", Value: bson.D{
				{Key: "roomType", Value: bson.D{
					{Key: "$sortArray", Value: bson.D{
						{Key: "input", Value: "$roomType"},
						{Key: "sortBy", Value: 1},
					}},
				}},
				{Key: "propertyType", Value: bson.D{
					{Key: "$sortArray", Value: bson.D{
						{Key: "input", Value: "$propertyType"},
						{Key: "sortBy", Value: 1},
					}},
				}},
				{Key: "bedType", Value: bson.D{
					{Key: "$sortArray", Value: bson.D{
						{Key: "input", Value: "$bedType"},
						{Key: "sortBy", Value: 1},
					}},
				}},
			}},
		},
		{
			{Key: "$project", Value: bson.D{
				{Key: "_id", Value: 0},
			}},
		},
	}

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)

	defer cancel()

	cursor, err := ms.db.Collection("listingsAndReviews").Aggregate(ctx, pipeline)

	if err != nil {
		return ListCategoryRes{}, err
	}

	defer cursor.Close(ctx)

	var res []ListCategoryRes

	if err = cursor.All(ctx, &res); err != nil {
		return ListCategoryRes{}, nil
	}

	if len(res) == 0 {
		return ListCategoryRes{}, nil
	}

	return res[0], nil
}

func (ms *MongoStorer) CreateUser(ctx context.Context, user *User) (*mongo.InsertOneResult, error) {
	collection := ms.db.Collection("users")

	result, err := collection.InsertOne(ctx, user)

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ms *MongoStorer) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	query := bson.M{
		"email": email,
	}
	var user User

	collection := ms.db.Collection("users")

	err := collection.FindOne(ctx, query).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (ms *MongoStorer) GetSession(ctx context.Context, id string) (*Session, error) {
	var session Session
	collection := ms.db.Collection("sessions")

	err := collection.FindOne(ctx, bson.D{{Key: "id", Value: id}}).Decode(&session)

	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (ms *MongoStorer) CreateSession(ctx context.Context, s *Session) (*Session, error) {
	collection := ms.db.Collection("sessions")

	_, err := collection.InsertOne(ctx, s)

	if err != nil {
		return nil, err
	}

	return s, nil
}

func (ms *MongoStorer) DeleteSession(ctx context.Context, id string) error {
	collection := ms.db.Collection("sessions")

	_, err := collection.DeleteOne(ctx, bson.D{{Key: "_id", Value: id}})

	if err != nil {
		return err
	}

	return nil
}

// bookings
func (ms *MongoStorer) CreateBooking(ctx context.Context, b *Booking) (*Booking, error) {
	collection := ms.db.Collection("bookings")

	_, err := collection.InsertOne(ctx, b)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func (ms *MongoStorer) ListBookings(ctx context.Context, userId primitive.ObjectID) ([]ListBookingRes, error) {
	var results []ListBookingRes

	bookings := ms.db.Collection("bookings")

	pipeline := mongo.Pipeline{
		{
			{Key: "$match", Value: bson.D{
				{Key: "user_id", Value: userId},
			}},
		},
		{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "listingsAndReviews"},
				{Key: "localField", Value: "room_id"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "room"},
			}},
		},
		{
			{Key: "$unwind", Value: bson.D{
				{Key: "path", Value: "$room"},
			}},
		},
		{
			{Key: "$project", Value: bson.D{
				{Key: "_id", Value: 0},
				{Key: "id", Value: "$_id"},
				{Key: "room", Value: bson.D{
					{Key: "id", Value: "$room._id"},
					{Key: "name", Value: "$room.name"},
					{Key: "address", Value: "$room.address.street"},
				}},
				{Key: "checkIn", Value: "$check_in"},
				{Key: "startDate", Value: "$start_date"},
				{Key: "endDate", Value: "$end_date"},
				{Key: "createdAt", Value: "$created_at"},
				{Key: "updatedAt", Value: "$updated_at"},
			}},
		},
	}

	cursor, err := bookings.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (ms *MongoStorer) DeleteBooking(ctx context.Context, params DeleteBookingReq) error {
	collection := ms.db.Collection("bookings")

	filter := bson.D{
		{Key: "user_id", Value: params.UserID},
		{Key: "_id", Value: params.ID},
	}

	_, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	return nil
}

func (ms *MongoStorer) UpsertWishList(ctx context.Context, params WishListPayload) error {
	collection := ms.db.Collection("wishLists")

	filter := bson.D{{Key: "user_id", Value: params.UserID}, {Key: "room_id", Value: params.RoomID}}

	update := bson.D{{
		Key: "$set", Value: bson.D{{Key: "liking", Value: params.Liking}},
	}}

	opts := options.Update().SetUpsert(true)

	_, err := collection.UpdateOne(ctx, filter, update, opts)

	if err != nil {
		return err
	}

	return nil
}

func (ms *MongoStorer) ListWishListIds(ctx context.Context, userId primitive.ObjectID) ([]string, error) {
	collection := ms.db.Collection("wishLists")

	filter := bson.D{{Key: "user_id", Value: userId}, {Key: "liking", Value: true}}

	opts := options.Find().SetProjection(bson.D{
		{Key: "roomId", Value: "$room_id"},
	})
	cursor, err := collection.Find(ctx, filter, opts)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var result []RoomID

	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	roomIds := make([]string, len(result))

	for i, res := range result {
		roomIds[i] = res.RoomID
	}

	return roomIds, nil
}

func (ms *MongoStorer) GetWishListItems(ctx context.Context, userId primitive.ObjectID) ([]WishListItem, error) {
	var res []WishListItem

	wishlists := ms.db.Collection("wishLists")

	pipeline := mongo.Pipeline{
		{
			{Key: "$match", Value: bson.D{
				{Key: "user_id", Value: userId},
				{Key: "liking", Value: true},
			}},
		},
		{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "listingsAndReviews"},
				{Key: "localField", Value: "room_id"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "room"},
			}},
		},
		{
			{Key: "$unwind", Value: bson.D{
				{Key: "path", Value: "$room"},
			}},
		},
		{
			{Key: "$project", Value: bson.D{
				{Key: "_id", Value: 1},
				{Key: "room._id", Value: 1},
				{Key: "room.name", Value: 1},
				{Key: "room.summary", Value: 1},
				{Key: "room.space", Value: 1},
				{Key: "room.description", Value: 1},
				{Key: "room.property_type", Value: 1},
				{Key: "room.room_type", Value: 1},
				{Key: "room.bed_type", Value: 1},
				{Key: "room.accommodates", Value: 1},
				{Key: "room.bedrooms", Value: 1},
				{Key: "room.beds", Value: 1},
				{Key: "room.bathrooms", Value: 1},
				{Key: "room.amenities", Value: 1},
				{Key: "room.price", Value: 1},
				{Key: "room.image", Value: "$room.images.picture_url"},
				{Key: "room.address", Value: "$room.address.street"},
			}},
		},
	}

	cursor, err := wishlists.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (ms *MongoStorer) GetRoomById(ctx context.Context, id string) (*RoomWithReviews, error) {
	var res []RoomWithReviews

	pipeline := mongo.Pipeline{
		{
			{Key: "$match", Value: bson.D{
				{Key: "_id", Value: id},
			}},
		},
		{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "reviews"},
				{Key: "localField", Value: "_id"},
				{Key: "foreignField", Value: "room_id"},
				{Key: "as", Value: "reviews"},
			}},
		},
		{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "users"},
				{Key: "localField", Value: "reviews.user_id"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "review_users"},
			}},
		},
		{
			{Key: "$addFields", Value: bson.D{
				{Key: "reviews", Value: bson.D{
					{Key: "$map", Value: bson.D{
						{Key: "input", Value: "$reviews"},
						{Key: "as", Value: "review"},
						{Key: "in", Value: bson.D{
							{Key: "_id", Value: "$$review._id"},
							{Key: "comment", Value: "$$review.comment"},
							{Key: "rating", Value: "$$review.rating"},
							{Key: "created_at", Value: "$$review.created_at"},
							{Key: "updated_at", Value: "$$review.updated_at"},
							{Key: "user", Value: bson.D{
								{Key: "$arrayElemAt", Value: bson.A{
									"$review_users",
									bson.D{{Key: "$indexOfArray", Value: bson.A{"$review_users._id", "$$review.user_id"}}},
								}},
							}},
						}},
					}},
				}},
			}},
		},
		{
			{Key: "$group", Value: bson.D{
				{Key: "_id", Value: "$_id"},
				{Key: "name", Value: bson.D{
					{Key: "$first", Value: "$name"},
				}},
				{Key: "summary", Value: bson.D{
					{Key: "$first", Value: "$summary"},
				}},
				{Key: "space", Value: bson.D{
					{Key: "$first", Value: "$space"},
				}},
				{Key: "description", Value: bson.D{
					{Key: "$first", Value: "$description"},
				}},
				{Key: "property_type", Value: bson.D{
					{Key: "$first", Value: "$property_type"},
				}},
				{Key: "room_type", Value: bson.D{
					{Key: "$first", Value: "$room_type"},
				}},
				{Key: "bed_type", Value: bson.D{
					{Key: "$first", Value: "$bed_type"},
				}},
				{Key: "accommodates", Value: bson.D{
					{Key: "$first", Value: "$accommodates"},
				}},
				{Key: "bedrooms", Value: bson.D{
					{Key: "$first", Value: "$bedrooms"},
				}},
				{Key: "beds", Value: bson.D{
					{Key: "$first", Value: "$beds"},
				}},
				{Key: "bathrooms", Value: bson.D{
					{Key: "$first", Value: "$bathrooms"},
				}},
				{Key: "amenities", Value: bson.D{
					{Key: "$first", Value: "$amenities"},
				}},
				{Key: "price", Value: bson.D{
					{Key: "$first", Value: "$price"},
				}},
				{Key: "image", Value: bson.D{
					{Key: "$first", Value: "$images.picture_url"},
				}},
				{Key: "address", Value: bson.D{
					{Key: "$first", Value: "$address.street"},
				}},
				{Key: "reviews", Value: bson.D{
					{Key: "$first", Value: "$reviews"},
				}},
			}},
		},
	}

	collection := ms.db.Collection("listingsAndReviews")

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &res); err != nil {
		log.Println("err", err)
		return nil, err
	}

	if len(res) == 0 {
		return nil, nil
	}

	return &res[0], nil
}

// review
func (ms *MongoStorer) CreateReview(ctx context.Context, payload Review) error {
	collection := ms.db.Collection("reviews")

	_, err := collection.InsertOne(ctx, payload)

	if err != nil {
		return err
	}

	return nil
}

func (ms *MongoStorer) DeleteReview(ctx context.Context, payload DeleteReviewReq) error {
	collection := ms.db.Collection("reviews")

	filter := bson.D{
		{Key: "_id", Value: payload.ID},
		{Key: "user_id", Value: payload.UserID},
	}
	_, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	return nil
}

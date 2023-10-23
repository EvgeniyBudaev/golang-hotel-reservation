package db

import (
	"context"
	"github.com/EvgeniyBudaev/golang-hotel-reservation/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookingStore interface {
	InsertBooking(context.Context, *models.Booking) (*models.Booking, error)
	GetBookings(context.Context, bson.M) ([]*models.Booking, error)
}

type MongoBookingStore struct {
	client *mongo.Client
	coll   *mongo.Collection
	HotelStore
}

func NewMongoBookingStore(client *mongo.Client) *MongoBookingStore {
	return &MongoBookingStore{
		client: client,
		coll:   client.Database(DBNAME).Collection("bookings"),
	}
}

func (s *MongoBookingStore) GetBookings(ctx context.Context, filter bson.M) ([]*models.Booking, error) {
	curr, err := s.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var bookings []*models.Booking
	if err := curr.All(ctx, &bookings); err != nil {
		return nil, err
	}
	return bookings, nil
}

func (s *MongoBookingStore) InsertBooking(ctx context.Context, booking *models.Booking) (*models.Booking, error) {
	resp, err := s.coll.InsertOne(ctx, booking)
	if err != nil {
		return nil, err
	}
	booking.ID = resp.InsertedID.(primitive.ObjectID)
	return booking, nil
}

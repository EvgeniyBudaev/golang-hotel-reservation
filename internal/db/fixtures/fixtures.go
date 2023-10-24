package fixtures

import (
	"context"
	"fmt"
	"github.com/EvgeniyBudaev/golang-hotel-reservation/internal/db"
	"github.com/EvgeniyBudaev/golang-hotel-reservation/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

func AddBooking(store *db.Store, userID, roomID primitive.ObjectID, from, till time.Time) *models.Booking {
	booking := &models.Booking{
		UserID:   userID,
		RoomID:   roomID,
		FromDate: from,
		TillDate: till,
	}
	insertedBooking, err := store.Booking.InsertBooking(context.Background(), booking)
	if err != nil {
		log.Fatal(err)
	}
	return insertedBooking
}

func AddRoom(store *db.Store, size string, ss bool, price float64, hotelID primitive.ObjectID) *models.Room {
	room := &models.Room{
		Size:    size,
		Seaside: ss,
		Price:   price,
		HotelID: hotelID,
	}
	insertedRoom, err := store.Room.InsertRoom(context.Background(), room)
	if err != nil {
		log.Fatal(err)
	}
	return insertedRoom
}

func AddHotel(store *db.Store, name, location string, rating int, rooms []primitive.ObjectID) *models.Hotel {
	var roomIDS = rooms
	if rooms == nil {
		roomIDS = []primitive.ObjectID{}
	}
	hotel := models.Hotel{
		Name:     name,
		Location: location,
		Rooms:    roomIDS,
		Rating:   rating,
	}
	insertedHotel, err := store.Hotel.InsertHotel(context.TODO(), &hotel)
	if err != nil {
		log.Fatal(err)
	}
	return insertedHotel
}

func AddUser(store *db.Store, firstName, lastName string, admin bool) *models.User {
	user, err := models.NewUserFromParams(models.CreateUserParams{
		Email:     fmt.Sprintf("%s@%s.com", firstName, lastName),
		FirstName: firstName,
		LastName:  lastName,
		Password:  fmt.Sprintf("%s %s", firstName, lastName),
	})
	if err != nil {
		log.Fatal(err)
	}
	user.IsAdmin = admin
	insertedUser, err := store.User.InsertUser(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	return insertedUser
}

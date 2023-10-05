package main

import (
	"context"
	"fmt"
	"github.com/EvgeniyBudaev/golang-hotel-reservation/internal/db"
	"github.com/EvgeniyBudaev/golang-hotel-reservation/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func main() {
	ctx := context.Background()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	hotelStore := db.NewMongoHotelStore(client, db.DBNAME)
	roomStore := db.NewMongoRoomStore(client, db.DBNAME)
	hotel := models.Hotel{
		Name:     "Bellucia",
		Location: "France",
	}
	rooms := []models.Room{
		{
			Type:      models.SingleRoomType,
			BasePrice: 99.9,
		},
		{
			Type:      models.DoubleRoomType,
			BasePrice: 122.9,
		},
		{
			Type:      models.DeluxeRoomType,
			BasePrice: 1999.9,
		},
	}
	insertedHotel, err := hotelStore.InsertHotel(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}
	for _, room := range rooms {
		room.HotelID = insertedHotel.ID
		insertedRoom, err := roomStore.InsertRoom(ctx, &room)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(insertedRoom)
	}
}

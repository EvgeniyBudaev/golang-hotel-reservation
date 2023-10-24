package api

import (
	"fmt"
	"github.com/EvgeniyBudaev/golang-hotel-reservation/internal/db/fixtures"
	"testing"
	"time"
)

func TestGetBookings(t *testing.T) {
	db := setup(t)
	defer db.teardown(t)
	user := fixtures.AddUser(db.Store, "james", "foo", false)
	hotel := fixtures.AddHotel(db.Store, "bar hotel", "a", 4, nil)
	room := fixtures.AddRoom(db.Store, "small", true, 4.4, hotel.ID)
	from := time.Now()
	till := from.AddDate(0, 0, 5)
	booking := fixtures.AddBooking(db.Store, user.ID, room.ID, from, till)
	fmt.Println(booking)
}

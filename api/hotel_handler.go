package api

import (
	"fmt"
	"github.com/EvgeniyBudaev/golang-hotel-reservation/internal/db"
	"github.com/gofiber/fiber/v2"
)

type HotelHandler struct {
	roomStore  db.RoomStore
	hotelStore db.HotelStore
}

func NewHotelHandler(hs db.HotelStore, rs db.RoomStore) *HotelHandler {
	return &HotelHandler{
		hotelStore: hs,
		roomStore:  rs,
	}
}

type HotelQueryParams struct {
	Rooms  bool
	Rating int
}

func (h *HotelHandler) HandleGetHotels(ctx *fiber.Ctx) error {
	var qparams HotelQueryParams
	if err := ctx.QueryParser(&qparams); err != nil {
		return err
	}
	fmt.Println(qparams)
	hotels, err := h.hotelStore.GetHotels(ctx.Context(), nil)
	if err != nil {
		return err
	}
	return ctx.JSON(hotels)
}

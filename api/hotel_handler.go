package api

import (
	"github.com/EvgeniyBudaev/golang-hotel-reservation/internal/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HotelHandler struct {
	store *db.Store
}

func NewHotelHandler(store *db.Store) *HotelHandler {
	return &HotelHandler{
		store: store,
	}
}

func (h *HotelHandler) HandleGetRooms(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidID()
	}
	filter := bson.M{"hotelID": oid}
	rooms, err := h.store.Room.GetRooms(ctx.Context(), filter)
	if err != nil {
		return ErrNotResourceNotFound("hotel")
	}
	return ctx.JSON(rooms)
}

func (h *HotelHandler) HandleGetHotel(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	hotel, err := h.store.Hotel.GetHotelByID(ctx.Context(), id)
	if err != nil {
		return ErrNotResourceNotFound("hotel")
	}
	return ctx.JSON(hotel)
}

type ResourceResp struct {
	Results int `json:"results"`
	Data    any `json:"data"`
	Page    int `json:"page"`
}

type HotelQueryParams struct {
	db.Pagination
	Rating int
}

func (h *HotelHandler) HandleGetHotels(ctx *fiber.Ctx) error {
	var params HotelQueryParams
	if err := ctx.QueryParser(&params); err != nil {
		return ErrBadRequest()
	}
	filter := db.Map{
		"rating": params,
	}
	hotels, err := h.store.Hotel.GetHotels(ctx.Context(), filter, &params.Pagination)
	if err != nil {
		return ErrNotResourceNotFound("hotels")
	}
	resp := ResourceResp{
		Data:    hotels,
		Results: len(hotels),
		Page:    int(params.Page),
	}
	return ctx.JSON(resp)
}

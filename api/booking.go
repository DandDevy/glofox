package api

import (
	"github.com/DandDevy/glofox/internal"
	"github.com/labstack/echo/v4"
	"net/http"
)

// bookingRoutes for routing bookings
type bookingRoutes struct {
	serviceProvider *internal.ServiceProvider
}

// newBookingRoutes returns new bookingRoutes
func newBookingRoutes(serviceProvider *internal.ServiceProvider) *bookingRoutes {
	return &bookingRoutes{serviceProvider: serviceProvider}
}


// createBookingRequest represents a JSON request for creating a booking.
type createBookingRequest struct {
	CreateBooking struct{
		Name string `json:"name"`
		Date string `json:"date"`
	} `json:"booking"`
}

func (r *createBookingRequest) ToBooking() internal.Booking {
	return internal.Booking{
		Name: r.CreateBooking.Name,
		Date: r.CreateBooking.Date,
	}
}


func (r bookingRoutes) handle(e *echo.Echo) {
	e.POST("bookings", r.create)
}

func (r bookingRoutes) create(context echo.Context) error {
	var req createBookingRequest

	if err := context.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	booking := req.ToBooking()
	add, err := r.serviceProvider.Booking().Add(booking)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusCreated, add)
}

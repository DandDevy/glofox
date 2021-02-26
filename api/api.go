package api

import (
	"github.com/DandDevy/glofox/internal"
	"github.com/labstack/echo/v4"
)

// InitRoutes sets up routing.
func InitRoutes(e *echo.Echo, serviceProvider *internal.ServiceProvider) {
	class(serviceProvider).handle(e)
	booking(serviceProvider).handle(e)
}

// class sets up routes for class handlers.
func class(serviceProvider *internal.ServiceProvider) *classRoutes {
	return newClassRoutes(serviceProvider)
}

// booking sets up routes for booking handlers.
func booking(serviceProvider *internal.ServiceProvider) *bookingRoutes {
	return newBookingRoutes(serviceProvider)
}

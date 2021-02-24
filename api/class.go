package api

import (
	"fmt"
	"github.com/DandDevy/glofox/internal"
	"github.com/labstack/echo/v4"
	"net/http"
)

// classRoutes for routing classes
type classRoutes struct {
	serviceProvider *internal.ServiceProvider
}

// newClassRoutes returns new classRoutes
func newClassRoutes(serviceProvider *internal.ServiceProvider) *classRoutes {
	return &classRoutes{serviceProvider: serviceProvider}
}

type createClassRequest struct {
	createClass struct {
		name      string `json:"name"`
		startDate string `json:"start_date"`
		endDate   string `json:"end_date"`
		capacity  uint `json:"capacity"`
	} `json:"class"`
}



func (* classRoutes) serve(e *echo.Echo)  {
	e.POST("/classes", create)
}

func create(context echo.Context) error {

	var req createClassRequest
	if err := context.Bind(&req); err != nil {
		return err
	}

	fmt.Println(req)
	return context.JSON(http.StatusOK, req)
}

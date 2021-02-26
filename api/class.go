package api

import (
	"github.com/DandDevy/glofox/internal"
	"github.com/labstack/echo/v4"
	"net/http"
)

// classRoutes for routing classes.
type classRoutes struct {
	serviceProvider *internal.ServiceProvider
}

// newClassRoutes returns new classRoutes.
func newClassRoutes(serviceProvider *internal.ServiceProvider) *classRoutes {
	return &classRoutes{serviceProvider: serviceProvider}
}

// createClassRequest represents a JSON request for creating a class.
type createClassRequest struct {
	CreateClass struct {
		Name      string `json:"name"`
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date" `
		Capacity  uint   `json:"capacity"`
	} `json:"class"`
}

func (r *createClassRequest) toClass() internal.Class {
	return internal.Class{
		Name: r.CreateClass.Name,
		StartDate: r.CreateClass.StartDate,
		EndDate: r.CreateClass.EndDate,
		Capacity: r.CreateClass.Capacity,
	}
}


func (c* classRoutes) handle(e *echo.Echo)  {
	e.POST("/classes", c.create)
	e.GET("/classes", c.list)
}

// create class endpoint
func (c *classRoutes) create(context echo.Context) error {

	var req createClassRequest

	if err := context.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	class := req.toClass()
	add, err := c.serviceProvider.Class().Add(class)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusCreated, add)
}

// list classes endpoint.
func (c *classRoutes) list(context echo.Context) error {

	classes := c.serviceProvider.Class().List()

	return context.JSON(http.StatusOK, classes)
}

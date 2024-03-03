// Description: Main entry point for the application.
package main

import (
	"cern/api"
	"cern/view/Home"
	"cern/view/components"
	"encoding/json"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	// TODO: Add Swagger API documentation

	// Status
	app.GET("/api/health", func(c echo.Context) error {
		data, err := json.MarshalIndent(app.Routes(), "", "  ")
		if err != nil {
			panic(err)
		}

		return c.String(200, string(data))
	})

	app.GET("/", func(c echo.Context) error {
		data, err := api.ReadAllUsers()
		if err != nil {
			panic(err)
		}

		members := Home.TableDataProps{}

		for _, member := range data {
			members.Data = append(members.Data, Home.Member{
				ID:        fmt.Sprint(member["id"].(float64)),
				FirstName: member["firstName"].(string),
				LastName:  member["lastName"].(string),
				Email:     member["email"].(string),
				Age:       fmt.Sprintf("%b", member["age"].(float64)),
				// ExperimentId: member["experimentId"].(string),
			})
		}

		return render(c, Home.Load(members))
	})

	app.GET("/component/add-member-form", func(c echo.Context) error {
		return render(c, components.AddMemberForm())
	})

	app.GET("/component/member-card/:id", func(c echo.Context) error {
		api.ReadUser(1)

		return render(c, components.MemberCard("Rubens", "Machado", "rubensmachado@gmail.com", "18", "1"))
	})

	// API
	app.POST("/api/create-member", api.CreateUser)
	app.PUT("/api/update-user/:id", api.UpdateUser)
	app.DELETE("/api/delete-user/:id", api.DeleteUser)

	// Static
	app.Static("/static", "static")

	app.Start(":8080")
}

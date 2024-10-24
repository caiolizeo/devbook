package routes

import (
	"api/src/controllers"
	"net/http"
)

var UserRoutes = []Route{
	{
		Uri:                   "/users",
		Method:                http.MethodPost,
		Function:              controllers.CreateUser,
		RequireAuthentication: false,
	},
	{
		Uri:                   "/users",
		Method:                http.MethodGet,
		Function:              controllers.FindUsers,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/users/{id}",
		Method:                http.MethodGet,
		Function:              controllers.FindUser,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/users/{id}",
		Method:                http.MethodPut,
		Function:              controllers.UpdateUser,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/users/{id}",
		Method:                http.MethodDelete,
		Function:              controllers.DeleteUser,
		RequireAuthentication: true,
	},
}

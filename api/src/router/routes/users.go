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
	{
		Uri:                   "/users/{id}/follow",
		Method:                http.MethodPost,
		Function:              controllers.Follow,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/users/{id}/unfollow",
		Method:                http.MethodPost,
		Function:              controllers.Unfollow,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/users/{id}/followers",
		Method:                http.MethodGet,
		Function:              controllers.Followers,
		RequireAuthentication: false,
	},
	{
		Uri:                   "/users/{id}/following",
		Method:                http.MethodGet,
		Function:              controllers.Following,
		RequireAuthentication: false,
	},
}

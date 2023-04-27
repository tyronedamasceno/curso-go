package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesLogin = []Route{
	{
		URI:          "/login",
		Method:       http.MethodPost,
		Function:     controllers.Login,
		RequiresAuth: false,
	},
	{
		URI:          "/reset-password",
		Method:       http.MethodPost,
		Function:     controllers.ResetPassword,
		RequiresAuth: true,
	},
}

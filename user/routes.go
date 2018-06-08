package user

import "github.com/hellojebus/go-envoz-api/router"

var Routes = router.RoutePrefix{
	"/users",
	[]router.Route{
		router.Route{
			"UsersIndex",
			"GET",
			"",
			IndexHandler,
			false,
		},
		router.Route{
			"UsersShow",
			"GET",
			"/{userId}",
			ShowHandler,
			true,
		},
		router.Route{
			"UsersCreate",
			"POST",
			"",
			CreateHandler,
			false,
		},
		router.Route{
			"UsersLogin",
			"POST",
			"/login",
			LoginHandler,
			false,
		},
		router.Route{
			"DeleteHandler",
			"DELETE",
			"/{userId}",
			DeleteHandler,
			true,
		},
		router.Route{
			"UpdateHandler",
			"PUT",
			"/{userId}",
			UpdateHandler,
			true,
		},
	},
}
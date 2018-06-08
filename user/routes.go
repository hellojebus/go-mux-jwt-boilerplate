package user

import "github.com/hellojebus/go-envoz-api/router"

var Routes = router.RoutePrefix{
	"/users",
	[]router.Route{
		router.Route{
			"UsersIndex",
			"GET",
			"",
			UsersIndexHandler,
			false,
		},
		router.Route{
			"UsersShow",
			"GET",
			"/{userId}",
			UsersShowHandler,
			true,
		},
		router.Route{
			"UsersCreate",
			"POST",
			"",
			UsersCreateHandler,
			false,
		},
		router.Route{
			"UsersLogin",
			"POST",
			"/login",
			UsersLoginHandler,
			false,
		},
		router.Route{
			"UsersDelete",
			"DELETE",
			"/{userId}",
			UsersDelete,
			true,
		},
		router.Route{
			"UsersUpdate",
			"PUT",
			"/{userId}",
			UsersUpdate,
			true,
		},
	},
}
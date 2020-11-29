package routes

import (
	. "github.com/lanvard/routing"
	"lanvard/app/http/controllers"
	"lanvard/app/http/middleware"
)

/*
|--------------------------------------------------------------------------
| Web routes
|--------------------------------------------------------------------------
|
| Here is where you can register Web routes for your website. By default
| these routes are loaded within a group which is assigned the "Web"
| middleware. Enjoy building your website!
|
*/
var Web = Group(
	Get("post/create", controllers.PostCreate),
	Post("post", controllers.PostStore),
	Get("/", controllers.Homepage),
).Middleware(middleware.Web...)

package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.GET("/", rt.getHelloWorld)
	rt.router.GET("/context", rt.wrap(rt.getContextReply))
	rt.router.GET("/users", rt.wrap(rt.getAllUsers))
	rt.router.PUT("/users/:userName", rt.wrap(rt.changeUsername))
	rt.router.PUT("/users/:userName/followers/:followerId", rt.wrap(rt.addFollower))
	rt.router.DELETE("/users/:userName/followers/:followerId", rt.wrap(rt.deleteFollower))
	rt.router.GET("/users/:userName", rt.wrap(rt.getUserProfile))
	rt.router.POST("/session", rt.wrap(rt.postSession))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}

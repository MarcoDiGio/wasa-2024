package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.GET("/", rt.getHelloWorld)
	rt.router.GET("/context", rt.wrap(rt.getContextReply))
	rt.router.POST("/session", rt.wrap(rt.postSession))
	rt.router.GET("/users", rt.wrap(rt.getAllUsers))
	rt.router.GET("/users/:userName", rt.wrap(rt.getUserProfile))
	rt.router.PUT("/users/:userName", rt.wrap(rt.changeUsername))
	rt.router.PUT("/users/:userName/followers/:followerId", rt.wrap(rt.addFollower))
	rt.router.DELETE("/users/:userName/followers/:followerId", rt.wrap(rt.deleteFollower))
	rt.router.PUT("/users/:userName/banned/:bannedId", rt.wrap(rt.addBan))
	rt.router.DELETE("/users/:userName/banned/:bannedId", rt.wrap(rt.deleteBan))
	// rt.router.GET("/users/:userName/photos", rt.wrap(rt.getPhotos))
	// rt.router.POST("/users/:userName/photos", rt.wrap(rt.addPhoto))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
